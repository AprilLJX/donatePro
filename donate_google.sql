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

 Date: 05/06/2020 11:32:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for acount
-- ----------------------------
DROP TABLE IF EXISTS `acount`;
CREATE TABLE `acount` (
  `id` int NOT NULL,
  `account` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`id`,`account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of acount
-- ----------------------------
BEGIN;
INSERT INTO `acount` VALUES (1, 'admin', 'admin');
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of demand_list
-- ----------------------------
BEGIN;
INSERT INTO `demand_list` VALUES (1, 1, 'test', '测试', 1, '口罩：2', 0, 0, '北京', '2020-06-08 10:55:57.000000', 0, '2020-06-05 10:56:16.000000');
COMMIT;

-- ----------------------------
-- Table structure for dona_project
-- ----------------------------
DROP TABLE IF EXISTS `dona_project`;
CREATE TABLE `dona_project` (
  `pro_id` int NOT NULL,
  `demand_id` int NOT NULL,
  `donation_id` int DEFAULT NULL,
  `rec_donation_num` int NOT NULL DEFAULT '0',
  `if_end` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`pro_id`),
  KEY `需求单` (`demand_id`),
  KEY `捐赠单` (`donation_id`),
  CONSTRAINT `捐赠单` FOREIGN KEY (`donation_id`) REFERENCES `target_donation` (`target_id`),
  CONSTRAINT `需求单` FOREIGN KEY (`demand_id`) REFERENCES `demand_list` (`demand_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of dona_project
-- ----------------------------
BEGIN;
INSERT INTO `dona_project` VALUES (1, 1, NULL, 0, 0);
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of donor
-- ----------------------------
BEGIN;
INSERT INTO `donor` VALUES (1, '123456', '123456', 'test', 'test', '123456', '北京', '北京', NULL, '100', NULL);
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of recipient
-- ----------------------------
BEGIN;
INSERT INTO `recipient` VALUES (1, '1234567', '123', 'test', '123', '一个医院', '医院', '1', '北京', '医院', 0);
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of target_donation
-- ----------------------------
BEGIN;
INSERT INTO `target_donation` VALUES (1, 1, 1, '口窄：2', 0, 0, '2020-06-05 10:57:40.000000', 1, 0, NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
