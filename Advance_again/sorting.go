package main

import (
	"fmt"
	"sort"
)


type Person struct{
	name string
	age int
}
type By func(p1, p2 *Person) bool

type personSorter struct{
	people []personSorter

	by func(p1, p2 *Person) bool

}

func (s *personSorter) Len() int{
	return len(s.people)

}

func (s *personSorter) Less(i, j int) bool{
	return s.by(&s.people[i], &s.people[i])
	
}


type ByAge []Person
type ByName []Person

func (a ByAge) Len() int{
	return len(a)

}
func (a ByAge) Less(i, j int) bool {
	return a[i].age< a[j].age
}
func (a ByAge) Swap(i, j int){
	a[i], a[j]=a[j], a[i]
}

func (a ByName) Len()int{
return len(a)

}
func (a ByName) Swap(i, j int )  {
	a[i].name, a[j].name=a[j].name, a[i].name
}
func (a ByName) Less(i, j int )bool{
	return a[i].name<a[j].name 

}


func main() {
people:=[]Person{{name: "Ram", age: 20}, {name:"Hari", age:21}, {name: "sita", age: 22}}

sort.Sort(ByAge(people))
fmt.Println(people)
fmt.Println("Length:", )
// 	numbers := []int{4, 5, 3, 2, 4, 2, 5}
// 	sort.Ints(numbers)
// 	fmt.Println(numbers)
// stringSlice:=[]string{"John", "Anthony", "Anthemy", "Ram", "Shyam"}

// sort.Strings(stringSlice)
// fmt.Println(stringSlice)




}