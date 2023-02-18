package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/zbsss/glox/errors.go"
	"github.com/zbsss/glox/scanner"
)

func runFile(path string) {
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("Error opening file: ", err)
		os.Exit(74)
	}

	run(string(bytes))

	if errors.HadError {
		os.Exit(65)
	}
}

func runPrompt() {
	fmt.Println("Glox REPL")
	fmt.Println("Ctrl+C to exit")

	var input string
	for {
		fmt.Print("> ")
		fmt.Scanln(&input)

		run(input)

		errors.HadError = false
	}
}

func run(source string) {
	sc := scanner.NewScanner(source)

	tokens := sc.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token)
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
