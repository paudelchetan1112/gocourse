package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() float64
}
type fullTimeEmployee struct {
	name          string
	monthlySalary float64
}
type contractEmployee struct {
	name        string
	HoursWorked int
	HourlyRate  float64
}

func (e fullTimeEmployee) CalculateSalary() float64 {
	return e.monthlySalary
}
func (e contractEmployee) CalculateSalary() float64 {
	return e.HourlyRate*float64(e.HoursWorked)
}

type payroll struct {
	Employee SalaryCalculator
}

func (p payroll) ProcessSalary()float64 {
	return p.Employee.CalculateSalary()
}

func main() {
	employee := payroll{Employee: fullTimeEmployee{name: "chetan", monthlySalary: 20000}}

	fmt.Println("The salary will be processed:",employee.ProcessSalary())


}