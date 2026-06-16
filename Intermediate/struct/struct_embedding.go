package main

import "fmt"

type person struct {
	name string
	age  int
}
type Employee struct {
	EmployeeInfo person
	empId  string
	salary float64
}
func (p person) introduce(){
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.name, p. age)
}
func (e Employee) introduce(){
	fmt.Printf("Hi, I'm %s, employee Id : %s, and I earn %.2f.\n", e.EmployeeInfo.name,e.empId, e.salary)
}

func main() {
	emp := Employee{
		EmployeeInfo: person{
			name: "chetan",
			age:  20,
		},
		empId: "E01",

		salary: 250000,
	}
	fmt.Println("Name:", emp.EmployeeInfo.name)
	fmt.Println("Age:", emp.EmployeeInfo.age)
	fmt.Println("Emp Id:", emp.empId)
	fmt.Println("Salary:", emp.salary)
	emp.introduce()



}