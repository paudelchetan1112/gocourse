package main

import "fmt"

func main() {

	//...Elipsis

	// //func functionName(param1 type1, param2 type2, param3 ....type3) returnType{
	// }
	// fmt.Println("sum of 1,2,3:", sum(1,2,3))

	statement,total:=sum("The sum of 1,2,3, is:",1,2,3)
	fmt.Println(statement, total)
	numbers:=[]int{1,2,3,4,5,6,7,8,9}

sequence, total:=sums(1, numbers...)
fmt.Println("sequence:",sequence,"Total:",total)

}
// func sum(nums ...int)int {
// 	total := 0
// 	for _, v := range nums {
// 		total += v
// 	}
// 	return total
// }
func sum(returnString string, nums ...int)(string,int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return returnString, total
}
func sums(sequence int, nums ...int)(int ,int){
	total:=0
	for _,v:=range nums{
		total+=v
	}
	return sequence, total
}