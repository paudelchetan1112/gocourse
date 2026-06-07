package main

import "fmt"

func main() {
	//var sliceName[]elementType
	// var numbers []int
	// var numbers1 = []int{1, 2, 3}
	// numbers2 := []int{9, 8, 7}
	// a := [5]int{1, 2, 3, 4, 5}
	// slice := a[1:4]
	// fmt.Println(slice)
	// slice1:=append(slice, 6,7)
	// fmt.Println(slice1)
	// sliceCopy:=make([]int, len(slice1))
	// copy(sliceCopy, slice1)

	// fmt.Println(sliceCopy)
	// // var nilSlice []int
	a := [5]int{1, 2, 3, 4, 5}
	slice := a[1:4]
	fmt.Println(slice)
	slice1:=append(slice, 6,7)
	for i, v:=range slice1{
		fmt.Println(i, v)

	}


}