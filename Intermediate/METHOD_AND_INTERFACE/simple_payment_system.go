package main

import "fmt"

type paymentMethod interface {
	pay(amount float64)
}
type Esewa struct {
	username string
	password string
}

func (e Esewa) pay(amount float64) {
	fmt.Printf("%.2f amount is paid through Esewa", amount)
}


type Khalti struct {
	username string
	password string
}
func (e Khalti) pay(amount float64) {
	fmt.Println(amount, "amount is paid through Khalti")
}
type customer struct {
	c_id           int
	customer_name  string
	paymentGateway paymentMethod
}
func ( c customer) makePayment(amount float64){
	c.paymentGateway.pay(amount)
	
}

func main() {
	customer1:=customer{c_id: 101, customer_name: "chetan", paymentGateway: Esewa{username: "98680934", password:"lskj"}}
	amount:=23324.34
	customer1.makePayment(amount)

}