// package main

// import (
// 	"fmt"
// 	"time"
// )

// func worker(id int, tasks <-chan int, results chan<- int) {
// 	for task := range tasks {
// 		fmt.Printf("Worer %d processing task %d\n", id, task)
// 		//simulate some work
// 		time.Sleep(time.Second)
// 		results<-task*2

// 	}
// }
// func main(){
// 	numWorkers:=3
// 	numjobs:=10
// 	tasks:=make(chan int, numjobs)
// 	results:=make(chan int, numjobs)
// 	//create worker
// 	for i:=range numWorkers{
// 		go worker(i, tasks, results)
// 	}
// 	//send the values to the tasks channel
// 	for i:=range numjobs{
// 		tasks<-i

// 	}

// 	for range numjobs{
// 		value:=<-results
// 		fmt.Println(value)
// 	}

// 	//receive the value fr

// }

package main

import (
	"fmt"
	"time"
)

type ticketRequests struct{
	personId int
	numTicket int
	cost int
}
func ticketProcessor(request <-chan ticketRequests, result chan <- int){
	for req:=range request{
		fmt.Printf("Processing %d ticket(s) of personId:%d", req.numTicket, req.personId)
		//simulate processing time
	time.Sleep(time.Second)
	result<-req.personId
	}
	
}

func main(){

numRequests:=5
price:=5
ticketRequest:=make(chan ticketRequests)
ticketResult:=make(chan int)

//start ticket processor/worker

for range 3 {
	go ticketProcessor(ticketRequest, ticketResult)

}

//send ticket requests

for i:=range numRequests{
	ticketRequest<-ticketRequests{personId: i, numTicket:i+1, cost: (i+1)*price }



}
close(ticketRequest)
for range numRequests{
	fmt.Printf("Ticket for personId %d processed successfully!\n", <-ticketResult)

}
}