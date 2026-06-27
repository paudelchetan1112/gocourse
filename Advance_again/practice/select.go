package main

import (
	"fmt"
	// "strings"
	"time"
)


func main(){
ch1 := make(chan string,2)
	ch2 := make(chan string,2)

	go func() {
		for i := range 10 {
			ch1 <- "msg: " + string(i + '0')

			time.Sleep(10 * time.Millisecond)

		}
		close(ch1)

	}()

	go func() {
		for i := range 10 {
			ch2 <- "msg: " + string(i + '0')
			time.Sleep(20 * time.Millisecond)

		}
		close(ch2)
	}()
	

	for {
		select {
		case msg, ok := <-ch1:
			if !ok {
				fmt.Println("channel 1 close")
				return
			}
			fmt.Println("message received from channel 1:", msg)
		case msg, ok := <-ch2:
			if !ok {
				fmt.Println("channel 2 close")
				return
			}
			fmt.Println("message received from channel 2:", msg)
		

		}
	}

}


