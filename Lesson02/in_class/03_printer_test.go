package printer

import (
	"math"
	"testing"
)

// Printer cartridges

// A wholesaler sells printer cartridges. The minimum order quantity is 5.
// There is a 20% discount for orders of 100 or more printer cartridges.
// You have been asked to prepare test cases using various values for the number of printer cartridges ordered.

// Use black-box analysis to identify a comprehensive series of test cases:

//     - Identify the corresponding equivalence partitions and propose test cases based on them
//     - Use 3-value boundary value analysis to identify further test cases
//     - Write down the full resulting list of test cases
//     - Implement the discount calculation function in code and write the corresponding unit tests
// 	   	 in the language and unit test framework of your choice

// Partitions:
// Valid: 5-99, 100-MAX_INT
// Invalid: MIN_INT - (-1), 0, 1-4



// POSITIVE TESTS
func TestValidAmountOrdered(t *testing.T) {
	testcases := map[string]struct {
		amount_ordered int
		want           float64
	}{
		"min_valid_amount":              {5, 1},
		"min_valid_amount_plus_1":       {6, 1},
		"safe_non_discount":             {50, 1},
		"max_valid_no_discount_minus_1": {98, 1},
		"max_valid_no_discount":         {99, 1},
		"min_valid_discount":            {100, 0.8},
		"min_valid_discount_plus_1":     {101, 0.8},
		"safe_discount_amount":          {500, 0.8},
		"max_int":                       {math.MaxInt, 0.8},
	}
	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			rev, _ := Discount_multiplier(test.amount_ordered)
			if rev != test.want {
				t.Errorf("Result: %v, expected %v", rev, test.want)
			}
		})
	}
}


// NEGATIVE TESTS
func TestBadOrders(t *testing.T) {
	testcases := map[string]struct {
		amount_ordered int
		want           string
	}{
		"max_invalid_amount":         {4, "too few ordered"},
		"max_invalid_amount_minus_1": {3, "too few ordered"},
		"order_1":                    {1, "too few ordered"},
		"negative_amount":            {-1, "too few ordered"},
		"min_int":                    {math.MinInt, "too few ordered"},
		"order_0":                    {0, "too few ordered"},
		"order_3.0_cast_to_int":      {int(3.0), "too few ordered"},
	}
	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			_, err := Discount_multiplier(test.amount_ordered)
			if err.Error() != test.want {
				t.Errorf("Result: %v, expected %v", err.Error(), test.want)
			}
		})
	}
}
