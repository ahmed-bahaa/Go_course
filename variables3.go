package main

import "fmt"

func main() {

	var myint1 int = 10
	myint2 := 20 // short declaraation

	var mystring, myboolean = "Ok", true

	fmt.Println("myint1 is : ", myint1)
	fmt.Println("myint2 is : ", myint2)
	fmt.Println("mystring is : ", mystring)
	fmt.Println("myboolean is : ", myboolean)

	//================================
	// this incase we have function
	// return multiple variables and
	// we only care about one
	//================================

	var myvar1, _ = "yes", "No"
	fmt.Println("My var : ", myvar1)

	// can use this way
	// var id1, id2, id3 type = val1, val2, val3

}
