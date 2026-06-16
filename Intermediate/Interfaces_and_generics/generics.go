package main

import "fmt"

// func swap[T any](a, b T) (T, T) {
// 	return b, a
// }

type Stack[T any] struct {
	elements []T
}
func (s *Stack[T] ) push (element T){
	s.elements=append(s.elements, element)

}
func (s *Stack[T]) Pop() (T, bool) {
    if len(s.elements) == 0 {
        var zero T
        return zero, false
    }
    element := s.elements[len(s.elements)-1]
    s.elements = s.elements[:len(s.elements)-1] // remove last element
    return element, true
}

func (s *Stack[T]) isEmpty() bool{
	return len(s.elements)==0
}
func(s *Stack[T]) printAll(){
	if len(s.elements)==0 {
		fmt.Println("The stack i smepty. ")
	}
	fmt.Println("Stack element:") 
	for _, element := range s.elements{
		fmt.Println(element)

	}
}

func main() {
// 	x, y := 1, 2
// 	fmt.Println("Before swap:", x, y)
// 	x, y = swap(x, y)
// 	fmt.Println("After swap:", x,y)
// x1,y1:="john", "jane"
// 	fmt.Println("Before swap:", x1, y1)
// 	x1,y1=swap(x1,y1)
// 		fmt.Println("After swap:", x1, y1)
intStack:=Stack[int]{}
intStack.push(50)
intStack.push(40)
intStack.push(60)
intStack.printAll()
fmt.Println(intStack.pop())
intStack.printAll()



stringStack:=Stack[string]{}
stringStack.push("Hello")
stringStack.push("world")
stringStack.push("John")
stringStack.printAll()
stringStack.pop()
stringStack.printAll()

}