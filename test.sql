/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 50735
 Source Host           : localhost:3306
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 50735
 File Encoding         : 65001

 Date: 27/10/2021 16:28:58
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '管理员id',
  `nickname` varchar(32) NOT NULL DEFAULT '' COMMENT '姓名',
  `username` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
  `telephone` varchar(20) NOT NULL DEFAULT '' COMMENT '联系方式',
  `pwd` varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
  `salt` varchar(32) NOT NULL DEFAULT '' COMMENT '密码加盐',
  `last_ip` varchar(64) NOT NULL DEFAULT '' COMMENT '最后登录id',
  `created_time` bigint(20) NOT NULL DEFAULT '0',
  `updated_time` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `phone` (`telephone`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of admin
-- ----------------------------
BEGIN;
INSERT INTO `admin` VALUES (1, 'admin', 'admin', '', 'b81b20da3fdeba192fa9ba37a447bd8a', 'qRikMWBM4RRmypyq', '127.0.0.1', 1624501078, 1624504494);
COMMIT;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ptype` varchar(255) DEFAULT NULL,
  `v0` varchar(255) DEFAULT NULL,
  `v1` varchar(255) DEFAULT NULL,
  `v2` varchar(255) DEFAULT NULL,
  `v3` varchar(255) DEFAULT NULL,
  `v4` varchar(255) DEFAULT NULL,
  `v5` varchar(255) DEFAULT NULL,
  `v6` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` VALUES (5, 'g', 'super', 'admin', '', 'false', '管理员', NULL, NULL);
INSERT INTO `casbin_rule` VALUES (6, 'p', 'admin', '/api/admin/user*', 'GET|POST|PUT|DELETE', '.*', '管理员', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for file
-- ----------------------------
DROP TABLE IF EXISTS `file`;
CREATE TABLE `file` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '文件ID',
  `name` varchar(200) NOT NULL DEFAULT '' COMMENT '原始文件名',
  `savename` char(100) NOT NULL DEFAULT '' COMMENT '保存名称',
  `savepath` varchar(100) NOT NULL DEFAULT '' COMMENT '文件保存路径',
  `savepathp` varchar(100) NOT NULL DEFAULT '' COMMENT '转码后路径',
  `ext` char(10) NOT NULL DEFAULT '' COMMENT '文件后缀',
  `mime` char(200) NOT NULL DEFAULT '' COMMENT '文件mime类型',
  `size` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '文件大小 单位 B',
  `width` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '尺寸 宽度 图片/视频',
  `height` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '尺寸 高度 图片/视频',
  `duration` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '时长 音频/视频',
  `md5` char(32) NOT NULL DEFAULT '' COMMENT '文件md5',
  `category` char(32) NOT NULL DEFAULT '' COMMENT '所属分类',
  `location` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '文件保存位置 0 阿里云',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '远程地址',
  `created_time` bigint(20) NOT NULL DEFAULT '0',
  `updated_time` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of file
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
