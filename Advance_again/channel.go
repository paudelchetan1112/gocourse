package main

import (
	"fmt"
	"time"
)

func Channel() {
	greeting := make(chan string)
	// greetString := "Hello"
	go func() {
		// greeting <- greetString //blocking because it is continuously trying to receive value, it is ready to receive continuous flow of data
		// greeting<-"chetan"
		for _,e:=range "abcde"{
			greeting<-"Alphabet "+string(e)

			
		}

	}()  
	// go func() {
	// 		// value := <-greeting
	// 		// fmt.Println(value)
	// 		// 	value=<-greeting
	// 		// fmt.Println(value)

	// }()

	for range 5{
		value:=<-greeting
		fmt.Println(value)
		time.Sleep(10*time.Millisecond)
	}


	
	
	
	time.Sleep(100*time.Millisecond)
	fmt.Println("End of the program")


}