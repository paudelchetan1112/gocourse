package main

import (
	"fmt"
	
)

func main() {

	// //var aarayName [size]elementType
	// var numbers [5]int
	// fmt.Println(numbers)
	// numbers[4]=10
	// fmt.Println(numbers)
	// numbers[len(numbers)-1]=11
	// fmt.Println(numbers)

	// numbers[0]=9
	// fmt.Println(numbers)

// 	// var fruits =[4]string{"apple","banana", "orange", "grape"}
// numbers:=[]int{1,2,3,4,5}
// 	orginalArray:=[3]int {1,2,3}
// 	copiedArray:=orginalArray
// 	copiedArray[0]=100
// 	fmt.Println(orginalArray)
// 	fmt.Println(copiedArray)
// 	for i:=0; i<len(numbers);i++{
// 		fmt.Println("Element at index,", i, ":", numbers[i])
// 	}
// 	//underscore is blank identifier and used to store unused value
// for _, value:=range numbers{
// 	fmt.Printf(" value:%d\n",  value)
// a,_:=someFunction()
// fmt.Println(a)
// }

// //comparing Arrays

// array1:=[3]int{1,2,3}
// array2:=[3]int{1,2,3}
// fmt.Println("Array1 is equal to array2:", array1==array2)

// var matrix =[3][3]int {
// 	{1,2,3},{1,2,3},{1,2,3},
// }
// fmt.Println(matrix)

orginalArray:=[3]int{1,2,3}
var copiedArray *[3]int
copiedArray=&orginalArray
fmt.Printf("%v",copiedArray)
fmt.Println(*copiedArray)
}
func someFunction()(int, int){
	return 1, 2
}