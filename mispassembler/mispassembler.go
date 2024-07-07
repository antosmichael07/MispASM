package main

import "fmt"

func Compile(data []byte) {
	lines := parse_file(string(data))

	header := get_header(lines)
	functions := functions_to_bytes(lines)
	fmt.Println(append(header, functions...))
}
