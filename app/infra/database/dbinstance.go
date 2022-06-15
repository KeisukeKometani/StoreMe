package database

import (
  "errors"
  "gorm.io/gorm"
  "sync"
  "example/online_shop/pkg/singleton"
)

var singleton *dbContext = &dbContext{}
var once sync.Once
var connecter DBConnecter

type DBConnecter interface {
  Connect() *gorm.DB
}

type dbContext struct {
  singleton.Singleton
}

func (d *dbContext) SetDBInstance(i *gorm.DB) {
  d.SetInstance(once, i)
}

func (d *dbContext) GetDBInstance() (*gorm.DB, error) {
  i := d.GetInstance()
  if i == nil {
    return nil, errors.New("Not a DB Instance")
  }

  if db, ok := i.(*gorm.DB); ok {
    return db, nil
  }
  return nil, errors.New("Not a DB instance")
}

func SetDBConnecter(p DBConnecter) {
  connecter = p
}

func getDBInstance() (db *gorm.DB, err error) {
  db, err := singleton.GetDBInstance()
  if err != nil {
    singleton.SetDBInstance(connecter.Connect())
    db, err := singleton.GetDBInstance()
    return
  }
  return
}
