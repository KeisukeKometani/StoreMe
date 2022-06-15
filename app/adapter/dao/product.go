package dao

import (
  "example/adapter/repo"
)

type ProductDAO struct {
  ac Accessor
}

func (c *ProductDAO) Save(product repo.Product) (err error) {
  c.ac.Save(&product)
  return c.ac.LastError()
}

func (c *ProductDAO) Find(product repo.Product) (err error) {
  c.ac.Find(&product)
  return c.ac.LastError()
}

func NewProductDAO(ac *Accessor) ProductDAO {
  return &ProductDAO{*ac}
}
