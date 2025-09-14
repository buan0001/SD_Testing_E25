package lesson2

import (
	"testing"
)

func TestDimensionsValid(t *testing.T) {
	testcases := map[string]struct {
		width  float64
		height float64
		want   float64
	}{
		"min_valid":         {30, 30, 3000},
		"min_valid+1":       {31, 31, 3000},
		"max_size_discount": {40, 40, 3000},
		"safe_middle":       {50, 50, 3500},
		"upper_allowed-1":   {99, 59, 3500},
		"upper_allowed":     {100, 60, 3500},
	}
	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			rev, _ := PriceForFrame(test.width, test.height)
			if rev != test.want {
				t.Errorf("Result: %v, expected %v", rev, test.want)
			}
		})
	}
}

func TestDimensionsInvalid(t *testing.T) {
	testcases := map[string]struct {
		width  float64
		height float64
		want   string
	}{
		"negative_width-50":  {-50, 30, "width out of bounds"},
		"negative_width-2":   {-2, 30, "width out of bounds"},
		"negative_width-1":   {-1, 30, "width out of bounds"},
		"width_0":            {0, 31, "width out of bounds"},
		"width_1":            {1, 31, "width out of bounds"},
		"width_2":            {2, 31, "width out of bounds"},
		"width_15":           {15, 31, "width out of bounds"},
		"width_28":           {28, 31, "width out of bounds"},
		"width_29":           {29, 31, "width out of bounds"},
		"width_101":          {101, 31, "width out of bounds"},
		"width_102":          {102, 31, "width out of bounds"},
		"width_150":          {150, 31, "width out of bounds"},
		"negative_height-50": {30, -50, "height out of bounds"},
		"negative_height-2":  {30, -2, "height out of bounds"},
		"negative_height-1":  {30, -1, "height out of bounds"},
		"height_0":           {30, 0, "height out of bounds"},
		"height_1":           {30, 1, "height out of bounds"},
		"height_2":           {30, 2, "height out of bounds"},
		"height_15":          {30, 15, "height out of bounds"},
		"height_28":          {30, 29, "height out of bounds"},
		"height_29":          {30, 29, "height out of bounds"},
		"height_61":          {61, 15, "height out of bounds"},
		"height_62":          {62, 29, "height out of bounds"},
		"height_150":         {150, 29, "height out of bounds"},
	}
	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			_, err := PriceForFrame(test.width, test.height)
			if err.Error() != test.want {
				t.Errorf("Result: %v, expected %v", err.Error(), test.want)
			}
		})
	}
}
