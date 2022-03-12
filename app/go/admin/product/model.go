package product

import (
	"gorm.io/gorm"
)

//make type product struct
type Product struct {
	gorm.Model
	Name	  string  `json:"name" `
	Price	  int  `json:"price" `
	Description string  `json:"description" `
	Image	  string  `json:"image" `
}

// CREATE TABLE products add created_at, updated_at, deleted_at
// CREATE TABLE products (
// 	id INT NOT NULL AUTO_INCREMENT,
// 	name VARCHAR(255) NOT NULL,
// 	price INT NOT NULL,
// 	description VARCHAR(255) NOT NULL,
// 	image VARCHAR(255) NOT NULL,
// 	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 	deleted_at TIMESTAMP NULL DEFAULT NULL,
// 	PRIMARY KEY (id)
// );