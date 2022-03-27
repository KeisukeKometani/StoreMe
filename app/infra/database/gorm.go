package database

import (
  "errors"
  "gorm.io/gorm"
  "sync"
  "example/online_shop/pkg/singleton"
)


type GormAccess struct {
  Error error
}

func (c *GormAccess) withErrHandle(fn func(*gorm.DB) error) {
  db, err := getDBInstance()
  if err != nil {
    c.Error = err
    return
  }
  if err := fn(db); err != nil {
    c.Error = err
  }
  return
}

func (c *GormAccess) Create(value interface{}) (c *GormAccess){
  c.withErrHandle(func (db *gorm.DB){
    return db.Create(&value).Error
  })
  return
}

func (c *GormAccess) Save(value interface{}) (c *GormAccess){
  c.withErrHandle(func (db *gorm.DB){
    return db.Save(&value).Error
  })
  return
}

func (c *GormAccess) First(value interface{}, conds ...interface{}) (c *GormAccess){
  c.withErrHandle(func (db *gorm.DB){
    return db.First(&value, conds...).Error
  })
  return
}

func (c *GormAccess) LastError() error {
  return c.Error
}

func NewGormAccess() GormAccess{
  return &GormAccess{}
}
