package main

import (
	"fmt"
	"time"
)

type rateLimiter struct {
	token      chan struct{}
	refillTime time.Duration

}
func newRateLimiter(rateLimit int, refillTime time.Duration)*rateLimiter{
	rl:=&rateLimiter{
		token: make(chan struct{}, rateLimit),
		refillTime: refillTime,

	}
	for range rateLimit{
		rl.token<-struct{}{}

	}
	return rl

}
func (rl *rateLimiter) startRefill(){
	ticker :=time.NewTicker(rl.refillTime)
	defer ticker.Stop()
	for{
		select{
		case <-ticker.C:
			select{
			case rl.token<-struct{}{}:
			default:
			}
		}
	}
}
func (rl *rateLimiter) allow()bool{
	select{
	case <-rl.token:
		return true
	default:
		return false 
	}
}
	
	func tokenBucket() {
		rateLimiter:=newRateLimiter(5, time.Second)
		for range 10{
			if rateLimiter.allow(){

				fmt.Println("Request allowed")

			}else{
				fmt.Println("Request denied")
			}
			time.Sleep(100*time.Millisecond)
		}
}