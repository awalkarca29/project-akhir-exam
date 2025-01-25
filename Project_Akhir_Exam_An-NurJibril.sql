create database project_akhir_exam;
drop database project_akhir_exam;
use project_akhir_exam;

INSERT INTO roles (name, created_at, updated_at) 
VALUES ("Admin", NOW(), NOW()), ("User", NOW(), NOW());

INSERT INTO events (name, description, slug, price, stock, created_at, updated_at) 
VALUES 
	("Konser", "Tiket Konser", "konser", "100000", 300, NOW(), NOW()), 
	("Pesawat", "Tiket Pesawat", "pesawat", "500000", 250, NOW(), NOW()),
	("Kereta", "Tiket Kereta", "kereta", "150000", 500, NOW(), NOW());

CREATE TABLE `roles` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `users` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `role_id` integer,
  `name` varchar(255),
  `email` varchar(255),
  `password` varchar(255),
  `photo` varchar(255),
  `phone` varchar(255),
  `address` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `events` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(255),
  `description` varchar(255),
  `slug` varchar(255),
  `image` varchar(255),
  `price` integer,
  `stock` integer,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `tickets` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` integer,
  `event_id` integer,
  `quantity` integer,
  `total` integer,
  `payment_method` varchar(255),
  `status` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);

ALTER TABLE `users` ADD FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);

ALTER TABLE `tickets` ADD FOREIGN KEY (`event_id`) REFERENCES `events` (`id`);

ALTER TABLE `tickets` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);