CREATE DATABASE `ecom` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
-- ecom.carts definition

CREATE TABLE `carts` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `price` bigint DEFAULT NULL,
  `total` bigint DEFAULT NULL,
  `date` datetime(3) DEFAULT NULL,
  `user_id` bigint DEFAULT NULL,
  `product_id` bigint DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_products_carts` (`product_id`),
  KEY `fk_users_carts` (`user_id`),
  CONSTRAINT `fk_products_carts` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `fk_users_carts` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ecom.categories definition

CREATE TABLE `categories` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `description` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `user_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_users_categories` (`user_id`),
  CONSTRAINT `fk_users_categories` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ecom.confrimations definition

CREATE TABLE `confrimations` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `transfer_amount` bigint DEFAULT NULL,
  `image_url` longtext,
  `description` longtext,
  `date` datetime(3) DEFAULT NULL,
  `order_id` bigint DEFAULT NULL,
  `user_id` bigint DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_orders_confrimations` (`order_id`),
  KEY `fk_users_confrimations` (`user_id`),
  CONSTRAINT `fk_orders_confrimations` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `fk_users_confrimations` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ecom.order_details definition

CREATE TABLE `order_details` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `harga` bigint DEFAULT NULL,
  `jumlah` bigint DEFAULT NULL,
  `product_id` bigint DEFAULT NULL,
  `order_id` bigint DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_orders_order_details` (`order_id`),
  KEY `fk_products_order_details` (`product_id`),
  CONSTRAINT `fk_orders_order_details` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`),
  CONSTRAINT `fk_products_order_details` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ecom.orders definition

CREATE TABLE `orders` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `recipients_name` longtext,
  `order_date` datetime(3) DEFAULT NULL,
  `city` longtext,
  `address` longtext,
  `status_payment` enum('paid','unpaid') DEFAULT NULL,
  `phone_number` longtext,
  `user_id` bigint DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_users_orders` (`user_id`),
  CONSTRAINT `fk_users_orders` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ecom.products definition

CREATE TABLE `products` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `condition` longtext,
  `description` longtext,
  `image_url` longtext,
  `stock` bigint DEFAULT NULL,
  `price` bigint DEFAULT NULL,
  `heavy` longtext,
  `user_id` bigint DEFAULT NULL,
  `category_id` bigint DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_users_products` (`user_id`),
  KEY `fk_categories_products` (`category_id`),
  CONSTRAINT `fk_categories_products` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `fk_users_products` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ecom.ratings definition

CREATE TABLE `ratings` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `star` bigint DEFAULT NULL,
  `description` longtext,
  `user_id` bigint DEFAULT NULL,
  `product_id` bigint DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_products_rating` (`product_id`),
  KEY `fk_users_rating` (`user_id`),
  CONSTRAINT `fk_products_rating` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `fk_users_rating` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ecom.ratings definition

CREATE TABLE `ratings` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `star` bigint DEFAULT NULL,
  `description` longtext,
  `user_id` bigint DEFAULT NULL,
  `product_id` bigint DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_products_rating` (`product_id`),
  KEY `fk_users_rating` (`user_id`),
  CONSTRAINT `fk_products_rating` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `fk_users_rating` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
-- ecom.reviews definition

CREATE TABLE `reviews` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `star` bigint DEFAULT NULL,
  `description` longtext,
  `user_id` bigint DEFAULT NULL,
  `product_id` bigint DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_products_review` (`product_id`),
  KEY `fk_users_review` (`user_id`),
  CONSTRAINT `fk_products_review` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `fk_users_review` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ecom.users definition

CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(191) NOT NULL,
  `email` varchar(191) NOT NULL,
  `name` longtext NOT NULL,
  `password` longtext NOT NULL,
  `date_of_birth` longtext NOT NULL,
  `gender` enum('man','woman') DEFAULT NULL,
  `phone_number` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `role` enum('customer','admin') DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;