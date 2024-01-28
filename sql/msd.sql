/*
SQLyog Professional v12.09 (64 bit)
MySQL - 8.0.30 : Database - msd
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`msd` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

/*Table structure for table `svc_config` */

CREATE TABLE `svc_config` (
  `id` int NOT NULL AUTO_INCREMENT,
  `svc_name` varchar(20) NOT NULL COMMENT '微服务名字',
  `desc` varchar(200) DEFAULT NULL COMMENT '描述',
  `env` int NOT NULL COMMENT '环境',
  `conf` varchar(5000) DEFAULT NULL COMMENT '配置',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `svc_config` */

insert  into `svc_config`(`id`,`svc_name`,`desc`,`env`,`conf`) values (1,'gorm-demo','gorm-demo开发环境的配置',0,NULL);

/*Table structure for table `user` */

CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `age` int DEFAULT NULL,
  `sex` enum('male','female','other') DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `user` */

insert  into `user`(`id`,`name`,`age`,`sex`) values (7,'tamaki kouji',61,'male');
insert  into `user`(`id`,`name`,`age`,`sex`) values (8,'jkjj',22,'female');
insert  into `user`(`id`,`name`,`age`,`sex`) values (9,'wang2',12,'other');
insert  into `user`(`id`,`name`,`age`,`sex`) values (10,'ming',21,'other');
insert  into `user`(`id`,`name`,`age`,`sex`) values (11,'dk',21,'male');
insert  into `user`(`id`,`name`,`age`,`sex`) values (12,'bgk',21,'female');
insert  into `user`(`id`,`name`,`age`,`sex`) values (13,'tokugawa',61,'male');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
