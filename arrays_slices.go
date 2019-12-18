package main

import "fmt"

func main() {

	// first way
	names := [3]string{"ahmed", "bahaa", "mahmoud"}
	fmt.Println("My first array : ", names)

	// second way
	var names2 [3]string
	names2[0] = "ahmed"
	names2[1] = "Bahaa"
	names2[2] = "Mahmoud"
	fmt.Println("\n")
	fmt.Println("My second array : ", names2)

	// third way
	var names3 [3]string
	names2[0] = "ahmed2"
	names2[1] = "Bahaa2"
	//names2[1] = "Mahmoud"
	fmt.Println("\n")
	fmt.Println("My third element in the array is empty? : ", names3[2] == "" )

	// this if we are not have idea about array's length
	// it requires us to allocate a new slice every time we insert an item.
	// called slices
	names4 := []string{}
	names4 = append( names4 , "hi ")
	names4 = append( names4 , "there ")
	names4 = append( names4 , "i"," am " , "here.")
	fmt.Println("My fourth array with slicing method : " , names4)

	// fifth way
	// this one if we know the minimum number of elements, so rather than every time we append we will create a new one
	//,, this will save some effort
	//If we know what the starting length of a slice is going to be then we can use the make function to pre-allocate the space.
	// called make
	names5 := make([]string , 2)
	names5[0] = "ahmed"
	names5[1] = "Bahaa"
	names5 = append(names5, "Mahmoud")
	fmt.Println("My fifth array with Make method : " , names5)

}
