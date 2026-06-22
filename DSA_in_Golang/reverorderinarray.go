// Create an array and print its elements in reverse order.
package main

import "fmt"


func main(){
	var array[5] int =[5]int{5,6,7,8,9}
	// for i,v:=range array{
	// 	defer fmt.Println("index:",i, "value:",v)

	// }
	for i:=len(array)-1; i>=0;i-- {
		fmt.Println("index:",i, "value:",array[i])
	}


}