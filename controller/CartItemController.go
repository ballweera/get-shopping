package controller

import (
	"net/http"

	"github.com/ballweera/get-shopping/service"
)

type CartItemController struct {
	CartService service.CartService
}

func (this *CartItemController) AddCartItem(w http.ResponseWriter, r *http.Request) {
}

func (this *CartItemController) RemoveCartItem(w http.ResponseWriter, r *http.Request) {
}
