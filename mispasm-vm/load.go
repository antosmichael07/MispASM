package main

func (p *Program) load_program() {
	i := 0
	for ; p.data[i] != 255; i++ {
		const_name := []byte{}
		if p.data[i] == 0 {
			const_name = []byte{0}
			i++
		}
		for ; p.data[i] != 0; i++ {
			const_name = append(const_name, p.data[i])
		}
		i++
		type_size := 0
		if p.data[i] != t_string {
			type_size = int(type_sizes[p.data[i]])
		} else {
			for j := i + 1; j < len(p.data); j++ {
				if p.data[j] == 0 {
					type_size = j - i
					break
				}
			}
		}

		p.constants = append(p.constants, constant{name: const_name, value: p.data[i : i+type_size+1]})
		i += type_size
	}
	i++
	for ; p.data[i] != 255; i++ {
		var_name := []byte{}
		if p.data[i] == 0 {
			var_name = []byte{0}
			i++
		}
		for ; p.data[i] != 0; i++ {
			var_name = append(var_name, p.data[i])
		}
		i++
		type_size := 0
		if p.data[i] != t_string {
			type_size = int(type_sizes[p.data[i]])
		} else {
			for j := i + 1; j < len(p.data); j++ {
				if p.data[j] == 0 {
					type_size = j - i
					break
				}
			}
		}

		p.variables = append(p.variables, variable{name: var_name, value: p.data[i : i+type_size+1]})
		i += type_size
	}

	p.global, p.funcs = get_functions(p.data[i+1:])
}
