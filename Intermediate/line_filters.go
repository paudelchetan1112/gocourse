package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

file,err:=os.Open("example.txt")
if err!=nil{
	fmt.Println("Error openiing file:", err)
	return 
}
defer file.Close()
scanner:=bufio.NewScanner(file)
//keyword to filter lines
keyword:="important"

//Read and filter lines

for scanner.Scan(){
	line:=scanner.Text()
	if strings.Contains(line, keyword){
		updateLine:=strings.ReplaceAll(line, keyword, "necessary")

	fmt.Println("filtered Line:", line)
	fmt.Println(updateLine)


	}
}
err=scanner.Err()
if err!=nil{
	fmt.Println("Error ",err)
}

}