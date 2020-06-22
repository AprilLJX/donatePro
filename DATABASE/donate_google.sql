/*
 Navicat Premium Data Transfer

 Source Server         : jacyn
 Source Server Type    : MySQL
 Source Server Version : 80020
 Source Host           : localhost:3306
 Source Schema         : donate_google

 Target Server Type    : MySQL
 Target Server Version : 80020
 File Encoding         : 65001

 Date: 22/06/2020 21:45:56
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for administrator
-- ----------------------------
DROP TABLE IF EXISTS `administrator`;
CREATE TABLE `administrator` (
  `id` int NOT NULL,
  `account` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`id`,`account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of administrator
-- ----------------------------
BEGIN;
INSERT INTO `administrator` VALUES (1, 'admin', 'admin');
COMMIT;

-- ----------------------------
-- Table structure for demand_list
-- ----------------------------
DROP TABLE IF EXISTS `demand_list`;
CREATE TABLE `demand_list` (
  `demand_id` int NOT NULL AUTO_INCREMENT,
  `recipient_id` int NOT NULL COMMENT '发起单位',
  `pro_name` varchar(255) NOT NULL,
  `introduction` varchar(255) NOT NULL,
  `category` tinyint NOT NULL COMMENT '医疗/生活/教育',
  `materials` longtext NOT NULL,
  `if_standard` tinyint(1) NOT NULL DEFAULT '0',
  `if_audit` tinyint(1) NOT NULL DEFAULT '0',
  `rec_address` varchar(255) NOT NULL,
  `cut_off_time` datetime(6) DEFAULT NULL,
  `emergency_degree` tinyint(1) NOT NULL DEFAULT '0',
  `initiation_time` datetime(6) NOT NULL ON UPDATE CURRENT_TIMESTAMP(6),
  PRIMARY KEY (`demand_id`) USING BTREE,
  KEY `发起单位` (`recipient_id`),
  CONSTRAINT `发起单位` FOREIGN KEY (`recipient_id`) REFERENCES `recipient` (`recipient_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of demand_list
-- ----------------------------
BEGIN;
INSERT INTO `demand_list` VALUES (1, 1, '急需口罩志愿', '医院a现急需口罩等医疗物资', 1, '口罩：200；手套：600', 1, 1, '北京', '2020-06-30 21:05:49.000000', 0, '2020-06-22 21:06:00.431798');
INSERT INTO `demand_list` VALUES (2, 2, '为老人献爱心', '养老院a现需要电热毯，热水袋若干，为老人的冬天带来温暖', 2, '电热毯：50；热水袋：50；保暖衣：100', 1, 1, '广州', '2020-06-30 21:09:14.000000', 0, '2020-06-22 21:09:21.000000');
INSERT INTO `demand_list` VALUES (3, 3, '孩子们的是艺术道路', '给自闭症儿童一只展现自己的画笔', 2, '水彩笔：30套；蜡笔：30套；画本：100本', 1, 1, '长沙', '2020-06-30 21:11:17.000000', 0, '2020-06-22 21:11:21.000000');
COMMIT;

-- ----------------------------
-- Table structure for dona_project
-- ----------------------------
DROP TABLE IF EXISTS `dona_project`;
CREATE TABLE `dona_project` (
  `pro_id` int NOT NULL,
  `demand_id` int NOT NULL,
  `rec_donation_num` int NOT NULL DEFAULT '0',
  `if_end` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`pro_id`),
  KEY `需求单` (`demand_id`),
  CONSTRAINT `需求单` FOREIGN KEY (`demand_id`) REFERENCES `demand_list` (`demand_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of dona_project
-- ----------------------------
BEGIN;
INSERT INTO `dona_project` VALUES (1, 1, 1, 0);
INSERT INTO `dona_project` VALUES (2, 2, 2, 0);
INSERT INTO `dona_project` VALUES (3, 3, 1, 0);
COMMIT;

-- ----------------------------
-- Table structure for donor
-- ----------------------------
DROP TABLE IF EXISTS `donor`;
CREATE TABLE `donor` (
  `donor_id` int NOT NULL AUTO_INCREMENT,
  `account` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `password` varchar(255) NOT NULL,
  `nickname` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `id_number` varchar(255) NOT NULL,
  `cur_residence` varchar(255) DEFAULT NULL,
  `city` varchar(255) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像，后期再改格式',
  `love_value` char(128) NOT NULL DEFAULT '100',
  `profile` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`donor_id`,`account`) USING BTREE,
  KEY `donor_id` (`donor_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of donor
-- ----------------------------
BEGIN;
INSERT INTO `donor` VALUES (1, '123456', '123456', 'alice', 'alice', '123456', '北京', '北京', NULL, '100', NULL);
INSERT INTO `donor` VALUES (2, '123457', '111', 'bob', 'bob', '123457', '湖南长沙', '长沙', NULL, '100', NULL);
INSERT INTO `donor` VALUES (3, '123458', '111', 'cindy', 'cindy', '123458', NULL, NULL, NULL, '100', NULL);
COMMIT;

-- ----------------------------
-- Table structure for logistics
-- ----------------------------
DROP TABLE IF EXISTS `logistics`;
CREATE TABLE `logistics` (
  `logistics_id` int NOT NULL AUTO_INCREMENT,
  `logi_com` varchar(255) NOT NULL,
  `logi_num` char(32) NOT NULL,
  `logi_time` datetime(6) NOT NULL,
  `if_anonymous` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`logistics_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of logistics
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for medical_supplies
-- ----------------------------
DROP TABLE IF EXISTS `medical_supplies`;
CREATE TABLE `medical_supplies` (
  `id` int NOT NULL,
  `category` varchar(255) NOT NULL,
  `standard` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of medical_supplies
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for no_target_donation
-- ----------------------------
DROP TABLE IF EXISTS `no_target_donation`;
CREATE TABLE `no_target_donation` (
  `notarget_id` int NOT NULL AUTO_INCREMENT,
  `donor_id` int NOT NULL,
  `category` tinyint NOT NULL,
  `donate_materials` longtext NOT NULL,
  `if_standard` tinyint(1) NOT NULL DEFAULT '0',
  `if_audit` tinyint(1) NOT NULL DEFAULT '0',
  `donate_time` datetime(6) NOT NULL ON UPDATE CURRENT_TIMESTAMP(6),
  `match_pro` int DEFAULT NULL,
  `if_anonymous` tinyint(1) NOT NULL DEFAULT '0',
  `message` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`notarget_id`) USING BTREE,
  KEY `捐赠方id` (`donor_id`),
  CONSTRAINT `捐赠方id` FOREIGN KEY (`donor_id`) REFERENCES `donor` (`donor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of no_target_donation
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for project_donation
-- ----------------------------
DROP TABLE IF EXISTS `project_donation`;
CREATE TABLE `project_donation` (
  `id` int NOT NULL,
  `project_id` int DEFAULT NULL,
  `donation_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `project_id` (`project_id`),
  KEY `donation_id` (`donation_id`),
  CONSTRAINT `donation_id` FOREIGN KEY (`donation_id`) REFERENCES `target_donation` (`target_id`),
  CONSTRAINT `project_id` FOREIGN KEY (`project_id`) REFERENCES `dona_project` (`pro_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of project_donation
-- ----------------------------
BEGIN;
INSERT INTO `project_donation` VALUES (1, 1, 1);
INSERT INTO `project_donation` VALUES (2, 2, 2);
INSERT INTO `project_donation` VALUES (3, 3, 3);
INSERT INTO `project_donation` VALUES (4, 2, 4);
COMMIT;

-- ----------------------------
-- Table structure for recipient
-- ----------------------------
DROP TABLE IF EXISTS `recipient`;
CREATE TABLE `recipient` (
  `recipient_id` int NOT NULL AUTO_INCREMENT,
  `account` char(32) NOT NULL,
  `password` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `id_number` varchar(256) NOT NULL,
  `company` varchar(255) NOT NULL,
  `com_category` varchar(255) NOT NULL,
  `credit_code` varchar(255) NOT NULL COMMENT '统一信用代码',
  `com_address` varchar(255) NOT NULL,
  `com_profile` varchar(255) DEFAULT NULL,
  `recipient_num` int DEFAULT '0',
  PRIMARY KEY (`recipient_id`,`account`) USING BTREE,
  KEY `recipient_id` (`recipient_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of recipient
-- ----------------------------
BEGIN;
INSERT INTO `recipient` VALUES (1, '123456', '123', '医院a', '123456', '医院a', '医院', '1', '北京', '医院', 0);
INSERT INTO `recipient` VALUES (2, '123457', '123', '养老院a', '123457', '养老院a', '养老院', '2', '广州', NULL, 0);
INSERT INTO `recipient` VALUES (3, '123458', '123', '福利院a', '123458', '福利院a', '福利院', '3', '长沙', NULL, 0);
COMMIT;

-- ----------------------------
-- Table structure for target_donation
-- ----------------------------
DROP TABLE IF EXISTS `target_donation`;
CREATE TABLE `target_donation` (
  `target_id` int NOT NULL AUTO_INCREMENT,
  `donor_id` int NOT NULL,
  `category` tinyint NOT NULL,
  `donate_materials` longtext NOT NULL,
  `if_standard` tinyint(1) NOT NULL DEFAULT '0',
  `if_audit` tinyint(1) NOT NULL DEFAULT '0',
  `donate_time` datetime(6) NOT NULL ON UPDATE CURRENT_TIMESTAMP(6),
  `match_pro` int DEFAULT NULL,
  `if_anonymous` tinyint(1) NOT NULL DEFAULT '0',
  `message` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`target_id`) USING BTREE,
  KEY `捐赠方id` (`donor_id`),
  CONSTRAINT `target_donation_ibfk_1` FOREIGN KEY (`donor_id`) REFERENCES `donor` (`donor_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of target_donation
-- ----------------------------
BEGIN;
INSERT INTO `target_donation` VALUES (1, 1, 1, '口罩：200；手套：600', 1, 1, '2020-06-22 21:25:00.290653', 1, 0, '加油');
INSERT INTO `target_donation` VALUES (2, 2, 2, '电热毯：20', 1, 1, '2020-06-22 21:27:23.863418', 2, 0, NULL);
INSERT INTO `target_donation` VALUES (3, 2, 3, '蜡笔：10', 1, 1, '2020-06-22 21:26:39.000000', 3, 0, NULL);
INSERT INTO `target_donation` VALUES (4, 3, 2, '电热毯：10', 1, 1, '2020-06-22 21:27:14.000000', 2, 0, NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
