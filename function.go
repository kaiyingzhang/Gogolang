package main

import (  
    "fmt"
)

// func functionname(parametername type) returntype {  
//  //function body
// }

func calculateBill(price, no int) int {  
    var totalPrice = price * no
    return totalPrice
}

// func rectProps(length, width float64)(float64, float64) {  
//     var area = length * width
//     var perimeter = (length + width) * 2
//     return area, perimeter
// }


func rectProps(length, width float64)(area, perimeter float64) {  
    area = length * width
    perimeter = (length + width) * 2
    return //no explicit return value
}

func main() {  
    price, no := 90, 6
    totalPrice := calculateBill(price, no)
    fmt.Println("Total price is", totalPrice)

    area , _:= rectProps(50.5, 45.3)
    fmt.Println("rectProps is : " , area)

}