SET
@@auto_increment_increment=9;

CREATE TABLE `city` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `code` int DEFAULT NULL COMMENT '行政区划代码',
  `name` varchar(32) DEFAULT NULL COMMENT '名称',
  `pid` int DEFAULT NULL COMMENT '上级id',
  `type` varchar(32) DEFAULT NULL COMMENT '类型',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='中国地图';