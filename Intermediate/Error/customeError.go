package main

import (
	"errors"
	"fmt"
)

func main() {
	err := doSomething()
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("Operation completed successfully !.... ")
}

type customError struct {
	code    int
	message string
	err     error
}

// Error returns the error message . Implementing Error() method of error interface
func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.code, e.message)

}

//Function that return a custome error
// func doSomething() error {
// 	return &customError{code: 500, message: "Something went wrong"}

// }

func doSomething() error{
	err:=doSomethingElse()
	if err!=nil{
		return &customError{code:500, message:"Something went wrong", 
	err:err, }
	}
}
func doSomethingElse() error{
	return errors.New("Internal Error")
}