package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	age       int
	Address
	PhoneHomeCell
}

func (p *person) fullName() string {
	return p.firstName + " " + p.lastName
}
func (p *person) incrementAgeByOne() {
	p.age++
}

type Address struct {
	city    string
	country string
}
type PhoneHomeCell struct{
	home string 
	cell string
}

func main() {

	p := person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
		Address: Address{
			city:"London", 
			country:"UK",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "2535328525",
			cell:"253532535385", 
		},
	}

	p1 := person{
		firstName: "chetan",
		lastName:  "paudel",
	}
	fmt.Println(p)

	fmt.Println(p1.firstName)

	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "pseudoemail@gmail.com",
	}
	fmt.Println(user.email)
	fmt.Println(p.fullName())
	fmt.Println("Before Increment:", p1.age)

	p1.incrementAgeByOne()
	fmt.Println("After Increment", p1.age)
	fmt.Println(p.city)


}
