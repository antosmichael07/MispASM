package main

import "fmt"

func functions_to_bytes(lines [][]string, var_types map[string][]byte) []byte {
	functions := []byte{}
	labels := map[string]string{}

	for i := range lines {
		if lines[i][0] == "label" {
			labels[lines[i][1]] = fmt.Sprintf("u8-%d", len(labels))
			lines[i] = []string{lines[i][0]}
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
