package main

import (
	"fmt"
)

func Compile(data []byte) []byte {
	lines := parse_file(string(data))

	var_types := map[string][]byte{}

	header := get_header(lines, &var_types)
	functions := functions_to_bytes(lines, var_types)
	fmt.Println(append(header, functions...))

	return append(header, functions...)
}
