package main

import "fmt"

func main() {
	//panic(interface{})
	//example of valid input
	process(10)
	process(-3)

}

func process(input int) {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred2")

	if input < 0 {
		panic("Input must be a non-negative number")
		
// defer fmt.Println("deferred 3")
	}
	fmt.Println("Processing input:", input)

}