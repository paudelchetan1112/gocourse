// package main

// import (
// 	"fmt"
// 	"sync"
// )

// type counter struct {
// 	mu sync.Mutex
// 	value int

// }

// func (c *counter) increment(){
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.value++

// }
// func (c *counter) getValue()int{
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.value

// }

// func main() {
// var wg sync.WaitGroup

// counter:=&counter{}

// numgoRoutine:=10

// for range numgoRoutine{
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for range 1000{
// 			counter.increment()
// 		}

// 	}()
// }
// wg.Wait()
// fmt.Println("Get value:", counter.getValue())

// }

package main

import (
	"fmt"
	"sync"
)


func main(){
var counter int
var wg sync.WaitGroup
var mu sync.Mutex 

numgoRoutine:=5

wg.Add(numgoRoutine)
increment:=func(){
	defer wg.Done()
	for range 1000{
		mu.Lock()
		counter++
		mu.Unlock()

	}
}
for range numgoRoutine{
	go increment()

}
wg.Wait()
fmt.Printf("Final counter value:%d\n", counter)

}