package main

type variable struct {
	name  []byte
	value []byte
}

func (program Program) get_var_value(name []byte) []byte {
	for _, v := range program.variables {
		for i := 0; i < len(v.name); i++ {
			if v.name[i] != name[i] {
				continue
			}
		}
		return v.value
	}
	return []byte{}
}

func (program Program) set_var_value(name []byte, value []byte) {
	for i, v := range program.variables {
		for j := 0; j < len(v.name); j++ {
			if v.name[j] != name[j] {
				continue
			}
		}
		program.variables[i].value = value
	}
}
