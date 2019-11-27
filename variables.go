package main 

import (
	"fmt"
)

var x int = 7

func main(){
	// the first way
	var x int 
	x = 2

	// the second way 
	var y int = 3
	
	// the third way ,, we can use it if we don't know the type
	z := 4
	
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%v,%T",z,z)
	fmt.Println(55)
}

// notes: 
// we can't declare a variable and don't use it ,,, it will through error 
// we cannot declare the same variable twice 
