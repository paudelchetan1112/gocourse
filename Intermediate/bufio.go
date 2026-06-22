package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	reader := bufio.NewReader(strings.NewReader("Hello this is chetan\n4rghxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"))
	// data:=make([]byte, 30)
// n, err:=reader.Read(data)
// if err!=nil{
// 	fmt.Println("Error while reading", err)

// }
// fmt.Println(string(data[:n]))
line, err:=reader.ReadString('\n')
if err!=nil{
	fmt.Println("Error", err)


}
fmt.Println(line)




}