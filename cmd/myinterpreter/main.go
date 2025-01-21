package main

import (
	"fmt"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	const (
		LEFT_PAREN  rune = '('
		RIGHT_PAREN rune = ')'
		LEFT_BRACE rune = '{'
		RIGHT_BRACE rune = '}'
		LEFT_BRACKET rune = '['
		RIGHT_BRACKET rune = ']'
		COMMA rune = ','
		DOT rune = '.'
		MINUS rune = '-'
		PLUS rune = '+'
		STAR rune = '*'
		SLASH rune = '/'
		SEMICOLON rune = ';'
	)
	
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
	for _, current := range fileContents {
		switch current {
		case LEFT_PAREN:
			fmt.Println("LEFT_PAREN ( null")
		case RIGHT_PAREN:
			fmt.Println("RIGHT_PAREN ) null")
		case LEFT_BRACE:
			fmt.Println("LEFT_BRACE { null")
		case RIGHT_BRACE:
			fmt.Println("RIGHT_BRACE } null")
		case LEFT_BRACKET:
			fmt.Println("LEFT_BRACKET [ null")
		case RIGHT_BRACKET:
			fmt.Println("RIGHT_BRACKET ] null")
		case COMMA:
			fmt.Println("COMMA , null")
		case DOT:
			fmt.Println("DOT . null")
		case MINUS:
			fmt.Println("MINUS - null")
		case PLUS:
			fmt.Println("PLUS + null")
		case STAR:
			fmt.Println("STAR * null")
		case SLASH:
			fmt.Println("SLASH / null")
		case SEMICOLON:
			fmt.Println("SEMICOLON ; null")
		}
	}
	fmt.Println("EOF  null")
}
