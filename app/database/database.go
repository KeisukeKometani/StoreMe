package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

// Confirm ping database.
func confirmPingDatabase(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database.")
}

// Open a database connection.
func ConnectDatabase() *sql.DB{
	// Capture connection properties.
	jst, _ := time.LoadLocation("Asia/Tokyo")
	cfg := mysql.Config{
		User: "root",
		Passwd: "1234",
		Net: "tcp",
		Addr: "online_shop_db:3306", // Specify the container name
		DBName: "online_shop",
		ParseTime: true,
		Loc: jst,
	}
	
	// Get a database handle.
	
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	confirmPingDatabase(db)
	return db
}