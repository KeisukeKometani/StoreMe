package product

import (
	"database/sql"
	"log"
	"time"
)

// Read the record.
func readRecord(db *sql.DB, p *Product, id string) {
	// Read existing records.
	row := db.QueryRow("SELECT * FROM product WHERE id = ?", id)
	if err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.Image, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt); err != nil {
		log.Fatal(err)
	}
}

// Read all records.
func readAllRecords(db *sql.DB, products *[]Product) {
	// Read existing records.
	rows, err := db.Query("SELECT * FROM product")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate through the result set.
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.Image, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt); err != nil {
			log.Fatal(err)
		}
		*products = append(*products, p)
	}
}

// Create the record.
func createRecord(db *sql.DB, p *Product) {
	// Insert the record.
	stmt, err := db.Prepare("INSERT INTO product (name, price, description, image) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	_ , err = stmt.Exec(p.Name, p.Price, p.Description, p.Image)
	if err != nil {
		log.Fatal(err)
	}
}

// Update the record.
func updateRecord(db *sql.DB, p *Product) {
	// Update the record.
	stmt, err := db.Prepare("UPDATE product SET name = ?, price = ?, description = ?, image = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	_ , err = stmt.Exec(p.Name, p.Price, p.Description, p.Image, p.ID)
	if err != nil {
		log.Fatal(err)
	}
}

// Delete the record.
func deleteRecord(db *sql.DB, p *Product) {
	// Update the DeletedAt column. time now.
	stmt, err := db.Prepare("UPDATE product SET deleted_at = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	
	_ , err = stmt.Exec(time.Now(), p.ID)
	if err != nil {
		log.Fatal(err)
	}

	p.DeletedAt = &[]time.Time{time.Now()}[0]
}