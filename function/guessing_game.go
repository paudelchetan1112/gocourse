package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main(){
source:=rand.NewSource(time.Now().UnixNano())
random:=rand.New(source)
target:=random.Intn(100)+1
var guess int
//some message:
fmt.Println("This a simple game , a computer generate a number between 1-100, user should guess:")
for{
	fmt.Println("\nEnter a Guess:")
	fmt.Scanln(&guess)
	if target==guess {
		fmt.Println("Congratulation! your guess matched with target")
break;
		
	}else if(target>guess){
		fmt.Print("Too low, please guess with another large Number:")
	}else{
		fmt.Print("Too high, Please try again with another lower number:")
	}

}


}