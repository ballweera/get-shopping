package model

// Cart is shopping cart
type Cart struct {
	CartID     string
	CartItems  []CartItem
	Customer   Customer
	TotalPrice float64
}

func (this *Cart) GetCartItems() []CartItem {
	return this.CartItems
}
