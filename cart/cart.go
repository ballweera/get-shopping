package cart

import "github.com/ballweera/get-shopping/customer"

// Cart is shopping cart
type Cart struct {
	CartID     string
	CartItems  []CartItem
	Customer   customer.Customer
	TotalPrice float64
}

type CartItem struct{}
