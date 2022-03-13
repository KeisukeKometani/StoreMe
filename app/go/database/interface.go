package database

import (
	"gorm.io/gorm"
)

type DBConnecter interface {
	ConnectDatabase() *gorm.DB
}
