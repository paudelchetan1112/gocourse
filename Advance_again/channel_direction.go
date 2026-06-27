package main

import "fmt"

func main() {

	ch := make(chan int)

	 

	// for range 5 {
	// 	fmt.Println(<-ch)
	// }
	producer(ch)
	consumer(ch)

}

func producer(ch chan<-int){
	go func(ch chan<- int) {
		for i := range 5 {
			ch <- i
		}
		close(ch)

	}(ch)
}
func consumer(ch <- chan int){
	for range 5 {
		fmt.Println(<-ch)
	}
}