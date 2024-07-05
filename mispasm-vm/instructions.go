package main

const (
	add byte = iota
	sub
	fcal
	mul
	div
	call
	push
	pop
	mov
	label
	jmp
	mod
	cmp
	je
	jne
	jg
	jge
	jl
	jle
	inc
	dec
	ret
	def
)

var arg_sizes = []byte{2, 2, 1, 2, 2, 1, 1, 1, 2, 0, 1, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 0, 0}

var instructions = [23]func([]byte, []byte, *function, *int, *int){}

func init_instructions(p *Program, should_close *bool) {
	instructions[add] = func(arg1 []byte, arg2 []byte, _ *function, _ *int, _ *int) {
		add_math_operation[arg1[0]](arg1, arg2, p)
	}
	instructions[sub] = func(arg1 []byte, arg2 []byte, _ *function, _ *int, _ *int) {
		sub_math_operation[arg1[0]](arg1, arg2, p)
	}
	instructions[fcal] = func(arg1 []byte, _ []byte, _ *function, _ *int, _ *int) {
		run_function(p.funcs[string(arg1[1:len(arg1)-1])], should_close)
	}
	instructions[mul] = func(arg1 []byte, arg2 []byte, _ *function, _ *int, _ *int) {
		mul_math_operation[arg1[0]](arg1, arg2, p)
	}
	instructions[div] = func(arg1 []byte, arg2 []byte, _ *function, _ *int, _ *int) {
		div_math_operation[arg1[0]](arg1, arg2, p)
	}
	instructions[call] = func(arg1 []byte, _ []byte, _ *function, _ *int, _ *int) {
		calls[arg1[1]]([]byte{}, []byte{})
	}
	instructions[push] = func(arg1 []byte, _ []byte, _ *function, _ *int, _ *int) {
		stack_push(arg1[1], int(arg1[2]), p)
	}
	instructions[pop] = func(arg1 []byte, _ []byte, _ *function, _ *int, _ *int) {
		stack_pop(arg1[1], int(arg1[2]), p)
	}
	instructions[mov] = func(arg1 []byte, arg2 []byte, _ *function, _ *int, _ *int) {
		if arg2[0] != t_reg {
			register_force_set[arg2[0]][arg1[1]](arg1[2], convert_to_value[arg2[0]](arg2, *p), p)
		} else {
			register_force_set[arg2[1]%11][arg1[1]](arg1[2], convert_to_value[arg2[0]](arg2, *p), p)
		}
	}
	instructions[label] = func(_ []byte, _ []byte, _ *function, _ *int, _ *int) {}
	instructions[jmp] = func(arg1 []byte, _ []byte, _ *function, i *int, arg_size *int) {
		*i = int(arg1[1])
		*arg_size = 0
	}
	instructions[mod] = func(arg1 []byte, arg2 []byte, _ *function, _ *int, _ *int) {
		mod_math_operation[arg1[0]](arg1, arg2, p)
	}
	instructions[cmp] = func(arg1 []byte, arg2 []byte, _ *function, _ *int, _ *int) {
		p.register_cmp[0] = arg1
		p.register_cmp[1] = arg2
	}
	instructions[je] = func(arg1 []byte, arg2 []byte, function *function, i *int, arg_size *int) {
		compare[0](arg1, arg2, function, i, *p)
		*arg_size = 0
	}
	instructions[jne] = func(arg1 []byte, arg2 []byte, function *function, i *int, arg_size *int) {
		compare[1](arg1, arg2, function, i, *p)
		*arg_size = 0
	}
	instructions[jg] = func(arg1 []byte, arg2 []byte, function *function, i *int, arg_size *int) {
		compare[2](arg1, arg2, function, i, *p)
		*arg_size = 0
	}
	instructions[jge] = func(arg1 []byte, arg2 []byte, function *function, i *int, arg_size *int) {
		compare[3](arg1, arg2, function, i, *p)
		*arg_size = 0
	}
	instructions[jl] = func(arg1 []byte, arg2 []byte, function *function, i *int, arg_size *int) {
		compare[4](arg1, arg2, function, i, *p)
		*arg_size = 0
	}
	instructions[jle] = func(arg1 []byte, arg2 []byte, function *function, i *int, arg_size *int) {
		compare[5](arg1, arg2, function, i, *p)
		*arg_size = 0
	}
	instructions[inc] = func(arg1 []byte, _ []byte, _ *function, _ *int, _ *int) {
		inc_math_operation[arg1[1]](p)
	}
	instructions[dec] = func(arg1 []byte, _ []byte, _ *function, _ *int, _ *int) {
		dec_math_operation[arg1[1]](p)
	}
	instructions[def] = func(_ []byte, _ []byte, _ *function, _ *int, _ *int) {
		*should_close = true
	}
}

func get_args(f_instructions []byte, i int) (arg1 []byte, arg2 []byte, is_arg1 bool, is_arg2 bool, arg_size int) {
	offset := 0
	if arg_sizes[f_instructions[i]] >= 1 {
		is_arg1 = true
		if f_instructions[i+1] == t_string || f_instructions[i+1] == t_const {
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
		if f_instructions[offset] == t_string || f_instructions[offset] == t_const {
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
