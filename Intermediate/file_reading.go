package main

import (
	"fmt"
	"os"
)

func main() {

	file,err:=os.Open("chetan.txt")
	if err!=nil{
		fmt.Println("Error while opening file:", err)

	}
	data:=make([]byte,100)

	n, err:=file.Read(data)
	
	if err!=nil{
		fmt.Println("Error while opening file:", err)

	}
	fmt.Println(string(data[:n]))
	
	
}