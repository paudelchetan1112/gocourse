package main 
import "fmt"
func main (){
	//switch statement
var number int
	fmt.Println("Enter any number:");
	fmt.Scanln(&number)
	switch number {
	case 1:
		fmt.Println("Sunday")
	break
	
default:
	fmt.Println("Invalid input")
	}

}