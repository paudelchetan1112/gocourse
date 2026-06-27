package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	Balance int
	mu      sync.Mutex
}
func (b *BankAccount) Deposit(amount int){
	b.mu.Lock()
	defer b.mu.Unlock()

	b.Balance+=amount

}
func (b *BankAccount) Withdraw(amount int){
	b.mu.Lock()
	defer b.mu.Unlock()

	if(amount<0){
		fmt.Println("amount must be positive")
		return
	}
	
	if(b.Balance>=amount){
	b.Balance-=amount
	}


}
func(b *BankAccount) GetBalance() int{
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.Balance
}


func main() {
	Account:=&BankAccount{}
var wg sync.WaitGroup

depositeGoroutine:=100
withdrawGoroutine:=50


for range depositeGoroutine{
		wg.Add(1)

	go func() {
			defer wg.Done()
		Account.Deposit(100)
	}()
	
}

for range withdrawGoroutine{
	wg.Add(1)

go func() {
	defer wg.Done()
	Account.Withdraw(100)
}()
}

wg.Wait()
fmt.Println("Current Balance is:", Account.GetBalance())




}
