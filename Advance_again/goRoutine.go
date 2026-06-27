package main

import (
	"fmt"
	"sync"
	"time"
	// "time"
)

func goRoutine() {
	fmt.Println("Beginning Program.")

		var wg sync.WaitGroup
		wg.Add(1)
	go sayHello(&wg)
	fmt.Println("Hello from main thread")
	wg.Add(1)
		go printNumber(&wg)
		wg.Add(1)
	go printLetter(&wg)
	wg.Wait()
	fmt.Println("Done")


var err error
go func() {
	err=doWork()
}()
time.Sleep(time.Second)
if err!=nil{
	fmt.Println("Error:", err)

}else{
	fmt.Println("Work completed successfully")

}





}

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	// time.Sleep(1*time.Second)
	fmt.Println("Hello from Goroutine")

}
func printNumber(wg *sync.WaitGroup){
	defer wg.Done()
	for i:=0;i<5;i++{
		fmt.Println(i, time.Now())
		time.Sleep(100*time.Millisecond)
	}
}
func printLetter(wg *sync.WaitGroup){
	defer wg.Done()
	for _, letter:=range "abcde"{
		fmt.Println(string(letter), time.Now())
		time.Sleep(100*time.Millisecond)
	}
}
func doWork()error{
	//simulate work
	time.Sleep(time.Second)
	return fmt.Errorf("Error occur")
}