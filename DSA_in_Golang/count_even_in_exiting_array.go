package main

import "fmt"

func main() {
	array := []int{5, 4, 5, 7, 5, 4, 3, 2}
	count := 0
	for _, v := range array {
		if v%2 == 0 {
			count++
		}
	}
	fmt.Println("no of even number in the given array is", count)

}
