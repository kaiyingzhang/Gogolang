/*

The following are the basic types available in go

bool
Numeric Types
int8, int16, int32, int64, int
uint8, uint16, uint32, uint64, uint
float32, float64
complex64, complex128
byte
rune
string

*/

package main

import "fmt" 
import "unsafe"

func main() {

	// bool
	a := true
	

	//integers
	/*
	int8: represents 8 bit signed integers 
	size: 8 bits 
	range: -128 to 127

	int16: represents 16 bit signed integers 
	size: 16 bits 
	range: -32768 to 32767

	int32: represents 32 bit signed integers 
	size: 32 bits 
	range: -2147483648 to 2147483647

	int64: represents 64 bit signed integers 
	size: 64 bits 
	range: -9223372036854775808 to 9223372036854775807
	*/

	b := 89

	//%T is the format specifier to print the type and %d is used to print the size.
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a)) //type and size of a
	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) //type and size of b    

	//float
	no1, no2 := 56, 89
    fmt.Println("sum", no1+no2, "diff", no1-no2)


	//Complex types
	/*
	The complex function has the following definition
	-func complex(r, i FloatType) ComplexType-
	*/
	c1 := complex(5, 7)
    c2 := 8 + 27i
    cadd := c1 + c2
    fmt.Println("sum:", cadd)
    cmul := c1 * c2
    fmt.Println("product:", cmul)


    //string
    first := "kaii"
    last := "zan"

    name := first + " " + last
    fmt.Println(name)

    

}


