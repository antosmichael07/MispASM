package main

import "fmt"

func functions_to_bytes(lines [][]string, var_types map[string][]byte) []byte {
	functions := []byte{}
	labels := map[string]string{}
	label_count := 0

	for i := range lines {
		if lines[i][0] == "label" {
			if len(lines[i]) > 1 {
				labels[lines[i][1]] = fmt.Sprintf("u8-%d", label_count)
				lines[i] = []string{lines[i][0]}
			}
			label_count++
		}
	}

	for i := range lines {
		if lines[i][0][0] == '_' {
			functions = append(append(functions, []byte(lines[i][0])...), 0)
			i++
			for ; i < len(lines) && lines[i][0][0] != '_'; i++ {
				functions = append(functions, instruction_to_bytes(lines[i], var_types, labels)...)
			}
			functions = append(functions, 255)
		}
	}

	return functions
}
