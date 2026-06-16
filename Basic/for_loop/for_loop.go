package main

import "fmt"

func main() {

	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)

	// }
	//iterate over collection
	// numbers:= []int {1,2,3,4,5}
	// for index, value:=range numbers{
	// 	fmt.Printf("Index:%d, Value:%d\n", index,value)

	// }

	// for i:=1; i<=10;i++{
	// 	if i%2==0{
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// 	if(i==5){
	// 		break
	// 	}
	// }

	rows:=5 
	for i:=1;i<=rows;i++ {
		for j:=i;j<=rows;j++ {
			fmt.Print(" ")

		}
		for j:=1;j<=2*i-1;j++ {
			fmt.Print("*")
		}
	fmt.Printf("\n")
	}



}