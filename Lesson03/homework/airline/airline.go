package airline

import (
	"math"
)

func AirlineDiscountCalc(INdia, monOrFri, sixDayStay bool, yo int) float64 {
	discount := 1.0
	if yo <= 2 {
		return 0
	} else if yo < 18 {
		discount -= 0.4
	}
	if !monOrFri {
		if INdia {
			if yo >= 18 {
				discount -= 0.2 // 20% discount if older than 18 inside india
			}
		} else {
			discount -= 0.25
		}
	}
	if sixDayStay {
		discount -= 0.1
	}
	ratio := 1000.0
	intermediate := math.Round(discount * ratio) / float64(ratio) // Round to two decimals
	return intermediate
}
