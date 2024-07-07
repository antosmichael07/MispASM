package main

func functions_to_bytes(lines [][]string, var_types map[string][]byte) []byte {
	functions := []byte{}

	for i := range lines {
		if lines[i][0][0] == '_' {
			functions = append(append(functions, []byte(lines[i][0])...), 0)
			i++
			for ; i < len(lines) && lines[i][0][0] != '_'; i++ {
				functions = append(functions, instruction_to_bytes(lines[i], var_types)...)
			}
			functions = append(functions, 255)
		}
	}

	return functions
}
