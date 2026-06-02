package main

import "fmt"
const GRAVITY = 9.81


func main() {
	const PI = 3.14
	const days int =7
	const (
		monday=1
		tuesday=2
		wednesday=3
		thursday int =4

	)
	name:="john"  // ":= only available for variable not for constant"

	fmt.Println(name)
	var a,b,c int=2,3,4
	fmt.Println(a,b,c)

	fmt.Println(monday, tuesday, wednesday, thursday)
	fmt.Println(PI, GRAVITY)
}