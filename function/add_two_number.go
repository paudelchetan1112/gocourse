package main

import "fmt"

func main() {
	var num1, num2 int
	num1 = 50
	num2 = 60
	result := int(add(num1, num2))
	fmt.Println("Addition of two number is", result)
	result2 := int(sub(num1, num2))

	fmt.Println("Subtraction of two number is:", result2)
}
func add(x int, y int) int {
	return x + y
}
func sub(x int, y int) int {
	return x - y
}
