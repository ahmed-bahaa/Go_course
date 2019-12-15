package main

import "fmt"

func main() {

	fmt.Println("hello one line")
	fmt.Println(`
	hello 
	multiple lines 
	:)
	`)

	// we can also put emojis here put i dunno the code :DD
	fmt.Println(":D")
	fmt.Println("\u2272")
	// runes ,, it takees only one character and print the ascii code for it
	fmt.Println('L')
}
