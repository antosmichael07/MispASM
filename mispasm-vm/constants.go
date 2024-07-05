package main

type constant struct {
	name  []byte
	value []byte
}

func (program Program) get_const_value(name []byte) []byte {
	for _, c := range program.constants {
		for i := 0; i < len(c.name); i++ {
			if c.name[i] != name[i] {
				continue
			}
		}
		return c.value
	}
	return []byte{}
}
