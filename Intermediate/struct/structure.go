package main

import "fmt"

// import "fmt"
type method interface{
	area() 
	perimeter() 
}

type rectangle struct {
	length float64
	width  float64
}

func (r rectangle) area()  {
	fmt.Println("Area of rectangle:",r.length * r.width) 
} 
type square struct{
	length float64
}
func (s square) area() float64{
	return s.length*s.length
}

func methodrunner(a method){
	a.area()
	a.perimeter()
}
func (r rectangle) perimeter(){
	fmt.Println("Perimeter of rectangle:",2*(r.length+r.width)) 

}

func main() {
	rectangle1 := rectangle{length: 20.345, width: 20.234}

methodrunner(rectangle1)





}