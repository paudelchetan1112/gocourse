package main

import (
	"fmt"
	foo "net/http"
)

func main() {
	fmt.Println("Hello, Go Standard Library")
	resp, err := foo.Get("http://jsonPlaceholder.typicode.com/posts/1")

	if err != nil {
		fmt.Println("Error:", err)
		return

	}
	defer resp.Body.Close()
	fmt.Println("Http Response Status:", resp.status)

}
