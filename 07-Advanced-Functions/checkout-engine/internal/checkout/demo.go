package checkout

func RunDemo() {
	PrintHeader("Hello Checkout Engine :)")
	order := NewOrder("ORDER-001", "Andre")

	order.AddItem(Item{
		SKU:   "ABC",
		Price: 100,
		Qty:   2,
	})

	order.AddItem(Item{
		SKU:   "XYZ",
		Price: 250,
		Qty:   1,
	})

	order.AddItem(Item{
		SKU:   "BCD",
		Price: 500,
		Qty:   2,
	})

	order.AddItem(Item{
		SKU:   "FGH",
		Price: 1000,
		Qty:   1,
	})

	order.AddItem(Item{
		SKU:   "GHI",
		Price: 1500,
		Qty:   4,
	})

	// Validating Order
	PrintKV("Validate Order", ValidateOrder(order))

	PrintKV("OrderID", order.ID)
	PrintKV("Customer", order.Customer)
	PrintKV("Items", len(order.Items))

	remove := order.RemoveItem("ABC")

	PrintKV("Removed", remove)
	PrintKV("Items", len(order.Items))

	PrintDivider()

	sub := order.CalcSubtotal()
	qty := order.CalcTotalQty()

	PrintKV("Subtotal", sub)
	PrintKV("Total Qty", qty)

	PrintDivider()

	TryChangeCustomerByValue(order, "New name 1")
	PrintKV("Customer", order.Customer)

	TryChangeCustomerByPointer(&order, "New name 2")
	PrintKV("Customer", order.Customer)

	setCity(&order, "ON")
	PrintKV("City", order.Meta["city"])

	setZone(&order, "NATIONAL")

	PrintDivider()
	items := []Item{
		{
			SKU:   "ABC",
			Name:  "Item A",
			Price: 1000,
			Qty:   1,
		},
		{
			SKU:   "BCD",
			Name:  "Item B",
			Price: 100,
			Qty:   3,
		},
	}
	order.AddItems(items...)
	PrintKV("Total: ", order.CalcTotalQty())
	PrintKV("Items: ", order.Items)

	PrintDivider()

	findItem, ok := order.FindItem("ABC")

	PrintKV2("Item", findItem, ok)

	getMeta, ok := GetMeta(order, "city")

	PrintKV2("Meta", getMeta, ok)

	indexOfItem, ok := IndexOfItem(order, "ABC")

	PrintKV2("Index of Item", indexOfItem, ok)

	indexofItem2, ok := IndexOfItem(order, "FGH")

	PrintKV2("Index of Item", indexofItem2, ok)

	PrintDivider()

	coupon, err := ParseCoupon("SAVE10")

	PrintKV2("Coupon", coupon, err)

	coupon2, err2 := ParseCoupon("UNKNOWN")

	PrintKV2("Coupon", coupon2, err2)

	PrintDivider()

	// computeValue, _ := Compute(order)

	// PrintKV2("Compute", computeValue, computeError)
	// PrintKV("Total: ", computeValue)

	PrintDivider()

	PrintKV("Discount: ", StringUSD(FlatDiscount(200)(order)))
	th := ThresholdPercentDiscount(2000, 20)
	PrintKV("Threshold Discount: ", th(order))

	PrintDivider()

	cityDiscount := func(o Order) Money {
		city, _ := GetMeta(o, "city")

		if city == "Jakarta" {
			return 200
		}

		return 0
	}

	PrintKV("Special Discount: ", cityDiscount(order))

	PrintDivider()

	discountKeyboard := MakeSKUDiscount("FGH", 100)

	discountABC := MakeSKUDiscount("ABC", 200)

	PrintKV("Discount Keyboard: ", discountKeyboard(order))
	PrintKV("Discount ABC: ", discountABC(order))

	PrintDivider()

	state, _ := GetMeta(order, "city")
	zone, _ := GetMeta(order, "zone")

	taxFn := NewTaxByState(state)
	shippingFn := NewShippingByZone(zone)

	promo := CompositeDiscount{
		Name: "Feb Promotions",
		Fns: []DiscountFn{
			FlatDiscount(100),
			ThresholdPercentDiscount(2000, 10),
			MakeSKUDiscount("ABC", 200),
		},
	}

	bundle := promo.Apply(order)
	PrintKV("Discount", StringUSD(bundle))

	computeValue, computeError := Compute(order, bundle, taxFn, shippingFn, FlatDiscount(5000), ThresholdPercentDiscount(2000, 20))

	PrintKV2("Compute", computeValue, computeError)
	PrintKV("Total: ", StringUSD(computeValue.Total))
}
