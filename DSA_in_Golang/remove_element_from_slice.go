package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50}
	removeelement := 30
	var removeIndex int
	for i, v := range slice {

		if v == removeelement {
			removeIndex = i
		}

	}
	slice = append(slice[:removeIndex], slice[removeIndex+1:]...)
	fmt.Println(slice)

}
