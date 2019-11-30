package main

import (
	"fmt"
	"strconv"
)

// package scope 
// if we used a variable name with upper case that's mean this variable accessable 
// by all packages 
// if it's lower case meaning it will be package level only files inside this package can use it 

// it's better to use camel case oneVariable but incase acronyms use it upper ->  varHTTPRequest  
// this only for readability 

var x int = 5
var Y int = 10

// type convertion 

func main(){

	// from int to float 
	var i int = 42
	fmt.Printf("%v , %T\n",i,i)

	var j float32
	j = float32(i)
	fmt.Printf("%v , %T\n",j,j)

	//------------------------------------------

	// from float to int  
	var i2 float32 = 42.5
	fmt.Printf("%v , %T\n",i2,i2)

	// data loss
	var j2 int
	j2 = int(i2)
	fmt.Printf("%v , %T\n",j2,j2)

	//------------------------------------------
	// from int to string   
	var i3 int = 42
	fmt.Printf("%v , %T\n",i3,i3)

	var j3 string 
	j3 = strconv.Itoa(i3)	// integer to ascii string
	fmt.Printf("%v , %T\n",j3,j3)

}
