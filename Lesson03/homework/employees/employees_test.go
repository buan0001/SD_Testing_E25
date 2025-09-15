package employees

import (
	"fmt"
	"testing"
	"time"
)

func TestValidCPR(t *testing.T) {
	testcases := []string{
		"1234567890",
		"0000000000",
		"9999999999",
	}
	for _, test := range testcases {
		t.Run(test, func(t *testing.T) {
			t.Parallel()
			emp := Employee{}
			err := emp.SetCPR(test)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
func TestInvalidCPR(t *testing.T) {
	testcases := []string{
		"-123456789",
		"00000000000",
		"99999999999",
		"123456789",
		"abcdefghij",
	}
	for _, test := range testcases {
		t.Run(test, func(t *testing.T) {
			t.Parallel()
			emp := Employee{}
			err := emp.SetCPR(test)
			if err == nil {
				t.Errorf("wanted error, none found with: %v", test)
			}
		})
	}
}

func TestValidName(t *testing.T) {
	testcases := []string{
		"a",
		" ",
		"-",
		"b a",
		"b-a",
		"b m a f k a-easdf",
	}
	var thirtyCharName string
	var twentyNineChar string
	for i := 0; i < 30; i++ {
		thirtyCharName += "a"
		if i == 29 {
			continue
		}
		twentyNineChar += "a"
	}
	testcases = append(testcases, thirtyCharName)
	testcases = append(testcases, twentyNineChar)
	for _, test := range testcases {
		t.Run(test, func(t *testing.T) {
			t.Parallel()
			emp := Employee{}
			err := emp.SetFirstName(test)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
func TestInvalidName(t *testing.T) {
	testcases := []string{
		"",
		"1",
		"b_",
		"_",
		"£",
		"b m a f k a_easdf",
	}
	var thirtyOneChar string
	var thirtyTwoChar string
	for i := 0; i < 32; i++ {
		thirtyTwoChar += "a"
		if i == 31 {
			continue
		}
		thirtyOneChar += "a"
	}
	testcases = append(testcases, thirtyOneChar)
	testcases = append(testcases, thirtyTwoChar)
	for _, test := range testcases {
		t.Run(test, func(t *testing.T) {
			t.Parallel()
			emp := Employee{}
			err := emp.SetFirstName(test)
			if err == nil {
				t.Errorf("wanted error, none found with: %v", test)
			}
		})
	}
}

func TestValidBaseSalary(t *testing.T) {
	testcases := []float64{
		20000,
		20000.01,
		20000.02,
		60000.02,
		99999.99,
		100000,
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%f", test), func(t *testing.T) {
			t.Parallel()
			emp := Employee{}
			err := emp.SetBaseSalary(test)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
func TestInvalidBaseSalary(t *testing.T) {
	testcases := []float64{
		0,
		19999.98,
		19999.99,
		100000.01,
		100000.02,
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%f", test), func(t *testing.T) {
			t.Parallel()
			emp := Employee{}
			err := emp.SetBaseSalary(test)
			if err == nil {
				t.Errorf("wanted error, got none from: %v", test)
			}
		})
	}
}

var eighteenYearsAgo = time.Now().AddDate(-18, 0, 0)

func TestValidBday(t *testing.T) {
	testcases := []time.Time{
		time.Date(1998, 5, 3, 0, 0, 0, 0, time.UTC),
		eighteenYearsAgo,
		eighteenYearsAgo.Add(time.Nanosecond),
		eighteenYearsAgo.Add(time.Microsecond),
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			t.Parallel()
			emp := Employee{}
			err := emp.SetBirthDate(test)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
func TestInvalidBday(t *testing.T) {
	testcases := []time.Time{
		time.Now().Add(time.Second * 5),
		time.Now().AddDate(-17, -11, -29),
		time.Now().AddDate(-17, -11, -30),
		time.Now().AddDate(1, 0, 0),
		eighteenYearsAgo.Add(time.Hour),
		eighteenYearsAgo.Add(time.Hour * 2),
	}
	for _, test := range testcases {
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			t.Parallel()
			emp := Employee{}
			err := emp.SetBirthDate(test)
			if err == nil {
				t.Errorf("expected error, got none from: %v", test)
			}
		})
	}
}

func TestValidEmpDay(t *testing.T) {
	testcases := []time.Time{
		time.Date(1998, 5, 3, 0, 0, 0, 0, time.UTC),
		time.Now(),
		time.Now().Add(-time.Second),
		time.Now().AddDate(0, 0, -1),
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			t.Parallel()
			emp := Employee{}
			err := emp.SetEmploymentDate(test)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
func TestInvalidEmpDay(t *testing.T) {
	testcases := []time.Time{
		time.Now().AddDate(1, 1, 1),
		time.Now().AddDate(0, 0, 1),
		time.Now().Add(time.Second),
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			t.Parallel()
			emp := Employee{}
			err := emp.SetEmploymentDate(test)
			if err == nil {
				t.Errorf("expected error, got none from: %v", test)
			}
		})
	}
}

func TestGetSalary(t *testing.T) {
	testcases := map[string]struct {
		salary   float64
		eduLevel EducationLevel
		want     float64
	}{
		"min_sal_tert": {20000, Tertiary, 23660},
		"min_sal_sec":  {20000, Secondary, 22440},
		"min_sal_prim": {20000, Primary, 21220},
		"min_sal_none": {20000, None, 20000},
		"max_sal_tert": {100000, Tertiary, 103660},
		"max_sal_sec":  {100000, Secondary, 102440},
		"max_sal_prim": {100000, Primary, 101220},
		"max_sal_none": {100000, None, 100000},
	}
	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			emp := Employee{baseSalary: test.salary, educationLevel: test.eduLevel}
			actualSalary := emp.GetSalary()
			if actualSalary != test.want {
				t.Errorf("expected %v, got: %v", test.want, actualSalary)
			}
		})
	}
}

func TestGetDiscount(t *testing.T) {
	testcases := map[string]struct {
		empDay time.Time
		want   float64
	}{
		"no_disc":     {time.Now(), 0},
		"1_yr_disc":   {time.Now().AddDate(-1, 0, 0), 0.5},
		"100_yr_disc": {time.Now().AddDate(-100, 0, 0), 50},
	}
	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			emp := Employee{employmentDate: test.empDay}
			actualSalary := emp.GetDiscount()
			if actualSalary != test.want {
				t.Errorf("expected %v, got: %v", test.want, actualSalary)
			}
		})
	}
}

func TestGetShippingCosts(t *testing.T) {
	testcases := map[string]struct {
		country string
		want    float64
	}{
		"denmark":                     {"denmark", 0},
		"sweden":                      {"sweden", 0},
		"norway":                      {"norway", 0},
		"iceland":                     {"iceland", 0.5},
		"finland":                     {"finland", 0.5},
		"literally_any_other_country": {"asdfhasdfa", 1},
		"mixed_cases_0_disc":          {"SwedEn", 0},
		"mixed_cases_50_disc":         {"icELanD", 0.5},
		"only numbers":                {"1234567", 1},
		"negative number":             {"-1234567", 1},
	}
	for name, test := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			emp := Employee{country: test.country}
			actualSalary := emp.GetShippingCosts()
			if actualSalary != test.want {
				t.Errorf("expected %v, got: %v", test.want, actualSalary)
			}
		})
	}
}
