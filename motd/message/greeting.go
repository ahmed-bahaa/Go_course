package message

import "fmt"

func Greeting(name string, welcoming string) string {
	return fmt.Sprintf("%s, %s", welcoming, name)
}
