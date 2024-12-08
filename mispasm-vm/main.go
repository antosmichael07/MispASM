package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <executable>\n", os.Args[0])
		os.Exit(1)
	}

	program := NewProgram(read_exec(os.Args[1]))
	program.Run()
}

func read_exec(location string) (data []byte) {
	data, err := os.ReadFile(location)
	if err != nil {
		fmt.Printf("Error reading executable file\n")
		os.Exit(1)
	}
	return data
}
