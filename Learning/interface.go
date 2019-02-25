package main

import (  
    "fmt"
)

//interface definition
type useInterface interface {  
    useInterfaceFunction()
    print()
}

type myInt int

func (mi myInt) useInterfaceFunction() {
	fmt.Println(mi + 2)
}

func (mi myInt) print() {
	fmt.Println("interface function")
}

func describe(i interface{}) {
	fmt.Println("describe : " , i)
}


func findType(i interface{}) {  
    switch i.(type) {
    case string:
        fmt.Printf("I am a string and my value is %s\n", i.(string))
    case int:
        fmt.Printf("I am an int and my value is %d\n", i.(int))
    default:
        fmt.Printf("Unknown type\n")
    }


func main() {
	ex := (myInt)(2)

	var me useInterface

	me = ex

	fmt.Println(me)
	me.useInterfaceFunction()
	me.print()

	num := 55
	describe(num)

	str := "adb"
	describe(str)	

	stru := struct {
		firstname string
		lastname string
	}{
		firstname : "kay",
		lastname : "zan",
		}
	describe(stru.firstname)
	describe(stru)// describe :  {kay zan}

	func findType(i interface{}) {  
    switch i.(type) {
    case string:
        fmt.Printf("I am a string and my value is %s\n", i.(string))
    case int:
        fmt.Printf("I am an int and my value is %d\n", i.(int))
    default:
        fmt.Printf("Unknown type\n")
    }


}


