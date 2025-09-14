package lesson2

import (
	"testing"
)

func TestPurchaseValid(t *testing.T) {
	testcases := map[string]struct {
		purchased float64
		want      float64
	}{
		"zero_valid":              {0, 1},
		"valid_0.01":              {0.01, 1},
		"valid_1":                 {1, 1},
		"valid_2":                 {2, 1},
		"middle_no_discount":      {150, 1},
		"max_val_no_disc_minus_1": {299, 1},
		"max_val_no_disc_minus":   {300, 1},
		"min_val_discount":        {300.01, 0.95},
		"min_val_discount+1":      {301, 0.95},
		"min_val_discount+2":      {302, 0.95},
		"safe_discount":           {600, 0.95},
		"middle_discount-1":       {799, 0.95},
		"middle_discount":         {800.00, 0.95},
		"min_val_max_disc":        {800.01, 0.9},
		"max_disc+1":              {801, 0.9},
		"max_disc":                {1500, 0.9},
	}
	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			rev, _ := GetDiscountMultiplier(test.purchased)
			if rev != test.want {
				t.Errorf("Result: %v, expected %v", rev, test.want)
			}
		})
	}
}

func TestPurchaseReturnsError(t *testing.T) {
	testcases := map[string]struct {
		purchased float64
		want      string
	}{
		"minus1":  {-1, "negative payment amount registered"},
		"minus2":  {-1, "negative payment amount registered"},
		"minus20": {-20, "negative payment amount registered"},
	}
	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			_, err := GetDiscountMultiplier(test.purchased)
			if err.Error() != test.want {
				t.Errorf("Result: %v, expected %v", err.Error(), test.want)
			}
		})
	}
}
