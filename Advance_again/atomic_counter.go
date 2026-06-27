package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Atomiccounter struct {
	count int64
}

func (c *Atomiccounter) increment() {
	atomic.AddInt64(&c.count, 1)

}
func (c *Atomiccounter) getValue() int64{
	return atomic.LoadInt64(&c.count)


}
func main() {
	// value:=0

	Counter:=&Atomiccounter{}
	var wg sync.WaitGroup
	NoOfGoRoutine:=10
	for range NoOfGoRoutine{
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000{
				Counter.increment()
				// value++

			}
			
		}()
	}
	wg.Wait()
	fmt.Println("Current value:=", Counter.getValue())
	// fmt.Println("Current value:=", value)
	
}