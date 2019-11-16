package handler

import (
	"net/http"

	"github.com/ballweera/get-shopping/service"
)

// CartController is handler of cart
type CartController struct {
	csvr service.CartService
}

// GetCartID return CartID
func (ctrl *CartController) GetCartID(w http.ResponseWriter, r *http.Request) {
}

// CartItems returns the list of CartItem
func (ctrl *CartController) CartItems(w http.ResponseWriter, r *http.Request) {
}
