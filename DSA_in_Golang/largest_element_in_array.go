//find the largest element in Array

package main

import "fmt"

func main() {

	array := [10]int{4, 6, 7, 4, 3, 12, 6, 5, 9, 10}
	larger, index := findLargest(array[:])

	fmt.Println("Larger value", larger, "index", index)

}
func findLargest(a []int) (int, int) {
	larger := a[0]
	var index int

	for i, v := range a {
		if v > larger {
			larger = v
			index = i
		}

	}
	return larger, index
}
