//string are immutable, sequence of byte

package main

import (
	"fmt"
	"unicode/utf8"
)
func main(){
	message:= "Hello,\n Go!"
	rawMessage:=`Hello\nGo`
	message1:="Hello\rGo!" //Go!lo
	fmt.Println(message)
	fmt.Println(rawMessage)
	fmt.Println(message1)
fmt.Println("Length of message variable is", len(message))
fmt.Println("Lenght of the message 2 variable is:", len(rawMessage))
fmt.Println("The first character in message var is", message[0])
greeting:="Hello"
name:="Alice"
fmt.Println(greeting+" "+name)

str1:="Apple"
str2:="banana"
str3:="app"
str:="apple"

fmt.Println(str1<str2)
fmt.Println(str3<str1)
fmt.Println(str>str1)
fmt.Println(str>str3)

// for i, char:=range message{
// 	fmt.Printf("Character at index %d is %c\n", i, char)
// 	fmt.Printf("%x\n", char)

// }
fmt.Println("Rune count:", utf8.RuneCountInString(greeting))
greetingwithname:=greeting+name
fmt.Println(greetingwithname)
var ch rune='a'
nch:='क'
fmt.Println(ch, nch)
fmt.Printf("%c",nch)
cstr:=string(nch)
fmt.Printf("Type of cstr is %T", cstr)
fmt.Println(cstr)

const NEPALI="नेपाली" //nepali language
fmt.Println(NEPALI)
for _, char :=range NEPALI{
	fmt.Printf("%c\n",char)
}
}