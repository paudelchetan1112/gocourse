/*Problem: Given an array and a target sum, return the indices of two numbers that add up to the target.

Example:

Input: nums = [2,7,11,15], target = 9
Output: [0,1] */

package main          
import "fmt"
func findTarget(a [4]int, target int)(int, int){

	for i := 0; i < 4; i++ {
	for j := i+1; j < 4; j++ {
		if a[i]+a[j]==target{
		return i,j
		}

	}
}
return -1, -1
}


func main(){
	nums:=[4]int{2,7,11,15}
	var target int = 9

a,b:=findTarget(nums, target)
fmt.Printf(" sum of %d and %d is: %d", nums[a], nums[b], target)


	


}