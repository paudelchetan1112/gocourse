package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math error : Square root of negative number")

	}
	//compute the square root
	return math.Sqrt(x), nil
}

func process(data []byte) error{
	if len(data)==0{
		return errors.New("Error: Empty data")
	}
	return nil 
}
func main(){
// 	result, error:= sqrt(16)
// 	if error!=nil{
// 		fmt.Println(error)
// 	return 
// 	}
// 	fmt.Println(result)
// result, error= sqrt(-16)
// 	if error!=nil{
// 		fmt.Println(error)
// 	return 
// 	}
// 	fmt.Println(result)
	
// 	data:= []byte{}
// 	err:=process(data)
// 	if err!=nil{
// 		fmt.Println("Error:", err)
// return 
// 	}
// 	fmt.Println("Data processed successfully ")

// err1:=eprocess();
// if err1!=nil{
// 	fmt.Println(err1)
// }


err:=readData()
if err!=nil{
	fmt.Println(err)
return 
}


}



type myError struct {
	message string

}
func (m *myError) Error()string{
	return fmt.Sprintf("Error:%s", m.message)

}
func eprocess() error{
	return &myError{"custome error message"}
}
func readData() error{
	error:=readConfig()
	if error!=nil{
		return fmt.Errorf("readdata: %w", error)
	}
	return nil
}
func readConfig() error{
	return errors.New("config error")
}