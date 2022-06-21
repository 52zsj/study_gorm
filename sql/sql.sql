/*
 Navicat Premium Data Transfer

 Source Server         : http://www.52zsj.com
 Source Server Type    : MySQL
 Source Server Version : 50737
 Source Host           : localhost:3306
 Source Schema         : study_gorm

 Target Server Type    : MySQL
 Target Server Version : 50737
 File Encoding         : 65001

 Date: 20/06/2022 14:28:49
*/

SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT,
    `name`       varchar(32) DEFAULT NULL COMMENT '角色名称',
    `created_at` datetime    DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`       int(11) NOT NULL AUTO_INCREMENT,
    `username` varchar(255) DEFAULT NULL,
    `nickname` varchar(255) DEFAULT NULL,
    `password` varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for user_other
-- ----------------------------
DROP TABLE IF EXISTS `user_other`;
CREATE TABLE `user_other`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT,
    `user_id`    int(11) DEFAULT NULL,
    `other_info` varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY          `q` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role`
(
    `id`      int(11) NOT NULL AUTO_INCREMENT,
    `role_id` int(11) DEFAULT NULL COMMENT 'role表主键',
    `user_id` int(11) DEFAULT NULL COMMENT 'user表主键',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='用户角色表 1对多关系';
