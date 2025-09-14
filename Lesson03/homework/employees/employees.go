package employees

import (
	"fmt"
	"log"
	"regexp"
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

func (d Department) String() string {
	switch d {
	case HR:
		return "HR"
	case Finance:
		return "finance"
	case IT:
		return "IT"
	case Sales:
		return "sales"
	case GeneralServices:
		return "general services"
	}
	return "unknown"
}

func (d EducationLevel) String() string {
	switch d {
	case Primary:
		return "primary"
	case Secondary:
		return "secondary"
	case Tertiary:
		return "tertiary"
	}
	return "none"
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

func isValidName(name string) bool {
	re := regexp.MustCompile(`^[A-Za-z -]{1,30}$`)
	return re.MatchString(name)
}

// GETTERS / SETTERS BOILERPLATE

func (e *Employee) GetCPR() string {
	return e.cpr
}

func (e *Employee) SetCPR(cpr string) {
	e.cpr = cpr
}

func (e *Employee) GetFirstName() string {
	return e.firstName
}

func (e *Employee) SetFirstName(firstName string) error {
	if isValidName(firstName) {
		e.firstName = firstName
		return nil
	}
	return fmt.Errorf("invalid first name")
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
	if time.Now().Year()-bday.Year() >= 18 {
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
		return fmt.Errorf("you cannot start in the future (starts in %d hours)", int(diff.Hours()/24))
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

// func setCountry(e *Employee, country string) {
// 	e.country = country
// }

// END OF BORING STUFF

func getSalary(emp *Employee) float64 {
	return emp.baseSalary + (float64(emp.educationLevel) * 1220)
}

func getDiscount(emp *Employee) float64 {
	return float64(time.Now().Year()-emp.employmentDate.Year()) * 0.5
}

func getShippingCosts(emp *Employee) float64 {
	switch emp.country {
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
		cpr:            cpr,
		department:     department,
		educationLevel: educationLevel,
		country:        country,
	}
	// Manually call the setters on the fields with validation. That is, first name, last name, base salary, date of birth, date of employment

	setters := []func() error {
		func() error { return emp.SetFirstName(firstName) },
		func() error { return emp.SetLastName(lastName) },
		func() error { return emp.SetBaseSalary(baseSalary) },
		func() error { return emp.SetBirthDate(birthDate) },
		func() error { return emp.SetEmploymentDate(employmentDate) },
	}

	for _, set := range setters {
		if err := set(); err != nil {
			return nil, err
		}
	}

	return &emp, nil
}

func CallEmployeesFuncs() {
	emp, err := NewEmployee("12345678", "Hornet", "Shaw", IT, 23456.78, Secondary,
		time.Date(1998, 5, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 2, 9, 0, 0, 0, 0, time.UTC),
		"Denmark",
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(emp)
}
