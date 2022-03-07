package database

import (
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

// Open a database connection.
func ConnectDatabase() *gorm.DB{
	// Capture connection properties.
	dsn := "root:1234@tcp(online_shop_db:3306)/online_shop?charset=utf8&parseTime=True&loc=Local"
	// Get a database handle.
	sqlDB, err := sql.Open("mysql", dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}