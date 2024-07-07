package main

import (
	"fmt"
	"os"
	"strings"
)

func Compile(data []byte) {
	lines := parse_file(string(data))

	var_types := map[string][]byte{}

	header := get_header(lines, &var_types)
	functions := functions_to_bytes(lines, var_types)
	fmt.Println(append(header, functions...))

	os.WriteFile(os.Args[1][:strings.LastIndex(os.Args[1], ".")]+".mexe", append(header, functions...), 0644)
}
