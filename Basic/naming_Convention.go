package main

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}
type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	//PascalCase
	//eg: calculateArea, UserInfo, NewHTTPRequest
	//Structs, interface enums

	//snake_case
	//eg: user_id, first_name, http_request
	//variable which are multiple word

	//UPPERCASE
	//used for declaring constant

	//mixedCase:
	//Eg. javaScript, htmlDocument, isValid
	const MAXRETRIES = 5
	var employeeId = 1001
	fmt.Println("Employee Id:", employeeId)

}
