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