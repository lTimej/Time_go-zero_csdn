SET
@@auto_increment_increment=9;

CREATE TABLE `order` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` varchar(20) NOT NULL COMMENT '用户id',
  `address_id` int unsigned NOT NULL COMMENT '地址',
  `total_count` int COMMENT '商品总数',
  `total_price` decimal(10,2) DEFAULT 0 COMMENT '商品总金额',
  `freight` decimal(10,2) DEFAULT 0 COMMENT '运费',
  `version` int default 0 COMMENT '乐观锁版本号',
  `sn` char(32) NOT NULL COMMENT '流水单号',
  `pay_status` tinyint(4) default 1 COMMENT '支付状态',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '支付创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '支付修改时间',
  `is_deleted` tinyint(1) default 0  COMMENT '逻辑删除',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `address_id` (`address_id`),
  CONSTRAINT `user_id_1` FOREIGN KEY (`user_id`) REFERENCES `user_basic` (`user_id`) ON DELETE CASCADE,
  CONSTRAINT `address_id_2` FOREIGN KEY (`address_id`) REFERENCES `address` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单';

CREATE TABLE `user_order` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `order_id` int unsigned NOT NULL COMMENT '订单id',
  `sku_id` bigint unsigned NOT NULL COMMENT '商品id',
  `spec_id` char(20) COMMENT '商品属性id;"1,2"',
  `specs` varchar(255) COMMENT '商品属性',
  `comment` varchar(200) COMMENT '商品评价',
  `score` int COMMENT '商品评分',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '支付创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '支付修改时间',
  `is_anonymous` tinyint(1) default 0  COMMENT '是否匿名',
  `is_commented` tinyint(1) default 0  COMMENT '是否评论',
  PRIMARY KEY (`id`),
  KEY `order_id` (`order_id`),
  KEY `sku_id` (`sku_id`),
  CONSTRAINT `order_id_1` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`) ON DELETE CASCADE,
  CONSTRAINT `sku_id_1` FOREIGN KEY (`sku_id`) REFERENCES `tb_sku` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品订单';