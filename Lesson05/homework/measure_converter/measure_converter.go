package measureconverter

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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

type GradeSystem string

const (
	Danish   GradeSystem = "DK"
	American GradeSystem = "USA"
)

//_______________________ HELPERS ____________________

func roundToDecimals(num float64, decimals int) float64 {
	if decimals < 0 {
		return roundToDecimals(num, 0)
	}
	ratio := math.Pow(10, float64(decimals))
	return math.Round(num*ratio) / ratio
}

func twoDecimals(num float64) float64 {
	return roundToDecimals(num, 2)
}

//_______________________ LENGTH CONVERTER ____________________

type LengthConverter struct {
	fromSystem MetricImpSys
	value      float64
}

func LengthConverterConstructor(baseSystem MetricImpSys, val float64) *LengthConverter {
	val = twoDecimals(val)

	return &LengthConverter{baseSystem, val}
}

func (lconv *LengthConverter) Convert() float64 {

	ONE_INCH_TO_CM := 2.54
	var val float64

	if lconv.fromSystem == Metric {
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

func WeightConverterConstructor(sys MetricImpSys, val float64) *WeightConverter {
	val = twoDecimals(val)

	return &WeightConverter{sys, val}
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

func TemperatureConverterConstructor(sys TempSystem, val float64) *TemperatureConverter {
	val = twoDecimals(val)

	return &TemperatureConverter{sys, val}
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
	return twoDecimals(5.0 / 9 * (val - celFahrDiff))
}

//_______________________ CURRENCY CONVERTER ____________________

type CurrencyConverter struct {
	baseCurrency string
	currencyUrl  string
}

type Currency struct {
	Code  string  `json:"code"`
	Value float64 `json:"value"`
}

type ApiResponse struct {
	Data map[string]Currency `json:"data"`
}

func getDataFromAPI(URLString string) (ApiResponse, error) {
	resp, err := http.Get(URLString)
	if err != nil {
		return ApiResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ApiResponse{}, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var result ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return ApiResponse{}, fmt.Errorf("error decoding json: %s", err)
	}

	return result, nil
}

func extractAmountFromRequest(currency string, body ApiResponse) (float64, error) {
	curr, ok := body.Data[currency]
	if !ok {
		return 0, fmt.Errorf("currency %s not found", currency)
	}
	return curr.Value, nil
}

func (conv *CurrencyConverter) Convert(amount float64, destCurrency string) (float64, error) {
	godotenv.Load() // attempt to load the ENV. Ideally this would happen at server start but it's going to have to stand alone
	API_KEY := os.Getenv("CURRENCY_API_KEY")
	BASE_URL := conv.currencyUrl
	if API_KEY == "" {
		return -1, fmt.Errorf("couldn't load the API key")
	} else if BASE_URL == "" {
		return -1, fmt.Errorf("no API URL found")
	} else if conv.baseCurrency == "" {
		return -1, fmt.Errorf("no base currency provided")
	} else if conv.baseCurrency == destCurrency {
		return amount, nil // Spare the API on pointless calls
	}


	api_url := fmt.Sprintf("%slatest?apikey=%s&base_currency=%s&currencies=%s", BASE_URL, API_KEY, conv.baseCurrency, destCurrency)
	data, err := getDataFromAPI(api_url)
	if err != nil {
		return -1, err
	}

	// https://api.currencyapi.com/v3/latest?apikey=cur_live_wbr6mOwSnH23NJdXL1Ckbx6TZ3fFP1mk3VeHJp3Y&base_currency=DKK&currencies=EUR,USD,CAD
	// BASE= // https://api.currencyapi.com/v3/
	// Use the following route:
	// latest?base_currency=&apikey=&currencies=
	return extractAmountFromRequest(destCurrency, data)
}

//_______________________ GRADES CONVERTER ____________________

func ConvertGrades(grade string, gradingSys GradeSystem) (string, error) {
	// Excecuting SQL script from GO
	// 1. Create database if not exists
	rootDB, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/")
	if err != nil {
		return "", err
	}
	defer rootDB.Close()

	if _, err := rootDB.Exec(`CREATE DATABASE IF NOT EXISTS converter
		DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci`); err != nil {
		return "", fmt.Errorf("create db: %w", err)
	}
	fmt.Println("Database ensured.")

	// 2. Connect specifically to converter DB
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/converter")
	if err != nil {
		return "", err
	}
	defer db.Close()

	content, err := os.ReadFile("grades.sql")
	if err != nil {
		return "", fmt.Errorf("read file: %w", err)
	}

	// 4. Start transaction
	tx, err := db.Begin()
	if err != nil {
		return "", fmt.Errorf("begin tx: %w", err)
	}

	queries := strings.Split(string(content), ";")
	for _, q := range queries {
		q = strings.TrimSpace(q)
		if q == "" || strings.HasPrefix(q, "CREATE DATABASE") {
			continue // already handled
		}
		if _, err := tx.Exec(q); err != nil {
			_ = tx.Rollback()
			return "", fmt.Errorf("exec %q: %w", q, err)
		}
	}

	// Commit if all successful
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Dummy data loaded in transaction.")

	rows, err := db.Query("SELECT * FROM grades")
	if err != nil {
		return "", fmt.Errorf("select: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var gradeDK string
		var gradeUSA string
		if err := rows.Scan(&id, &gradeDK, &gradeUSA); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id %d. DK: %s. USA: %s\n", id, gradeDK, gradeUSA)
	}
	return "abc", nil
}
