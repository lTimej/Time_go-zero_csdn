SET
@@auto_increment_increment=9;

CREATE TABLE `news_article_basic`
(
    `article_id`     bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '文章ID',
    `user_id`        bigint(20) unsigned NOT NULL COMMENT '用户ID',
    `channel_id`     int(11) unsigned NOT NULL COMMENT '频道ID',
    `title`          varchar(128) NOT NULL COMMENT '标题',
    `is_advertising` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否投放广告，0-不投放，1-投放',
    `create_time`    datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`    datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `status`         tinyint(1) NOT NULL DEFAULT '0' COMMENT '贴文状态，0-草稿，1-待审核，2-审核通过，3-审核失败，4-已删除',
    `reviewer_id`    int(11) NULL COMMENT '审核人员ID',
    `review_time`    datetime NULL COMMENT '审核时间',
    `delete_time`    datetime NULL COMMENT '删除时间',
    `reject_reason`  varchar(200) COMMENT '驳回原因',
    `comment_count`  int(11) unsigned NOT NULL DEFAULT '0' COMMENT '累计评论数',
    `allow_comment`  tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否允许评论，0-不允许，1-允许',
    PRIMARY KEY (`article_id`),
    KEY              `user_id` (`user_id`),
    KEY              `article_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章基本信息表';

CREATE TABLE `news_article_content`
(
    `article_id` bigint(20) unsigned NOT NULL COMMENT '文章ID',
    `content`    longtext NOT NULL COMMENT '文章内容',
    PRIMARY KEY (`article_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='文章内容表';

CREATE TABLE `news_article_statistic`
(
    `article_id`         bigint(20) unsigned NOT NULL COMMENT '文章ID',
    `read_count`         int(11) unsigned NOT NULL DEFAULT '0' COMMENT '阅读量',
    `like_count`         int(11) unsigned NOT NULL DEFAULT '0' COMMENT '点赞数',
    `dislike_count`      int(11) unsigned NOT NULL DEFAULT '0' COMMENT '不喜欢数',
    `repost_count`       int(11) unsigned NOT NULL DEFAULT '0' COMMENT '转发数',
    `collect_count`      int(11) unsigned NOT NULL DEFAULT '0' COMMENT '收藏数',
    `fans_comment_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '粉丝评论数',
    PRIMARY KEY (`article_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章统计表';

CREATE TABLE `news_collection`
(
    `collection_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `user_id`       bigint(20) unsigned NOT NULL COMMENT '用户ID',
    `article_id`    bigint(20) unsigned NOT NULL COMMENT '文章ID',
    `create_time`   datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `is_deleted`    tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否取消收藏, 0-未取消, 1-已取消',
    `update_time`   datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`collection_id`),
    UNIQUE KEY `user_article` (`user_id`, `article_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户收藏表';