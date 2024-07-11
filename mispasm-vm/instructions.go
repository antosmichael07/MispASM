package main

import (
	"strings"
	"time"
)

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
	set
	and
	or
	xor
	shr
	shl
	not
	hlt

	instruction_count
)

var arg_sizes = []byte{2, 2, 1, 2, 2, 1, 1, 1, 2, 0, 1, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 0, 0, 2, 2, 2, 2, 2, 2, 1, 0}

func (p *Program) init_instructions() {
	p.instructions[add] = func(arg1 []byte, arg2 []byte, _ *function, _ *int, _ *int) {
		add_math_operation[arg1[0]](arg1, arg2, p)
	}
	p.instructions[sub] = func(arg1 []byte, arg2 []byte, _ *function, _ *int, _ *int) {
		sub_math_operation[arg1[0]](arg1, arg2, p)
	}
	p.instructions[fcal] = func(arg1 []byte, _ []byte, _ *function, _ *int, _ *int) {
		run_function(p.instructions, p.funcs[string(arg1[1:len(arg1)-1])], &p.should_close)
	}
	p.instructions[mul] = func(arg1 []byte, arg2 []byte, _ *function, _ *int, _ *int) {
		mul_math_operation[arg1[0]](arg1, arg2, p)
	}
	p.instructions[div] = func(arg1 []byte, arg2 []byte, _ *function, _ *int, _ *int) {
		div_math_operation[arg1[0]](arg1, arg2, p)
	}
	p.instructions[call] = func(arg1 []byte, _ []byte, _ *function, _ *int, _ *int) {
		cal := convert_to_value[t_string](arg1, *p).(string)
		p.call(cal[:strings.Index(cal, ".")], cal[strings.Index(cal, ".")+1:])
	}
	p.instructions[push] = func(arg1 []byte, _ []byte, _ *function, _ *int, _ *int) {
		stack_push(arg1[1], arg1[2], p)
	}
	p.instructions[pop] = func(arg1 []byte, _ []byte, _ *function, _ *int, _ *int) {
		stack_pop(arg1[1], arg1[2], p)
	}
	p.instructions[mov] = func(arg1 []byte, arg2 []byte, _ *function, _ *int, _ *int) {
		if arg2[0] < t_reg {
			register_force_set[arg2[0]][arg1[1]](arg1[2], convert_to_value[arg2[0]](arg2, *p), p)
		} else if arg2[0] == t_reg {
			register_force_set[arg2[1]%11][arg1[1]](arg1[2], convert_to_value[arg2[0]](arg2, *p), p)
		} else {
			val := convert_to_value[arg2[0]](arg2, *p).([]byte)
			register_force_set[arg2[1]][arg1[1]](arg1[2], convert_to_value[val[0]](val, *p), p)
		}
	}
	p.instructions[label] = func(_ []byte, _ []byte, _ *function, _ *int, _ *int) {}
	p.instructions[jmp] = func(arg1 []byte, _ []byte, function *function, i *int, arg_size *int) {
		*i = function.labels[int(arg1[1])]
		*arg_size = 0
	}
	p.instructions[mod] = func(arg1 []byte, arg2 []byte, _ *function, _ *int, _ *int) {
		mod_math_operation[arg1[0]](arg1, arg2, p)
	}
	p.instructions[cmp] = func(arg1 []byte, arg2 []byte, _ *function, _ *int, _ *int) {
		p.register_cmp[0] = arg1
		p.register_cmp[1] = arg2
	}
	p.instructions[je] = func(arg1 []byte, arg2 []byte, function *function, i *int, arg_size *int) {
		compare[0](arg1, arg2, function, i, *p)
		*arg_size = 0
	}
	p.instructions[jne] = func(arg1 []byte, arg2 []byte, function *function, i *int, arg_size *int) {
		compare[1](arg1, arg2, function, i, *p)
		*arg_size = 0
	}
	p.instructions[jg] = func(arg1 []byte, arg2 []byte, function *function, i *int, arg_size *int) {
		compare[2](arg1, arg2, function, i, *p)
		*arg_size = 0
	}
	p.instructions[jge] = func(arg1 []byte, arg2 []byte, function *function, i *int, arg_size *int) {
		compare[3](arg1, arg2, function, i, *p)
		*arg_size = 0
	}
	p.instructions[jl] = func(arg1 []byte, arg2 []byte, function *function, i *int, arg_size *int) {
		compare[4](arg1, arg2, function, i, *p)
		*arg_size = 0
	}
	p.instructions[jle] = func(arg1 []byte, arg2 []byte, function *function, i *int, arg_size *int) {
		compare[5](arg1, arg2, function, i, *p)
		*arg_size = 0
	}
	p.instructions[inc] = func(arg1 []byte, _ []byte, _ *function, _ *int, _ *int) {
		inc_math_operation[arg1[1]](p)
	}
	p.instructions[dec] = func(arg1 []byte, _ []byte, _ *function, _ *int, _ *int) {
		dec_math_operation[arg1[1]](p)
	}
	p.instructions[def] = func(_ []byte, _ []byte, _ *function, _ *int, _ *int) {
		p.should_close = true
	}
	p.instructions[set] = func(arg1 []byte, arg2 []byte, _ *function, _ *int, _ *int) {
		p.set_var_value(arg1, arg2)
	}
	p.instructions[and] = func(arg1 []byte, arg2 []byte, _ *function, _ *int, _ *int) {
		and_bitwise_operation[arg1[0]](arg1, arg2, p)
	}
	p.instructions[or] = func(arg1 []byte, arg2 []byte, _ *function, _ *int, _ *int) {
		or_bitwise_operation[arg1[0]](arg1, arg2, p)
	}
	p.instructions[xor] = func(arg1 []byte, arg2 []byte, _ *function, _ *int, _ *int) {
		xor_bitwise_operation[arg1[0]](arg1, arg2, p)
	}
	p.instructions[shr] = func(arg1 []byte, arg2 []byte, _ *function, _ *int, _ *int) {
		shr_bitwise_operation[arg1[0]](arg1, arg2, p)
	}
	p.instructions[shl] = func(arg1 []byte, arg2 []byte, _ *function, _ *int, _ *int) {
		shl_bitwise_operation[arg1[0]](arg1, arg2, p)
	}
	p.instructions[not] = func(arg1 []byte, _ []byte, _ *function, _ *int, _ *int) {
		not_bitwise_operation[arg1[0]](arg1, p)
	}
	p.instructions[hlt] = func(_ []byte, _ []byte, _ *function, _ *int, _ *int) {
		time.Sleep(100 * 365 * 24 * time.Hour)
	}
}

func get_args(f_instructions []byte, i int) (arg1 []byte, arg2 []byte, is_arg1 bool, is_arg2 bool, arg_size int) {
	offset := 0
	if arg_sizes[f_instructions[i]] >= 1 {
		is_arg1 = true
		if f_instructions[i+1] == t_string || f_instructions[i+1] >= t_const {
			for j := i + 1; j < len(f_instructions); j++ {
				if f_instructions[j] == 0 {
					if j == i+2 && f_instructions[i+1] >= t_const {
						continue
					}
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
		if f_instructions[offset] == t_string || f_instructions[offset] >= t_const {
			for j := offset; j < len(f_instructions); j++ {
				if f_instructions[j] == 0 {
					if j == offset+1 && f_instructions[offset] >= t_const {
						continue
					}
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
