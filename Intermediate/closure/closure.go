package main

import "fmt"

func main() {
sequence:=adder()
fmt.Println(sequence())
fmt.Println(sequence())
fmt.Println(sequence())
fmt.Println(sequence())
fmt.Println(sequence())
fmt.Println(sequence())
fmt.Println(sequence())
fmt.Println(sequence())
subtractor:=func() func(int)int {
	countDown:=99
	return func(x int)int{
		countDown-=x
		return countDown
	}
}()
fmt.Println(subtractor(1))
}

func adder() func() int {
	i := 0
	fmt.Println("previous value:",i)
	return func() int{
		i++
		 
		return i
	}
}