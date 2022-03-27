package model

type Product struct {
  ID          uint
	Name        string
	Price       int
	Description string
	Image       string
}
type Products []Product
