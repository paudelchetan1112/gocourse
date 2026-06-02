package main

import "fmt"
// middleName:="cane" only used := inside the function locally "local variable initialization"
var middleName string="cane"

func main() {
	var age int
	age = 111
	var name string ="john"
	var name1 string="jane"
	count:=10


	fmt.Println(age, name, name1,count)
	//default values:
	//for Numeric Type:0
	//Boolean Type:false
	//String type: ""
	//pointer, slices, maps, functions, and struts:nil
//---------scope
printname();
fmt.Println(middleName)

}
func printname(){
	firstName:="Michael"
	fmt.Println(firstName)

}