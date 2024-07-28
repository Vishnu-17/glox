package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	cmdArg := os.Args[1:]
	if len(cmdArg) > 1 {
		err := fmt.Errorf("error passing multiple args %q", cmdArg)
		fmt.Println(err)
	} else if len(cmdArg) == 1 {
		runFile(cmdArg[0])
	} else {
		runPrompt()
	}
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
