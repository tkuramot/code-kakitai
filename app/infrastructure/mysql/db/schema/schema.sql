-- Active: 1689853340670@@127.0.0.1@3306@code_kakitai
CREATE TABLE users (
  id VARCHAR(255) PRIMARY KEY NOT NULL,
  email VARCHAR(255) NOT NULL,
  firebaseUid VARCHAR(255) NOT NULL,
  phone_number VARCHAR(255) NOT NULL,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  postal_code VARCHAR(255) NOT NULL COMMENT "郵便番号",
  prefecture VARCHAR(255) NOT NULL COMMENT "都道府県",
  city VARCHAR(255) NOT NULL COMMENT "市区町村",
  address_extra VARCHAR(255) NOT NULL COMMENT "住所詳細",
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE owners (
  id VARCHAR(255) PRIMARY KEY NOT NULL,
  email VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE products (
  id VARCHAR(255) PRIMARY KEY NOT NULL,
  owner_id VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  description VARCHAR(255) NOT NULL,
  price BIGINT NOT NULL,
  stock INT NOT NULL,
  CONSTRAINT foreign_key_products_owner_id FOREIGN KEY (`owner_id`) REFERENCES `owners` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;