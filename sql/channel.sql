SET
@@auto_increment_increment=9;

CREATE TABLE `news_channel`
(
    `channel_id`   int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '频道ID',
    `channel_name` varchar(32) NOT NULL COMMENT '频道名称',
    `create_time`  datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `sequence`     int(11) unsigned NOT NULL DEFAULT '0' COMMENT '序号',
    `is_visible`   tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否可见',
    `is_default`   tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否默认',
    PRIMARY KEY (`channel_id`),
    UNIQUE KEY `channel_name` (`channel_name`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='新闻频道表';

CREATE TABLE `news_user_channel`
(
    `user_channel_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `user_id`         bigint(20) unsigned NOT NULL COMMENT '用户ID',
    `channel_id`      int(11) unsigned NOT NULL COMMENT '频道ID',
    `create_time`     datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `is_deleted`      tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除, 0-未删除, 1-已删除',
    `update_time`     datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `sequence`        int(11) unsigned NOT NULL DEFAULT '0' COMMENT '序号',
    PRIMARY KEY (`user_channel_id`),
    UNIQUE KEY `user_channel` (`user_id`, `channel_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户关注频道表';