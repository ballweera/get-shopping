package handler

import "net/http"

// API returns handler
func API() http.Handler {
	cartCtrl := &CartController{}
	mux := http.NewServeMux()
	mux.HandleFunc("/cart", cartCtrl.CartItems)

	return mux
}
