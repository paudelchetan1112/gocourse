package main
import "fmt"

func main(){
	// for i := 1; i <=5; i++ {
		
	// 	fmt.Println(i)

	// }
	// //iterate over collection
	// numbers:=[]int {1,2,3,4,5,6}
	// for index, value:=range numbers {
	// 	fmt.Printf("Index:%d , Value:%d\n",  index, value)

	// }

	// for i := 1; i <=10; i++ {
	// 	if i%2==0{
	// 	continue //continue the loop but skip the rest of lines/statement
	// 	}
	// 	fmt.Println("odd Number:",i)
	// 	if(i==5){
	// 		break   //break out of the loop
	// 	}
		
	// }

	rows:=5
	for i:=1;i<=rows;i++ {
		//inner loop for space before start
		for j:=1; j<=rows-i; j++{
			fmt.Print(" ")

		}
	// 	//inner loop for starts
	// 	for k:=1;k<=i;k++{
	// 		fmt.Print("*")
	// 	}
	// for k:=1;k<i;k++{
	// 	fmt.Print("*")
	// }
	for k:=1;k<=2*i-1;k++{
		fmt.Print("*")
	}
fmt.Println()	
}
	

for i:=range 10{
	fmt.Println(i)
}

}