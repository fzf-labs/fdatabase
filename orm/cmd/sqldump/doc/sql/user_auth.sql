CREATE TABLE `user_auth` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '用户ID',
  `identity_type` tinyint NOT NULL DEFAULT '0' COMMENT '1 微信 2 苹果',
  `identity_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '业务登录key',
  `identifier_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '标识码',
  `identity_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '昵称',
  `other` json DEFAULT NULL COMMENT '其他数据',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态 1绑定 0解绑',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `identifier` (`identifier_code`) USING BTREE,
  KEY `member_id` (`uid`) USING BTREE,
  KEY `identity_key` (`identity_key`) USING BTREE,
  KEY `identity_type` (`identity_type`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户权限表'