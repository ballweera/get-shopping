package service

import "github.com/ballweera/get-shopping/model"

// CartItemServiceImpl is implementation of CartItemService
type CartItemServiceImpl struct{}

// AddCartItem adds item to cart
func (service *CartItemServiceImpl) AddCartItem(cartItem model.CartItem) {}

// RemoveCartItem removes item from cart
func (service *CartItemServiceImpl) RemoveCartItem(cartItemID string) {}

// RemoveCartItems removes all items
func (service *CartItemServiceImpl) RemoveCartItems(cart model.Cart) {}
