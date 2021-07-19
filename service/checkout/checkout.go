package checkout

import (
	"errors"

	"go.uber.org/zap"
)

type specialPrice struct {
	Quantity int     `json:"quantity,omitempty"`
	Price    float32 `json:"price,omitempty"`
}
type PriceRule struct {
	Sku          string       `json:"sku"`
	UnitPrice    float32      `json:"unit_Price"`
	SpecialPrice specialPrice `json:"special_price,omitempty"`
}

type item struct {
	count int
	rule  PriceRule
}

var PriceRules []PriceRule
var Cart = map[string]item{}

// Scan will add the Sku in the cart and manage the counts
func Scan(logger *zap.Logger, skuName string) error {
	// check if SKU Exist then add to cart
	itemDetail, checkSku := IsSkuExist(skuName)
	if checkSku {
		c := string(skuName)
		if Cart[c].count > 0 {
			Cart[c] = item{
				count: Cart[c].count + 1,
				rule:  itemDetail,
			}
		} else {
			Cart[c] = item{
				count: 1,
				rule:  itemDetail,
			}
		}

	} else {
		logger.Error("Sku not Found")
		return errors.New("sku Not found")
	}
	return nil
}

// This will check if the SKu exist in the price List
func IsSkuExist(sku string) (PriceRule, bool) {
	for _, p := range PriceRules {
		if p.Sku == sku {
			return p, true
		}
	}
	return PriceRule{}, false
}

// Final call to get Total amount
func CalculateTotal() (float32, error) {
	var total float32

	if len(Cart) > 0 {
		for _, value := range Cart {
			if value.rule.SpecialPrice.Quantity > 0 {
				total += float32(value.count/value.rule.SpecialPrice.Quantity) * value.rule.SpecialPrice.Price
				total += float32(value.count%value.rule.SpecialPrice.Quantity) * value.rule.UnitPrice
			} else {
				total += float32(value.count) * value.rule.UnitPrice
			}
		}
	} else {
		return total, errors.New("no item added")
	}
	return total, nil
}
