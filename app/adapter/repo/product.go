package repo

import (
  "example/online_shop/domain/model"
  "example/online_shop/adapter/dao"
)

type ProductDAOer interface {
  Save(Product) error
  Find(Product) error
}

type Product struct {
  BaseModel
	Name        string  `json:"name"`
	Price       int     `json:"price"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
}

func mapProductEntity(product model.Product) (ret Product){
  ret := new(product)
  ret.ID          = product.ID
  ret.Name        = product.Name
  ret.Price       = product.Price
  ret.Description = product.Description
  ret.Image       = product.Image
  return
}

type ProductRepository struct {
  dao *ProductDAOer
}

func NewProductRepository(dao *ProductDAOer) {
  return &ProductDAOer{dao}
}

func (repo *ProductRepository) Save(product model.Product) error {
  p := mapProductEntity(product)
  err := repo.dao.Save(&p)
  if err != nil {
    return errors.New("Save Error Product")
  }
  return nil
}

func (repo *ProductRepository) Find(productId uint) (model.Product, error) {

}

// Save(product model.Product) error
// Find(productId uint) (product model.Product, error)
