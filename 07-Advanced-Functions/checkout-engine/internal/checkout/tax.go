package checkout

type TaxFn func(Order) Money

func NoTax(Order) Money {
	return 0
}

func HST(order Order) Money {
	sub := order.CalcSubtotal()
	return sub * 13 / 100
}

func NewTaxByState(state string) TaxFn {
	switch state {
	case "ON":
		return func(o Order) Money { return o.CalcSubtotal() * 13 / 100 }
	case "QC":
		return func(o Order) Money { return o.CalcSubtotal() * 15 / 100 }
	case "BC":
		return func(o Order) Money { return o.CalcSubtotal() * 12 / 100 }
	case "AB":
		return func(o Order) Money { return o.CalcSubtotal() * 5 / 100 }
	case "MB":
		return func(o Order) Money { return o.CalcSubtotal() * 12 / 100 }
	case "NB":
		return func(o Order) Money { return o.CalcSubtotal() * 15 / 100 }
	case "NS":
		return func(o Order) Money { return o.CalcSubtotal() * 15 / 100 }
	case "PE":
		return func(o Order) Money { return o.CalcSubtotal() * 15 / 100 }
	case "SK":
		return func(o Order) Money { return o.CalcSubtotal() * 11 / 100 }
	case "NL":
		return func(o Order) Money { return o.CalcSubtotal() * 15 / 100 }
	case "NT":
		return func(o Order) Money { return o.CalcSubtotal() * 5 / 100 }
	case "NU":
		return func(o Order) Money { return o.CalcSubtotal() * 5 / 100 }
	case "YT":
		return func(o Order) Money { return o.CalcSubtotal() * 5 / 100 }
	default:
		return NoTax
	}
}
