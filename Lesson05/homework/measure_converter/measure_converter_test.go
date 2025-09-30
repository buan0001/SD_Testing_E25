//go:build !integration

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
		{"Metric positive", Metric, 10, 3.94},
		{"Metric negative", Metric, -10.1, -3.98},
		{"Imperial positive", Imperial, 3.93, 9.98},
		{"Imperial negative", Imperial, -3.97, -10.08},
	}

	for _, test := range tcs {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, LengthConverterConstructor(test.system, test.value).Convert(), "%v == %v\n", test.value, test.expect)
		})
	}
}
func TestWeightConverter(t *testing.T) {
	tcs := []struct {
		name   string
		system MetricImpSys
		value  float64
		expect float64
	}{
		{"Metric to Imperial positive", Metric, 1.0, 2.20},
		{"Metric to Imperial negative", Metric, -5.5, -12.13},
		{"Metric to Imperial zero", Metric, 0.0, 0.0},
		{"Metric to Imperial large", Metric, 100.0, 220.46},
		{"Imperial to Metric positive", Imperial, 2.20, 1},
		{"Imperial to Metric negative", Imperial, -12.13, -5.5},
		{"Imperial to Metric zero", Imperial, 0.0, 0.0},
		{"Imperial to Metric large", Imperial, 220.46, 100},
		{"Metric decimal precision", Metric, 2.555, 5.64},
		{"Imperial decimal precision", Imperial, 5.63, 2.55},
	}

	for _, test := range tcs {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, WeightConverterConstructor(test.system, test.value).Convert())
		})
	}
}

func TestTemperatureConverter(t *testing.T) {
	tcs := []struct {
		name   string
		from   TempSystem
		to     TempSystem
		value  float64
		expect float64
	}{
		// Celsius conversions
		{"Celsius to Fahrenheit freezing", Celsius, Fahrenheit, 0.0, 32.0},
		{"Celsius to Fahrenheit boiling", Celsius, Fahrenheit, 100.0, 212.0},
		{"Celsius to Fahrenheit negative", Celsius, Fahrenheit, -40.0, -40.0},
		{"Celsius to Kelvin freezing", Celsius, Kelvin, 0.0, 273.15},
		{"Celsius to Kelvin absolute zero", Celsius, Kelvin, -273.15, 0.0},
		{"Celsius to Kelvin positive", Celsius, Kelvin, 25.0, 298.15},

		// Fahrenheit conversions - Fixed values based on actual fToC implementation
		{"Fahrenheit to Celsius freezing", Fahrenheit, Celsius, 0, -17.78},
		{"Fahrenheit to Celsius boiling", Fahrenheit, Celsius, 212.0, 100.00},
		{"Fahrenheit to Celsius negative", Fahrenheit, Celsius, -40.0, -40.00},
		{"Fahrenheit to Kelvin freezing", Fahrenheit, Kelvin, 32.0, 273.15},
		{"Fahrenheit to Kelvin absolute zero", Fahrenheit, Kelvin, -459.67, 0.0},
		{"Fahrenheit to Kelvin room temp", Fahrenheit, Kelvin, 68.0, 293.15},

		// Kelvin conversions - Fixed rounding issues
		{"Kelvin to Celsius absolute zero", Kelvin, Celsius, 0, -273.15},
		{"Kelvin to Celsius freezing", Kelvin, Celsius, 273.15, 0},
		{"Kelvin to Celsius room temp", Kelvin, Celsius, 298.15, 25.0},
		{"Kelvin to Fahrenheit absolute zero", Kelvin, Fahrenheit, 0, -459.67},
		{"Kelvin to Fahrenheit freezing", Kelvin, Fahrenheit, 273.15, 32.0},
		{"Kelvin to Fahrenheit boiling", Kelvin, Fahrenheit, 373.15, 212.0},

		// Same system conversions (edge case)
		{"Celsius to Celsius", Celsius, Celsius, 25.0, 25.0},
		{"Fahrenheit to Fahrenheit", Fahrenheit, Fahrenheit, 77.0, 77.0},
		{"Kelvin to Kelvin", Kelvin, Kelvin, 300.0, 300.0},

		// Decimal precision tests - Fixed based on actual calculations
		{"Celsius to Fahrenheit precision", Celsius, Fahrenheit, 36.67, 98.01},
		{"Kelvin to Celsius precision", Kelvin, Celsius, 255.37, -17.78},
	}

	for _, test := range tcs {
		t.Run(test.name, func(t *testing.T) {
			converter := TemperatureConverterConstructor(test.from, test.value)
			assert.InDelta(t, test.expect, converter.Convert(test.to), 0.01)
		})
	}
}
