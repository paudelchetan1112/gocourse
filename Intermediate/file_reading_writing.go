package main

import (
	"fmt"
	"os"
)

func main() {
	file, err:=os.Create("chetan.txt")
defer fmt.Println("File close successfully")
	defer file.Close()
	
	if err!=nil{
		fmt.Println("Error while creating file,", err)

	}
data:=[]byte("Hello my name is chetan paudel")

n, err:=file.Write(data);
if err!=nil{
	fmt.Println("Error while writing data into the file",err )
return
}
fmt.Println("data which is write into the file is:", string(data[:n]))

n, err=file.WriteString("This is a string which is added to the file");


if err!=nil{
	fmt.Println("Error while writing string to the file",err )
return 


}
fmt.Println("string writing successfully");




}
