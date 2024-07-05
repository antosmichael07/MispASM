package main

type function struct {
	instructions []byte
	labels       []int
}

func get_functions(data []byte) (global string, funcs map[string]function) {
	i := 0
	for ; data[i] != 0; i++ {
		global += string(data[i])
	}
	funcs = make(map[string]function)
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

		funcs[name] = function{
			instructions: instructions,
			labels:       labels,
		}
	}

	return global, funcs
}

func run_function(f function, should_close *bool) {
	for i := 0; i < len(f.instructions); i++ {
		if *should_close {
			return
		}
		if f.instructions[i] == ret {
			return
		}
		arg1, arg2, is_arg1, is_arg2, arg_size := get_args(f.instructions, i)
		if is_arg1 && is_arg2 {
			instructions[f.instructions[i]](arg1, arg2, &f, &i, &arg_size)
		} else if is_arg1 {
			instructions[f.instructions[i]](arg1, []byte{}, &f, &i, &arg_size)
		} else {
			instructions[f.instructions[i]]([]byte{}, []byte{}, &function{}, &i, &arg_size)
		}
		i += arg_size
	}
}
