-- MySQL dump 10.13  Distrib 5.7.21, for Linux (x86_64)
--
-- Host: localhost    Database: onehouronedayphoto
-- ------------------------------------------------------
-- Server version	5.7.21-0ubuntu0.16.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `event_submission_actions`
--

DROP TABLE IF EXISTS `event_submission_actions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `event_submission_actions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `time` timestamp NULL DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `user_id` int(10) unsigned DEFAULT NULL,
  `event_submission_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_event_submission_actions_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `event_submission_actions`
--

LOCK TABLES `event_submission_actions` WRITE;
/*!40000 ALTER TABLE `event_submission_actions` DISABLE KEYS */;
/*!40000 ALTER TABLE `event_submission_actions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `event_submissions`
--

DROP TABLE IF EXISTS `event_submissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `event_submissions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `image_name` varchar(255) DEFAULT NULL,
  `likes_nr` int(10) unsigned DEFAULT NULL,
  `score` double DEFAULT NULL,
  `latitude` double DEFAULT NULL,
  `longitude` double DEFAULT NULL,
  `submission_date` timestamp NULL DEFAULT NULL,
  `event_id` int(10) unsigned DEFAULT NULL,
  `user_id` int(10) unsigned DEFAULT NULL,
  `image_uid` varchar(255) DEFAULT NULL,
  `thumb_url` varchar(255) DEFAULT NULL,
  `image_url` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_event_submissions_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=60 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `event_submissions`
--

LOCK TABLES `event_submissions` WRITE;
/*!40000 ALTER TABLE `event_submissions` DISABLE KEYS */;
INSERT INTO `event_submissions` VALUES (43,'2017-09-05 15:52:40','2017-09-05 15:52:40',NULL,'file:///storage/emulated/0/DCIM/IMG_20170905_175239.jpg',0,0,0,0,'2017-09-05 15:52:40',1,4,'oujxHBRRUOfHNKVnuEIL0HeXEQhhQ1fw','',''),(44,'2017-09-05 15:52:41','2017-09-05 15:52:41',NULL,'file:///storage/emulated/0/DCIM/IMG_20170905_175240.jpg',0,0,0,0,'2017-09-05 15:52:41',3,4,'2rWLTFVs4m0QXXZa5otUTL9Izdwj-DsC','',''),(45,'2017-09-05 16:04:38','2017-09-05 16:04:38',NULL,'file:///storage/emulated/0/DCIM/IMG_20170905_180429.jpg',0,0,0,0,'2017-09-05 16:04:38',1,4,'U80FR00GoTdxasexTgCP8JXGg2Ptbmee','',''),(46,'2017-09-05 16:06:36','2017-09-05 16:06:36',NULL,'file:///storage/emulated/0/DCIM/IMG_20170905_180629.jpg',0,0,0,0,'2017-09-05 16:06:36',1,4,'QCvCuW9yGRI6O4JMijb92KBxgecmwZly','',''),(47,'2017-09-05 16:11:22','2017-09-05 16:11:22',NULL,'file:///storage/emulated/0/DCIM/IMG_20170905_181112.jpg',0,0,0,0,'2017-09-05 16:11:22',3,4,'zbORaGJug3M1JBFJtKEg6vxyh-sKFQvM','',''),(48,'2017-09-05 16:16:14','2017-09-05 16:16:14',NULL,'file:///storage/emulated/0/DCIM/IMG_20170905_181605.jpg',0,0,0,0,'2017-09-05 16:16:14',1,4,'D3muPnoTaO81VMjqj1noBBfLMxbz3sTF','',''),(49,'2017-09-05 16:16:16','2017-09-05 16:16:16',NULL,'file:///storage/emulated/0/DCIM/IMG_20170905_181504.jpg',0,0,0,0,'2017-09-05 16:16:16',1,4,'F-B02sK7EUpS2n6iQHyjmY1CdffTeo86','',''),(50,'2017-09-06 18:47:52','2017-09-06 18:47:52',NULL,'YuanPeng_China_Professional_Sport_2017.jpg',0,0,0,0,'2017-09-06 18:47:52',3,4,'QFkOwr0pk7g7SMDVE2IkNqzGkSHrhIcR','',''),(51,'2017-09-06 20:05:49','2017-09-06 20:05:49',NULL,'YuanPeng_China_Professional_Sport_2017.jpg',0,0,0,0,'2017-09-06 20:05:49',3,4,'gJUYwrIUuEBw0L-wjtbKQ0oh7XPpcBBl','',''),(52,'2017-09-06 20:07:56','2017-09-06 20:07:56',NULL,'YuanPeng_China_Professional_Sport_2017.jpg',0,0,0,0,'2017-09-06 20:07:56',2,4,'vF9HciD11l7ZZ7ksdQQgQwjZq1lkPCQ7','',''),(53,'2017-09-06 20:16:17','2017-09-06 20:16:17',NULL,'YuanPeng_China_Professional_Sport_2017.jpg',0,0,0,0,'2017-09-06 20:16:17',3,4,'oQfwFKapmXAnZmvUyNqWFSDmcHvOOQ8I','',''),(54,'2017-09-07 13:18:03','2017-09-07 13:18:03',NULL,'file:///storage/emulated/0/DCIM/IMG_20170907_151757.jpg',0,0,0,0,'2017-09-07 13:18:03',2,5,'aliRcevAivdXCNmJBfZGw17zFILazeIo','',''),(55,'2017-09-07 13:25:08','2017-09-07 13:25:08',NULL,'file:///storage/emulated/0/DCIM/IMG_20170907_152502.jpg',0,0,0,0,'2017-09-07 13:25:08',3,5,'nSEQV3vAgUvfeJksshw6ZggqXcFVRelo','',''),(56,'2017-09-07 13:29:47','2017-09-07 13:29:47',NULL,'file:///storage/emulated/0/DCIM/IMG_20170907_152938.jpg',0,0,0,0,'2017-09-07 13:29:47',3,5,'MMc37pzaGDMxwV7cigWCD-OnfuMU24Zg','',''),(57,'2017-09-15 12:28:19','2017-09-15 12:28:19',NULL,'file:///storage/emulated/0/DCIM/IMG_20170915_142811.jpg',0,0,0,0,'2017-09-15 12:28:19',3,5,'ZajBlyCAFJu5h09xH0KdiB7py0Blj4W0','',''),(58,'2017-09-15 15:45:17','2017-09-15 15:45:17',NULL,'file:///storage/emulated/0/DCIM/IMG_20170915_174508.jpg',0,0,0,0,'2017-09-15 15:45:17',3,5,'alfkxLvXmPKkGGSbjnIYFSMWVhVdUWKF','',''),(59,'2017-09-15 15:45:45','2017-09-15 15:45:45',NULL,'file:///storage/emulated/0/DCIM/IMG_20170915_174532.jpg',0,0,0,0,'2017-09-15 15:45:45',3,5,'yZiaBx0wF3kj2JFc1QeFCXLCOaWXCy1N','','');
/*!40000 ALTER TABLE `event_submissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `events`
--

DROP TABLE IF EXISTS `events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `events` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `start` timestamp NULL DEFAULT NULL,
  `end` timestamp NULL DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_events_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events`
--

LOCK TABLES `events` WRITE;
/*!40000 ALTER TABLE `events` DISABLE KEYS */;
INSERT INTO `events` VALUES (1,NULL,NULL,NULL,'2017-08-12 13:53:53','2017-08-12 14:53:53','First event!'),(2,NULL,NULL,NULL,'2017-08-21 19:38:08','2017-08-21 19:38:08','Second event'),(3,NULL,NULL,NULL,'2017-09-06 14:38:08','2017-10-06 20:38:08','Inizia la scuola!');
/*!40000 ALTER TABLE `events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_app_tokens`
--

DROP TABLE IF EXISTS `user_app_tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_app_tokens` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `app_token` varchar(128) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `valid` tinyint(1) DEFAULT NULL,
  `validation_code` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_app_tokens_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=120 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_app_tokens`
--

LOCK TABLES `user_app_tokens` WRITE;
/*!40000 ALTER TABLE `user_app_tokens` DISABLE KEYS */;
INSERT INTO `user_app_tokens` VALUES (28,4,'4cGM0WFSZTpNdAEG3zEyz47Kb9A37RRf','2017-08-09 13:53:53','2017-08-09 15:09:52',NULL,1,'J107X'),(30,4,'4cGM0WFSZTpNdAEG3zEyz47Kb9A37RRf','2017-08-09 15:15:52','2017-08-09 15:15:52',NULL,0,'J107X'),(31,4,'YmFPAVdykz2cKMfRU7Vt3PSR52tu2nZa','2017-08-09 15:16:36','2017-08-09 15:16:55',NULL,1,'wNm7o'),(32,4,'r4WOSlVops9yMtbys-Sh3fJ09wmdG39L','2017-08-21 16:40:22','2017-08-21 16:40:22',NULL,0,'3HIKz'),(33,4,'EwkGVb74-oXeeGtf15yM0Kfe0X1mnJ4c','2017-08-21 16:56:13','2017-08-21 16:56:13',NULL,0,'JGI1A'),(34,4,'CqPC7Dy5dUWhio3ZasN5ensHJMdN8ztk','2017-08-21 16:57:24','2017-08-21 16:57:24',NULL,0,'64NA2'),(35,4,'3czPviwx5jGTf7NlZCmYvrmbmntnBSx2','2017-08-21 17:02:38','2017-08-21 17:02:38',NULL,0,'FZch1'),(36,4,'nVSSq9gp307fiXsIfFWltrp3W170Jvuz','2017-08-21 17:05:58','2017-08-21 17:05:58',NULL,0,'i3Sif'),(37,4,'JtzsflkEWwIajiRZhmlzMyZTZcAW42sB','2017-08-21 17:05:59','2017-08-21 17:05:59',NULL,0,'ZQCAC'),(38,4,'iXjWtpE7tTGs3q3jM0H1Z5EK-sxb00hr','2017-08-21 17:06:14','2017-08-21 17:06:14',NULL,0,'h5zWo'),(39,4,'te1VXhwbDA1Q09wYT6IsG9obUKphYQ0M','2017-08-21 17:06:55','2017-08-21 17:06:55',NULL,0,'nuALy'),(40,4,'h8hNcRYiJgnSjuO8VXYwTNPgA049x-DP','2017-08-21 17:09:31','2017-08-21 17:09:31',NULL,0,'ogUHP'),(41,4,'vavaQKw2jCrPY66DaS7GIOqHJnX02gwE','2017-08-21 17:09:45','2017-08-21 17:09:45',NULL,0,'JOm9g'),(42,4,'ndB0YEhLR7zxtlC02JzL1rIxbSpFQ8GC','2017-08-21 17:13:33','2017-08-21 17:13:33',NULL,0,'MDVtq'),(43,4,'FElHluJD0QR7HZ5XzHFnDx8B3HAjYON1','2017-08-21 17:13:36','2017-08-21 17:13:36',NULL,0,'z8ckt'),(44,4,'p3jma-IYbmS6O0NpysjmMaPEyA1pkErE','2017-08-21 17:27:18','2017-08-21 17:27:18',NULL,0,'Ny4Kl'),(45,5,'QpZRAbHg0uJUIsYplBVEBf2-daT50etX','2017-08-21 17:35:06','2017-08-21 17:35:06',NULL,0,'0kpSa'),(46,6,'h0NLTkju3ixAqOxIjcG5I4jA2UNSRW1C','2017-08-21 17:36:49','2017-08-21 17:36:49',NULL,0,'X-12w'),(47,6,'oPAc1PS4OjUmnuzq43kuuRSHq03r7BfP','2017-08-21 17:36:56','2017-08-21 17:36:56',NULL,0,'hN0pV'),(48,5,'fVBlzhGcadLipOVVlkPIlyteG6y3nrCp','2017-08-21 17:37:48','2017-08-21 17:37:48',NULL,0,'eZKQK'),(49,4,'x3fa6VLLMs3gXKDvrK0p6f2d7pSI-2lx','2017-08-21 17:39:46','2017-08-21 17:39:46',NULL,0,'5CbR3'),(50,5,'ahWXeFgk2Xk75pqhA0zmrXj4uwbD2q1S','2017-08-21 17:41:54','2017-08-21 17:42:19',NULL,1,'evCAo'),(51,4,'MMN4Hr5dI3d-JupOSzOZ6FttDdxhySEp','2017-08-21 17:46:35','2017-08-21 17:46:35',NULL,0,'dQAoQ'),(52,4,'Jh2GkWkmMTevpxOC-LunCduPMWIyhb6H','2017-08-21 17:47:03','2017-08-21 17:47:03',NULL,0,'Ey9qd'),(53,5,'KfZ-wobPzJ3EW0Nzhu0rzulHQkukjt2e','2017-08-21 19:55:34','2017-08-21 19:56:08',NULL,1,'WNqhO'),(54,5,'XNOjeg7sdkVz1fuD9aEtlNdYTdS1a9P5','2017-08-21 20:20:59','2017-08-21 20:22:11',NULL,1,'sdLiQ'),(55,5,'EFH8XxZNnwpXztjXVxQewNZApmKOMFQr','2017-08-21 20:34:56','2017-08-21 20:35:13',NULL,1,'QWcYD'),(56,5,'17IosQ6CbLAcAmXhDlV3bd65e78IF0ou','2017-08-21 20:37:54','2017-08-21 20:38:10',NULL,1,'GBCM7'),(57,5,'DCTLQvGV-uluKUU58-Lo5HSoxbzJu8jz','2017-08-21 20:41:17','2017-08-21 20:41:57',NULL,1,'ojfyZ'),(58,5,'aA4wD16KJfoa9-qr9dycgFRLEznVM7P9','2017-08-21 20:43:26','2017-08-21 20:43:45',NULL,1,'mvA1j'),(59,7,'7O13eaPmcAfhIdq8cWgG8vKRfDsnbtKv','2017-08-21 22:50:49','2017-08-21 22:50:49',NULL,0,'hfCkY'),(60,5,'Mvf0ULNuWi6aDxBnLHbAgFM03X3tQb67','2017-08-21 22:52:42','2017-08-21 22:52:42',NULL,0,'DOJ0s'),(61,4,'ofgCbLN1rdOSzF30VxqO024oV-fi905W','2017-08-21 22:55:49','2017-08-21 22:55:49',NULL,0,'MKOT0'),(62,5,'Kz09TmB6BCgV42f0WM7D-Z3fde9wbZ4G','2017-08-21 22:56:08','2017-08-21 22:56:35',NULL,1,'qfvty'),(63,5,'35HAFmWqhu4eXrqMJs4d9vkQyQ3oV1FZ','2017-08-21 22:59:03','2017-08-21 22:59:24',NULL,1,'otDFR'),(64,5,'oTYvcFgAqDryqKLm0lLr2mRGWD8VNni6','2017-08-21 23:02:47','2017-08-21 23:02:47',NULL,0,'0lT2P'),(65,5,'jVzSHfE0GOTq4VjPT39ccnGyZJDOyq8w','2017-08-21 23:03:48','2017-08-21 23:03:48',NULL,0,'y5bIm'),(66,5,'Rcq-RmbwGrPBVzQvwqlV-Vx-CbDfPsue','2017-08-21 23:05:36','2017-08-21 23:05:53',NULL,1,'s52TQ'),(67,5,'f9ap7DOwXQNroI5tr74mMPe3ImMcnoz4','2017-08-21 23:19:23','2017-08-21 23:19:40',NULL,1,'jBcuh'),(68,5,'Hi3RCeSdDvdOfstKedAK9mh7nrGOtyHJ','2017-08-21 23:21:33','2017-08-21 23:21:49',NULL,1,'EZM99'),(69,5,'fUgsLou1npJHQ0VBL8wwsUHzZQ6AI9vN','2017-08-22 10:14:34','2017-08-22 10:14:51',NULL,1,'byC09'),(70,5,'KM6JyOAaDVpt2wsmnN6igarjTIuZj5OB','2017-08-22 11:41:57','2017-08-22 11:42:17',NULL,1,'A2brT'),(71,5,'ymJYvWhxFSuvFRdUI5MCUcx2r6Y06-l2','2017-08-22 12:27:55','2017-08-22 12:28:13',NULL,1,'a7VvL'),(72,4,'DN-Sb5A6JX3IeTZZ23DpOBOVJ1z4ZG2h','2017-08-22 16:21:34','2017-08-22 16:21:34',NULL,0,'Tgx65'),(73,4,'1gh7O1k4rxrYZlTDWe4up10cTO25ZLLZ','2017-08-22 16:22:58','2017-08-22 16:22:58',NULL,0,'ZHIoh'),(74,4,'poVHWW2ZfcoGbHc0ko1fw-87Y8YhDYMP','2017-08-22 16:23:18','2017-08-22 16:23:18',NULL,0,'EhR-b'),(75,4,'g27g2TVi-PyRA3WZQfHtdIitOdwS6DQz','2017-08-22 16:26:08','2017-08-22 16:26:20',NULL,1,'4ZwI-'),(76,5,'VheAwjEZ-B2ZsPoed9sbfAr8uAFG4pGq','2017-08-22 22:22:13','2017-08-22 22:22:33',NULL,1,'U89RM'),(77,5,'CUs3hDDLYDuWU89PyE0bwTfD9kY41NNw','2017-08-24 10:20:03','2017-08-24 10:21:24',NULL,1,'GQg34'),(78,5,'1cYINOdUaNWLJxFMBzcYVuPVVpdrYYXr','2017-08-24 10:28:14','2017-08-24 10:28:59',NULL,1,'ecPKf'),(79,5,'D0ppeceszXn-6gol86AhgdNDc4Ifm3YD','2017-08-24 11:35:56','2017-08-24 11:36:17',NULL,1,'wFvZF'),(80,5,'4wSof5bi1UAlG3gMulpOFzS16tviyqLJ','2017-08-24 14:35:37','2017-08-24 14:36:08',NULL,1,'PEp8L'),(81,5,'PJjQM7ht51A1gKU4lOLY--ycjps4P3wI','2017-08-24 15:54:27','2017-08-24 15:55:23',NULL,1,'BqTrR'),(82,4,'ngDa1MrjANK6gEUvwaXUvFqcVsdMDI14','2017-08-24 16:51:43','2017-08-24 16:51:43',NULL,0,'Nw6lt'),(83,5,'MCvFhFXH0UrAlEi0o6pXlH8N5UwnbGyD','2017-08-24 16:51:46','2017-08-24 16:52:06',NULL,1,'BVPM6'),(84,4,'3VotDcFX-7bdt1EvHs5lb8Gab353M8ob','2017-08-24 16:53:29','2017-08-24 16:53:42',NULL,1,'hV4OS'),(85,4,'GXlK4AIGWJGx-QYLfOqyK8cDWk6lPZwO','2017-08-25 13:02:58','2017-08-25 13:03:10',NULL,1,'HJr3U'),(86,5,'EKiJhRox-gW5hzePIXLw7w0Bb8f6CWnP','2017-08-29 22:07:14','2017-08-29 22:07:33',NULL,1,'EJJUi'),(87,4,'ppwuId-dlqo3FgxfGEbYaAGQsqFL8rmr','2017-08-30 11:19:00','2017-08-30 11:19:00',NULL,0,'HeTHc'),(88,5,'16CT-7m7HBIRbdQOoXvzPFaNwZ6LR8nK','2017-08-30 11:19:14','2017-08-30 11:19:40',NULL,1,'SD4r6'),(89,5,'18A-W08okL6CHZSxAD1I-d4R2wfeUKFI','2017-08-30 12:30:19','2017-08-30 12:31:00',NULL,1,'K5ZNd'),(90,5,'oCRjIYBzSmgnLJTZkXekbxScGBorgie2','2017-08-30 12:57:51','2017-08-30 12:58:13',NULL,1,'VnSMu'),(91,5,'F9civSXip01xmJDPCeZmq6IUJwRNuBuY','2017-08-30 19:17:39','2017-08-30 19:18:11',NULL,1,'M3l83'),(92,4,'5FK8WUClTntHKB9-eB-x-kiSmFVHJhAo','2017-09-05 15:02:42','2017-09-05 15:03:05',NULL,1,'BnQWH'),(93,8,'RJANmelvg9tGwVOsGahI1Pp4bge-ZKv8','2017-09-05 15:51:22','2017-09-05 15:51:22',NULL,0,'1-7NP'),(94,4,'MGAxDbWbbS31436M7GC0g01pAcdbX-Eg','2017-09-05 15:51:49','2017-09-05 15:52:12',NULL,1,'akcmG'),(95,4,'-xv8mzNqp2xwt0VTdLYC3Dc330XsdUMh','2017-09-05 16:03:12','2017-09-05 16:03:44',NULL,1,'mUPJe'),(96,4,'apx84jYUr08tH0yuJ0AdWtW3r3y-hARr','2017-09-06 09:12:55','2017-09-06 09:12:55',NULL,0,'BYGrF'),(97,4,'G49iKR8tB5WKse1Wkr19WDPoByzf4mrO','2017-09-06 09:33:56','2017-09-06 10:19:31',NULL,1,'yy6uh'),(98,5,'iN0rAFgXu2t70DwKGcYQL2z3802dPIEm','2017-09-06 18:23:00','2017-09-06 18:23:20',NULL,1,'B1rUv'),(99,4,'9fnDCbSXY5MKU9PWi3TMjs23On6cPSMb','2017-09-06 18:25:57','2017-09-06 18:26:26',NULL,1,'dcDpu'),(100,4,'1xK38najUfBO7zKb6vffwjZXVI9vYr4w','2017-09-06 18:47:12','2017-09-06 18:47:32',NULL,1,'mpgTX'),(101,5,'Tnbvaz3gciL1QmY6Qfry-wrn0vdb15eA','2017-09-06 20:01:42','2017-09-06 20:01:58',NULL,1,'Pajug'),(102,4,'7E2VLCL92z5AokBIhdzbceuzV-rdZ2PB','2017-09-06 20:05:11','2017-09-06 20:05:36',NULL,1,'B6rhF'),(103,4,'goHtW3vYcfeIU5C8EjFY-t7G7s2eayUJ','2017-09-06 20:15:42','2017-09-06 20:16:07',NULL,1,'SIHNR'),(104,5,'c6RLnPMk3ZDfB4E6xwCyRsjXVRBpjPuk','2017-09-07 12:21:54','2017-09-07 12:22:19',NULL,1,'GQ3bF'),(105,5,'XbUzaglM1Z7dr87Q7-jXkK1TQkS4i-aS','2017-09-07 13:17:35','2017-09-07 13:17:50',NULL,1,'ZDzWg'),(106,4,'gOrdqxzrDW43s-JjT-qIDZ2FiDB2hp7S','2017-09-08 13:23:30','2017-09-08 13:26:37',NULL,1,'ge35k'),(107,9,'KB9t1om4Cs7H6dh5yonIBsuymcs9XjoW','2017-09-08 13:43:33','2017-09-08 13:43:33',NULL,0,'ndz4K'),(108,5,'DMZnRAjaT6FxusTqKYPcpIeYMD4k4Wtk','2017-09-08 13:44:07','2017-09-08 13:45:07',NULL,1,'D9l64'),(109,5,'4I9AWv32uKZh2FkL3QP-vwf4m6i4lorC','2017-09-08 18:02:09','2017-09-08 18:02:36',NULL,1,'BL7uJ'),(110,5,'QWzbkW0THcaLuPC3lHn30whI1twVglFW','2017-09-15 11:32:46','2017-09-15 11:32:46',NULL,0,'lMR3l'),(111,5,'8cf3hNrNAmT9iv0drk39r57r7eR327mH','2017-09-15 11:40:54','2017-09-15 11:41:29',NULL,1,'W6N05'),(112,5,'9nARnEJJk-rWeZJUCNrGPLUAVao95cHF','2017-09-15 12:27:33','2017-09-15 12:28:03',NULL,1,'OpAHZ'),(113,5,'RRpAzrRqXjZb4O33oUXqomhumaJXF8tA','2017-09-15 15:44:15','2017-09-15 15:44:53',NULL,1,'VIeHM'),(114,5,'yqqk4lH65uJ5EdEiWRv-Zn6JG5dDqK51','2017-11-02 19:47:54','2017-11-02 19:48:17',NULL,1,'Fp3KF'),(115,5,'Z9wyF23IBd39k63DRwE1XHQErxDIlDRX','2018-04-11 11:29:37','2018-04-11 11:29:37',NULL,0,'3fWPI'),(116,5,'3I8N11nDhFS87VE5C0mS-p4cLTvSqABn','2018-04-11 11:30:36','2018-04-11 11:30:36',NULL,0,'ia1Or'),(117,5,'ocLdKBL9xW00mtShkGALAU6eRG7FWiOK','2018-04-11 11:39:21','2018-04-11 11:39:41',NULL,1,'tzRF2'),(118,4,'Xrs1BZ3PhWBEVt0n0HtODXADnaNJv7G5','2018-04-11 11:53:40','2018-04-11 11:53:40',NULL,0,'fWvzY'),(119,5,'BGxl3ksl1aylvfwzQ8oluOafnxDDkboz','2018-04-11 11:54:18','2018-04-11 11:54:44',NULL,1,'8JD9v');
/*!40000 ALTER TABLE `user_app_tokens` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) DEFAULT NULL,
  `email` varchar(128) DEFAULT NULL,
  `EMAIL_VALID` tinyint(1) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (3,'virgolous','ezio.vergine@gmail.com',0,'2017-08-07 12:22:37','2017-08-07 12:22:37',NULL),(4,'enricod','enrico.donelli@gmail.com',0,'2017-08-07 12:31:07','2018-04-11 11:53:40',NULL),(5,'virgolus','virgolus@gmail.com',0,'2017-08-21 17:35:06','2018-04-11 11:54:18',NULL),(7,'virgolus','virfolus@gmail.com',0,'2017-08-21 22:50:49','2017-08-21 22:50:49',NULL),(8,'asdf','asdfasdf',0,'2017-09-05 15:51:22','2017-09-05 15:51:22',NULL),(9,'As','As',0,'2017-09-08 13:43:33','2017-09-08 13:43:33',NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-04-13 10:03:50
