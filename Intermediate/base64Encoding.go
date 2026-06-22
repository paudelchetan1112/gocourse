package main

import (
	"encoding/base64"
	"fmt"

)

func main() {

	data := []byte("Hello, Base64 Encoding")
	//Encode Base 64

	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)
	decoded, err:=base64.StdEncoding.DecodeString((encoded))
if err!=nil{
	fmt.Println("Decoding error", err)

}
fmt.Println(decoded)
fmt.Printf("decoded string %s", decoded)

//URL safe, avoid '/' and '+'
urlSafeEncoded:=base64.URLEncoding.EncodeToString(data)

fmt.Println("url safe encoded:", string(urlSafeEncoded))
}