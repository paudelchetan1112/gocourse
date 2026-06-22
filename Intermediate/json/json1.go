package main

import (
	"encoding/json"
	"fmt"
)
type person struct{
	Name string `json:"name"`
	Age int `json:age`

}

func main() {
	person1:=person{Name:"Chetan", Age:10}
	data, err:=json.Marshal(person1)
	if err!=nil{
		fmt.Println("Error while marshing json", err)

	}
	fmt.Println(string(data))


}