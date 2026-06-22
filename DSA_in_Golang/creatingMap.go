package main

import "fmt"

func main() {
	student := make(map[string]int)
	student["Alice"] = 90
	student["Bob"] = 85
	student["charlie"] = 95
	for k, v:=range student{
		fmt.Println(k, ":", v)
	}
	_, ok := student["chetan"]
	if ok{
		fmt.Println("Chetan is exist")

	}else{
		fmt.Println("not exist")

	}


}