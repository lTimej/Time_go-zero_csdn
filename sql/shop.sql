-- MySQL dump 10.13  Distrib 8.0.27, for Linux (x86_64)
--
-- Host: localhost    Database: shop
-- ------------------------------------------------------
-- Server version       8.0.27

CREATE TABLE `tb_product_id` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `pid` varchar(32) DEFAULT NULL COMMENT '商品编号',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

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

CREATE TABLE `tb_goods_list` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `pid_id` bigint unsigned NOT NULL COMMENT '商品pid',
  `category_id` bigint unsigned NOT NULL COMMENT '商品类别',
  `url` varchar(256) DEFAULT NULL COMMENT '商品url',
  `sequeue` bigint unsigned DEFAULT NULL COMMENT '顺序',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `category_id` (`category_id`),
  KEY `pid_id` (`pid_id`),
  CONSTRAINT `tb_goods_list_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `tb_goods_category` (`id`) ON DELETE CASCADE,
  CONSTRAINT `tb_goods_list_ibfk_2` FOREIGN KEY (`pid_id`) REFERENCES `tb_product_id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

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

CREATE TABLE `tb_spu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(64) DEFAULT NULL COMMENT '名称',
  `category1_id` bigint unsigned NOT NULL COMMENT '一级类别',
  `category2_id` bigint unsigned NOT NULL COMMENT '二级类别',
  `sales` bigint unsigned DEFAULT NULL COMMENT '销量',
  `cfavs` bigint unsigned DEFAULT NULL COMMENT '收藏数',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `category1_id` (`category1_id`),
  KEY `category2_id` (`category2_id`),
  CONSTRAINT `tb_spu_ibfk_1` FOREIGN KEY (`category1_id`) REFERENCES `tb_goods_category` (`id`) ON DELETE CASCADE,
  CONSTRAINT `tb_spu_ibfk_2` FOREIGN KEY (`category2_id`) REFERENCES `tb_goods_category` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `tb_sku` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(64) DEFAULT NULL COMMENT '名称',
  `spu_id` bigint unsigned NOT NULL COMMENT '商品',
  `category_id` bigint unsigned NOT NULL COMMENT '从属类别',
  `price` decimal(10,2) DEFAULT NULL COMMENT '单价',
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

CREATE TABLE `tb_spu_desc` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `spu_id` bigint unsigned NOT NULL COMMENT '商品SPU',
  `detail_info` varchar(888) DEFAULT NULL COMMENT '商品详情',
  `desc_image` varchar(255) DEFAULT NULL COMMENT '图片',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `spu_id` (`spu_id`),
  CONSTRAINT `tb_spu_desc_ibfk_1` FOREIGN KEY (`spu_id`) REFERENCES `tb_spu` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

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


CREATE TABLE `tb_spu_specification` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `spu_id` bigint unsigned NOT NULL COMMENT '商品spu',
  `name` varchar(20) DEFAULT NULL COMMENT '规格名称',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `spu_id` (`spu_id`),
  CONSTRAINT `tb_spu_specification_ibfk_1` FOREIGN KEY (`spu_id`) REFERENCES `tb_spu` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
