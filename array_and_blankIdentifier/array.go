package main

import "fmt"

func main() {
	//var arrayName [size] element type
// 	var numbers [5]int
// 	var b int
// 	_=b
// 	fmt.Println(numbers)
// 	numbers[4]=20
// 	fmt.Println(numbers)
// 	numbers[0]=25
// 	fmt.Println(numbers)


// 	fruits:=[4]string{"Apple", "Banana", "Orange", "Grapes"}
// 	fmt.Println(fruits)
// 	fruitSlice:=[4]string{"Apple","Banana", "Orange", "Grapes"}
// 	fmt.Println(fruitSlice)
// 	fmt.Println("Third element is:",fruits[2])

// originalArray:=[3]int{1,2,3}
// copiedArray:=originalArray
// copiedArray[0]=100
// fmt.Println("Original Array",originalArray)
// fmt.Println("copiedArray",copiedArray)
// for i:=0;i<len(numbers);i++{
// 	fmt.Println("Element at index ",i,"is:",numbers[i])
// }
// // for index, value:=range numbers{
// // 	fmt.Printf("Index:%d Value:%d\n",index,value)
// // }

// for _, value:=range numbers{
// 	fmt.Println(value)
// }

// //length of the array
// fmt.Println("The length of the number array is:", len(numbers))
// //comparing two array
// array1:=[3]int{10,2,3}
// array2:=[3]int{1,2,3}
// fmt.Println("Array1 is equal to Array2:",array1==array2)

// //multidimensional array

// var matrix[3][3]int =[3][3]int{
// 	{1,2,3},
// 	{4,5,6},
// 	{7,8,9},
// }
// fmt.Println(matrix)
array1:=[3]int{10,2,3}
var array2 *[3]int
array2=&array1
fmt.Println(*array2)
}