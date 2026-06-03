package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// source:=rand.NewSource(time.Now().UnixNano())
	// random:=rand.New(source)
source:=rand.NewSource(time.Now().UnixNano())

random:=rand.New(source)
fmt.Println(random)

	//Generate a random number between 1 and 100
	// target:=random.Intn(100)+1
	target:=random.Intn(100)+1
	//welcome message
	fmt.Println("Welcome to the Guessing Game !")
	fmt.Println("I have chosen 10a number between 1 and 100")
	fmt.Println("Can you guess what it is?")
	var guess int 
	for{
		fmt.Println("Enter your Guess:");
		fmt.Scanln(&guess)
		//check if the guess is correct 
		if guess==target{
			fmt.Println("Congratulation you guessed the correct number!")
			break
		}else if guess<target {
			fmt.Println("Too low! try higher number.")

		}else {
			fmt.Println("too high!, guessing the lower number ");
		}
	}
}