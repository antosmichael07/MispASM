package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <mispasm file>\n", os.Args[0])
		os.Exit(1)
	}

	data, _ := os.ReadFile(os.Args[1])
	Compile(data)
}
