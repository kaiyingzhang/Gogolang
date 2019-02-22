package main

import (
		"fmt"
	)


func main() {
	
	// if else
	if( false ){
		fmt.Println("1")
	}else if( false ){
		fmt.Println("2")
	}else{
		fmt.Println("3")
	}

	//loops
	//for initialisation; condition; post {  ...  }
	for i := 1; i < 10; i++ {
		fmt.Println("i : " , i)
		//break
		if( i == 5 ){
			break
		}
	}

	//switch
	finger := 5
	switch finger{
	case 1 :
		fmt.Println("wrong")
	case 5, 3, 6, 7:
		fmt.Println("right")
	}

 
}