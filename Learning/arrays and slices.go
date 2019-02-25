package main

import "fmt"

func main() {
	a := [3]int{1,2,3}

	fmt.Println(a)// [1 2 3]
	for i := 0; i < 3; i++ {
		fmt.Println(a[i])
	}// 1 2 3

	fmt.Println("length of array : " , len(a))

	//range return both of the index and value
	for i , v := range a {
		fmt.Printf("%d the element of a is %d \n", i , v)
	}

	b := [3][2]int{
		{1,3},
		{2,4},	
		{5,6},//this comma is necessary
	}

	fmt.Println(b)

	//slice
	var c []int = a[1 : 2]
	fmt.Println(c)

	c2 := append(c ,5)
	fmt.Println(c2)

	copy(a[0 : 2], c2)//first element should be slice..
	fmt.Println(c2)

}
