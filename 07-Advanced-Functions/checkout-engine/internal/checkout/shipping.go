package checkout

type ShippingFn func(Order) Money

func FreeShipping(Order) Money {
	return 0
}

func NewShippingByZone(zone string) ShippingFn {
	switch zone {
	case "LOCAL":
		return FreeShipping
	case "NATIONAL":
		return func(o Order) Money { return 500 }
	case "INTERNATIONAL":
		return func(o Order) Money { return 1000 }
	default:
		return FreeShipping
	}
}
