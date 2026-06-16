package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	user := os.Getenv("USER")
	home:=os.Getenv("HOME")
	fmt.Println(user)
	fmt.Println(home)

	os.Setenv("Fruit", "Apple")
	for _, e=range os.Environ(){
		kvpair:=strings.SplitN()
	}
}