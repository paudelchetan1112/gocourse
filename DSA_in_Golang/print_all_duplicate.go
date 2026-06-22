//duplicate value print using map 

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 2, 4, 5, 1}
	count:=make(map[int]int)
for _, v:=range slice{
	count[v]++
}
for k, v:=range count{
	if v>1{
		fmt.Println(k)
	}
}
	
}

