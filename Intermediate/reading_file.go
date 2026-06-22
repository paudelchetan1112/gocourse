package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file,err:=os.Open("output.txt")
		if err!=nil{
		fmt.Println("Error openingfile",err)
		return
	}
	defer func(){
		fmt.Println("Closing open file")
file.Close()
	}()
	
	fmt.Println("File opened successfully.")
data:=make([]byte, 1024)
_,err=file.Read(data)
	if err!=nil{
		fmt.Println("Error reading",err)
		return
	}
	fmt.Println("File content:", string(data))
scanner:=bufio.NewScanner(file)
for scanner.Scan(){
	line:=scanner.Text()
	fmt.Println(line)

}

}