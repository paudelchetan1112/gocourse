package main

import (
	"fmt"
	"strings"
)

func main() {
	var string1 string= "go is fun and go is fast"
word:=strings.Split(string1, " ")
fmt.Println(word)
count:=make(map[string]int)
for _, w:=range word{
count[w]++
}
for k, v:=range count{
	fmt.Println(k,":",v)
	
}

}