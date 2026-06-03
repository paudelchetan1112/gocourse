package main

import "fmt"

//Print numbers from 10 to 1.

// func main() {
// for i := 10; i > 0; i-- {
// 	fmt.Println(i)

// }
// }

//3.	Print all even numbers between 1 and 100.
// func main(){
// 	for i := 1; i < 100; i++ {
// 		if i%2==0 {
// 			fmt.Println(i)
// 		}
// 	}
// }

//4. print all odd number between 1 and 100
// func main(){
// 	for i := 1; i < 100; i++ {
// 		if i%2!=0 {
// 			fmt.Println(i)
// 		}
// 	}
// }

//5. Find the sum of numbers from 1 to 100.
// func main(){
// 	var sum int=0;
// 	for i := 1; i <= 100; i++ {
// sum+=i;

// 	}
// 	fmt.Println(sum)
// }

//6. 6.	Find the product of numbers from 1 to 10.
// func main(){
// 	var product int=1;
// 	for i := 1; i <= 10; i++ {
// product*=i;

// 	}
// 	fmt.Println(product)
// }

//7. 7.	Print the multiplication table of 5.
// func main(){
// 	for i := 1; i <= 10; i++ {
// 		fmt.Println("5 *",i,"=",5*i)

// 	}
// }

//8.Count the number of digits in a given number.
// func main(){
// 	var number int=34232323;
// 	var  count int

// 	for i := number; i > 0; i=i/10 {
// 		count++
// 	}
// 	fmt.Println(count)

// }

//9.Multiplication Table from 1-10
// func main(){
// 	for i := 1; i <=10; i++ {
// 		fmt.Println("Multiplication table of ",i)
// 		for j := 1; j <=10; j++ {
// 			fmt.Println(i,"*",j,"=",i*j)
// 		}

// 	}
// }

//10. Reverse the given number

// func main(){
// 	var number int =234234
// 	var result int=0
// 	var rem int
// 	for number>0 {
// 		rem=number%10
// 		result=result*10+rem
// 		number=number/10
// 	}
// 	fmt.Println(result)

// }

//11. Check the given number is palindrome or not
// func main(){
// 	var number int=2322
// 	var rem, result, temp int
// 	temp=number
// 	result=0
// 	for temp>0{
// 		rem=temp%10
// 		result=result*10+rem
// 		temp=temp/10

// 	}
// 	if number==result {
// 		fmt.Println("Number is palindrome")

// 	}else {
// 		fmt.Println("Number is not palindrome")
// 	}
// }

//12. Find the sum of digit of the given number

// func main(){
// 	var number int=2323
// 	var rem , result int
// 	result =0
// 	for number>0{
// 		rem=number%10
// 		result+=rem
// 		number=number/10
// 	}
// 	fmt.Println(result)
// }

//13.Find the largest digit in a number.
// func main(){
// 	var number int=9328
// 	var rem , largest int
// 	largest=0
// 	for number>0{
// 		rem=number%10
// 		if rem>largest {
// 			largest=rem
// 		}
		
// 		number=number/10
// 	}
// 	fmt.Println(largest)
// }

//14.	Find the smallest digit in a number. 
// func main(){
// 	var number int=9328
// 	var rem , smallest int
// 	smallest=9
// 	for number>0{
// 		rem=number%10
// 		if rem<smallest {
// 			smallest=rem
// 		}
		
// 		number=number/10
// 	}
// 	fmt.Println(smallest)
// }

//15.	Count how many even digits are in a number. 
// func main(){
// 	var number int =34522352680
// 	var count,rem int
// 	count=0
// 	for number>0{
// rem=number%10
// if rem%2==0 {
// 	count++
// }
// number=number/10
// 	}
// 	fmt.Println(count)
// }

// //17.	Print squares of numbers from 1 to 20. 
// func main(){
	
// 	for i :=1; i <= 20; i++ {
// 	fmt.Println(i*i)
// 	}
// }

// 22.	Print numbers divisible by 3 between 1 and 100. 
// func main(){
// 	for i := 1; i <=100; i++ {
// 		if i%3==0 {
// 			fmt.Println(i)
// 		}
		
// 	}
// }

//23. print numbers divisible by 3 and 5
// func main(){
// 	for i := 1; i <=100; i++ {
// 		if i%3==0 && i%5==0 {
// 			fmt.Println(i)
// 		}
		
// 	}
// }
//24.81.	Check whether a number is prime. 
// func main(){
// 	var c1, number int
// 	c1=0
// 	number=4
// for i := 1; i <=number; i++ {
// 	if number%i==0 {
// 		c1++
		
// 	}

// }
// if(c1==2){
// 	fmt.Println("number is prime")
// }else{
// 	fmt.Println("number is not prime")
// }
// }






// }
