package main

import (
	"fmt"
	"os"
)

func main() {
	program := NewProgram(read_exec(os.Args[1]))
	program.Run()
}

func read_exec(location string) (data []byte) {
	data, err := os.ReadFile(location)
	if err != nil {
		fmt.Printf("Error reading executable file\n")
		panic(err)
	}
	return data
}
