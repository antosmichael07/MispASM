package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile(os.Args[1])
	fmt.Print(split_file(string(data)))
}
