//Create an array of 5 integers and calculate the sum of all element
package main   

import "fmt"

func main(){
Array1:=[]int{10, 20, 30, 40, 50}
total:=sums(Array1...)
fmt.Println(total)

}
func sums(num ...int)int{
	total:=0
	for _, v:=range num{
		total+=v;
	}
	return total
}