package main

type Constant struct {
	Name  []byte
	Value []byte
}

var constants = []Constant{}

func get_const_value(name []byte) []byte {
	for _, c := range constants {
		for i := 0; i < len(c.Name); i++ {
			if c.Name[i] != name[i] {
				continue
			}
		}
		return c.Value
	}
	return []byte{}
}
