package measureconverter

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/joho/godotenv"
)

ENVS := godotenv.Read()

// INdia, monOrFri, sixDayStay bool, yo int
func TestDiscountDT(t *testing.T) {
	testcases := map[string]struct {
		INdia, monOrFri, sixDayStay bool
		yo                          int
		expected                    float64
	}{
		"R1":    {true, true, true, 18, 0.9},
		"R2":    {true, true, false, 18, 1},
		"R3":    {false, false, true, 18, 0.65},
		"R4":    {false, false, false, 18, 0.75},
		"R5":    {true, false, true, 18, 0.7},
		"R6":    {true, false, false, 18, 0.8},
		"R7":    {true, true, true, 17, 0.5},
		"R8":    {true, false, false, 3, 0.6},
		"R9":    {false, false, true, 15, 0.25},
		"R10":   {false, false, false, 15, 0.35},
		"R11":   {true, true, true, 2, 0},
		"R11v2": {false, false, false, 2, 0},
	}
	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			result := AirlineDiscountCalc(test.INdia, test.monOrFri, test.sixDayStay, test.yo)
			if result != test.expected {
				t.Errorf("got: %v. wanted %v. INdia: %v, monOrFri: %v, sixDayStay: %v, yo: %v", result, test.expected, test.INdia, test.monOrFri, test.sixDayStay, test.yo)
			}
		})
	}
}
