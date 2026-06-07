/* FMT Package
	Printing Function
		Print()
		Println()
		Printf()
	Formatting Functions
		Sprint()
		Sprintf()
		Sprintln()
	Scanning Functions
		Scan()
		Scanf()
		Scanln()
	Error Formatting Functions
		Errorf()



*/


package main
import "fmt"
func main(){

	//Printing Action
// 	fmt.Print("Hello")
// 	fmt.Print("World! ")
// 	fmt.Print(12,456)


// 	fmt.Println("Hello")
// 	fmt.Println("World! ")
// 	fmt.Println(12,456)

// 	name:= "john"
// 	age:=25
// 	fmt.Printf("Name:%s, Age:%d", name, age)
// 	fmt.Printf("Binary: %b, Hex:%x\n", age, age)



// 	//Formatting function

// 	s:=fmt.Sprint("Hello", "World!", 123, 456 )
// 	fmt.Print(s)
// 	s=fmt.Sprintln("Hello", "world!", 123, 456)
// 	fmt.Print(s)


// sf:=fmt.Sprintf("Name:%s, Age:%d", name, age)
// fmt.Println(sf)



//Scanning Functions

var name string
var age int
fmt.Print("Enter your name and age:")
// fmt.Scan(&name,&age)
// fmt.Scanln(&name,&age)   //it stop scan in next line
fmt.Scanf("%s %d", &name, &age)


fmt.Printf("Name:%s, Age:%d", name, age)

//Error Formatting Function

error:=checkAge(15)
if error!=nil{
	fmt.Println(error)
}





}
func checkAge(age int) error{
	if age<18{
		return fmt.Errorf("Age %d is too young to drive", age)

	}
	return nil
}