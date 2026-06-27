package main

import (
	"context"
	"fmt"
	"time"
)

func checkEvenOdd(ctx context.Context, num int)string{
	select{
	case <-ctx.Done():
		return "Operation canceled"
	default:
		if num%2==0{
			return fmt.Sprintf("%d is even", num)
			
		}else{
			return fmt.Sprintf("%d id odd", num)

		}
	}
}
func main(){

	ctx:=context.TODO()
	result=checkEvenOdd(ctx, 5)
	fmt.Println("Result with context.Todo():", result)
	ctx:=context.Background()
	ctx, cancel:=context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	



}


//===difference Between context.todo and context.background

// func main() {
// 	todoContext := context.TODO()
// 	contextBkg:=context.Background()

	
// 	ctx:=context.WithValue(todoContext, "name", "john")
// 	fmt.Println(ctx)
// 	fmt.Println(ctx.Value("name"))
// 	ctx1:=context.WithValue(contextBkg, "City", "New york")
// 	fmt.Println(ctx1)


}