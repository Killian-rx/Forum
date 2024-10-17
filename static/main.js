const register = document.getElementById('register')
const login = document.getElementById('login')
const forum = document.getElementById('forum')
const profil = document.getElementById('profilcontainer')
const liked = document.getElementById('like')
const profilpage = document.getElementById('profilpage')
const Chooseavatar = document.getElementById('chooseavatar')
const messages = document.getElementById('messages')


function SignIn() {
    register.style.display = "none"
    login.style.display = "block"
}

function SignUp() {
    register.style.display = "block"
    login.style.display = "none"
}

function PostMessage() {
    if (forum.style.display === "block") {
        messages.style.filter = "blur(0px)"

        forum.style.display = "none"
    } else {
        messages.style.filter = "blur(5px)"
        forum.style.display = "block"
    }

}

function LikedMessage(button) {
    if (likeIcon.style.color === "red") {
        likeIcon.style.color = "black";
    } else {
        likeIcon.style.color = "red";
    }
}


function Profil(){
    if (profil.style.display === "block") {
        profil.style.display = "none"
    } else {
        profil.style.display = "block"
    }
}


function selectAvatar(imageName) {
    document.getElementById('selectedAvatar').value = imageName;
    Chooseavatar.style.display = "none"
}

function Avatar(){
    if (Chooseavatar.style.display === "flex") {
        Chooseavatar.style.display = "none"
    } else {
        Chooseavatar.style.display = "flex"
    }
}

function ProfilPage(){
    if (profilpage.style.display === "block") {
        profilpage.style.display = "none"
        messages.style.filter = "blur(0px)"
    } else {
        profilpage.style.display = "block"
        messages.style.filter = "blur(5px)"
    }
}

function Return() {
    forum.style.display = "none"
    profilpage.style.display = "none"
    messages.style.filter = "blur(0px)"
}

function logout(){
    window.location.href = "/connexion.html";
}


