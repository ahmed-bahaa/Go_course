package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	// define Flags or options
	var name string
	var greeting string
	var prompt bool
	var preview bool

	flag.StringVar(&name, "name", "", "name to use within the message")
	flag.StringVar(&greeting, "greeting", "", "phrase to use within the greeting")
	flag.BoolVar(&prompt, "prompt", false, "use prompt to input name and message")
	flag.BoolVar(&preview, "preview", false, "use preview to output message without writing to /temp/temp_file.txt")

	// Parse Flags
	flag.Parse()

	// Show how to use flags if the flags aren't valid and exit
	if prompt == false && (name == "" || greeting == "") {
		flag.Usage()
		os.Exit(1)
	}

	// Optionally print flags and exit if DEBUG is set
	//$ DEBUG=true go run main.go -name "ahmed" -greeting "hi"
	if os.Getenv("DEBUG") != "" {
		fmt.Println("Name:", name)
		fmt.Println("Greeting:", greeting)
		fmt.Println("Prompt:", prompt)
		fmt.Println("Preview:", preview)

		os.Exit(0)
	}

	// Conditionally read from stdin
	if prompt {
		name, greeting = renderPrompt()
	}

	// Generate message
	m := fmt.Sprintf("%s, %s", greeting, name)

	// Either preview message or write to file
	if preview {
		fmt.Println(m)
	} else {
		// Write content
		f, err := os.OpenFile("/etc/motd", os.O_WRONLY, 0644)

		if err != nil {
			fmt.Println("Error: Unable to open to /etc/motd")
			os.Exit(1)
		}

		defer f.Close()

		_, err = f.Write([]byte(m))

		if err != nil {
			fmt.Println("Error: Failed to write to /etc/motd")
			os.Exit(1)
		}
	}

}

func renderPrompt() (name, greeting string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("your Greeting: ")
	greeting, _ = reader.ReadString('\n')
	greeting = strings.TrimSpace(greeting)

	fmt.Print("your Name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	return
}
