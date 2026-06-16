package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	reader:=bufio.NewReader(strings.NewReader("Hello, bufio packageee!\n How are you doing"))

data:=make([]byte, 20)
n, err:=reader.Read(data)
if err!=nil{
	fmt.Println("Error reading:", err)
	return 
}
fmt.Printf("Read %d bytes:%s", n, data[:n])
fmt.Println(len(data))

str, err:=reader.ReadString('\n')
if err!=nil{
	fmt.Println("Error reading:", err)

return
}
fmt.Printf("Read String:%s", str)


	
}