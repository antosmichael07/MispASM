package main

import (
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile(os.Args[1])
	program := Compile(data)

	os.WriteFile(os.Args[1][:strings.LastIndex(os.Args[1], ".")]+".mexe", program, 0644)
}
