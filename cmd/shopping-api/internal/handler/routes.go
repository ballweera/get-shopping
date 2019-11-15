package handler

import "net/http"

func API() http.Handler {
	ctrl := &CartController{}

	mux := http.NewServeMux()
	mux.HandleFunc("/cart", ctrl.GetCartItems)
	return mux
}
