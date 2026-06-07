package main

import "fmt"

func main() {
sequence:=adder()
fmt.Println(sequence())
fmt.Println(sequence())
fmt.Println(sequence())
fmt.Println(sequence())
}
func adder() func() int{
	i:=0
	fmt.Println("Initialized value:", i)
	return func() int {
		i++
		return i
	}


}