CREATE TABLE `context_models`
(
    `id`        INT ( 11 ) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`      VARCHAR(20) NOT NULL DEFAULT '' COMMENT '姓名',
    `timestamp` bigint(20) NOT NULL DEFAULT 0 COMMENT '时间戳',
    `date`      VARCHAR(10) NOT NULL DEFAULT '' COMMENT '日期',
    `msg`       BLOB        NOT NULL COMMENT '存储行数据',
    PRIMARY KEY (`id`),
    KEY         `d_idx` ( `date` ) USING BTREE
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = '数据内容模型';