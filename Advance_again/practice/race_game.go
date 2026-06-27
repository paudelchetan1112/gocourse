package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
type person struct{
	name string
	duration int

}
func runner(p person, winner chan person, wg *sync.WaitGroup) {
	defer wg.Done()
	delay := int(rand.Intn(5))
	time.Sleep(time.Duration(delay)*time.Second)
	p.duration=delay
	winner <- p

}

func main() {
	var wg sync.WaitGroup
	winner:=make(chan person)
	
	runner1:=person{name: "chetan",}
	runner2:=person{name: "Dinesh", }
	runner3:=person{name: "Mohan", }
	wg.Add(1)
	go runner(runner1,winner,&wg )
	wg.Add(1)
	go runner(runner2,winner,&wg )
	wg.Add(1)
	go runner(runner3,winner,&wg )
	go func() {
		wg.Wait()
		close(winner)
	}()
for value:= range winner{
	fmt.Println(value.name, ": finished at:", value.duration)

}



}