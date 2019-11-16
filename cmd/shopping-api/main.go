package main

import "net/http"

import "github.com/ballweera/get-shopping/cmd/shopping-api/internal/handler"

func main() {
	http.ListenAndServe(":9090", handler.API())
}
