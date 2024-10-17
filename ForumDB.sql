CREATE DATABASE  IF NOT EXISTS `forum` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `forum`;
-- MySQL dump 10.13  Distrib 8.0.36, for Win64 (x86_64)
--
-- Host: localhost    Database: forum
-- ------------------------------------------------------
-- Server version	8.3.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `liked`
--

DROP TABLE IF EXISTS `liked`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `liked` (
  `id` int NOT NULL AUTO_INCREMENT,
  `message_id` int DEFAULT NULL,
  `user_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `message_id` (`message_id`,`user_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `liked_ibfk_1` FOREIGN KEY (`message_id`) REFERENCES `messages` (`id`),
  CONSTRAINT `liked_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `utilisateurs` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=183 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `liked`
--

LOCK TABLES `liked` WRITE;
/*!40000 ALTER TABLE `liked` DISABLE KEYS */;
INSERT INTO `liked` VALUES (66,8,1),(47,11,1),(177,19,1),(172,20,1),(109,21,1),(27,22,1),(170,24,1),(105,25,1),(56,26,1),(65,27,1),(132,28,1),(85,28,2),(123,29,1),(130,29,2),(160,30,2),(164,31,1),(166,32,1),(171,33,1),(175,34,1),(174,34,2),(176,35,1),(178,35,2),(182,36,1),(180,38,1);
/*!40000 ALTER TABLE `liked` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `messages`
--

DROP TABLE IF EXISTS `messages`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `messages` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `message_text` text,
  `created_at` datetime DEFAULT NULL,
  `likes_count` int DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `messages_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `utilisateurs` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `messages`
--

LOCK TABLES `messages` WRITE;
/*!40000 ALTER TABLE `messages` DISABLE KEYS */;
INSERT INTO `messages` VALUES (4,1,'adcazcazcazc','2024-03-15 10:13:06',7),(6,1,'salut la team snapchat','2024-03-15 11:07:39',3),(7,1,'ezkfonzilsdvmIUEJVBLUISVJildbvMUIEBVm','2024-03-15 11:35:36',4),(8,2,'zefizeaofhmzobhvmqzrl','2024-03-15 12:53:35',4),(9,3,'je vous nique kazd,fjaiopnqlzcuibhfAIUZHFZOQUILHG','2024-03-15 13:00:06',2),(10,1,'azkjfnalkfbakubvej','2024-03-16 12:50:26',3),(11,4,',ecblazebckzuec','2024-03-16 12:54:22',3),(12,1,'AIFUHeliufvhiseugvikZ','2024-03-20 13:34:59',4),(13,1,'aeiufhzaefhiuezhfviezbyluif\"avy ifyviu zbyfliugybli ygnilu','2024-03-20 13:39:47',3),(14,1,'jhfjuyfjyufyjf','2024-03-20 15:39:38',4),(15,1,'aFafaEGVEfva','2024-03-21 08:05:43',4),(17,1,'eczcezcze','2024-03-21 09:00:09',4),(18,1,'sdgb<srhrdwrwenhr<ehne<','2024-03-21 09:46:13',3),(19,1,'cazc','2024-03-21 10:09:19',5),(20,1,'COUCOU JE SUIS LE KANGOO A KYLIAN ET JE ME PETE LE BIDE AU RICARD','2024-03-21 10:10:35',12),(21,1,'QHUILLIANN','2024-03-21 10:11:29',20),(22,1,'CRINGELIAN','2024-03-21 10:11:40',20),(24,1,'j jh nhunuobnhyubuo','2024-03-22 09:27:08',3),(25,1,'zadadazda','2024-03-22 09:33:13',1),(26,1,'wdnqnqtnqten','2024-03-22 10:23:36',1),(27,1,'zefzfzfzef','2024-03-22 10:40:29',1),(28,2,'GErbEHBhqrene','2024-03-25 09:41:16',2),(29,1,'egzqegegsqgqrqrgqrqer','2024-03-25 09:44:41',2),(30,1,'zefzfeffefa','2024-03-25 10:16:00',1),(31,2,'jnibiub','2024-03-28 08:31:03',1),(32,1,'bonjour la team snap','2024-04-05 08:44:34',1),(33,6,'deéfé\"f','2024-04-16 07:55:53',1),(34,7,'efzijnvrzrzrai','2024-04-16 08:14:39',2),(35,1,'zefzfvzev','2024-04-19 13:12:58',2),(36,2,'kzeofnzeoaiougvaezgqviozeqtvunà_\"açp\'yhglfzpçeflzliubvmuieZHF0I¨FUÄUifàç,ezfIAUEKFVHO8GFBzlvbevhSJEFGVo_fievgBUEFYEZFZEFZFZEFZEFZEFZEFZEFZEFZEFZEFEZFZEFZEFEZ','2024-04-19 13:41:55',1),(37,2,'ijbzdifezluzgehvbqelrhuqikyrfvbquhqk<bgqkruegbqikyvqreougybOQURSGBVqujgvbquokebvoquyrbhvousbvoqzeurkbveqhrbqueyvuevyevoyqrbuoqeyrbquoervbzHzgvjhbzeqiusvhbqijvnzeuiqvbihduihvibfhqi','2024-04-19 13:44:03',0),(38,1,'ebkechavekjq','2024-04-19 18:54:15',1),(39,1,'zjekjsbvksqjevqljv','2024-04-28 13:04:00',0);
/*!40000 ALTER TABLE `messages` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `utilisateurs`
--

DROP TABLE IF EXISTS `utilisateurs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `utilisateurs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `country` varchar(255) DEFAULT NULL,
  `profile_photo` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `utilisateurs`
--

LOCK TABLES `utilisateurs` WRITE;
/*!40000 ALTER TABLE `utilisateurs` DISABLE KEYS */;
INSERT INTO `utilisateurs` VALUES (1,'kiki','kiki','France','static/images/avatar (2).svg'),(2,'Patlamenace','pat','Pologne','static/images/avatar (3).svg'),(3,'lemoche','az','Botswana','static/images/avatar (4).svg'),(4,'test1','tase1','Autriche','static/images/avatar (5).svg'),(5,'test3','test3','France','static/images/avatar (6).svg'),(6,'test4','test4','France','static/images/avatar (7).svg'),(7,'nonbinary','non','France','static/images/avatar (8).svg');
/*!40000 ALTER TABLE `utilisateurs` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-05-30 10:47:54
