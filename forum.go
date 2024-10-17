package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db          *sql.DB
	UserMessage string
	User        string
)

type Message struct {
	ID           int
	Username     string
	Message      string
	CreatedAt    time.Time
	TimeAgo      string
	Country      string
	ProfilePhoto string
	LikesCount   int
	Liked        bool
}

var UserID = ""

func main() {

	// Handler pour servir la page HTML
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/connexion.html", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "connexion", nil)
	})

	// Gestionnaire pour la soumission du formulaire d'enregistrement
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		// Récupérer les informations d'enregistrement depuis le formulaire
		newUsername := r.FormValue("new_username")
		newPassword := r.FormValue("new_password")
		Pays := r.FormValue("pays")
		// Récupérer le fichier photo de profil
		profilePhoto := r.FormValue("selected_avatar")
		var message string

		// Vérifier si le nom d'utilisateur existe déjà dans la base de données
		if isUsernameExists(newUsername) {
			message := "Le nom d'utilisateur existe déjà."
			renderTemplate(w, "connexion", message)
			return
		} else {
			addUser(newUsername, newPassword, Pays, profilePhoto)
		}

		renderTemplate(w, "connexion", message)
	})

	var username string // Stocker le nom d'utilisateur actuel en variable globale

	// Gestionnaire pour la soumission du formulaire de connexion
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		// Récupérer les informations de connexion depuis le formulaire
		username = r.FormValue("username")
		password := r.FormValue("password")

		// Vérifier les informations d'identification dans la base de données
		if authenticate(username, password) {

			// récupérer l'id de l'utilisateur dans la base de données
			err := db.QueryRow("SELECT id FROM utilisateurs WHERE username = ?", username).Scan(&UserID)
			if err != nil {
				return
			}

			// Rediriger l'utilisateur vers la page du forum
			http.Redirect(w, r, "/forum.html", http.StatusFound)
		} else {
			// Afficher un message d'erreur si les informations d'identification sont incorrectes
			message := "Nom d'utilisateur ou mot de passe incorrect."
			renderTemplate(w, "connexion", message)
		}
	})

	// Handler pour servir la page du forum
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if username != "" && !isUsernameExists(username) {
			username = "" // Réinitialiser le nom d'utilisateur s'il n'existe pas dans la base de données
		}

		// Récupérer les messages de la base de données
		messages := getMessages()

		// Récupérer la photo de profil de l'utilisateur actuel
		profilePhoto := getProfilePhoto(username)

		// Créer une structure de données pour les informations de la page
		type PageData struct {
			Name     string
			Photo    string
			Messages []Message
		}

		// Créer une instance de la structure de données avec les informations nécessaires
		pageData := PageData{
			Name:     username,
			Photo:    profilePhoto,
			Messages: messages,
		}

		renderTemplate(w, "forum", pageData)
	})

	// Gestionnaire pour soumettre un nouveau message sur le forum
	http.HandleFunc("/postmessage", func(w http.ResponseWriter, r *http.Request) {

		// Récupérer le message à partir des paramètres de la requête
		message := r.FormValue("user_message")

		// Ajouter le message à la base de données avec le nom d'utilisateur correspondant
		addMessage(username, message)

		// Rediriger l'utilisateur vers la page du forum avec le même cookie de session
		http.Redirect(w, r, "/forum.html", http.StatusSeeOther)
	})

	// Gestionnaire pour soumettre un nouveau like sur le forum
	http.HandleFunc("/updatelike", func(w http.ResponseWriter, r *http.Request) {
		// Récupérer l'ID du message à partir des paramètres de la requête
		messageID := r.FormValue("messageID")

		// récupérer l'id de l'utilisateur

		if username != "" {
			addlike(username, messageID)
		}

		// Rediriger l'utilisateur vers la page du forum avec le même cookie de session
		http.Redirect(w, r, "/forum.html", http.StatusSeeOther)
	})

	// Démarrer le serveur sur le port 8080
	fmt.Println("Serveur écoutant sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func GetDatabase() *sql.DB {
	if db != nil {
		return db
	}

	// Charger les variables d'environnement à partir du fichier .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erreur lors du chargement du fichier .env:", err)
		return nil
	}

	// Si la connexion n'est pas établie, établissez-la
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	if dbPassword == "" {
		fmt.Println("Erreur: Mot de passe MySQL non défini.")
		return nil
	}

	// Définir les autres informations de connexion
	dbUsername := "root"
	hostname := "localhost"
	port := "3306"
	dbname := "Forum"

	// Former la chaîne de connexion
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, hostname, port, dbname)

	// Se connecter à la base de données
	conn, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	// Vérifier si la connexion est réussie
	err = conn.Ping()
	if err != nil {
		panic(err.Error())
	}

	// Affecter la connexion à la variable db globale
	db = conn
	return db
}

func authenticate(username, password string) bool {
	db := GetDatabase()
	if db == nil {
		return false
	}
	// Vérifier si les informations d'identification sont valides
	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM utilisateurs WHERE username = ? AND password = ?", username, password)
	err := row.Scan(&count)
	if err != nil {
		panic(err.Error())
	}
	return count > 0
}

func addMessage(username, message string) {
	db := GetDatabase()
	if db == nil {
		fmt.Println("Erreur lors de la connexion à la base de données.")
		return
	}

	currentTime := time.Now()

	// Récupérer l'ID de l'utilisateur à partir de la base de données
	var userID int
	err := db.QueryRow("SELECT id FROM utilisateurs WHERE username = ?", username).Scan(&userID)
	if err != nil {
		fmt.Println("Erreur lors de la récupération de l'ID de l'utilisateur:", err)
		return
	}

	// Insérer le nouveau message dans la base de données avec l'ID de l'utilisateur et l'heure actuelle
	_, err = db.Exec("INSERT INTO messages (user_id, message_text, created_at) VALUES (?, ?, ?)", userID, message, currentTime)
	if err != nil {
		fmt.Println("Erreur lors de l'insertion du message:", err)
	}
}

func getProfilePhoto(username string) string {
	db := GetDatabase()
	if db == nil {
		return ""
	}

	// Récupérer la photo de profil de l'utilisateur à partir de la base de données
	var profilePhoto string
	err := db.QueryRow("SELECT profile_photo FROM utilisateurs WHERE username = ?", username).Scan(&profilePhoto)
	if err != nil {
		return ""
	}

	return profilePhoto
}

func getMessages() []Message {
	db := GetDatabase()
	if db == nil {
		return nil
	}

	rows, err := db.Query(`
        SELECT utilisateurs.id, utilisateurs.username, utilisateurs.country, utilisateurs.profile_photo,
               messages.id, messages.message_text, messages.likes_count, messages.created_at
        FROM messages
        INNER JOIN utilisateurs ON messages.user_id = utilisateurs.id
        ORDER BY messages.created_at DESC
    `)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des messages:", err)
		return nil
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var userID, messageID, likesCount int
		var username, country, message, createdAtStr, profilePhoto string
		err := rows.Scan(&userID, &username, &country, &profilePhoto, &messageID, &message, &likesCount, &createdAtStr)
		if err != nil {
			fmt.Println("Erreur lors de la lecture des lignes des messages:", err)
			continue
		}

		createdAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			fmt.Println("Erreur lors du parsing de la date/heure:", err)
			continue
		}

		diff := time.Since(createdAt)
		hours := int(diff.Hours())
		minutes := int(diff.Minutes())
		var messageTime string
		if hours >= 24 {
			messageTime = fmt.Sprintf("il y a %d jours", hours/24)
		} else if hours >= 1 {
			messageTime = fmt.Sprintf("il y a %d heures", hours)
		} else if minutes >= 1 {
			messageTime = fmt.Sprintf("il y a %d minutes", minutes)
		} else {
			messageTime = "il y a quelques secondes"
		}

		// Tansformer UserID en int
		UserID, err := strconv.Atoi(UserID)
		if err != nil {
			continue
		}

		liked, err := isLiked(UserID, messageID) // Appel de isLiked avec l'ID de l'utilisateur
		if err != nil {
			fmt.Println("Erreur lors de la vérification si l'utilisateur a déjà aimé ce message:", err)
			continue
		}

		messages = append(messages, Message{
			ID:           messageID,
			Username:     username,
			Country:      country,
			ProfilePhoto: profilePhoto,
			Message:      message,
			LikesCount:   likesCount,
			TimeAgo:      messageTime,
			Liked:        liked,
		})
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Erreur lors de la récupération des messages:", err)
		return nil
	}

	return messages
}

func addlike(username, messageID string) {
	db := GetDatabase()
	if db == nil {
		fmt.Println("Erreur lors de la connexion à la base de données.")
		return
	}

	// Récupérer l'ID de l'utilisateur à partir de la base de données
	var userID int
	err := db.QueryRow("SELECT id FROM utilisateurs WHERE username = ?", username).Scan(&userID)
	if err != nil {
		return
	}

	messageIDInt, err := strconv.Atoi(messageID)
	if err != nil {
		fmt.Println("Erreur lors de la conversion de l'ID du message en entier:", err)
		return
	}

	// Vérifier si l'utilisateur a déjà aimé ce message
	var likeCount int
	err = db.QueryRow("SELECT COUNT(*) FROM liked WHERE user_id = ? AND message_id = ?", userID, messageIDInt).Scan(&likeCount)
	if err != nil {
		fmt.Println("Erreur lors de la vérification si l'utilisateur a déjà aimé ce message:", err)
		return
	}

	if likeCount > 0 {

		// Si l'utilisateur a déjà aimé le message, supprimer le like de la base de données
		_, err = db.Exec("DELETE FROM liked WHERE user_id = ? AND message_id = ?", userID, messageIDInt)
		if err != nil {
			fmt.Println("Erreur lors de la suppression du like de la base de données:", err)
			return
		}

		// Mettre à jour le nombre de likes_count dans la table messages pour retirer le like
		_, err = db.Exec("UPDATE messages SET likes_count = likes_count - 1 WHERE id = ?", messageIDInt)
		if err != nil {
			fmt.Println("Erreur lors de la mise à jour du nombre de likes du message:", err)
			return
		}

		// Arrêtez ici car nous avons supprimé le like
		return
	}

	// Ajouter un like en insérant une nouvelle ligne dans la table liked
	_, err = db.Exec("INSERT INTO liked (user_id, message_id) VALUES (?, ?)", userID, messageIDInt)
	if err != nil {
		fmt.Println("Erreur lors de l'ajout du like dans la table liked:", err)
		return
	}

	// Mettre à jour le nombre de likes_count dans la table messages
	_, err = db.Exec("UPDATE messages SET likes_count = likes_count + 1 WHERE id = ?", messageIDInt)
	if err != nil {
		fmt.Println("Erreur lors de la mise à jour du nombre de likes du message:", err)
		return
	}
}

func isLiked(userID, messageID int) (bool, error) {
	db := GetDatabase()
	if db == nil {
		return false, fmt.Errorf("erreur lors de la connexion à la base de données")
	}

	// Vérifier si l'utilisateur a aimé ce message
	query := "SELECT EXISTS(SELECT 1 FROM liked WHERE user_id = ? AND message_id = ?)"
	var exists bool
	err := db.QueryRow(query, userID, messageID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("erreur lors de la vérification si l'utilisateur a aimé ce message: %v", err)
	}

	// Retourner vrai si l'utilisateur a aimé ce message, sinon faux
	return exists, nil
}

func addUser(username, password, country string, profilePhoto string) {
	db := GetDatabase()

	// Insérer l'utilisateur dans la base de données avec les données binaires de l'image
	_, err := db.Exec("INSERT INTO utilisateurs (username, password, country, profile_photo) VALUES (?, ?, ?, ?)",
		username, password, country, profilePhoto)
	if err != nil {
		fmt.Println("Erreur lors de l'insertion de l'utilisateur:", err)
	}
}

func isUsernameExists(username string) bool {
	db := GetDatabase()

	// Exécuter la requête pour vérifier si le nom d'utilisateur existe déjà
	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM utilisateurs WHERE username = ?", username)
	err := row.Scan(&count)
	if err != nil {
		panic(err.Error())
	}
	return count > 0
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}
}
