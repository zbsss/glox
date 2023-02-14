package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func run(source string) {
	fmt.Println(source)
}

func runFile(path string) {
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("Error opening file: ", err)
		os.Exit(74)
	}

	run(string(bytes))
}

func runPrompt() {
	fmt.Println("Glox REPL")
	fmt.Println("Ctrl+C to exit")

	var input string
	for {
		fmt.Print("> ")
		fmt.Scanln(&input)
		run(input)
	}
}

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: glox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		runPrompt()
	}
}
