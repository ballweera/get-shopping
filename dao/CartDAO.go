package dao

import "github.com/ballweera/get-shopping/model"

// CartDAO is dao of Cart
type CartDAO interface {
	GetCartByCartID(cartID string) model.Cart
	Validate(cartID string) error
	Update(cart model.Cart)
}
