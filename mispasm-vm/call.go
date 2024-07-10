package main

func (p *Program) init_calls() {
	for _, f := range p.funcs {
		for j := 0; j < len(f.instructions); j++ {
			arg1, _, _, _, arg_size := get_args(f.instructions, j)
			if f.instructions[j] == call {
				lib := convert_to_value[t_string](arg1[:len(arg1)-1], *p).(string)
				if arg1[1] != '#' && !p.is_lib_loaded(lib) {
					p.load_lib(lib)
				}
			}
			j += arg_size
		}
	}
}

/*func (p *Program) sprint_stack() string {
	msg := ""
	for _, v := range p.stack {
		msg = fmt.Sprint(msg, v.data)
	}
	return msg
}*/

/*var scanf_func = [11]func(*Program){
	func(p *Program) {
		var data int8
		fmt.Scanf(p.stack[0].data.(string), &data)
		register_set[p.stack[1].name](0, data, p)
	},
	func(p *Program) {
		var data int16
		fmt.Scanf(p.stack[0].data.(string), &data)
		register_set[p.stack[1].name](0, data, p)
	},
	func(p *Program) {
		var data int32
		fmt.Scanf(p.stack[0].data.(string), &data)
		register_set[p.stack[1].name](0, data, p)
	},
	func(p *Program) {
		var data int64
		fmt.Scanf(p.stack[0].data.(string), &data)
		register_set[p.stack[1].name](0, data, p)
	},
	func(p *Program) {
		var data uint8
		fmt.Scanf(p.stack[0].data.(string), &data)
		register_set[p.stack[1].name](0, data, p)
	},
	func(p *Program) {
		var data uint16
		fmt.Scanf(p.stack[0].data.(string), &data)
		register_set[p.stack[1].name](0, data, p)
	},
	func(p *Program) {
		var data uint32
		fmt.Scanf(p.stack[0].data.(string), &data)
		register_set[p.stack[1].name](0, data, p)
	},
	func(p *Program) {
		var data uint64
		fmt.Scanf(p.stack[0].data.(string), &data)
		register_set[p.stack[1].name](0, data, p)
	},
	func(p *Program) {
		var data float32
		fmt.Scanf(p.stack[0].data.(string), &data)
		register_set[p.stack[1].name](0, data, p)
	},
	func(p *Program) {
		var data float64
		fmt.Scanf(p.stack[0].data.(string), &data)
		register_set[p.stack[1].name](0, data, p)
	},
	func(p *Program) {
		var data string
		fmt.Scanf(p.stack[0].data.(string), &data)
		register_set[p.stack[1].name](0, data, p)
	},
}*/
