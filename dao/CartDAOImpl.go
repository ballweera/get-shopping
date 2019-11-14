package dao

import "github.com/ballweera/get-shopping/model"

// CartDAOImpl is implementation of CartDAO
type CartDAOImpl struct{}

// GetCartByCartID return Cart
func (c *CartDAOImpl) GetCartByCartID(cartID string) model.Cart {
	return model.Cart{}
}

// Validate validates input
func (c *CartDAOImpl) Validate(cartID string) error {
	return nil
}

// Update a cart
func (c *CartDAOImpl) Update(cart model.Cart) {

}
