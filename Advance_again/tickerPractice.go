package main

import (
	"fmt"
	"time"
)

func main(){
	ticker:=time.NewTicker(time.Second)

	defer ticker.Stop()
	starting:=5
	
	for(starting>0){
		select{
	case <-ticker.C:
			fmt.Println("Remaining:",starting)
			starting=starting-1
		
	}
}
	fmt.Println("Time's up")

}


//print time in every second
// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	stop:=time.After(5*time.Second)
// 	defer ticker.Stop()
	
// for{
// 	select{
// 	case <-ticker.C:
// 		fmt.Println("Current Time:", time.Now().Hour(),":",time.Now().Minute(),":",time.Now().Second())
// 	case <-stop:
// 		fmt.Println("Stopping ticker")
// 		return

// 	}
// }

// }