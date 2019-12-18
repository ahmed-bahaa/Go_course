package main

import "fmt"

func main(){

	// if, else
	ages := map[string]int{}
	ages["Kevin"] = 57

	if ages["Kevin"] < 18 {
					fmt.Println("Kevin can't vote yet")
	} else if ages["Kevin"] < 67 {
					fmt.Println("Kevin is not of retirement age just yet")
	} else {
					fmt.Println("Enjoy retirement Kevin!")
	}


	// switch

	// switch {
	// case ages["Kevin"] < 18:
	// 				fmt.Println("Kevin can't vote yet")
	// case ages["Kevin"] < 67:
	// 				fmt.Println("Kevin is not of retirement age just yet")
	// default:
	// 				fmt.Println("Enjoy retirement Kevin!")
	// }

	//  or

	switch ages["Kevin"] {

		case 1, 2, 3, 5, 7, 11, 13, 17, 19:
						fmt.Println("Kevin's age is a prime number!")
		case 16:
						fmt.Println("Kevin can drive now")
		case 18:
						fmt.Println("Kevin can vote now!")
		case 67:
						fmt.Println("Kevin can retire now!")
		default:
						fmt.Println("There's nothing special about Kevin's current age.")
	}


}
