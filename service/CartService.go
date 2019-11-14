package service

import (
	"github.com/ballweera/get-shopping/model"
)

// CartService is service interface
type CartService interface {
	GetCartByCartID(cartID string) model.Cart
}
