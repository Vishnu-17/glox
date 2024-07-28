package main

import "fmt"

func errorReport(line uint, msg string) {
	fmt.Printf("Error in line %d: %s\n", line, msg)
}
