//Write a function runningSum() that returns a closure. Each call to the closure accepts an integer and returns the cumulative sum of all numbers seen so far.
package main
import "fmt"

func main(){
cummulativesum:=runningsum()
fmt.Println(cummulativesum(5))
fmt.Println(cummulativesum(7))
}
func runningsum() func(a int)int{
	sum:=0
	fmt.Println("Intially sum is:",sum)
	return func(a int)int{
sum+=a;
return sum
	}
}