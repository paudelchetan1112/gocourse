package main

import (
	"fmt"
	"sync"
	"time"
)

type ratelimiter struct {
	mu sync.Mutex
	count int
	limit int
	window time.Duration
	resetTime time.Time

}

func newRatelimiter(limit int, window time.Duration) *ratelimiter{
	return &ratelimiter{
		limit:limit, 
		window:window,
	}
}
func (rl *ratelimiter) allow()bool{
	rl.mu.Lock()
	defer rl.mu.Unlock()
	now:=time.Now()

	if now.After(rl.resetTime){
		rl.resetTime=now.Add(rl.window)
		rl.count=0

	}
	if rl.count<rl.limit{
		rl.count++
		return true
	}
	return false


}

func main() {
	rateLimiter:=newRatelimiter(5, time.Second)
	for range 10{
		if rateLimiter.allow(){
		fmt.Println("Request allwed")

		
	}else{
		fmt.Println("Request denied")
	}
	}

}