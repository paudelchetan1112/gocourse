package main

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "12345"
	num, err:=strconv.Atoi(numStr)
	if err!=nil{
		fmt.Println("Error parsing the value:", err)

	}
	fmt.Println("Parsed Integer:", num)
	fmt.Println("parsed Integer:", num+1)

numistr, err:=strconv.ParseInt(numStr, 10, 32)  // base(10, 2, 16) and bit size (32/64)
if err!=nil{
	fmt.Println("Error parsing the value:", err)

}
fmt.Println("Parsed Integer:", numistr)

floatstr:="3.14"
floatval, err:=strconv.ParseFloat(floatstr, 64)
if err!=nil{
	fmt.Println("Error parsing the value:", err)

}
fmt.Println(floatval)


binarystr:="1010"
decimal,err:=strconv.ParseInt(binarystr, 2, 64)
if err!=nil{
	fmt.Println("Error parsing the value:", err)

}
fmt.Println(decimal)
	

hexstr:="FF"
hex,err:=strconv.ParseInt(hexstr, 16, 64)
if err!=nil{
	fmt.Println("Error parsing the value:", err)

}
fmt.Println(hex)
	

	
}



// invalidNum := "345lkdj"
// invalid,err:=strconv.Atoi(invalidNum)
// if err!=nil{
// 	fmt.Println("Error parsing the value:", err)
// 	return

// }
fmt.Println(invalid)
	
