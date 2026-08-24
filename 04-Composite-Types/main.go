package main

import "fmt"

type Item struct {
	SKU   string
	Name  string
	Price float32
	Qty   int
}

type Order struct {
	ID       int
	Customer string
	Items    []Item
	Meta     map[string]string
}

func main() {

	item1 := Item{
		SKU:   "1234",
		Name:  "Shirt",
		Price: 10.99,
		Qty:   3,
	}

	item2 := Item{
		SKU:   "5678",
		Name:  "Shoes",
		Price: 20.99,
		Qty:   2,
	}

	order := Order{
		ID:       1,
		Customer: "John Doe",
		Items:    []Item{item1, item2},
		Meta: map[string]string{
			"city":   "Jakarta",
			"source": "Online",
			"coupon": "LOOOL",
		},
	}

	// Order
	fmt.Println(order.ID)
	fmt.Println(order.Customer)
	fmt.Println(order.Items[0])
	fmt.Println(order.Meta["coupon"])

	order.Meta["coupon"] = "HOLACARADEBOLA"

	fmt.Println(order.Meta["coupon"])

}
