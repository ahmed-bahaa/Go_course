package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Hello, World. \n")
	phrase, _ := reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)
	fmt.Println(phrase)

}
