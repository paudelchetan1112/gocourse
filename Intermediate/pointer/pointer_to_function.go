//Pointer

package main

import "fmt"

func main() {
	// var ptr *int
	var a int = 10
	// ptr = &a
	modifyValue(&a)
	fmt.Println(a)
}
func modifyValue(ptr *int) {
	*ptr++
}