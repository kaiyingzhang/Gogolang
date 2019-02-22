package main

import "fmt"

func main() {
	
	//create map
	var personSalary map[string]int
	personSalary["a"] = 1000
	personSalary["b"] = 2000

	delete(personSalary, "a")

	//map is reference
	
}