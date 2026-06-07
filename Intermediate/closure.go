package main

import "fmt"

func main() {
// sequence:=adder()
// fmt.Println(sequence())
// fmt.Println(sequence())
// fmt.Println(sequence())
// fmt.Println(sequence())
// fmt.Println(sequence())
// fmt.Println(sequence())
// fmt.Println(sequence())
// fmt.Println(sequence()) 
subtraction:=func() func(int) int{
	countdown:=99
	return func(x int)int{
countdown-=x
return countdown
	}
}()
//using the closure subtracter
fmt.Println(subtraction(1))
fmt.Println(subtraction(1))
fmt.Println(subtraction(1))
fmt.Println(subtraction(1))
fmt.Println(subtraction(1))
}

// func adder() func() int {
// 	i := 0
// 	fmt.Println("previous value of i:", i)
// 	return func() int {
// 		i++
// 		fmt.Println("Added 1 to i")
// 		return i
// 	}
// }