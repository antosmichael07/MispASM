package main

type Function struct {
	instructions []byte
	labels       []int
}

func get_functions(data []byte) (global string, funcs map[string]Function) {
	i := 0
	for ; data[i] != 0; i++ {
		global += string(data[i])
	}
	funcs = make(map[string]Function)
	for i += 1; i < len(data); i++ {
		name := ""
		for ; data[i] != 0; i++ {
			name += string(data[i])
		}

		instructions := []byte{}
		for i += 1; data[i] != 255; {
			_, _, _, _, arg_size := get_args(data, i)
			for j := i; j <= i+arg_size; j++ {
				instructions = append(instructions, data[j])
			}
			i += arg_size + 1
		}

		labels := []int{}
		for j := 0; j < len(instructions); {
			_, _, _, _, arg_size := get_args(instructions, j)
			if instructions[j] == label {
				labels = append(labels, j)
			}
			j += arg_size + 1
		}

		funcs[name] = Function{
			instructions: instructions,
			labels:       labels,
		}
	}

	return global, funcs
}

func run_function(function Function) {
	for i := 0; i < len(function.instructions); i++ {
		if should_close {
			return
		}
		if function.instructions[i] == ret {
			return
		}
		arg1, arg2, is_arg1, is_arg2, arg_size := get_args(function.instructions, i)
		if is_arg1 && is_arg2 {
			instructions[function.instructions[i]](arg1, arg2, &function, &i, &arg_size)
		} else if is_arg1 {
			instructions[function.instructions[i]](arg1, []byte{}, &function, &i, &arg_size)
		} else {
			instructions[function.instructions[i]]([]byte{}, []byte{}, &Function{}, &i, &arg_size)
		}
		i += arg_size
	}
}
