package checkout

import (
	"testing"

	"go.uber.org/zap/zaptest"
)

func TestIsSkuExist(t *testing.T) {
	test := []struct {
		skuName string
	}{
		{
			skuName: "A",
		},
		{
			skuName: "A",
		},
		{
			skuName: "B",
		},
		{
			skuName: "C",
		},
	}
	PriceRules = []PriceRule{
		{
			Sku:       "A",
			UnitPrice: 50,
			SpecialPrice: specialPrice{
				Quantity: 3,
				Price:    130,
			},
		},
		{
			Sku:       "B",
			UnitPrice: 30,
			SpecialPrice: specialPrice{
				Quantity: 2,
				Price:    45,
			},
		},
		{
			Sku:       "C",
			UnitPrice: 20,
		},
	}
	for _, tt := range test {
		t.Run(tt.skuName, func(t *testing.T) {
			_, status := IsSkuExist(tt.skuName)

			if !status {
				t.Errorf("Failed to found the value")
			}
		})
	}
}

func TestScan(t *testing.T) {
	test := []struct {
		skuName string
	}{
		{
			skuName: "A",
		},
		{
			skuName: "A",
		},
		{
			skuName: "B",
		},
		{
			skuName: "C",
		},
	}
	PriceRules = []PriceRule{
		{
			Sku:       "A",
			UnitPrice: 50,
			SpecialPrice: specialPrice{
				Quantity: 3,
				Price:    130,
			},
		},
		{
			Sku:       "B",
			UnitPrice: 30,
			SpecialPrice: specialPrice{
				Quantity: 2,
				Price:    45,
			},
		},
		{
			Sku:       "C",
			UnitPrice: 20,
		},
	}
	for _, tt := range test {
		t.Run(tt.skuName, func(t *testing.T) {
			logger := zaptest.NewLogger(t)
			err := Scan(logger, tt.skuName)

			if err != nil {
				t.Errorf("Failed to setup Cart %s", err)
			}
		})
	}
}
