package main

import "fmt"
type Shape struct{
	Rectangle
}
type Rectangle struct {
	length float64
	width  float64
}

//Method with value receiver

func (r Rectangle) Area() float64 {
	return r.length * r.width
}
//Method with pointer receiver 

func (r *Rectangle) Scale(factor float64){
r.length=r.length*factor
r.width=r.width*factor
}
func main() {
	rect := Rectangle{
		length: 10, width: 9,
	}
	area := rect.Area()
	rect.Scale(2)
	area=rect.Area();

	fmt.Printf("Area of rectangle witdth %.2f and length %.2f = %.2f",rect.length,rect.width, area)
num:=MyInt(-5)
num1:=MyInt(9)
fmt.Println(num.isPositive())
fmt.Println(num1.isPositive())
fmt.Println(num.welcomeMessage())
s:=Shape{
	Rectangle:Rectangle{
		length:15, 
		width: 20,
	},
}
fmt.Println(s.Area())
}
type MyInt int
func (m MyInt) isPositive() bool{
	return m>0
}


func (MyInt) welcomeMessage() string {
	return "Welcome MyInt Type"
}