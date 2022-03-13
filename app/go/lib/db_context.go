package lib

import (
	"errors"
	"sync"
	"gorm.io/gorm"
	"example/online_shop/go/database"
)

var singleton *DBContext = &DBContext{}
var once sync.Once
var connecter database.DBConnecter

// DBContext はアプリケーション全体で共通のDBコネクションを保持します
// 動作はいわゆるシングルトンパターンです。一応スレッドセーフな(はず)
// の実装です。
type DBContext struct {
	Singleton
}

func (d *DBContext) SetDBInstance(i *gorm.DB) {
	d.SetInstance(once, i)
}

func (d *DBContext) GetDBInstance() (*gorm.DB, error) {
	i := d.GetInstance()
	if i == nil {
		return nil, errors.New("Not a DB Instance")
	}
	if db, ok := i.(*gorm.DB); ok {
		return db, nil
	}
	return nil, errors.New("Not a DB Instance")
}

// SetDBConnecter はDBコンテキストシングルトンにコネクションを取得
// するコネクタを指定します。
func SetDBConnecter(p database.DBConnecter) {
	connecter = p
}

// GetDBInstance はDB(コネクション)インスタンスを返します。
// 通常コネクションはアプリケーション内で同一インスタンスです
func GetDBInstance() *gorm.DB {
	db, err := singleton.GetDBInstance()
	if err != nil {
		singleton.SetDBInstance(connecter.ConnectDatabase())
		db, _ := singleton.GetDBInstance()
		return db
	}
	return db
}
