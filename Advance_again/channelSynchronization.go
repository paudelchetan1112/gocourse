package main

import (
	"fmt"
	"time"
)

// func main() {
// 	done := make(chan bool)
// 	go func() {
// 		fmt.Println("Working..")
// 		time.Sleep(2*time.Second)
// 		done<-false

// 	}()

// 	<-done
// 	fmt.Println("Finished")

// }


// func main(){
// 	ch:=make(chan int)
// 	go func ()  {
// 		ch<-9
// 		time.Sleep(time.Second)
// 		fmt.Println("sent value")

		
// 	}()
// 	value:=<-ch //Blocking until the value is received
// 	fmt.Println(value)
// }


//SYNCHRONIZATION MULTIPLE GOROUTINE AND ENSURING THAT ALL GOROUTINE ARE COMPLETE

// func main(){
// 	numGoroutine:=3
// 	done:=make(chan int)
// 	for i:=range numGoroutine{
// 		go func(id int){
// 			fmt.Printf("Go routine %d working\n",id)
// 			time.Sleep(2*time.Second)
// 			done<-id //SENDING SIGNAL OF COMPLETION

// 		}(i)
// 	}
// 	for range 3{
// 		<-done

// 	}
// 	fmt.Println("All go routine are finished")
// }
//===== synchronizing data exchange


func main(){
	data:=make(chan string)
	go func ()  {
for i:= range 5{
	data<-"hello "+ string(i+'0')
time.Sleep(100*time.Millisecond)



}
close(data)
		
	}()

for value:=range data{
	fmt.Println("Received value:", value,":", time.Now())


}


}