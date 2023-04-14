SET
@@auto_increment_increment=9;

-- CREATE TABLE `user_legalize_log`
-- (
--     `legalize_id`      bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '认证申请ID',
--     `user_id`          bigint(20) unsigned NOT NULL COMMENT '用户ID',
--     `type`             tinyint(1) NOT NULL COMMENT '认证类型',
--     `status`           tinyint(1) NOT NULL DEFAULT '1' COMMENT '申请状态',
--     `reject_reason`    varchar(200) COMMENT '驳回原因',
--     `qualification_id` bigint(20) unsigned NOT NULL COMMENT '资质认证材料ID',
--     `create_time`      datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
--     `update_time`      datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
--     PRIMARY KEY (`legalize_id`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户认证申请记录';

-- CREATE TABLE `user_qualification`
-- (
--     `qualification_id`  bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '资质认证材料ID',
--     `user_id`           bigint(20) unsigned NOT NULL COMMENT '用户ID',
--     `name`              varchar(20)  NOT NULL COMMENT '姓名',
--     `id_number`         varchar(20) NULL COMMENT '身份证号',
--     `industry`          varchar(200) NOT NULL COMMENT '行业',
--     `company`           varchar(200) NOT NULL COMMENT '公司',
--     `position`          varchar(200) NOT NULL COMMENT '职位',
--     `add_info`          varchar(200) COMMENT '补充信息',
--     `id_card_front`     varchar(200) COMMENT '身份证正面',
--     `id_card_back`      varchar(200) COMMENT '身份证背面',
--     `id_card_handheld`  varchar(200) COMMENT '手持身份证',
--     `qualification_img` varchar(200) COMMENT '证明资料',
--     `create_time`       datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
--     `update_time`       datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
--     PRIMARY KEY (`qualification_id`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户资质认证材料';
-- ###


-- CREATE TABLE `user_basic`
-- (
--     `user_id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
--     `account`         varchar(20) COMMENT '账号',
--     `email`           varchar(20) COMMENT '邮箱',
--     `status`          tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态，是否可用，0-不可用，1-可用',
--     `mobile`          char(11)    NOT NULL COMMENT '手机号',
--     `password`        varchar(93) NULL COMMENT '密码',
--     `user_name`       varchar(32) NOT NULL COMMENT '昵称',
--     `profile_photo`   varchar(128) NULL COMMENT '头像',
--     `last_login`      datetime NULL COMMENT '最后登录时间',
--     `is_media`        tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否是自媒体，0-不是，1-是',
--     `is_verified`     tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否实名认证，0-不是，1-是',
--     `introduction`    varchar(50) NULL COMMENT '简介',
--     `certificate`     varchar(30) NULL COMMENT '认证',
--     `article_count`   int(11) unsigned NOT NULL DEFAULT '0' COMMENT '发文章数',
--     `following_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '关注的人数',
--     `fans_count`      int(11) unsigned NOT NULL DEFAULT '0' COMMENT '被关注的人数',
--     `like_count`      int(11) unsigned NOT NULL DEFAULT '0' COMMENT '累计点赞人数',
--     `read_count`      int(11) unsigned NOT NULL DEFAULT '0' COMMENT '累计阅读人数',
--     `code_year`       int(11) unsigned NOT NULL DEFAULT '0' COMMENT '码龄',
--     PRIMARY KEY (`user_id`),
--     UNIQUE KEY `mobile` (`mobile`),
--     UNIQUE KEY `user_name` (`user_name`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户基本信息表';

-- CREATE TABLE `user_profile`
-- (
--     `user_id`             bigint(20) unsigned NOT NULL COMMENT '用户ID',
--     `gender`              tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别，0-男，1-女',
--     `birthday`            date NULL COMMENT '生日',
--     `real_name`           varchar(32) NULL COMMENT '真实姓名',
--     `id_number`           varchar(20) NULL COMMENT '身份证号',
--     `id_card_front`       varchar(128) NULL COMMENT '身份证正面',
--     `id_card_back`        varchar(128) NULL COMMENT '身份证背面',
--     `id_card_handheld`    varchar(128) NULL COMMENT '手持身份证',
--     `create_time`         datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
--     `update_time`         datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
--     `register_media_time` datetime NULL COMMENT '注册自媒体时间',
--     `area`                varchar(20) COMMENT '地区',
--     `company`             varchar(20) COMMENT '公司',
--     `career`              varchar(20) COMMENT '职业',
--     `tag`                 varchar(20) COMMENT '标签',
--     PRIMARY KEY (`user_id`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户资料表';

-- CREATE TABLE `user_relation`
-- (
--     `relation_id`    bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
--     `user_id`        bigint(20) unsigned NOT NULL COMMENT '用户ID',
--     `target_user_id` bigint(20) unsigned NOT NULL COMMENT '目标用户ID',
--     `relation`       tinyint(1) NOT NULL DEFAULT '0' COMMENT '关系，0-取消，1-关注，2-拉黑',
--     `create_time`    datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
--     `update_time`    datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
--     PRIMARY KEY (`relation_id`),
--     UNIQUE KEY `user_target` (`user_id`, `target_user_id`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户关系表';

-- CREATE TABLE `user_search`
-- (
--     `search_id`   bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
--     `user_id`     bigint(20) unsigned NOT NULL COMMENT '用户ID',
--     `keyword`     varchar(100) NOT NULL COMMENT '关键词',
--     `create_time` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
--     `is_deleted`  tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除, 0-未删除，1-已删除',
--     `update_time` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
--     PRIMARY KEY (`search_id`),
--     KEY           `user_id` (`user_id`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户搜索历史';

-- CREATE TABLE `user_visitors`
-- (
--     `visit_id`    bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
--     `user_id`     bigint(20) unsigned NOT NULL COMMENT '用户ID',
--     `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
--     `count`       bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '用户访问量',
--     PRIMARY KEY (`visit_id`),
--     KEY           `user_id` (`user_id`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户访问量';

-- CREATE TABLE `user_material`
-- (
--     `material_id`  bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '素材id',
--     `user_id`      bigint(20) unsigned NOT NULL COMMENT '用户ID',
--     `type`         tinyint(1) NOT NULL DEFAULT '0' COMMENT '素材类型，0-图片, 1-视频, 2-音频',
--     `hash`         varchar(128) NULL COMMENT '素材指纹',
--     `url`          varchar(128) NOT NULL COMMENT '素材链接地址',
--     `create_time`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
--     `status`       tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态，0-待审核，1-审核通过，2-审核失败，3-已删除',
--     `reviewer_id`  int(11) unsigned NULL COMMENT '审核人员ID',
--     `review_time`  datetime NULL COMMENT '审核时间',
--     `is_collected` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否收藏，0-未收藏，1-已收藏',
--     `update_time`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
--     PRIMARY KEY (`material_id`),
--     KEY            `user_id` (`user_id`),
--     UNIQUE KEY `user_material` (`user_id`, `hash`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户素材表';

CREATE TABLE `address` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` varchar(20) NOT NULL COMMENT '用户id',
  `title` varchar(20) DEFAULT NULL COMMENT '地址名称',
  `receiver` varchar(20) NOT NULL COMMENT '收货人',
  `province_id` int unsigned NOT NULL COMMENT '省',
  `city_id` int unsigned NOT NULL COMMENT '市',
  `district_id` int unsigned NOT NULL COMMENT '区/县',
  `place` varchar(50) NOT NULL COMMENT '地址',
  `mobile` varchar(11) DEFAULT NULL COMMENT '手机',
  `tel` varchar(20) DEFAULT NULL COMMENT '固定电话',
  `email` varchar(30) DEFAULT NULL COMMENT '电子邮箱',
  `is_default` tinyint(1)  COMMENT '默认地址',
  `is_deleted` tinyint(1)  COMMENT '逻辑删除',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `province_id` (`province_id`),
  KEY `city_id` (`city_id`),
  KEY `district_id` (`district_id`),
  CONSTRAINT `user_id` FOREIGN KEY (`user_id`) REFERENCES `user_basic` (`user_id`) ON DELETE CASCADE,
  CONSTRAINT `province_id` FOREIGN KEY (`province_id`) REFERENCES `city` (`id`) ON DELETE CASCADE,
  CONSTRAINT `city_id` FOREIGN KEY (`city_id`) REFERENCES `city` (`id`) ON DELETE CASCADE,
  CONSTRAINT `district_id` FOREIGN KEY (`district_id`) REFERENCES `city` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户地址';