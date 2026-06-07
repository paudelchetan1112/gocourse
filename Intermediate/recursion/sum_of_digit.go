//using recursion function, find the sum of the digit in the given number.

package main

import "fmt"

func main() {
	fmt.Println(sumOfDigit(45))
}
func sumOfDigit(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + sumOfDigit(n/10)
}