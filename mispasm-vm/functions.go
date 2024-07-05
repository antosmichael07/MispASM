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

func get_args(f_instructions []byte, i int) (arg1 []byte, arg2 []byte, is_arg1 bool, is_arg2 bool, arg_size int) {
	offset := 0
	if arg_sizes[f_instructions[i]] >= 1 {
		is_arg1 = true
		if f_instructions[i+1] == t_string {
			for j := i + 1; j < len(f_instructions); j++ {
				if f_instructions[j] == 0 {
					offset = j + 1
					arg_size = j - i
					arg1 = f_instructions[i+1 : j+1]
					break
				}
			}
		} else {
			offset = i + int(type_sizes[f_instructions[i+1]]) + 2
			arg_size = int(type_sizes[f_instructions[i+1]] + 1)
			arg1 = f_instructions[i+1 : offset]
		}
	}

	if arg_sizes[f_instructions[i]] >= 2 {
		is_arg2 = true
		if f_instructions[offset] == t_string {
			for j := offset; j < len(f_instructions); j++ {
				if f_instructions[j] == 0 {
					arg_size = j - i
					arg2 = f_instructions[offset : j+1]
					break
				}
			}
		} else {
			arg_size = offset + int(type_sizes[f_instructions[offset]]) - i
			arg2 = f_instructions[offset : offset+int(type_sizes[f_instructions[offset]])+1]
		}
	}

	return arg1, arg2, is_arg1, is_arg2, arg_size
}
