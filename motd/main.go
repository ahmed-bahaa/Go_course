package main

import (
	"fmt"
	"motd/message"
)

// first way

func main() {
	mesage := message.Greeting("ahmed", "hello")
	fmt.Println(mesage)
}

// note that this function should be under go paths
// /usr/local/go/src/motd/message (from $GOROOT)
// /home/vegon/go/src/motd/message (from $GOPATH
//
// package message
//
// import "fmt"
//
// func Greeting(name string, welcoming string) string {
// 	return fmt.Sprintf("%s, %s", welcoming, name)
// }

//second way

// func main() {
// 	mesage := greeting("ahmed", "hello")
// 	fmt.Println(mesage)
// }
//
// func greeting( name string , welcoming string) string{
// 	return fmt.Sprintf( "%s, %s" , welcoming , name)
// }
