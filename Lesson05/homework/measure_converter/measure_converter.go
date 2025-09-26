package measureconverter

import (
	"math"
)

//_______________________ TYPES ____________________

type MetricImpSys int

const (
	Metric MetricImpSys = iota
	Imperial
)

type TempSystem int

const (
	Celsius TempSystem = iota
	Fahrenheit
	Kelvin
)

//_______________________ HELPERS ____________________

func roundToDecimals(num float64, decimals int) float64 {
	if decimals < 0 {
		return roundToDecimals(num, 0)
	}
	decimal_const := math.Pow10(decimals)
	return math.Trunc(num*decimal_const) / decimal_const
}

func twoDecimals(num float64) float64 {
	return roundToDecimals(num, 2)
}

//_______________________ LENGTH CONVERTER ____________________

type LengthConverter struct {
	system MetricImpSys
	value  float64
}

func LengthConverterConstructor(sys MetricImpSys, val float64) (*LengthConverter, error) {
	val = twoDecimals(val)

	return &LengthConverter{sys, val}, nil
}

func (lconv *LengthConverter) Convert() float64 {
	ONE_INCH_TO_CM := 2.54
	var val float64

	if lconv.system == Imperial {
		val = lconv.value / ONE_INCH_TO_CM
	} else {
		val = lconv.value * ONE_INCH_TO_CM
	}

	return twoDecimals(val)
}

// _______________________ WEIGHT CONVERTER ____________________
type WeightConverter struct {
	system MetricImpSys
	value  float64
}

func WeightConverterConstructor(sys MetricImpSys, val float64) (*WeightConverter, error) {
	val = twoDecimals(val)

	return &WeightConverter{sys, val}, nil
}

func (lconv *WeightConverter) Convert() float64 {
	ONE_KG_TO_PUNDS := 2.204623
	var val float64

	if lconv.system == Imperial {
		val = lconv.value / ONE_KG_TO_PUNDS
	} else {
		val = lconv.value * ONE_KG_TO_PUNDS
	}

	return twoDecimals(val)
}

//_______________________ TEMPERATURE CONVERTER ____________________

type TemperatureConverter struct {
	system TempSystem
	value  float64
}

func TemperatureConverterConstructor(sys TempSystem, val float64) (*TemperatureConverter, error) {
	val = twoDecimals(val)

	return &TemperatureConverter{sys, val}, nil
}

func (lconv *TemperatureConverter) Convert(to TempSystem) float64 {
	from := lconv.system
	val := lconv.value
	
	switch {
	case to == Kelvin && from == Fahrenheit:
		return fToK(val)
	case to == Kelvin && from == Celsius:
		return cToK(val)
	case to == Fahrenheit && from == Celsius:
		return cToF(val)
	case to == Fahrenheit && from == Kelvin:
		return kToF(val)
	case to == Celsius && from == Kelvin:
		return kToC(val)
	case to == Celsius && from == Fahrenheit:
		return fToC(val)
	default:
		return twoDecimals(val) // Probably the same to and from
	}
}

// Consult: https://www.unitconverters.net/temperature-converter.html
const fahrKelDiff = 459.67
const celKelDiff = 273.15
const celFahrDiff = 32

func fToK(val float64) float64 {
	return twoDecimals(5.0 / 9 * (val + fahrKelDiff))
}
func cToK(val float64) float64 {
	return twoDecimals(val + celKelDiff)
}
func cToF(val float64) float64 {
	return twoDecimals(9.0/5*val + celFahrDiff)
}
func kToF(val float64) float64 {
	return twoDecimals(9.0/5*val - fahrKelDiff)
}
func kToC(val float64) float64 {
	return twoDecimals(val - celKelDiff)
}
func fToC(val float64) float64 {
	return twoDecimals(5.0 / 9 * (val + celFahrDiff))
}


//_______________________ CURRENCY CONVERTER ____________________

type CurrencyConverter struct {
	baseCurrency string
}

func (conv *CurrencyConverter) Convert(amount float64, destCurr string) {
	const BASE_URL = "https://freecurrencyapi.net/"
	// Use the following routes
	// latest?base_currency=
	// currencies
}