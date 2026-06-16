package main

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
)

func main() {
	// password := "password123"
	// hash := sha256.Sum256([]byte(password))
	// fmt.Println("plain password", password)
	// hash512:=sha512.Sum512([]byte(password))

	// fmt.Printf("SHA-256 Hash hex value %x\n",hash)
	// fmt.Printf("sha512:%x",hash512 )
}
func generateSalt() ([]byte, error){
	salt:=make([]byte, 16)
	_, err:=io.ReadFull(rand.Reader, salt)
	if err!=nil{
		return nil, err
	}

	
}