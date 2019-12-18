package main

import "fmt"

func main(){

	boyznoize := map[string]int{
		"alarm" : 1,
		"midnight" : 2,
	}

	boyznoize["loaded"] = 3

	fmt.Println(boyznoize)
	fmt.Println(boyznoize["loaded"])
	delete(boyznoize, "midnight")
	fmt.Println(boyznoize)
	
}
