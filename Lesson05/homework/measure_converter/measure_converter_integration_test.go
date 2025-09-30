//go:build integration

package measureconverter

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

const API_URL = "https://api.currencyapi.com/v3/"

func TestCurrencyConverter(t *testing.T) {
	tcs := map[string]struct {
		fromCurr string
		toCurr   string
		value    float64
	}{
		"dkk_to_usd": {"DKK", "USD", 20.123},
		"sek_to_dzd_many_decimals": {"SEK", "DZD", 20.12343523452345234523452345234},
		"tjs_to_eur_many_zeroes": {"TJS", "EUR", 0.000000000000000000},
		"tjs_to_eur_very_small_number": {"TJS", "EUR", 0.0000000000000000001},
	}

	for name, test := range tcs {
		t.Run(name, func(t *testing.T) {
			converter := CurrencyConverter{test.fromCurr, API_URL}
			val, err := converter.Convert(test.value, test.toCurr)
			fmt.Println("Value is:", val)
			assert.NoErrorf(t, err, "%s: error: %v\n", name, err)
		})
	}
}
