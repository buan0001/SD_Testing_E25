package employees

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"
)

// Solve the employees exercise
// Try to follow the specification to the letter
// Testing dates might be problematic. Give it some thought

type Department int

const (
	Undefined Department = iota
	HR
	Finance
	IT
	Sales
	GeneralServices
)

type EducationLevel int

const (
	None EducationLevel = iota
	Primary
	Secondary
	Tertiary
)

var departmentName = map[Department]string{
	HR:              "HR",
	Finance:         "finance",
	IT:              "IT",
	Sales:           "sales",
	GeneralServices: "general services",
}

var educationName = map[EducationLevel]string{
	Primary:   "primary",
	Secondary: "secondary",
	Tertiary:  "tertiary",
}

func (d Department) String() string {
	name, found := departmentName[d]
	if !found {
		return "unknown"
	}
	return name
}

func (e EducationLevel) String() string {
	name, found := educationName[e]
	if !found {
		return "none"
	}
	return name
}

type Employee struct {
	cpr            string
	firstName      string
	lastName       string
	department     Department
	baseSalary     float64
	educationLevel EducationLevel
	birthDate      time.Time
	employmentDate time.Time
	country        string
}

func (e Employee) String() string {
	return fmt.Sprintf(
		"Employee{CPR: %s, Name: %s %s, Department: %s, BaseSalary: %.2f, EducationLevel: %s, BirthDate: %s, EmploymentDate: %s, Country: %s}",
		e.cpr,
		e.firstName,
		e.lastName,
		e.department, // calls Department.String()
		e.baseSalary,
		e.educationLevel, // calls EducationLevel.String(), if implemented
		e.birthDate.Format(time.DateOnly),
		e.employmentDate.Format(time.DateOnly),
		e.country,
	)
}

// GETTERS / SETTERS BOILERPLATE

func (e *Employee) GetCPR() string {
	return e.cpr
}

func (e *Employee) SetCPR(cpr string) error {
	if isValidCPR(cpr) {
		e.cpr = cpr
		return nil
	}
	return fmt.Errorf("invalid cpr %v", cpr)
}

func (e *Employee) GetFirstName() string {
	return e.firstName
}

func (e *Employee) SetFirstName(firstName string) error {
	if isValidName(firstName) {
		e.firstName = firstName
		return nil
	}
	return fmt.Errorf("invalid first name %v", firstName)
}

func (e *Employee) GetLastName() string {
	return e.lastName
}

func (e *Employee) SetLastName(lastName string) error {
	if isValidName(lastName) {
		e.lastName = lastName
		return nil
	}
	return fmt.Errorf("invalid last name")
}

func (e *Employee) GetDepartment() Department {
	return e.department
}

func (e *Employee) SetDepartment(dep Department) {
	e.department = dep
}

func (e *Employee) GetBaseSalary() float64 {
	return e.baseSalary
}

func (e *Employee) SetBaseSalary(sal float64) error {
	if 20000 <= sal && sal <= 100000 {
		e.baseSalary = sal
		return nil
	}
	return fmt.Errorf("%v out of bounds (20000 - 100000)", sal)
}

func (e *Employee) GetEducationLevel() string {
	return e.educationLevel.String()
}

func (e *Employee) SetEducationLevel(ed EducationLevel) {
	e.educationLevel = ed
}

func (e *Employee) GetBirthDate() time.Time {
	return e.birthDate
}

func (e *Employee) SetBirthDate(bday time.Time) error {
	if bday.Before(time.Now().AddDate(-18, 0, 0)) {
		e.birthDate = bday
		return nil
	}
	return fmt.Errorf("we don't hire anyone under 18")
}

func (e *Employee) GetEmploymentDate() time.Time {
	return e.employmentDate
}

func (e *Employee) SetEmploymentDate(empD time.Time) error {
	if time.Now().Before(empD) {
		diff := time.Until(empD)
		return fmt.Errorf("you cannot start in the future (starts in %d days)", int(diff.Hours()/24))
	}
	e.employmentDate = empD
	return nil
}

func (e *Employee) GetCountry() string {
	return e.country
}

func (e *Employee) SetCountry(country string) {
	e.country = country
}

func isValidName(name string) bool {
	re := regexp.MustCompile(`^[A-Za-z -]{1,30}$`)
	return re.MatchString(name)
}

func isValidCPR(cpr string) bool {
	re := regexp.MustCompile(`^[0-9]{10}$`)
	return re.MatchString(cpr)
}

// emp.employmentDate OF BORING STUFF

func (emp *Employee) GetSalary() float64 {
	return emp.baseSalary + (float64(emp.educationLevel) * 1220)
}

func (emp *Employee) GetDiscount() float64 {
	now := time.Now()
	years :=  now.Year() - emp.employmentDate.Year()
	if emp.employmentDate.Month() < now.Month() ||
		(emp.employmentDate.Month() == now.Month() && emp.employmentDate.Day() < now.Day()) {
		years--
	}

	return float64(years) * 0.5
}

func (emp *Employee) GetShippingCosts() float64 {
	switch strings.ToLower(emp.country) {
	case "denmark", "norway", "sweden":
		return 0
	case "iceland", "finland":
		return 0.5
	default:
		return 1
	}
}

func NewEmployee(
	cpr, firstName, lastName string,
	department Department,
	baseSalary float64,
	educationLevel EducationLevel,
	birthDate, employmentDate time.Time,
	country string,
) (*Employee, error) {

	emp := Employee{
		department:     department,
		educationLevel: educationLevel,
		country:        country,
	}
	// Manually calling setters on fields that require validation
	setters := []func() error{
		func() error { return emp.SetFirstName(firstName) },
		func() error { return emp.SetLastName(lastName) },
		func() error { return emp.SetBaseSalary(baseSalary) },
		func() error { return emp.SetBirthDate(birthDate) },
		func() error { return emp.SetEmploymentDate(employmentDate) },
		func() error { return emp.SetCPR(cpr) },
	}

	for _, set := range setters {
		if err := set(); err != nil {
			return nil, err
		}
	}

	return &emp, nil
}

func CallEmployeesFuncs() {
	emp, err := NewEmployee("1234567890", "Hornet", "Shaw", IT, 23456.78, Secondary,
		time.Date(1998, 5, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 2, 9, 0, 0, 0, 0, time.UTC),
		"Denmark",
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(emp)
}
