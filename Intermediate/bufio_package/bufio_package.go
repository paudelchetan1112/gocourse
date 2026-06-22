package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

func main() {
// 	reader:=bufio.NewReader(strings.NewReader("Hello, bufioPackage\n How are you doing?"))

// 	//Reading byte slice
// 	data:=make([]byte, 20)
// 	n, err:=reader.Read(data)
// 	if err!=nil{
// 		fmt.Println("Error reading:", err)
// return
// 	}
// 	fmt.Printf("Reading %d bytes:%s", n, data[:n])


// 	line, err:=reader.ReadString('\n')

// if err!=nil{
// 	fmt.Println("Error reading string:", err)
// 	return
// }
// fmt.Println("Read string:", line)

writer:=bufio.NewWriter(os.Stdout)
//wirting byte slice
data:=[]byte("Hello bufio package")
n, err:=writer.Write(data)
if err!=nil{
	fmt.Println("error while writing", err)

}
fmt.Println(n)
err=writer.Flush();
if err!=nil{
	fmt.Println("Error flushing writer:",err)
}
//writing string
str:="This is a string"
n, err=writer.WriteString(str)
if err!=nil{
	fmt.Println("Error while writing string", err)
	return
}
fmt.Println("Wrote", n, "data")



}