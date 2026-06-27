package main

import (
	"fmt"
	"time"
)



func main(){
	//==============BLOCKING ON RECEIVE ONLY IF THE BUFFER IS EMPTY

	ch:=make(chan int, 2)
	ch<-1
	ch<-2
	fmt.Println("start")

	go func() {
		time.Sleep(10*time.Second)
		fmt.Println("received:",<-ch)
		fmt.Println("Blocking end")



	}()

	fmt.Println("Blocking start:")
	ch<-3
	// fmt.Println("end of program")
	// fmt.Println("value received:", <-ch)


	// fmt.Println("End of program")

}

// func main() {
// 	//make(chan Type, capacity)
// 	ch := make(chan int, 2)
// 	ch <- 2
// 	ch<-3
// 	// ch<-4 
// 	fmt.Println(<-ch)
// 	ch<-4
// 	fmt.Println(<-ch)
// 	ch<-5
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// 	// fmt.Println(<-ch)


// }