package main

import (
	"fmt"
	// "time"
	// "time"
)



func main(){
	ch:=make(chan int)

	go func() {
		ch<-1
		close(ch)

		
	}()

	for{
		select{
		case msg,ok:=<-ch:
			if !ok{
				fmt.Println("Channel closed")
				//clean up activities

				return
			}
			fmt.Println("received:",msg)
		}
	}
}




// func main(){
// 	ch:=make(chan int)
// 	go func() {
// 		time.Sleep(time.Second)
// 		ch<-1

// 	}()
// 	select{
// 	case msg,ok:=<-ch:
// 		if !ok{
// 			fmt.Println("not received")
// 		}
// 		fmt.Println("Received:",msg)
// 	case <-time.After(3*time.Second):
// 		fmt.Println("Timeout.")

// 	}
// }

// func main() {

// 	ch1 := make(chan int)
// 	ch2 := make(chan int)
// 	go func() {
		
// 		ch2<-2
// 	}()
// 	go func() {
		
// 		ch1<-1
// 	}()
// // time.Sleep(time.Second)
// for range 2{
// select{
// case msg:=<-ch1:
// 	fmt.Println("Received from ch1",msg)
// case msg:=<-ch2:
// 	fmt.Println("received from ch2:",msg)
// // default:
// // 	fmt.Println("no received")

// }
// }


// }