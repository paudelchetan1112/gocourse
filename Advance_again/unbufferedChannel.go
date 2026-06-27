package main

import (
	"fmt"
	"time"
	// "time"
)

func UnBufferedChannel() {
	ch := make(chan int)
	go func() {
			ch<-1
	}()

	go func() {
		value:= <-ch
		fmt.Println(value)
	}()
	time.Sleep(time.Second)

}