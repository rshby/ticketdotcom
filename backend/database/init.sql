CREATE DATABASE ticketdotcom;

USE ticketdotcom;

CREATE TABLE `account` (
    `email` varchar(255) UNIQUE PRIMARY KEY NOT NULL,
    `username` varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL
    );

CREATE TABLE `gender` (
                          `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
                          `code` varchar(1) NOT NULL,
    `name` varchar(255) NOT NULL
    );

CREATE TABLE `user` (
                        `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
                        `full_name` varchar(500) DEFAULT null,
    `gender_id` int,
    `phone` varchar(255) DEFAULT null,
    `address_id` int,
    `email_account` varchar(255)
    );

CREATE TABLE `province` (
                            `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
                            `name` varchar(255) NOT NULL
    );

CREATE TABLE `city` (
                        `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
                        `name` varchar(255) NOT NULL,
    `province_id` int
    );

CREATE TABLE `district` (
                            `id` int PRIMARY KEY AUTO_INCREMENT,
                            `name` varchar(255) NOT NULL,
    `city_id` int
    );

CREATE TABLE `sub_district` (
                                `id` int PRIMARY KEY AUTO_INCREMENT,
                                `name` varchar(255) NOT NULL,
    `postal_code` varchar(255) DEFAULT null,
    `district_id` int
    );

CREATE TABLE `address` (
                           `id` int PRIMARY KEY AUTO_INCREMENT,
                           `street` varchar(500) DEFAULT null,
    `sub_district` int
    );

CREATE INDEX `user_index_0` ON `user` (`id`);

CREATE INDEX `user_index_1` ON `user` (`email_account`);

ALTER TABLE `user` ADD FOREIGN KEY (`gender_id`) REFERENCES `gender` (`id`);

ALTER TABLE `user` ADD FOREIGN KEY (`address_id`) REFERENCES `address` (`id`);

ALTER TABLE `user` ADD FOREIGN KEY (`email_account`) REFERENCES `account` (`email`);

ALTER TABLE `city` ADD FOREIGN KEY (`province_id`) REFERENCES `province` (`id`);

ALTER TABLE `district` ADD FOREIGN KEY (`city_id`) REFERENCES `city` (`id`);

ALTER TABLE `sub_district` ADD FOREIGN KEY (`district_id`) REFERENCES `district` (`id`);

ALTER TABLE `address` ADD FOREIGN KEY (`sub_district`) REFERENCES `sub_district` (`id`);
