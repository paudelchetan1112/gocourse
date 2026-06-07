//formatting Verb

package main
import "fmt"
func main(){
	//------------General formatting verb
	/* 
	%v    print the value in the default format
	%#v    prints the value in Go-syntax format
	%T		prints the type of the value
	%% 		print the % sign

*/ 
i:=13_455.5
string:="Hello world!"
fmt.Printf("%v\n",i)
fmt.Printf("%#v\n",i)
fmt.Printf("%T\n",i)
fmt.Printf("%%\n",i)

fmt.Printf("%v\n",string)
fmt.Printf("%#v\n",string)
fmt.Printf("%T\n",string)
	//----Integer formatting Verbs

	// %b		Base 2
	// %d		Base 10
	// %+d		base 10 and always show sign
	// %o		base 8
	// %0		Base 8, with leading 0o
	// %x		Base 16 , lowercase
	// %X		Base 16 , Uppercase
	// %#x 	Base 16, with leading 0x
	// %4d		pad with the spaces (width 4, right justified)
	// %-4d	Pad with spaces (width 4, right justified)
	// %04d	Pad with zeroes (width 4)

int:=255
fmt.Printf("%b\n", int)
fmt.Printf("%d\n", int)
fmt.Printf("%+d\n", int)
fmt.Printf("%o\n", int)
fmt.Printf("%O\n", int)
fmt.Printf("%x\n", int)
fmt.Printf("%X\n", int)
fmt.Printf("%#x\n", int)
fmt.Printf("%4d\n", int)
fmt.Printf("%-4d\n", int)

	//---String formatting verb


	//%s Print the value aas plain string
	// %q		prints the vlaue as a double-quote string
	//%8s 		prints the value s plain string (width 8, right justified)
	// %-8s		prints the vlaue as plain string (widht 8, left justified)
	// %x		prints the value as hex dump 


// Boolean Formatting verbs
//%t    value of the boolean operator in true or false format (same as using %v)
t:=true
f:=false
fmt.Printf("%t\n", t)
fmt.Printf("%v\n", f)

	//------ Float formatting Verbs
	//Verb Description
	// %e Sceintific notation with 'e' as exponent
	// %f  Decimal point, no exponent
	// %.2f     Default widht, precision 2
	//  %6.2f   width 6, precision 2
	// %g Exponent as need only necessary digit


	flt:=918.00

	fmt.Printf("%e\n", flt)

	fmt.Printf("%f\n", flt)
	fmt.Printf("%.2f\n", flt)
fmt.Printf("%6.2f", flt)

}