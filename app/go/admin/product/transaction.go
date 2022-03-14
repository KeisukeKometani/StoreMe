package product

import (
	"gorm.io/gorm"
)


//subtract inventory 出庫
func subtractInventory(db *gorm.DB, product *Product, quantity int) {
	db.Model(&product).Update("inventory", product.Inventory - quantity)
}

//add inventory 入庫
func addInventory(db *gorm.DB, product *Product, quantity int) {
	db.Model(&product).Update("inventory", product.Inventory + quantity)
}

//get inventory 在庫取得
func getInventory(db *gorm.DB, product *Product) {
	db.Model(&product).Select("inventory").Scan(&product)
}

//record transaction 入出庫記録
func recordTransaction(db *gorm.DB, tp *TransactionProducts, quantity int, transactionType string) {
	db.Create(&Transaction{
		TransactionType: transactionType,
		Quantity: quantity,
	})
	db.Model(&tp).Update("quantity", tp.Quantity + quantity)
}

//get transaction within a period 期間内の入出庫記録取得
func getTransaction(db *gorm.DB, tp *TransactionProducts, startDate string, endDate string) {
	db.Where("created_at >= ? AND created_at <= ?", startDate, endDate).Find(&tp)
}