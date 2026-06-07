// Write a function multiplier(factor int) that returns a closure. The closure should multiply its input by factor.
package main

import "fmt"
func main(){
var num int
multiplyby:=multiplier()
	for{
fmt.Printf("enter any number to multiply:");
fmt.Scanln(&num)
fmt.Println("result=",multiplyby(num))
if(num==0){
	break
}

	}

}
func multiplier() func(a int)int{
result:=1
fmt.Println("Initially result is:",result)
fmt.Println("you can enter 0 for exit, thank you...")
return func(a int)int{
result=result*a
return result
}
}