package controller

import (
	"net/http"

	"github.com/ballweera/get-shopping/service"
)

type CartController struct {
	CartService service.CartService
}

func (this *CartController) GetCartId(w http.ResponseWriter, r *http.Request) {
}

func (this *CartController) GetCartItems(w http.ResponseWriter, r *http.Request) {
}
