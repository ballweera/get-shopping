package service

import (
	"github.com/ballweera/get-shopping/model"
)

// CartServiceImpl implements CartService
type CartServiceImpl struct {
}

// GetCartByCartID return cart
func (c *CartServiceImpl) GetCartByCartID(cartID string) model.Cart {
	return model.Cart{}
}
