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

var instructions = [23]func([]byte, []byte, *Function, *int, *int){}

func init_instructions(funcs *map[string]Function) {
	instructions[add] = func(arg1 []byte, arg2 []byte, _ *Function, _ *int, _ *int) {
		add_math_operation[arg1[0]](arg1, arg2)
	}
	instructions[sub] = func(arg1 []byte, arg2 []byte, _ *Function, _ *int, _ *int) {
		sub_math_operation[arg1[0]](arg1, arg2)
	}
	instructions[fcal] = func(arg1 []byte, _ []byte, _ *Function, _ *int, _ *int) {
		run_function((*funcs)[string(arg1[1:len(arg1)-1])])
	}
	instructions[mul] = func(arg1 []byte, arg2 []byte, _ *Function, _ *int, _ *int) {
		mul_math_operation[arg1[0]](arg1, arg2)
	}
	instructions[div] = func(arg1 []byte, arg2 []byte, _ *Function, _ *int, _ *int) {
		div_math_operation[arg1[0]](arg1, arg2)
	}
	instructions[call] = func(arg1 []byte, _ []byte, _ *Function, _ *int, _ *int) {
		calls[arg1[1]]([]byte{}, []byte{})
	}
	instructions[push] = func(arg1 []byte, _ []byte, _ *Function, _ *int, _ *int) {
		stack_push(arg1[1], int(arg1[2]))
	}
	instructions[pop] = func(arg1 []byte, _ []byte, _ *Function, _ *int, _ *int) {
		stack_pop(arg1[1], int(arg1[2]))
	}
	instructions[mov] = func(arg1 []byte, arg2 []byte, _ *Function, _ *int, _ *int) {
		if arg2[0] != t_reg {
			register_force_set[arg2[0]][arg1[1]](arg1[2], convert_to_value[arg2[0]](arg2))
		} else {
			register_force_set[arg2[1]%11][arg1[1]](arg1[2], convert_to_value[arg2[0]](arg2))
		}
	}
	instructions[label] = func(_ []byte, _ []byte, _ *Function, _ *int, _ *int) {}
	instructions[jmp] = func(arg1 []byte, _ []byte, _ *Function, i *int, arg_size *int) {
		*i = int(arg1[1])
		*arg_size = 0
	}
	instructions[mod] = func(arg1 []byte, arg2 []byte, _ *Function, _ *int, _ *int) {
		mod_math_operation[arg1[0]](arg1, arg2)
	}
	instructions[cmp] = func(arg1 []byte, arg2 []byte, _ *Function, _ *int, _ *int) {
		register_cmp[0] = arg1
		register_cmp[1] = arg2
	}
	instructions[je] = func(arg1 []byte, arg2 []byte, function *Function, i *int, arg_size *int) {
		compare[0](arg1, arg2, function, i)
		*arg_size = 0
	}
	instructions[jne] = func(arg1 []byte, arg2 []byte, function *Function, i *int, arg_size *int) {
		compare[1](arg1, arg2, function, i)
		*arg_size = 0
	}
	instructions[jg] = func(arg1 []byte, arg2 []byte, function *Function, i *int, arg_size *int) {
		compare[2](arg1, arg2, function, i)
		*arg_size = 0
	}
	instructions[jge] = func(arg1 []byte, arg2 []byte, function *Function, i *int, arg_size *int) {
		compare[3](arg1, arg2, function, i)
		*arg_size = 0
	}
	instructions[jl] = func(arg1 []byte, arg2 []byte, function *Function, i *int, arg_size *int) {
		compare[4](arg1, arg2, function, i)
		*arg_size = 0
	}
	instructions[jle] = func(arg1 []byte, arg2 []byte, function *Function, i *int, arg_size *int) {
		compare[5](arg1, arg2, function, i)
		*arg_size = 0
	}
	instructions[inc] = func(arg1 []byte, _ []byte, _ *Function, _ *int, _ *int) {
		inc_math_operation[arg1[1]]()
	}
	instructions[dec] = func(arg1 []byte, _ []byte, _ *Function, _ *int, _ *int) {
		dec_math_operation[arg1[1]]()
	}
	instructions[def] = func(_ []byte, _ []byte, _ *Function, _ *int, _ *int) {
		should_close = true
	}
}
