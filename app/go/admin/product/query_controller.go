package product

import (
	"log"

	"gorm.io/gorm"
)

// Read the record.
func readRecord(db *gorm.DB, p *Product, id string) {
	// Read existing records.
	db.First(&p, id)
}

// Read all records.
func readAllRecords(db *gorm.DB, products *[]Product) {
	// Read existing records.
	db.Find(&products)
}

// Create the record.
func createRecord(db *gorm.DB, p *Product) {
	// Insert the record.
	result := db.Create(&p)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}

// Update the record.
func updateRecord(db *gorm.DB, p *Product, update_product Product) {
	// Update the record.
	db.First(&p)
	db.Model(&p).Updates(update_product)
}

// Delete the record.
func deleteRecord(db *gorm.DB, p *Product) {
	// Update the DeletedAt column. time now.
	db.Delete(&p)
}