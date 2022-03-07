package product

import (
	"time"
	"gorm.io/gorm"
)

//make type product struct
type Product struct {
	gorm.Model
	ID		  int  `json:"id" `
	Name	  string  `json:"name" `
	Price	  int  `json:"price" `
	Description string  `json:"description" `
	Image	  string  `json:"image" `
	CreatedAt time.Time  `json:"created_at" `
	UpdatedAt time.Time  `json:"updated_at" `
	DeletedAt *time.Time  `json:"deleted_at" `
}