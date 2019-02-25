package main

import (
		"math"
		"fmt"
		)


func main() {
	var age int
	fmt.Println("how old are you? :", age)

	age = 10
	fmt.Println("how old are you? :", age)

	age = 100

	fmt.Println("how old are you? :", age)

	var momage int= 2000
	fmt.Println("how old are you? :", momage)

	var height, weight int = 100, 101
	fmt.Println("height : ", height," weight : " , weight)

	var (
		a = "hahaha"
		b = 14
		c = 0.001
	)

	fmt.Println("a : " , a , " b :" , b ," c " , c)

	//-------------------------------------------------
	a , b = "haha", 15

	fmt.Println("a : " , a , " b :" , b ," c " , c)
	//-------------------------------------------------

	//name := initialvalue is the short hand syntax to declare a variable.

 	aa, bb := 145.8, 543.8
    cc := math.Min(aa, bb)
    fmt.Println("minimum value is ", cc)

    fmt.Println("a : " , aa , " b :" , bb ," c " , c)
	//-------------------------------------------------

	//":=" replace the var and type!!! this is kind of short hand syntax. But it can only used inside the function. 
	//if you use ":=" outside a function, it will get an error!

}