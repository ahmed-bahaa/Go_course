package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	f, err := os.OpenFile("/home/vegon/Desktop/Go/Go_course/temp/temp_file.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error in opening the file")
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the world:  \n")
	phrase, _ := reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)
	_, err = f.Write([]byte(phrase))
	fmt.Println(phrase)

	if err != nil {
		fmt.Println("Error: Failed to write to /etc/motd")
		os.Exit(1)
	}

}
