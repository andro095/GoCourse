package checkout

import (
	"errors"
	"fmt"
	"strings"
)

func TryChangeCustomerByValue(o Order, name string) {
	o.Customer = name
}

func TryChangeCustomerByPointer(o *Order, name string) {
	o.Customer = name
}

func setCity(o *Order, city string) {
	if o.Meta == nil {
		o.Meta = map[string]string{}
	}

	o.Meta["city"] = city // Map, Slice, func, pointer, chan
}

func setZone(o *Order, zone string) {
	if o.Meta == nil {
		o.Meta = map[string]string{}
	}

	o.Meta["zone"] = zone
}

func ValidateOrder(o Order) error {
	if o.ID == "" {
		return errors.New("Order ID is mandatory")
	}

	if o.Customer == "" {
		return errors.New("Order Customer is mandatory")
	}

	if len(o.Items) == 0 {
		return errors.New("Order Items are mandatory")
	}

	for i, item := range o.Items {
		if item.SKU == "" {
			return fmt.Errorf("Element [%d]: SKU is mandatory", i)
		}
		if item.Qty <= 0 {
			return fmt.Errorf("Item[%s]: Quantity must be > 0", item.SKU)
		}

		if item.Price <= 0 {
			return fmt.Errorf("Item[%s]: Price must be > 0", item.SKU)
		}
	}

	return nil
}

func ParseCoupon(code string) (Coupon, error) {
	coupon := strings.TrimSpace(strings.ToUpper(code))

	if coupon == "" {
		return Coupon{}, errors.New("Coupon code is mandatory")
	}

	switch coupon {
	case "SAVE10":
		return Coupon{
			Code: coupon,
			Kind: "PERCENT",
			Val:  10,
		}, nil
	case "LESS500":
		return Coupon{
			Code: coupon,
			Kind: "FLAT",
			Val:  500,
		}, nil

	case "FREESHIPPING":
		return Coupon{
			Code: coupon,
			Kind: "SHIPPING",
			Val:  0,
		}, nil
	default:
		return Coupon{}, fmt.Errorf("Unknown coupon code: %s", coupon)
	}
}
