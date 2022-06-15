package repository

import "example/domain/model"

type Product interface {
  Save(product model.Product) error
  Find(productId uint) (product model.Product, error)
}
