package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/vishnu/glox/parser"
	"github.com/vishnu/glox/token"
)

func main() {
	// cmdArg := os.Args[1:]
	// if len(cmdArg) > 1 {
	// 	err := fmt.Errorf("error passing multiple args %q", cmdArg)
	// 	fmt.Println(err)
	// } else if len(cmdArg) == 1 {
	// 	runFile(cmdArg[0])
	// } else {
	// 	runPrompt()
	// }
	x := parser.Unary{Operator: token.Token{Type: token.MINUS, Lexeme: "-", Literal: nil, Line: 1}, Right: &parser.Literal{Value: 123}}
	y := parser.Binary{Left: &x, Operator: token.Token{Type: token.STAR, Lexeme: "*", Literal: nil, Line: 1}, Right: &parser.Grouping{Expression: &parser.Literal{Value: 45.67}}}
	fmt.Printf("x: %v\n", parser.AstPrinter(&y))

}

func runFile(filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error while reading the file: " + err.Error())
	}
	run(data)
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		data, _ := reader.ReadBytes('\n')
		run(data)
	}
}

func run(source []byte) {
	var sc = Scanner{source: string(source), line: 0, start: 0, current: 0}
	var tokens = sc.ScanTokens()
	for _, val := range tokens {
		val.ToString()
	}
}
