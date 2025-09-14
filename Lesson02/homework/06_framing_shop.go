package lesson2

import (
	"fmt"
)

func PriceForFrame(width, height float64) (float64, error) {
	if height < 30 || 60 < height {
		return -1, fmt.Errorf("height out of bounds")
	} else if width < 30 || 100 < width {
		return -1, fmt.Errorf("width out of bounds")
	}
	if width*height > 1600 {
		return 3500, nil
	}
	return 3000, nil
}
