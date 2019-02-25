package main

import "fmt"

type myInt int

type rectangle struct {  
    length int
    width  int
}
	

func (a myInt) add(b myInt) myInt {  
    return a + b
}

func perimeter(r *rectangle) {  
    fmt.Println("perimeter function output:", 2*(r.length+r.width))

}

func (r *rectangle) perimeter() {  
    fmt.Println("perimeter method output:", 2*(r.length+r.width))
}


func main() {  
    num1 := myInt(5)
    num2 := myInt(10)
    sum := num1.add(num2)
    fmt.Println("Sum is", sum)


    r := rectangle{
        length: 10,
        width:  5,
    }
    p := &r //pointer to r
    perimeter(p)
    p.perimeter()

    /*
        cannot use r (type rectangle) as type *rectangle in argument to perimeter
    */
    //perimeter(r)

    r.perimeter()//calling pointer receiver with a value

}