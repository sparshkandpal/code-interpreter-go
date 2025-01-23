package main

import (
	"fmt"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")
	hasError := false

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]
	
	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	rawfileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	
	fileContents := string(rawfileContents)
	for i := 0; i < len(fileContents); i++ {
		current := fileContents[i]
		var next = rune(0)
		if i+1 < len(fileContents) {
			next = rune(fileContents[i+1])
		}
		switch {
		case current == '=' && next == '=':
			fmt.Println("EQUAL_EQUAL == null")
			i++
		case current == '!' && next == '=':
			fmt.Println("BANG_EQUAL != null")
			i++
		case current == '!' && next == '=':
			fmt.Println("BANG_EQUAL != null")
			i++
		case current == '<' && next == '=':
			fmt.Println("LESS_EQUAL <= null")
			i++
		case current == '>' && next == '=':
			fmt.Println("GREATER_EQUAL >= null")
			i++
		case current == '!':
			fmt.Println("BANG ! null")	
		case current == '=':
			fmt.Println("EQUAL = null")
		case current == '<':
			fmt.Println("LESS < null")
		case current == '>':
			fmt.Println("GREATER > null")
		case current == '(':
			fmt.Println("LEFT_PAREN ( null")
		case current == ')':
			fmt.Println("RIGHT_PAREN ) null")
		case current == '{':
			fmt.Println("LEFT_BRACE { null")
		case current == '}':
			fmt.Println("RIGHT_BRACE } null")
		case current == '[':
			fmt.Println("LEFT_BRACKET [ null")
		case current == ']':
			fmt.Println("RIGHT_BRACKET ] null")
		case current == ',':
			fmt.Println("COMMA , null")
		case current == '.':
			fmt.Println("DOT . null")
		case current == '-':
			fmt.Println("MINUS - null")
		case current == '+':
			fmt.Println("PLUS + null")
		case current == '*':
			fmt.Println("STAR * null")
		case current == '/':
			fmt.Println("SLASH / null")
		case current == ';':
			fmt.Println("SEMICOLON ; null")
		default:
			fmt.Fprintf(os.Stderr, "[line 1] Error: Unexpected character: %c\n", current)
			hasError = true
		}
	}
	fmt.Println("EOF  null")

	if hasError {
		os.Exit(65)
	} else {
		os.Exit(0)
	}
}
