SET NAMES utf8;
SET FOREIGN_KEY_CHECKS=0;

#--------the user---start------#

DROP TABLE IF EXISTS `zd_uc_role_auth`;
CREATE TABLE `zd_uc_role_auth`(
  `role_id` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '角色ID',
  `auth_id` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '权限ID',
  PRIMARY KEY (`role_id`,`auth_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限和角色关系管理';

DROP TABLE IF EXISTS `zd_uc_role`;
CREATE TABLE `zd_uc_role` (
  `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `role_name` VARCHAR(32) NOT NULL DEFAULT '0' COMMENT '角色名称',
  `detail` VARCHAR(255) NOT NULL DEFAULT '0' COMMENT '备注',
  `create_id` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '修改ID',
  `status` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态1:正常,0:删除',
  `create_time` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
)ENGINE=Innodb DEFAULT CHARSET=utf8 COMMENT='角色表';

DROP TABLE IF EXISTS `zd_uc_auth`;
CREATE TABLE `zd_uc_auth`(
  `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `pid` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '上级ID,0为顶级',
  `auth_name` VARCHAR(64) NOT NULL DEFAULT '0' COMMENT '权限名称',
  `auth_url` VARCHAR(255) NOT NULL DEFAULT '0' COMMENT 'URL地址',
  `sort` INT(1) UNSIGNED NOT NULL DEFAULT '999' COMMENT '排序,从小到大',
  `icon` VARCHAR(255) NOT NULL DEFAULT '0' COMMENT '图标',
  `is_show` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否显示,0:隐藏,1:显示',
  `user_id` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '操作者ID',
  `create_id` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `status` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '状态,1:正常,0:删除',
  `create_time` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限因子';

DROP TABLE IF EXISTS `zd_uc_admin`;
CREATE TABLE `zd_uc_admin`(
  `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `login_name` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `real_name` VARCHAR(32) NOT NULL DEFAULT '0' COMMENT '真实姓名',
  `password` CHAR(32) NOT NULL DEFAULT '' COMMENT '密码',
  `role_ids` VARCHAR(255) NOT NULL DEFAULT '0' COMMENT '角色id字符串,eg:2,3,4,5',
  `phone` VARCHAR(20) NOT NULL DEFAULT '0' COMMENT '手机号码',
  `email` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `salt` CHAR(10) NOT NULL DEFAULT '' COMMENT '密码加盐',
  `last_login` INT(11) NOT NULL DEFAULT '0' COMMENT '最后登录时间',
  `last_ip` CHAR(15) NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `status` TINYINT(4) NOT NULL DEFAULT '0' COMMENT '状态,1:正常,0:禁用',
  `create_id` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_name` (`login_name`)
)ENGINE=InnoDB DEFAULT CHARSET =utf8mb4 COMMENT ='管理员表';

#--------the user---end------#

#--------the set---start------#

DROP TABLE IF EXISTS `zd_set_group`;
CREATE TABLE `zd_set_group`(
  `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增',
  `group_name` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '组名',
  `detail` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '说明',
  `status` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '状态:1:正常,0:删除',
  `create_id` INT(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `update_id` INT(11) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `create_time` INT(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` INT(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_create_id` (`create_id`)
)ENGINE =InnoDB DEFAULT CHARSET =utf8mb4;

DROP TABLE IF EXISTS `zd_set_evn`;
CREATE TABLE `zd_set_evn`(
  `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `env_name` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '环境名称',
  `env_host` VARCHAR(255) NOT NULL DEFAULT '0' COMMENT '主机',
  `detail` VARCHAR(255) NOT NULL DEFAULT '0' COMMENT '备注',
  `status` TINYINT(4) UNSIGNED NOT NULL DEFAULT '0' COMMENT '状态,1:正常,0:禁用',
  `create_id` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_env_name` (`env_name`)
)ENGINE =InnoDB DEFAULT CHARSET =utf8mb4 COMMENT ='环境分组表';

DROP TABLE IF EXISTS `zd_set_code`;
CREATE TABLE `zd_set_code`(
  `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `code` VARCHAR(50) NOT NULL DEFAULT '0' COMMENT '状态码',
  `descript` VARCHAR(255) NOT NULL DEFAULT '0' COMMENT '描述',
  `detail` VARCHAR(255) NOT NULL DEFAULT '0' COMMENT '备注',
  `status` TINYINT(4) UNSIGNED NOT NULL DEFAULT '0' COMMENT '状态,1:正常,0:禁用',
  `create_id` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE  KEY `idx_env_name` (`code`)
)ENGINE =InnoDB DEFAULT CHARSET =utf8mb4 COMMENT ='状态代码表';

#--------the set---end------#

#--------the api---start------#

DROP TABLE IF EXISTS `zd_api_src`;
CREATE TABLE `zd_api_src`(
  `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `group_id` INT(11) NOT NULL DEFAULT '0' COMMENT '分组ID',
  `src_name` VARCHAR(50) NOT NULL DEFAULT '0' COMMENT '接口名称',
  `status` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态,1:审核通过,0:暂停使用,2:草稿,3:审核中',
  `audit_id` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '审核人ID',
  `create_id` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `create_time` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间',
  `audit_time` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '审核时间',
  PRIMARY KEY (`id`),
  KEY `idx_group_id` (`group_id`)
)ENGINE =InnoDB DEFAULT CHARSET =utf8mb4 COMMENT ='API主表';

DROP TABLE IF EXISTS `zd_api_param`;
CREATE TABLE `zd_api_param`(
  `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `detail_id` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '附表ID',
  `api_key` VARCHAR(100) NOT NULL DEFAULT '0' COMMENT '参数名',
  `api_type` VARCHAR(100) NOT NULL DEFAULT '0' COMMENT '类型',
  `api_value` VARCHAR(500) NOT NULL DEFAULT '0' COMMENT '参数值',
  `api_detail` VARCHAR(500) NOT NULL DEFAULT '0' COMMENT '参数说明',
  `is_null` VARCHAR(10) NOT NULL DEFAULT '0' COMMENT '是否必填',
  `status` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态,1:正常,0:删除',
  `create_id` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_main_id` (`detail_id`)
)ENGINE =InnoDB DEFAULT CHARSET =utf8mb4 COMMENT ='API参数表';

DROP TABLE IF EXISTS `zd_api_detail`;
CREATE TABLE `zd_api_detail`(
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `src_id` INT(11) NOT NULL DEFAULT '0' COMMENT '主表ID',
  `method` TINYINT(1) NOT NULL DEFAULT '1' COMMENT '方法名称,1:GET,2:POST,3:PUT,4:PATCH,5:DELETE',
  `api_name` VARCHAR(100) NOT NULL DEFAULT '0' COMMENT '接口名称',
  `api_url` VARCHAR(100) NOT NULL DEFAULT '0' COMMENT '接口地址',
  `protocol_type` VARCHAR(20) NOT NULL DEFAULT '1' COMMENT '协议类型,1:http,2:https',
  `result` TEXT COMMENT '返回结果,正确或者错误',
  `example` TEXT COMMENT '接口示例',
  `detail` VARCHAR(1000) NOT NULL DEFAULT '0' COMMENT '注意事项',
  `audit_time` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '审核时间',
  `audit_id` INT(11) NOT NULL DEFAULT '0' COMMENT '审核人ID',
  `status` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态,0:暂停 使用,1:正在审核,2:审核通过',
  `create_id` INT(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` INT(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` INT(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` INT(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY  `idx_main_id` (`src_id`)
)ENGINE =InnoDB DEFAULT CHARSET =utf8mb4 COMMENT ='API附表';

#--------the api---end------#
SET FOREIGN_KEY_CHECKS=1;