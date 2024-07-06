package main

func load_program(data []byte) (global string, funcs map[string]function, constants []constant, variables []variable) {
	i := 0
	for ; data[i] != 255; i++ {
		const_name := []byte{}
		if data[i] == 0 {
			const_name = []byte{0}
			i++
		}
		for ; data[i] != 0; i++ {
			const_name = append(const_name, data[i])
		}
		i++
		type_size := 0
		if data[i] != t_string {
			type_size = int(type_sizes[data[i]])
		} else {
			for j := i + 1; j < len(data); j++ {
				if data[j] == 0 {
					type_size = j - i
					break
				}
			}
		}

		constants = append(constants, constant{name: const_name, value: data[i : i+type_size+1]})
		i += type_size
	}
	i++
	for ; data[i] != 255; i++ {
		var_name := []byte{}
		if data[i] == 0 {
			var_name = []byte{0}
			i++
		}
		for ; data[i] != 0; i++ {
			var_name = append(var_name, data[i])
		}
		i++
		type_size := 0
		if data[i] != t_string {
			type_size = int(type_sizes[data[i]])
		} else {
			for j := i + 1; j < len(data); j++ {
				if data[j] == 0 {
					type_size = j - i
					break
				}
			}
		}

		variables = append(variables, variable{name: var_name, value: data[i : i+type_size+1]})
		i += type_size
	}

	global, funcs = get_functions(data[i+1:])
	return global, funcs, constants, variables
}
