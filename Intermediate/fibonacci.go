//printing fibonacci using closure
package main

import "fmt"

func main() {
	fmt.Println(0)
	fmt.Println(1)
	fib := fibonacci()
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
}
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		result :=a+b
		a = b
		b = result
		return result
	}
}