package printer

import (
	"fmt"
)

func Discount_multiplier(ordered int) (float64, error) {
	if ordered < 5 {
		return -1, fmt.Errorf("too few ordered")
	}
	if ordered >= 100 {
		return 0.8, nil
	}
	return 1, nil
}
