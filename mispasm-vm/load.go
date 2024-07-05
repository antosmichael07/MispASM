package main

func load_program(data []byte) (global string, funcs map[string]Function) {
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

		constants = append(constants, Constant{Name: const_name, Value: data[i : i+type_size+1]})
		i += type_size
	}

	return get_functions(data[i+1:])
}
