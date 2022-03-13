package transaction

import (
	"example/online_shop/go/admin/product"
	"time"

	"gorm.io/gorm"
)

//make type product struct
type TransactionProduct struct {
	gorm.Model
	ProductId     int  						`json:"-"`
	TransactionAt time.Time  			`json:"price" `
	RecordAt      time.Time  			`json:"description" `
	Product       product.Product `json:"product"`
}
