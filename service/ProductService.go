package service

import "github.com/ballweera/get-shopping/model"

// ProductService is interface
type ProductService interface {
	GetAllProducts() []model.Product
	GetProductByID(productID string) model.Product
	DeleteProduct(productID string)
	AddProduct(p model.Product)
	EditProduct(p model.Product)
}
