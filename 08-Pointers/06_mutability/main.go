package main

import "fmt"

type Order struct {
	Total int
}

func ApplyDiscount(order *Order, percent int) {
	order.Total = order.Total - (order.Total*percent)/100
}

func DiscountedOrder(order Order, percent int) Order {
	order.Total = order.Total - (order.Total*percent)/100
	return order
}

func main() {
	order := Order{
		Total: 1000,
	}

	fmt.Println(order)

	ApplyDiscount(&order, 20)

	fmt.Println(order)

	order2 := Order{
		Total: 1000,
	}

	discounted := DiscountedOrder(order2, 10)

	fmt.Println("Original:", order2.Total)
	fmt.Println("Discounted:", discounted.Total)
}
