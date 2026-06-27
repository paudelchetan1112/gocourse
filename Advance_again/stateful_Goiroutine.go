package main

import (
	"fmt"
	"time"
)

type statefulworker struct {
	count int
	ch    chan int
}

func (w *statefulworker) Start() {
	go func() {
		for {
		select {
		case value:= <- w.ch:
			w.count += value
			fmt.Println("Current count:", w.count)
		}
	}
	}()
	
}

func (w *statefulworker) Send(value int){
w.ch<-value
}

func main(){
	stWorker:=&statefulworker{ch:make(chan int)}

	stWorker.Start()
	for i:=range 5{
		stWorker.Send(i)
		time.Sleep(500*time.Millisecond)
	}
	
}