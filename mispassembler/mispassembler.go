package main

import "fmt"

func Compile(data []byte) {
	lines := parse_file(string(data))

	header := get_header(lines)
	fmt.Println(header)
}
