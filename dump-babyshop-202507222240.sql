-- MySQL dump 10.13  Distrib 8.0.19, for Win64 (x86_64)
--
-- Host: localhost    Database: babyshop
-- ------------------------------------------------------
-- Server version	8.0.42

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `order_items`
--

DROP TABLE IF EXISTS `order_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `order_items` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `order_id` bigint unsigned DEFAULT NULL,
  `product_id` bigint unsigned DEFAULT NULL,
  `qty` bigint DEFAULT NULL,
  `unit_price` double DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `created_by_id` bigint unsigned DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `updated_by_id` bigint unsigned DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `deleted_by_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_order_items_deleted_at` (`deleted_at`),
  KEY `fk_order_items_product` (`product_id`),
  KEY `fk_orders_items` (`order_id`),
  CONSTRAINT `fk_order_items_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  CONSTRAINT `fk_orders_items` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_items`
--

LOCK TABLES `order_items` WRITE;
/*!40000 ALTER TABLE `order_items` DISABLE KEYS */;
INSERT INTO `order_items` VALUES (1,1,1,1,120000,'2025-07-12 00:38:44.083',NULL,'2025-07-12 00:38:44.083',NULL,NULL,NULL),(2,1,2,1,180000,'2025-07-12 00:38:44.083',NULL,'2025-07-12 00:38:44.083',NULL,NULL,NULL),(3,2,1,1,120000,'2025-07-12 00:54:44.824',NULL,'2025-07-12 00:54:44.824',NULL,NULL,NULL),(4,2,2,1,180000,'2025-07-12 00:54:44.824',NULL,'2025-07-12 00:54:44.824',NULL,NULL,NULL),(5,3,1,1,120000,'2025-07-12 00:56:00.569',NULL,'2025-07-12 00:56:00.569',NULL,NULL,NULL),(6,3,2,1,180000,'2025-07-12 00:56:00.569',NULL,'2025-07-12 00:56:00.569',NULL,NULL,NULL),(7,4,1,1,120000,'2025-07-12 00:57:29.290',NULL,'2025-07-12 00:57:29.290',NULL,NULL,NULL),(8,4,2,1,180000,'2025-07-12 00:57:29.290',NULL,'2025-07-12 00:57:29.290',NULL,NULL,NULL),(9,5,1,1,120000,'2025-07-12 01:01:47.457',NULL,'2025-07-12 01:01:47.457',NULL,NULL,NULL),(10,5,2,1,180000,'2025-07-12 01:01:47.457',NULL,'2025-07-12 01:01:47.457',NULL,NULL,NULL),(11,6,1,1,120000,'2025-07-12 01:07:14.253',NULL,'2025-07-12 01:07:14.253',NULL,NULL,NULL),(12,6,2,1,180000,'2025-07-12 01:07:14.253',NULL,'2025-07-12 01:07:14.253',NULL,NULL,NULL);
/*!40000 ALTER TABLE `order_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `orders`
--

DROP TABLE IF EXISTS `orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `orders` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned DEFAULT NULL,
  `total_price` double DEFAULT NULL,
  `status` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `created_by_id` bigint unsigned DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `updated_by_id` bigint unsigned DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `deleted_by_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_orders_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
INSERT INTO `orders` VALUES (1,1,300000,'pending','2025-07-12 00:38:44.080',NULL,'2025-07-12 00:38:44.080',NULL,NULL,NULL),(2,1,300000,'pending','2025-07-12 00:54:44.822',NULL,'2025-07-12 00:54:44.822',NULL,NULL,NULL),(3,1,300000,'pending','2025-07-12 00:56:00.566',NULL,'2025-07-12 00:56:00.566',NULL,NULL,NULL),(4,1,300000,'pending','2025-07-12 00:57:29.288',NULL,'2025-07-12 00:57:29.288',NULL,NULL,NULL),(5,1,300000,'pending','2025-07-12 01:01:47.455',NULL,'2025-07-12 01:01:47.455',NULL,NULL,NULL),(6,1,300000,'pending','2025-07-12 01:07:14.250',NULL,'2025-07-12 01:07:14.250',NULL,NULL,NULL);
/*!40000 ALTER TABLE `orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `price` double DEFAULT NULL,
  `size` longtext,
  `stock` bigint DEFAULT NULL,
  `image_url` longtext,
  `description` text,
  `deleted_at` datetime(3) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `created_by_id` bigint unsigned DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `updated_by_id` bigint unsigned DEFAULT NULL,
  `deleted_by_id` bigint unsigned DEFAULT NULL,
  `slug` varchar(191) DEFAULT NULL,
  `meta_title` longtext,
  `meta_description` longtext,
  `meta_keywords` longtext,
  `og_image` longtext,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_products_slug` (`slug`),
  KEY `idx_products_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (1,'Áo thun bé trai',120000,'M',10,'https://cdn.example.com/aothun.jpg','',NULL,NULL,NULL,NULL,NULL,NULL,'a',NULL,NULL,NULL,NULL),(2,'Váy công chúa',180000,'S',5,'https://cdn.example.com/vaycongchua.jpg','',NULL,NULL,NULL,NULL,NULL,NULL,'b',NULL,NULL,NULL,NULL),(3,'Áo khoác lông',250000,'L',7,'https://cdn.example.com/aokhoac.jpg','',NULL,NULL,NULL,NULL,NULL,NULL,'c',NULL,NULL,NULL,NULL),(4,'Quần jean bé gái',160000,'M',12,'https://cdn.example.com/quanjean.jpg','',NULL,NULL,NULL,NULL,NULL,NULL,'d',NULL,NULL,NULL,NULL),(5,'Bộ đồ ngủ trẻ em',90000,'S',20,'https://cdn.example.com/dongu.jpg','',NULL,NULL,NULL,NULL,NULL,NULL,'f',NULL,NULL,NULL,NULL);
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(191) DEFAULT NULL,
  `password` longtext,
  `role` longtext,
  `status` bigint DEFAULT NULL,
  `verified_at` datetime(3) DEFAULT NULL,
  `verified_by_id` bigint unsigned DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `created_by_id` bigint unsigned DEFAULT NULL,
  `updated_by_id` bigint unsigned DEFAULT NULL,
  `deleted_by_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_users_email` (`email`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'admin@example.com','$2a$10$XlA1vwM1kJs2uCM7BGjPEOuMznpqEawRiZbEX1OQJGptZcc2fCcku','admin',2,'2025-07-11 23:49:01.000',1,NULL,NULL,NULL,NULL,NULL,NULL),(2,'newuser@example.com','$2a$10$7n11oMQ2IKzIqPdPgoGqseu..7uYb711SETYjkCbNPdpcGpVn99jW','staff',1,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL),(3,'huutrong@example.com','$2a$10$.zsuV0wd8bSfey3IGqtZKuMFrpcDysSrKQUNQuE4gVohEw7h4E2jS','staff',1,NULL,NULL,'2025-07-11 23:49:01.000','2025-07-11 23:49:01.000',NULL,NULL,NULL,NULL),(4,'tuanh@example.com','$2a$10$6mUJa.Xg.mt1DTjTh5OW/uZldQ0x2RVn6.uRWOBJdaB/Ztm/Fif2i','staff',1,NULL,NULL,'2025-07-11 23:50:45.000','2025-07-11 23:50:45.000',NULL,NULL,NULL,NULL),(5,'nhi@example.com','$2a$10$mS3UrQQBiHTvbjw1CeJ3l.soH65jSSjyXgEFGpy6B/HlOPcCUj.q.','staff',1,NULL,NULL,'2025-07-11 23:52:27.000','2025-07-11 23:52:27.000',NULL,NULL,NULL,NULL),(6,'nhi1@example.com','$2a$10$OK/UGvrxGTt16MjKCASkc.uCq/cKo3O88509Ac9PKRWk6yJsCX7D.','staff',1,NULL,NULL,'2025-07-11 23:54:33.000','2025-07-11 23:54:33.000',NULL,NULL,NULL,NULL),(7,'nhi2@example.com','$2a$10$5J.9gj/67vbWBv7Ioxf55Oor8HoOs17cocnKjodVRm7lTtjiBT0fy','staff',1,NULL,NULL,'2025-07-11 23:54:58.000','2025-07-11 23:54:58.000',NULL,NULL,NULL,NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'babyshop'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-07-22 22:40:54
