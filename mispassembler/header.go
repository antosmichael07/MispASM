package main

func get_header(lines [][]string, var_types *map[string][]byte) []byte {
	constants := []byte{}
	variables := []byte{}
	global := []byte{}

	is_const := false
	is_var := false

	for i := range lines {
		if lines[i][0] == "const:" {
			constants = []byte{}
			is_const = true
			i++
			for ; lines[i][0] != "var:" && lines[i][0] != "global"; i++ {
				c_type := get_type(lines[i][1], *var_types)
				constants = append(constants, c_type)
				constants = append(constants, []byte(lines[i][0])...)
				constants = append(constants, 0)
				constants = append(constants, c_type)
				constants = append(constants, value_to_byte(lines[i][1], map[string][]byte{})...)

				(*var_types)[lines[i][0]] = []byte{12, c_type}
			}
			constants = append(constants, 255)
			i++
		} else if !is_const {
			constants = []byte{255}
		}
		if lines[i][0] == "var:" {
			variables = []byte{}
			is_var = true
			i++
			for ; lines[i][0] != "global"; i++ {
				c_type := get_type(lines[i][1], *var_types)
				variables = append(variables, c_type)
				variables = append(variables, []byte(lines[i][0])...)
				variables = append(variables, 0)
				variables = append(variables, c_type)
				if type_sizes[c_type] >= 1 {
					variables = append(variables, value_to_byte(lines[i][1], map[string][]byte{})...)
				}

				(*var_types)[lines[i][0]] = []byte{13, c_type}
			}
			variables = append(variables, 255)
			i++
		} else if !is_var {
			variables = []byte{255}
		}
		if lines[i][0] == "global" {
			global = append(global, []byte(lines[i][1])...)
			global = append(global, 0)
			i++
		}
	}

	return append(append(constants, variables...), global...)
}
