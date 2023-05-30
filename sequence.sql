CREATE TABLE `sequence`(
   `id` bigint(20) unsigned NOT NULL auto_increment,
   `stub` varchar(1) NOT NULL,
   `timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   PRIMARY kEY(id),
   UNIQUE KEY `idx_uniq_stub` (`stub`)
)ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='序号表';