package main

import "fmt"

func main() {
	num := 4234534525
	fmt.Printf("%05d\n", num)
	message:="Hello"
	fmt.Printf("|%10s|", message)
	fmt.Printf("|%-5s|\n", message)
	message1:="Hello \nworld!"
	message2:=`Hello \nworld!`
	fmt.Println(message1)
	fmt.Println(message2)
	sqlQuery:=`SELECT * FROM users WHERE age>30`
	fmt.Println(sqlQuery)

}