package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FullName string `json:"name"`

	Age          int     `json:"age,omitempty"`
	EmailAddress string  `json:"email,omitempty"`
	Address      Address `json:address`
}
type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {
	person := Person{FullName: "john", Age: 30}
	//Marshalling
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to json:", err)
		return
	}
	fmt.Println(string(jsonData))
	person1 := Person{FullName: "chetan", Age: 30, EmailAddress: "example@gmail.com", Address: Address{City: "nepalgunj", State: "lumbini"}}
	jsondata1,err1:=json.Marshal(person1)
	if err1!=nil{
		fmt.Println(err1)

	}
	fmt.Println(string(jsondata1))
	jsonDatas:= `{"full_name":"johny doe", "emp_id":"009", "age":"30", "address":{"City":"san jose", "state":"CA"}}`
	unmarshData:=json.Unmarshal([]byte(jsonDatas), )
	fmt.Println(unmarshData)







}
type Employee struct{
	FullName string `json:"full_name"`
	EmpId string `json:"emp_id"`
	Age int `json:"age"`
	Address Address `json:"address"`
}
