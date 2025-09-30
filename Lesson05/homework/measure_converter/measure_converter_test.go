package measureconverter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthConverter(t *testing.T) {
	tcs := []struct {
		name   string
		system MetricImpSys
		value  float64
		expect float64
	}{
		{"Metric positive", Metric, 10, 3.93},
		{"Metric negative", Metric, -10.1, -3.97},
		{"Imperial positive", Imperial, 3.93, 9.98},
		{"Imperial negative", Imperial, -3.97, -10.08},
	}

	for _, test := range tcs {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, LengthConverterConstructor(test.system, test.value).Convert())
		})
	}
}
