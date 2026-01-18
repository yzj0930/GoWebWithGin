drop database if exists goodsplateform;
create database goodsplateform;
use goodsplateform;

-- 用户表 
drop table if exists users;
CREATE TABLE `users` (
   `user_id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键 用户ID',
   `user_name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
   `user_code` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户唯一编码',
   `telephone` varchar(20) NOT NULL DEFAULT "" COMMENT '用户联系方式',
   `password` varchar(100) NOT NULL COMMENT '用户密码加密后存储',
   `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
   `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
   PRIMARY KEY (`user_id`),
   UNIQUE KEY `uk_user_code` (`user_code`)
 ) ENGINE=InnoDB COMMENT='用户表';
 
-- 权限表
drop table if exists `privileges`;
create table `privileges` (
	`privilege_id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键 功能权限ID',
    `privilege_name`  varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '功能权限名称',
    `privilege_code`  varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '功能权限编码',
    `description` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '权限描述',
    `remark` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    primary key(`privilege_id`),
    UNIQUE KEY `uk_privilege_code` (`privilege_code`)
) ENGINE=InnoDB COMMENT='权限表';

-- 用户权限表
drop table if exists `user_auth`;
create table `user_auth` (
	`user_id` bigint NOT NULL COMMENT '用户ID',
    `privilege_id` bigint NOT NULL COMMENT '权限ID',
    `status` ENUM('valid', 'invalid', 'close') DEFAULT 'valid' COMMENT '状态',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB COMMENT='用户权限表';

-- 项目表
drop table if exists projects;
create table `projects` (
	`project_id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键 项目ID',
    `project_name`  varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '项目名称',
	`project_code` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '项目唯一编码',
    `remark` varchar(500) NOT NULL DEFAULT '' COMMENT '备注',
    `extra_info` json NOT NULL DEFAULT ('{}') COMMENT '项目额外信息',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `expire_time` datetime NOT NULL DEFAULT (DATE_ADD(CURRENT_TIMESTAMP, INTERVAL 1 MONTH)),
    primary key(`project_id`),
    UNIQUE KEY `project_code` (`project_code`)
) ENGINE=InnoDB COMMENT='项目表';

-- 用户数据权限关联表
drop table if exists user_project;
create table `user_project` (
	`user_id` bigint NOT NULL COMMENT '用户ID',
    `project_id` bigint NOT NULL COMMENT '项目ID',
    `status` ENUM('valid', 'invalid', 'close') DEFAULT 'valid' COMMENT '状态',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB COMMENT='用户数据权限关联表';

-- 货品表
drop table if exists category;
create table `category` (
	`category_id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键 货品ID',
    `project_id` bigint NOT NULL COMMENT '项目ID',
	`category_name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '货品名称',
	`category_code` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '货品编码',  
    `is_valid` boolean NOT NULL DEFAULT true COMMENT '货品生效标识',
	`remark` varchar(500) NOT NULL DEFAULT '' COMMENT '备注',
    `extra_info` json NOT NULL DEFAULT ('{}') COMMENT '项目额外信息',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    primary key(`category_id`),
    UNIQUE KEY `uk_project_code`(`project_id`, `category_code`)
) ENGINE=InnoDB COMMENT='货品表';

-- 库存表
drop table if exists inventory;
create table `inventory` (
	`id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键 商品ID',
    `project_id` bigint NOT NULL COMMENT '项目ID',
    `category_id` bigint NOT NULL COMMENT '货品ID',
    `total_num` bigint NOT NULL DEFAULT 0 COMMENT '总数量',
    `available_num` bigint NOT NULL DEFAULT 0 COMMENT '可用数量',
    `sold_num` bigint NOT NULL DEFAULT 0 COMMENT '已售数量',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    primary key(`id`),
    UNIQUE KEY `uk_project_category` (`project_id`, `category_id`)
) ENGINE=InnoDB COMMENT='库存表';

-- 商品表
drop table if exists sku;
create table `sku` (
	`sku_id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键 商品ID',
    `project_id` bigint NOT NULL COMMENT '项目ID',
    `sku_name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '商品名称',
    `sku_code` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '商品编码', 
    `price` float8 NOT NULL DEFAULT 0 COMMENT '进价',
    `exp_price` float8 NOT NULL DEFAULT 0 COMMENT '预售价',
    `act_price` float8 NOT NULL DEFAULT 0 COMMENT '实际售价',
    `is_sold` boolean NOT NULL DEFAULT false COMMENT '已售',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_timeusersusers` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    primary key(`sku_id`),
    UNIQUE KEY `uk_project_sku` (`project_id`, `sku_code`)
) ENGINE=InnoDB COMMENT='库存表';
    
    
    
    