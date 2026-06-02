/*
Basic Arithmetic Operator
    Addition +
    Subtraction -
    Multiplication *
    Division /
    Remainder (Modulus %)
Operator Precedence:
    1. Parenthesis ()
    2. Multiplication *, Division /, Remainder %
    3. Addition +, Subtraction -
Overflow
Underflow

Why be mindful of Overflow and underflow
    program stability
    data integrituy
    Type Safety
Mitigation Strategies
    Range Checking
    Type Conversion
    Error handling

*/

package main

import (
	"fmt"
	"math"
)
func main(){
    //variable declaration
    var a, b int =10, 3
    var result int
    result =a+b
    fmt.Println("Addition:", result)
    result = a-b
    fmt.Println("Subtraction:", result)
    result=a*b
    fmt.Println("Multiplication:", result)
    result = a/b
    fmt.Println("Division:", result)
    result = a%b
    fmt.Println("Remainder:", result)
    const p float64 =22.0/7.0
    fmt.Println(p)
    //overflow with signed integers

    var maxInt int64 = 9223372036854775807
    fmt.Println(maxInt)
    maxInt=maxInt+1;
    fmt.Println(maxInt)
    //overflow with unsigned integers
    var uMaxInt uint64 =18446744073709551615 //max fvalue for uint64

    uMaxInt=uMaxInt+1
    fmt.Println(uMaxInt)

    var smallFloat float64 =1.0e-323
    fmt.Println(smallFloat)
    smallFloat=smallFloat/math.MaxFloat64
    fmt.Println(smallFloat)

}