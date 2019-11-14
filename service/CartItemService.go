package service

import(
	"github.com/ballweera/get-shopping/model"
)

// CartItemService is service
type CartItemService interface {
	AddCartItem(ci model.CartItem)
	RemoveCartItem(cartItemID string)
	RemoveAllCartItems(cart model.Cart)
}
