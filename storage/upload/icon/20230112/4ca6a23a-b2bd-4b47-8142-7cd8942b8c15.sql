-- -------------------------------------------------------------
-- TablePlus 5.1.2(472)
--
-- https://tableplus.com/
--
-- Database: fgzs_single
-- Generation Time: 2023-01-06 19:14:36.7320
-- -------------------------------------------------------------


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


DROP TABLE IF EXISTS `asset`;
CREATE TABLE `asset` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户id',
  `asset_type` tinyint NOT NULL COMMENT '资产类型',
  `amount` bigint NOT NULL COMMENT '资产值',
  `version` int DEFAULT NULL COMMENT '乐观锁',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态(1 正常 2冻结)',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `member_id_idx` (`uid`) USING BTREE,
  KEY `asset_type_idx` (`asset_type`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='资产';

DROP TABLE IF EXISTS `asset_business`;
CREATE TABLE `asset_business` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `key` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '流水号',
  `uid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户ID',
  `common_one` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '关联的业务的自定义ID',
  `common_two` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '关联的业务的自定义ID',
  `common_three` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '关联的业务的自定义ID',
  `business_type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '业务类型',
  `asset_type` tinyint DEFAULT NULL COMMENT '资产类型',
  `asset_opt` tinyint DEFAULT NULL COMMENT '操作类型(1增加，-1减少)',
  `amount` bigint DEFAULT NULL COMMENT '金额',
  `record` json DEFAULT NULL COMMENT '记录',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `consume_asset_type_idx` (`asset_type`) USING BTREE,
  KEY `key_idx` (`key`) USING BTREE,
  KEY `member_id_idx` (`uid`) USING BTREE,
  KEY `business_type_idx` (`business_type`) USING BTREE,
  KEY `common_one_idx` (`common_one`) USING BTREE,
  KEY `common_three_idx` (`common_three`) USING BTREE,
  KEY `common_two_idx` (`common_two`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='资产-消费记录';

DROP TABLE IF EXISTS `asset_log`;
CREATE TABLE `asset_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户ID',
  `key` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '流水号',
  `business_type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '业务类型',
  `asset_type` tinyint DEFAULT NULL COMMENT '资产类型',
  `asset_opt` tinyint NOT NULL COMMENT '操作类型(1增加，-1减少)',
  `amount` bigint NOT NULL COMMENT '变动的金额',
  `before_amount` bigint NOT NULL COMMENT '变更前的金额',
  `after_amount` bigint NOT NULL COMMENT '变更后的金额',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `member_id_idx` (`uid`) USING BTREE,
  KEY `key_idx` (`key`) USING BTREE,
  KEY `business_type_idx` (`business_type`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户资产变更记录表';

DROP TABLE IF EXISTS `file_upload`;
CREATE TABLE `file_upload` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `file_category` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件分类',
  `file_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件新名称',
  `original_file_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件原名称',
  `storage` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '存储方式',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件路径',
  `ext` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件类型',
  `size` bigint DEFAULT NULL COMMENT '文件大小',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态(1 正常 2冻结)',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='文件-上传';

DROP TABLE IF EXISTS `payment_record`;
CREATE TABLE `payment_record` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `key` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '流水号',
  `out_trade_no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '支付订单号',
  `trade_no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '第三方支付的流水号',
  `pay_method` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '支付方式',
  `total_fee` bigint NOT NULL DEFAULT '0' COMMENT '支付金额，整数方式保存(分)',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '商品描述',
  `client_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户的ip',
  `order_time` datetime DEFAULT NULL COMMENT '下单时间',
  `expire_time` datetime DEFAULT NULL COMMENT '订单过期时间',
  `pay_time` datetime DEFAULT NULL COMMENT '第三方支付成功的时间',
  `notify_time` datetime DEFAULT NULL COMMENT '收到异步通知的时间',
  `extension` json DEFAULT NULL COMMENT '扩展字段',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '0:待支付，1 :支付成功，2:已关闭，3:已退款, -1:支付失败',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `key_idx` (`key`) USING BTREE,
  KEY `out_trade_no_idx` (`out_trade_no`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='支付记录表';

DROP TABLE IF EXISTS `sys_admin`;
CREATE TABLE `sys_admin` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '编号',
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
  `password` char(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '昵称',
  `avatar` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '头像',
  `gender` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0=保密 1=女 2=男',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '邮件',
  `mobile` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '手机号',
  `job_id` bigint unsigned DEFAULT NULL COMMENT '岗位',
  `dept_id` bigint unsigned DEFAULT NULL COMMENT '部门',
  `role_ids` json DEFAULT NULL COMMENT '角色集',
  `salt` char(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '盐值',
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '0=禁用 1=开启',
  `sort` bigint unsigned NOT NULL DEFAULT '0' COMMENT '排序值',
  `motto` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '个性签名',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `username` (`username`),
  KEY `mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='系统-用户';

DROP TABLE IF EXISTS `sys_api`;
CREATE TABLE `sys_api` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '编号',
  `group` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '分组',
  `method` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '方法',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '路径',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '描述',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=75 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='系统-工作岗位';

DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '编号',
  `pid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '父级id',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '部门简称',
  `full_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '部门全称',
  `responsible` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '负责人',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '负责人电话',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '负责人邮箱',
  `type` tinyint unsigned NOT NULL DEFAULT '3' COMMENT '1=公司 2=子公司 3=部门',
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '0=禁用 1=开启',
  `sort` bigint unsigned NOT NULL DEFAULT '0' COMMENT '排序值',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='系统-部门';

DROP TABLE IF EXISTS `sys_dictionary`;
CREATE TABLE `sys_dictionary` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '编号',
  `pid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '0=配置集 !0=父级id',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '名称',
  `type` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '1文本 2数字 3数组 4单选 5多选 6下拉 7日期 8时间 9单图 10多图 11单文件 12多文件',
  `unique_key` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '唯一值',
  `value` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '配置值',
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '0=禁用 1=开启',
  `sort` bigint unsigned NOT NULL DEFAULT '0' COMMENT '排序值',
  `remark` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '备注',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_key` (`unique_key`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='系统-参数';

DROP TABLE IF EXISTS `sys_job`;
CREATE TABLE `sys_job` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '岗位名称',
  `code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '岗位编码',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  `sort` bigint unsigned NOT NULL DEFAULT '0' COMMENT '排序值',
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '0=禁用 1=开启 ',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='系统-工作岗位';

DROP TABLE IF EXISTS `sys_log`;
CREATE TABLE `sys_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '编号',
  `admin_id` bigint unsigned NOT NULL COMMENT '管理员ID',
  `ip` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'ip',
  `uri` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '请求路径',
  `useragent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '浏览器标识',
  `header` json DEFAULT NULL COMMENT 'header',
  `req` json DEFAULT NULL COMMENT '请求数据',
  `resp` json DEFAULT NULL COMMENT '响应数据',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='系统-日志';

DROP TABLE IF EXISTS `sys_perm_menu`;
CREATE TABLE `sys_perm_menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `pid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '上级菜单',
  `type` enum('menu_dir','menu','button') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'menu' COMMENT '类型:menu_dir=菜单目录,menu=菜单项,button=页面按钮',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标题',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '规则名称',
  `path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '路由路径',
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '图标',
  `menu_type` enum('tab','link','iframe') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '菜单类型:tab=选项卡,link=链接,iframe=Iframe',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'Url',
  `component` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '组件路径',
  `extend` enum('none','add_rules_only','add_menu_only') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'none' COMMENT '扩展属性:none=无,add_rules_only=只添加为路由,add_menu_only=只添加为菜单',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  `sort` bigint NOT NULL DEFAULT '0' COMMENT '权重(排序)',
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '0=禁用 1=开启',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `pid` (`pid`),
  KEY `sort` (`sort`)
) ENGINE=InnoDB AUTO_INCREMENT=72 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='菜单和权限规则表';

DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '编号',
  `pid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '父级id',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '名称',
  `perm_menu_ids` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '菜单权限集合',
  `remark` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '0=禁用 1=开启',
  `sort` bigint unsigned NOT NULL DEFAULT '0' COMMENT '排序值',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='系统-角色';

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `uid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '全局唯一ID(账户ID)',
  `nickname` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '昵称',
  `sex` tinyint DEFAULT '0' COMMENT '性别0未知1男2女',
  `phone` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '手机号',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户头像',
  `profile` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '简介',
  `level` tinyint DEFAULT '0' COMMENT '会员级别',
  `vip` tinyint DEFAULT '0' COMMENT '会员等级',
  `is_auth` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否认证 0:否 1:是',
  `other` json DEFAULT NULL COMMENT '其他数据',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1 正常 2锁定 3注销',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `member_id_idx` (`uid`) USING BTREE,
  KEY `status` (`status`) USING BTREE,
  KEY `phone` (`phone`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1403 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户信息表';

DROP TABLE IF EXISTS `user_auth`;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户信息表';

DROP TABLE IF EXISTS `user_cancellation`;
CREATE TABLE `user_cancellation` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '用户ID',
  `reason` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '申请理由',
  `apply_time` datetime NOT NULL COMMENT '申请时间',
  `confirm_time` datetime DEFAULT NULL COMMENT '确认时间',
  `confirm_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '确认备注',
  `status` tinyint DEFAULT NULL COMMENT '处理状态（1待处理，2注销通过，3注销驳回）',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户注销表';

DROP TABLE IF EXISTS `user_feedback`;
CREATE TABLE `user_feedback` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户ID',
  `content` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '反馈内容',
  `remark` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注内容',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `member_feedback_member_id_idx` (`uid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `sys_admin` (`id`, `username`, `password`, `nickname`, `avatar`, `gender`, `email`, `mobile`, `job_id`, `dept_id`, `role_ids`, `salt`, `status`, `sort`, `motto`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'admin', '$2a$10$fi2i8Si8SyLkPbSlS.I/QOl4SiPqbLEK.i5EDjC3Lwbo0ALi2OiCC', '超级管理员', 'https://avataaars.io/?clotheColor=Black&accessoriesType=Prescription02&avatarStyle=Circle&clotheType=ShirtCrewNeck&eyeType=Cry&eyebrowType=FlatNatural&facialHairColor=BrownDark&facialHairType=MoustacheFancy&hairColor=Brown&hatColor=Blue02&mouthType=Grimace&skinColor=Pale&topType=ShortHairShortCurly', 0, 'admin@163.com', '18666666666', 0, 0, '[1]', 'LYPqyi9wv64akuc5', 1, 0, '嘻嘻哈哈', '2022-11-27 01:17:21', '2022-12-09 16:35:05', NULL),
(3, 'fan', '$2a$10$fi2i8Si8SyLkPbSlS.I/QOl4SiPqbLEK.i5EDjC3Lwbo0ALi2OiCC', '凡', 'https://avataaars.io/?clotheColor=Black&accessoriesType=Prescription02&avatarStyle=Circle&clotheType=ShirtCrewNeck&eyeType=Cry&eyebrowType=FlatNatural&facialHairColor=BrownDark&facialHairType=MoustacheFancy&hairColor=Brown&hatColor=Blue02&mouthType=Grimace&skinColor=Pale&topType=ShortHairShortCurly', 1, '1264857660@qq.com', '13542864871', 0, 0, '[1]', 'LYPqyi9wv64akuc5', 1, 0, '是干啥干啥干啥', '2022-11-27 01:17:21', '2022-12-10 17:49:56', NULL),
(4, 'fei', '$2a$10$7ORWhu3r2tJTPB3Ormve8errI2QoZ6wDNB/WzV0aJT4SowdDg9br6', '飞飞飞', '', 2, '', '', 2, 5, '[8]', 'gggggggggggggggg', 1, 0, '', '2022-12-09 16:17:59', '2022-12-09 16:32:49', NULL);

INSERT INTO `sys_api` (`id`, `group`, `method`, `path`, `desc`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'system', 'GET', '/system/ping', 'ping', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(2, 'dashboard', 'GET', '/dashboard/speech', '言语', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(3, 'device/log', 'POST', '/device/log/list', '设备-日志列表', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(4, 'sys/dept', 'POST', '/sys/dept/list', '部门列表', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(5, 'sys/dept', 'POST', '/sys/dept/info', '单个部门', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(6, 'sys/dept', 'POST', '/sys/dept/store', '保存部门', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(7, 'sys/dept', 'POST', '/sys/dept/del', '删除部门', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(8, 'sys/job', 'POST', '/sys/job/list', '岗位列表', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(9, 'sys/job', 'POST', '/sys/job/info', '单个岗位', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(10, 'sys/job', 'POST', '/sys/job/store', '保存岗位', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(11, 'sys/job', 'POST', '/sys/job/del', '删除岗位', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(12, 'sys/permmenu', 'POST', '/sys/permmenu/list', '权限菜单列表', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(13, 'sys/permmenu', 'POST', '/sys/permmenu/info', '单个权限菜单', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(14, 'sys/permmenu', 'POST', '/sys/permmenu/store', '保存权限菜单', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(15, 'sys/permmenu', 'POST', '/sys/permmenu/del', '删除权限菜单', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(16, 'sys/permmenu', 'POST', '/sys/permmenu/status', '修改权限菜单状态', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(17, 'sys/role', 'POST', '/sys/role/list', '角色列表', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(18, 'sys/role', 'POST', '/sys/role/info', '单个角色', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(19, 'sys/role', 'POST', '/sys/role/store', '保存角色', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(20, 'sys/role', 'POST', '/sys/role/del', '删除角色', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(21, 'sys/admin', 'GET', '/sys/admin/login/captcha', '登录验证码', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(22, 'sys/admin', 'POST', '/sys/admin/login', '登录', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(23, 'sys/admin', 'POST', '/sys/admin/logout', '退出', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(24, 'sys/admin', 'GET', '/sys/admin/info', '管理员信息', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(25, 'sys/admin', 'POST', '/sys/admin/info/update', '管理员信息', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(26, 'sys/admin', 'GET', '/sys/admin/permmenu', '菜单权限', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(27, 'sys/admin', 'GET', '/sys/admin/avatar/generate', '生成头像', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(28, 'sys/manage', 'POST', '/sys/manage/list', '管理员列表', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(29, 'sys/manage', 'POST', '/sys/manage/info', '单个管理员', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(30, 'sys/manage', 'POST', '/sys/manage/store', '保存管理员', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(31, 'sys/manage', 'POST', '/sys/manage/del', '删除管理员', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(32, 'sys/log', 'POST', '/sys/log/ownlist', '自身日志列表', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(33, 'sys/log', 'POST', '/sys/log/list', '日志列表', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(34, 'sys/log', 'POST', '/sys/log/info', '单条日志', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(35, 'res/book', 'POST', '/res/book/list', '示例列表', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(36, 'res/book', 'POST', '/res/book/info', '单个示例', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(37, 'res/book', 'POST', '/res/book/store', '保存示例', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(38, 'res/book', 'POST', '/res/book/del', '删除示例', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(39, 'res/bookfile', 'POST', '/res/bookfile/list', '书籍文件列表', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(40, 'res/bookfile', 'POST', '/res/bookfile/info', '单个书籍文件', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(41, 'res/bookfile', 'POST', '/res/bookfile/store', '保存书籍文件', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(42, 'res/bookfile', 'POST', '/res/bookfile/del', '删除书籍文件', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(43, 'res/category', 'POST', '/res/category/list', '分类列表', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(44, 'res/category', 'POST', '/res/category/info', '单个分类', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(45, 'res/category', 'POST', '/res/category/store', '保存分类', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(46, 'res/category', 'POST', '/res/category/del', '删除分类', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(47, 'res/tag', 'POST', '/res/tag/list', '标签列表', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(48, 'res/tag', 'POST', '/res/tag/info', '单个标签', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(49, 'res/tag', 'POST', '/res/tag/store', '保存标签', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(50, 'res/tag', 'POST', '/res/tag/del', '删除标签', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(51, 'res/tagoption', 'POST', '/res/tagoption/list', '标签选项列表', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(52, 'res/tagoption', 'POST', '/res/tagoption/info', '单个标签选项', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(53, 'res/tagoption', 'POST', '/res/tagoption/store', '保存标签选项', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(54, 'res/tagoption', 'POST', '/res/tagoption/del', '删除标签选项', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(55, 'res/version', 'POST', '/res/version/list', '分类列表', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(56, 'res/version', 'POST', '/res/version/info', '单个分类', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(57, 'res/version', 'POST', '/res/version/store', '保存分类', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(58, 'res/version', 'POST', '/res/version/del', '删除分类', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(59, 'wordbook/catalog', 'POST', '/wordbook/catalog/list', '单词本目录列表', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(60, 'wordbook/catalog', 'POST', '/wordbook/catalog/info', '单个单词本目录', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(61, 'wordbook/catalog', 'POST', '/wordbook/catalog/store', '保存单词本目录', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(62, 'wordbook/catalog', 'POST', '/wordbook/catalog/del', '删除单词本目录', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(63, 'wordbook/data', 'POST', '/wordbook/data/list', '单词本数据列表', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(64, 'wordbook/data', 'POST', '/wordbook/data/info', '单个单词本数据', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(65, 'wordbook/data', 'POST', '/wordbook/data/store', '保存单词本数据', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(66, 'wordbook/data', 'POST', '/wordbook/data/del', '删除单词本数据', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(67, 'wordbook/version', 'POST', '/wordbook/version/list', '单词本版本列表', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(68, 'wordbook/version', 'POST', '/wordbook/version/info', '单个单词本版本', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(69, 'wordbook/version', 'POST', '/wordbook/version/store', '保存单词本版本', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(70, 'wordbook/version', 'POST', '/wordbook/version/del', '删除单词本版本', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(71, 'behavior/studystate', 'POST', '/behavior/studystate/list', '学习状态列表', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(72, 'behavior/studystate', 'POST', '/behavior/studystate/info', '单个学习状态', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(73, 'behavior/studystate', 'POST', '/behavior/studystate/store', '保存学习状态', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL),
(74, 'behavior/studystate', 'POST', '/behavior/studystate/del', '删除学习状态', '2022-12-10 16:09:40', '2022-12-10 16:09:40', NULL);

INSERT INTO `sys_dept` (`id`, `pid`, `name`, `full_name`, `responsible`, `phone`, `email`, `type`, `status`, `sort`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 0, '深圳公司', '深圳公司', NULL, NULL, NULL, 1, 1, 0, '2022-12-08 18:24:48', '2022-12-08 18:25:40', NULL),
(2, 1, '总经办', '总经办', NULL, NULL, NULL, 3, 1, 0, '2022-12-08 18:25:35', '2022-12-08 18:27:23', NULL),
(3, 1, '人力资源', '人力资源', NULL, NULL, NULL, 3, 1, 0, '2022-12-08 18:26:58', '2022-12-08 18:26:58', NULL),
(4, 1, '市场部', '市场部', NULL, NULL, NULL, 3, 1, 0, '2022-12-08 18:26:58', '2022-12-08 18:26:58', NULL),
(5, 1, '技术部', '技术部', NULL, NULL, NULL, 3, 1, 0, '2022-12-08 18:26:58', '2022-12-08 18:26:58', NULL),
(6, 5, '后端开发', '后端开发', '', '', '', 3, 1, 0, '2022-12-08 19:08:06', '2022-12-08 19:11:30', '2022-12-08 19:11:30'),
(7, 5, '前端开发', '前端开发', '小李2', '18665550229', '382391525@qq.com', 3, 1, 0, '2022-12-08 19:08:48', '2022-12-08 19:11:28', '2022-12-08 19:11:28'),
(8, 5, '安卓开发', '安卓开发', '', '', '', 3, 1, 0, '2022-12-08 21:48:55', '2022-12-08 21:52:37', '2022-12-08 21:52:38');

INSERT INTO `sys_job` (`id`, `name`, `code`, `remark`, `sort`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, '总经理', NULL, NULL, 0, 1, '2022-11-27 17:53:44', '2022-11-27 17:53:44', NULL),
(2, '后端开发', NULL, NULL, 0, 1, '2022-11-27 17:54:14', '2022-11-27 17:54:14', NULL),
(3, '前端开发', NULL, NULL, 0, 1, '2022-11-27 17:54:21', '2022-11-27 17:54:21', NULL),
(4, '技术经理', 'J0001', '1321312312', 0, 1, '2022-12-08 21:42:24', '2022-12-08 21:44:23', NULL);

INSERT INTO `sys_perm_menu` (`id`, `pid`, `type`, `title`, `name`, `path`, `icon`, `menu_type`, `url`, `component`, `extend`, `remark`, `sort`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 0, 'menu', '控制台', 'dashboard', 'dashboard/dashboard', 'fa fa-dashboard', 'tab', '', '/src/views/dashboard/dashboard.vue', 'none', 'remark_text', 999, 1, '2022-11-27 17:44:20', '2022-12-03 11:20:34', NULL),
(2, 0, 'menu_dir', '权限管理', 'auth', 'auth', 'fa fa-group', NULL, '', '', 'none', '', 100, 1, '2022-11-27 17:44:20', '2022-11-27 17:44:20', NULL),
(3, 2, 'menu', '管理员', 'auth/manage', 'auth/manage', 'el-icon-UserFilled', 'tab', '', '/src/views/auth/manage/index.vue', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-03 22:22:45', NULL),
(4, 3, 'button', '查看', 'auth/manage/index', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-03 22:28:17', NULL),
(5, 3, 'button', '添加', 'auth/manage/add', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-03 22:28:17', NULL),
(6, 3, 'button', '编辑', 'auth/manage/edit', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-03 22:28:17', NULL),
(7, 3, 'button', '删除', 'auth/manage/del', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-03 22:28:17', NULL),
(8, 2, 'menu', '角色管理', 'auth/role', 'auth/role', 'fa fa-group', 'tab', '', '/src/views/auth/role/index.vue', 'none', '', 99, 1, '2022-11-27 17:44:20', '2022-11-30 00:35:29', NULL),
(9, 8, 'button', '查看', 'auth/role/index', '', '', NULL, '', '', 'none', '', 99, 1, '2022-11-27 17:44:20', '2022-12-03 22:34:38', NULL),
(10, 8, 'button', '添加', 'auth/role/add', '', '', NULL, '', '', 'none', '', 99, 1, '2022-11-27 17:44:20', '2022-12-03 22:34:38', NULL),
(11, 8, 'button', '编辑', 'auth/role/edit', '', '', NULL, '', '', 'none', '', 99, 1, '2022-11-27 17:44:20', '2022-12-03 22:34:38', NULL),
(12, 8, 'button', '删除', 'auth/role/del', '', '', NULL, '', '', 'none', '', 99, 1, '2022-11-27 17:44:20', '2022-12-03 22:34:38', NULL),
(13, 2, 'menu', '菜单管理', 'auth/permmenu', 'auth/permmenu', 'el-icon-Grid', 'tab', '', '/src/views/auth/permmenu/index.vue', 'none', '', 97, 1, '2022-11-27 17:44:20', '2022-12-03 22:33:07', NULL),
(14, 13, 'button', '查看', 'auth/permmenu/index', '', '', NULL, '', '', 'none', '', 97, 1, '2022-11-27 17:44:20', '2022-12-03 11:32:32', NULL),
(15, 13, 'button', '添加', 'auth/permmenu/add', '', '', NULL, '', '', 'none', '', 97, 1, '2022-11-27 17:44:20', '2022-12-03 11:32:32', NULL),
(16, 13, 'button', '编辑', 'auth/permmenu/edit', '', '', NULL, '', '', 'none', '', 97, 1, '2022-11-27 17:44:20', '2022-12-03 11:32:32', NULL),
(17, 13, 'button', '删除', 'auth/permmenu/del', '', '', NULL, '', '', 'none', '', 97, 1, '2022-11-27 17:44:20', '2022-12-03 11:32:32', NULL),
(18, 13, 'button', '快速排序', 'auth/menu/sortable', '', '', NULL, '', '', 'none', '', 97, 1, '2022-11-27 17:44:20', '2022-11-27 17:44:20', NULL),
(19, 2, 'menu', '部门管理', 'auth/dept', 'auth/dept', 'fa fa-sitemap', 'tab', '', '/src/views/auth/dept/index.vue', 'none', '', 0, 1, '2022-11-27 17:44:20', '2022-12-03 00:20:20', NULL),
(20, 19, 'button', '查看', 'auth/dept/index', '', '', NULL, '', '', 'none', '', 0, 1, '2022-11-27 17:44:20', '2022-12-03 22:20:50', NULL),
(21, 19, 'button', '添加', 'auth/dept/add', '', '', NULL, '', '', 'none', '', 0, 1, '2022-11-27 17:44:20', '2022-12-03 22:20:50', NULL),
(22, 19, 'button', '编辑', 'auth/dept/edit', '', '', NULL, '', '', 'none', '', 0, 1, '2022-11-27 17:44:20', '2022-12-03 22:20:50', NULL),
(23, 19, 'button', '删除', 'auth/dept/del', '', '', NULL, '', '', 'none', '', 0, 1, '2022-11-27 17:44:20', '2022-12-03 22:20:50', NULL),
(24, 19, 'button', '快速排序', 'auth/dept/sortable', '', '', NULL, '', '', 'none', '', 0, 1, '2022-11-27 17:44:20', '2022-12-03 22:20:50', NULL),
(25, 2, 'menu', '岗位管理', 'auth/job', 'auth/job', 'fa fa-id-badge', 'tab', '', '/src/views/auth/job/index.vue', 'none', '', 0, 1, '2022-11-27 17:44:20', '2022-12-03 00:21:26', NULL),
(26, 25, 'button', '查看', 'auth/job/index', '', '', NULL, '', '', 'none', '', 0, 1, '2022-11-27 17:44:20', '2022-12-03 22:21:10', NULL),
(27, 25, 'button', '添加', 'auth/job/add', '', '', NULL, '', '', 'none', '', 0, 1, '2022-11-27 17:44:20', '2022-12-03 22:21:10', NULL),
(28, 25, 'button', '编辑', 'auth/job/edit', '', '', NULL, '', '', 'none', '', 0, 1, '2022-11-27 17:44:20', '2022-12-03 22:21:09', NULL),
(29, 25, 'button', '删除', 'auth/job/del', '', '', NULL, '', '', 'none', '', 0, 1, '2022-11-27 17:44:20', '2022-12-03 22:21:09', NULL),
(30, 25, 'button', '快速排序', 'auth/job/sortable', '', '', NULL, '', '', 'none', '', 0, 1, '2022-11-27 17:44:20', '2022-12-03 22:21:09', NULL),
(31, 0, 'menu_dir', '常规管理', 'routine', 'routine', 'fa fa-cogs', NULL, '', '', 'none', '', 89, 1, '2022-11-27 17:44:20', '2022-12-06 18:51:29', NULL),
(32, 31, 'menu', '系统配置', 'routine/config', 'routine/config', 'el-icon-Tools', 'tab', '', '/src/views/routine/config/index.vue', 'none', '', 88, 0, '2022-11-27 17:44:20', '2022-12-09 18:39:01', NULL),
(33, 32, 'button', '查看', 'routine/config/index', '', '', NULL, '', '', 'none', '', 88, 1, '2022-11-27 17:44:20', '2022-12-03 22:20:38', NULL),
(34, 32, 'button', '编辑', 'routine/config/edit', '', '', NULL, '', '', 'none', '', 88, 1, '2022-11-27 17:44:20', '2022-12-03 22:20:38', NULL),
(35, 31, 'menu', '操作日志', 'routine/adminLog', 'routine/adminLog', 'el-icon-List', 'tab', '', '/src/views/routine/adminLog/index.vue', 'none', '', 96, 1, '2022-11-27 17:44:20', '2022-12-03 22:21:49', NULL),
(36, 35, 'button', '查看', 'routine/adminLog/index', '', '', NULL, '', '', 'none', '', 96, 1, '2022-11-27 17:44:20', '2022-12-03 23:43:00', NULL),
(37, 0, 'menu_dir', '设备管理', 'device', 'device', 'fa fa-cubes', NULL, '', '', 'none', '', 0, 1, '2022-12-04 20:35:32', '2022-12-07 14:51:30', NULL),
(38, 37, 'menu', '设备日志', 'device/deviceLog', 'device/deviceLog', 'el-icon-List', 'tab', '', '/src/views/device/deviceLog/index.vue', 'none', '', 96, 1, '2022-11-27 17:44:20', '2022-12-10 11:19:28', NULL),
(39, 38, 'button', '查看', 'device/log/index', '', '', NULL, '', '', 'none', '', 96, 1, '2022-11-27 17:44:20', '2022-12-10 11:19:28', NULL),
(40, 0, 'menu_dir', '本地化资源', 'resource', 'resource', 'fa fa-music', NULL, '', '', 'none', '', 0, 1, '2022-12-07 14:52:41', '2022-12-07 14:55:22', NULL),
(41, 40, 'menu', '分类管理', 'resource/category', 'resource/category', 'fa fa-list', 'tab', '', '/src/views/resource/category/index.vue', 'none', '', 0, 1, '2022-12-07 14:57:36', '2022-12-10 11:20:00', NULL),
(42, 41, 'button', '查看', 'resource/category/index', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-10 11:20:00', NULL),
(43, 41, 'button', '添加', 'resource/category/add', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-10 11:20:00', NULL),
(44, 41, 'button', '编辑', 'resource/category/edit', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-10 11:20:00', NULL),
(45, 41, 'button', '删除', 'resource/category/del', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-10 11:20:00', NULL),
(46, 40, 'menu', '书籍管理', 'resource/book', 'resource/book', 'fa fa-book', 'tab', '', '/src/views/resource/book/index.vue', 'none', '', 0, 1, '2022-12-07 14:57:36', '2022-12-10 11:20:00', NULL),
(47, 46, 'button', '查看', 'resource/book/index', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-10 11:20:38', NULL),
(48, 46, 'button', '添加', 'resource/book/add', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-10 11:20:38', NULL),
(49, 46, 'button', '编辑', 'resource/book/edit', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-10 11:20:38', NULL),
(50, 46, 'button', '删除', 'resource/book/del', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-10 11:20:38', NULL),
(51, 40, 'menu', '版本管理', 'resource/version', 'resource/version', 'fa fa-picture-o', 'tab', '', '/src/views/resource/version/index.vue', 'none', '', 0, 1, '2022-12-07 14:57:36', '2022-12-10 11:20:38', NULL),
(52, 51, 'button', '查看', 'resource/version/index', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-10 11:20:38', NULL),
(53, 0, 'menu_dir', '单词本', 'wordbook', 'wordbook', 'fa fa-book', NULL, '', '', 'none', '', 0, 1, '2022-12-07 14:52:41', '2022-12-08 10:11:20', NULL),
(54, 53, 'menu', '目录管理', 'wordbook/catalog', 'wordbook/catalog', 'fa fa-list', 'tab', '', '/src/views/wordbook/catalog/index.vue', 'none', '', 0, 1, '2022-12-07 14:57:36', '2022-12-10 11:21:39', NULL),
(55, 54, 'button', '查看', 'wordbook/catalog/index', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-10 11:21:39', NULL),
(56, 54, 'button', '添加', 'wordbook/catalog/add', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-10 11:21:39', NULL),
(57, 54, 'button', '编辑', 'wordbook/catalog/edit', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-10 11:21:39', NULL),
(58, 54, 'button', '删除', 'wordbook/catalog/del', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-10 11:21:39', NULL),
(59, 53, 'menu', '单词管理', 'wordbook/data', 'wordbook/data', 'fa fa-file-word-o', 'tab', '', '/src/views/wordbook/data/index.vue', 'none', '', 0, 1, '2022-12-07 14:57:36', '2022-12-10 11:21:39', NULL),
(60, 59, 'button', '查看', 'wordbook/data/index', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-10 11:21:39', NULL),
(61, 59, 'button', '添加', 'wordbook/data/add', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-10 11:21:39', NULL),
(62, 59, 'button', '编辑', 'wordbook/data/edit', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-10 11:21:39', NULL),
(63, 59, 'button', '删除', 'wordbook/data/del', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-10 11:21:39', NULL),
(64, 53, 'menu', '版本管理', 'wordbook/version', 'wordbook/version', 'fa fa-picture-o', 'tab', '', '/src/views/wordbook/version/index.vue', 'none', '', 0, 1, '2022-12-07 14:57:36', '2022-12-10 11:21:39', NULL),
(65, 64, 'button', '查看', 'wordbook/version/index', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-10 11:21:39', NULL),
(66, 0, 'menu_dir', '用户行为', 'behavior', '', 'el-icon-Stamp', NULL, '', '', 'none', '', 0, 1, '2022-12-08 22:24:02', '2022-12-08 22:30:04', NULL),
(67, 66, 'menu', '学习状况', 'behavior/studystate', 'behavior/studystate', 'el-icon-TrophyBase', 'tab', '', '/src/views/behavior/studystate/index.vue', 'none', '', 0, 1, '2022-12-08 22:24:48', '2022-12-10 11:21:39', NULL),
(68, 67, 'button', '查看', 'behavior/studystate/index', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-10 11:21:39', NULL),
(69, 67, 'button', '添加', 'behavior/studystate/add', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-10 11:21:39', NULL),
(70, 67, 'button', '编辑', 'behavior/studystate/edit', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-10 11:21:39', NULL),
(71, 67, 'button', '删除', 'behavior/studystate/del', '', '', NULL, '', '', 'none', '', 98, 1, '2022-11-27 17:44:20', '2022-12-10 11:21:39', NULL);

INSERT INTO `sys_role` (`id`, `pid`, `name`, `perm_menu_ids`, `remark`, `status`, `sort`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 0, '超级管理员', '*', '最高权限角色', 1, 0, '2022-11-27 17:57:23', '2022-11-27 17:58:59', NULL),
(2, 0, '开发工程师', NULL, NULL, 1, 0, '2022-11-27 17:59:31', '2022-11-27 17:59:31', NULL),
(3, 1, '一级管理员', NULL, NULL, 1, 0, '2022-11-29 22:07:38', '2022-11-29 22:08:01', NULL),
(4, 3, '二级管理员', '1', '131231231', 1, 0, '2022-11-29 22:07:48', '2022-12-08 21:57:54', '2022-12-08 21:57:54'),
(5, 5, '三级管理员', NULL, NULL, 1, 0, '2022-11-29 22:07:54', '2022-11-29 22:08:03', NULL),
(6, 2, '12321321', '1', '12313123', 1, 0, '2022-12-02 12:01:51', '2022-12-02 14:04:01', '2022-12-02 14:04:01'),
(7, 3, '1321', '27,28,29,30,31,21', '1232', 1, 0, '2022-12-02 12:21:11', '2022-12-02 14:04:05', '2022-12-02 14:04:05'),
(8, 2, '后端开发工程师', '1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,88,89,90,91,92,93,94,95,96,97,98,99', '1232131', 1, 0, '2022-12-02 13:45:03', '2022-12-02 13:45:32', NULL),
(9, 2, '安卓工程师', '1,62,63,64,65,66,56', '哈哈哈', 1, 0, '2022-12-08 21:57:34', '2022-12-08 21:57:44', NULL),
(10, 0, '设备管理', '37,38,39', '设备管理', 1, 0, '2022-12-10 09:33:39', '2022-12-10 11:24:14', NULL),
(11, 2, '前端开发工程师', '1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70,71', '', 1, 0, '2022-12-20 12:03:11', '2022-12-20 12:03:11', NULL);



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;