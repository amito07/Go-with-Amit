package main

import "fmt"

// enum define
type OrderStatus string

// const (
// 	Pending OrderStatus = iota
// 	Shipped
// 	Delivered
// 	Cancelled
// )

const (
	Pending   OrderStatus = "pending"
	Shipped               = "shipped"
	Delivered             = "delivered"
	Cancelled             = "cancelled"
)

func ChangeOrderStatus(status OrderStatus) {
	fmt.Println("Order status changed to: ", status)
}

func main() {
	ChangeOrderStatus(Pending)
}
