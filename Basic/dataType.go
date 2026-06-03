/* some of the data types that go language support are as follow:
Basic Data Types
Integer: Represents whole numbers without a fractional component.
Examples: int, int8, int16, int32, int64
Floating-Point Number: Represents numbers with a fractional component.
Examples: float32, float64
Complex Number: Represents numbers with real and imaginary parts.
Go provides complex64 and complex128 for working with complex numbers and performing mathematical operations on them.
Boolean Type: Represents logical values.
Possible values: true and false
String: Represents a sequence of characters.
Example: "Hello, World"
Composite and Reference Types
Array: A fixed-size collection of elements of the same type.
Struct: A user-defined composite type that groups together variables (fields) of different types under a single name.
Pointer: Stores the memory address of another variable.
Map: A collection of key-value pairs, similar to dictionaries or hash tables in other languages.
Slice: A flexible, dynamically sized view into an array.
Functional and Concurrency Types
Function: A reusable block of code that performs a specific task.
Channel: A communication mechanism used by goroutines to safely exchange data and synchronize execution.
Standard Library Packages/Common Concepts
JSON: Go provides the encoding/json package for encoding and decoding JSON data.
Text Templates: The text/template package is used to generate text output from templates.
HTML Templates: The html/template package is used to generate safe HTML output and helps prevent cross-site scripting (XSS) attacks.
*/
package main
import "fmt"
func main(){
	fmt.Println("Hello"+"World")
	fmt.Println("9x10=", 9*10)
	fmt.Println("180.18/2.0", 180.18/2.0)
	fmt.Println(true&&false)
	fmt.Println(true||false)
	fmt.Println(!true)

}
