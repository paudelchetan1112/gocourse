//using recursion function , find the factorial of the given number
package main

import "fmt"

func main() {
	fmt.Println("Factorial of 5 is:", factorial(5))
}
func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		//recursive case: factorial of n is n*factorial(n-1)

		return n * factorial(n-1)

	}
}