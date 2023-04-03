-- MySQL dump 10.13  Distrib 8.0.27, for Linux (x86_64)
--
-- Host: localhost    Database: shop
-- ------------------------------------------------------
-- Server version	8.0.27

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
-- Table structure for table `tb_goods_category`
--

DROP TABLE IF EXISTS `tb_goods_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_goods_category` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(125) DEFAULT NULL COMMENT '类名',
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '目录',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `parent_id` (`parent_id`),
  CONSTRAINT `tb_goods_category_ibfk_1` FOREIGN KEY (`parent_id`) REFERENCES `tb_goods_category` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_goods_category`
--

LOCK TABLES `tb_goods_category` WRITE;
/*!40000 ALTER TABLE `tb_goods_category` DISABLE KEYS */;
/*!40000 ALTER TABLE `tb_goods_category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_goods_list`
--

DROP TABLE IF EXISTS `tb_goods_list`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_goods_list` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `pid` bigint unsigned NOT NULL COMMENT '商品pid',
  `category` bigint unsigned NOT NULL COMMENT '商品类别',
  `url` varchar(256) DEFAULT NULL COMMENT '商品url',
  `sequeue` bigint unsigned DEFAULT NULL COMMENT '顺序',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `category` (`category`),
  KEY `pid` (`pid`),
  CONSTRAINT `tb_goods_list_ibfk_1` FOREIGN KEY (`category`) REFERENCES `tb_goods_category` (`id`) ON DELETE CASCADE,
  CONSTRAINT `tb_goods_list_ibfk_2` FOREIGN KEY (`pid`) REFERENCES `tb_product_id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_goods_list`
--

LOCK TABLES `tb_goods_list` WRITE;
/*!40000 ALTER TABLE `tb_goods_list` DISABLE KEYS */;
/*!40000 ALTER TABLE `tb_goods_list` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_goods_visit`
--

DROP TABLE IF EXISTS `tb_goods_visit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_goods_visit` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `category_id` bigint unsigned NOT NULL COMMENT '商品分类',
  `count` bigint unsigned DEFAULT NULL COMMENT '访问量',
  `date` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '统计日期',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `category_id` (`category_id`),
  CONSTRAINT `tb_goods_visit_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `tb_goods_category` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_goods_visit`
--

LOCK TABLES `tb_goods_visit` WRITE;
/*!40000 ALTER TABLE `tb_goods_visit` DISABLE KEYS */;
/*!40000 ALTER TABLE `tb_goods_visit` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_product_id`
--

DROP TABLE IF EXISTS `tb_product_id`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_product_id` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `pid` varchar(32) DEFAULT NULL COMMENT '商品编号',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_product_id`
--

LOCK TABLES `tb_product_id` WRITE;
/*!40000 ALTER TABLE `tb_product_id` DISABLE KEYS */;
/*!40000 ALTER TABLE `tb_product_id` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_sku`
--

DROP TABLE IF EXISTS `tb_sku`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_sku` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(64) DEFAULT NULL COMMENT '名称',
  `spu_id` bigint unsigned NOT NULL COMMENT '商品',
  `category_id` bigint unsigned NOT NULL COMMENT '从属类别',
  `now_price` decimal(10,2) DEFAULT NULL COMMENT '进价',
  `stock` bigint unsigned DEFAULT NULL COMMENT '库存',
  `sales` bigint unsigned DEFAULT NULL COMMENT '销量',
  `comments` bigint unsigned DEFAULT NULL COMMENT '评价数',
  `is_launched` tinyint(1) DEFAULT NULL COMMENT '是否上架销售',
  `default_image` varchar(255) DEFAULT NULL COMMENT '默认图片',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `category_id` (`category_id`),
  KEY `spu_id` (`spu_id`),
  CONSTRAINT `tb_sku_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `tb_goods_category` (`id`) ON DELETE CASCADE,
  CONSTRAINT `tb_sku_ibfk_2` FOREIGN KEY (`spu_id`) REFERENCES `tb_spu` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_sku`
--

LOCK TABLES `tb_sku` WRITE;
/*!40000 ALTER TABLE `tb_sku` DISABLE KEYS */;
/*!40000 ALTER TABLE `tb_sku` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_sku_image`
--

DROP TABLE IF EXISTS `tb_sku_image`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_sku_image` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `sku_id` bigint unsigned NOT NULL COMMENT 'sku',
  `image` varchar(255) DEFAULT NULL COMMENT '图片',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `sku_id` (`sku_id`),
  CONSTRAINT `tb_sku_image_ibfk_1` FOREIGN KEY (`sku_id`) REFERENCES `tb_sku` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_sku_image`
--

LOCK TABLES `tb_sku_image` WRITE;
/*!40000 ALTER TABLE `tb_sku_image` DISABLE KEYS */;
/*!40000 ALTER TABLE `tb_sku_image` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_sku_specification`
--

DROP TABLE IF EXISTS `tb_sku_specification`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_sku_specification` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `sku_id` bigint unsigned NOT NULL COMMENT 'sku',
  `spec_id` bigint unsigned NOT NULL COMMENT '规格名称',
  `option_id` bigint unsigned NOT NULL COMMENT '规格值',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `option_id` (`option_id`),
  KEY `sku_id` (`sku_id`),
  KEY `spec_id` (`spec_id`),
  CONSTRAINT `tb_sku_specification_ibfk_1` FOREIGN KEY (`option_id`) REFERENCES `tb_specification_option` (`id`) ON DELETE CASCADE,
  CONSTRAINT `tb_sku_specification_ibfk_2` FOREIGN KEY (`sku_id`) REFERENCES `tb_sku` (`id`) ON DELETE CASCADE,
  CONSTRAINT `tb_sku_specification_ibfk_3` FOREIGN KEY (`spec_id`) REFERENCES `tb_spu_desc` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_sku_specification`
--

LOCK TABLES `tb_sku_specification` WRITE;
/*!40000 ALTER TABLE `tb_sku_specification` DISABLE KEYS */;
/*!40000 ALTER TABLE `tb_sku_specification` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_specification_option`
--

DROP TABLE IF EXISTS `tb_specification_option`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_specification_option` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `spec_id` bigint unsigned NOT NULL COMMENT '规格',
  `value` varchar(20) DEFAULT NULL COMMENT '选项值',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `spec_id` (`spec_id`),
  CONSTRAINT `tb_specification_option_ibfk_1` FOREIGN KEY (`spec_id`) REFERENCES `tb_spu_desc` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_specification_option`
--

LOCK TABLES `tb_specification_option` WRITE;
/*!40000 ALTER TABLE `tb_specification_option` DISABLE KEYS */;
/*!40000 ALTER TABLE `tb_specification_option` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_spu`
--

DROP TABLE IF EXISTS `tb_spu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_spu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(64) DEFAULT NULL COMMENT '名称',
  `category_id1` bigint unsigned NOT NULL COMMENT '一级类别',
  `category_id2` bigint unsigned NOT NULL COMMENT '二级类别',
  `sales` bigint unsigned DEFAULT NULL COMMENT '销量',
  `cfavs` bigint unsigned DEFAULT NULL COMMENT '收藏数',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `category_id1` (`category_id1`),
  KEY `category_id2` (`category_id2`),
  CONSTRAINT `tb_spu_ibfk_1` FOREIGN KEY (`category_id1`) REFERENCES `tb_goods_category` (`id`) ON DELETE CASCADE,
  CONSTRAINT `tb_spu_ibfk_2` FOREIGN KEY (`category_id2`) REFERENCES `tb_goods_category` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_spu`
--

LOCK TABLES `tb_spu` WRITE;
/*!40000 ALTER TABLE `tb_spu` DISABLE KEYS */;
/*!40000 ALTER TABLE `tb_spu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_spu_desc`
--

DROP TABLE IF EXISTS `tb_spu_desc`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_spu_desc` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `spu_id` bigint unsigned NOT NULL COMMENT '商品SPU',
  `name` varchar(20) DEFAULT NULL COMMENT '规格名称',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `spu_id` (`spu_id`),
  CONSTRAINT `tb_spu_desc_ibfk_1` FOREIGN KEY (`spu_id`) REFERENCES `tb_spu` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_spu_desc`
--

LOCK TABLES `tb_spu_desc` WRITE;
/*!40000 ALTER TABLE `tb_spu_desc` DISABLE KEYS */;
/*!40000 ALTER TABLE `tb_spu_desc` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_spu_specification`
--

DROP TABLE IF EXISTS `tb_spu_specification`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_spu_specification` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `spu_id` bigint unsigned NOT NULL COMMENT 'spu_descs',
  `detail_info` varchar(888) DEFAULT NULL COMMENT '商品详情',
  `desc_image` varchar(255) DEFAULT NULL COMMENT '图片',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `spu_id` (`spu_id`),
  CONSTRAINT `tb_spu_specification_ibfk_1` FOREIGN KEY (`spu_id`) REFERENCES `tb_spu` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_spu_specification`
--

LOCK TABLES `tb_spu_specification` WRITE;
/*!40000 ALTER TABLE `tb_spu_specification` DISABLE KEYS */;
/*!40000 ALTER TABLE `tb_spu_specification` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-04-03 15:55:35
