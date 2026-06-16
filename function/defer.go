package main

import "fmt"

func main() {
process(10)
}
func process(i int) {
	defer fmt.Println("Deffered i value:", i)

	defer fmt.Println("first Deferred Statement executed")
	defer fmt.Println("second Deferred Statement executed")
	defer fmt.Println("third Deferred Statement executed")
	fmt.Println("Normal execution statement")
fmt.Println("value of i", i)
}