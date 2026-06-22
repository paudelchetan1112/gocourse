package main

import (
	"fmt"
	"os"
)

func main() {

	file,err:=os.Create("output.txt")
	if err!=nil{
		fmt.Println("Error creating file",err)
		return
	}
	defer file.Close()
	data:=[]byte("Hello world")
	n,err:=file.Write(data)

	if err!=nil{
		fmt.Println("Error while wriiting to file",err)
		return
	}
	fmt.Println(n)
	fmt.Println(data[:n])

	file, err=os.Create("writeString.txt")
		if err!=nil{
		fmt.Println("Error creating file",err)
		return
	}
	defer file.Close()
	_, err=file.WriteString("Hello Go!\n")
	if err!=nil{
		fmt.Println("Error while writing string",err)
		return
	}
	fmt.Println("Writing to writestring.txt complete")
}