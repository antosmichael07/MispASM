package main

import (
	"strings"
)

func get_header(lines [][]string) []byte {
	constants := []byte{}
	variables := []byte{}
	global := []byte{}

	for i := range lines {
		if lines[i][0] == "const:" {
			i++
			for ; lines[i][0] != "var:" && lines[i][0] != "global"; i++ {
				c_type := get_type(lines[i][1])
				constants = append(constants, c_type)
				constants = append(constants, []byte(lines[i][0])...)
				constants = append(constants, 0)
				constants = append(constants, c_type)
				if type_sizes[c_type] >= 1 {
					constants = append(constants, value_to_byte[c_type](lines[i][1][strings.Index(lines[i][1], "-")+1:])...)
				}
			}
			constants = append(constants, 255)
			i++
		}
		if lines[i][0] == "var:" {
			i++
			for ; lines[i][0] != "global"; i++ {
				c_type := get_type(lines[i][1])
				variables = append(variables, c_type)
				variables = append(variables, []byte(lines[i][0])...)
				variables = append(variables, 0)
				variables = append(variables, c_type)
				if type_sizes[c_type] >= 1 {
					variables = append(variables, value_to_byte[c_type](lines[i][1][strings.Index(lines[i][1], "-")+1:])...)
				}
			}
			variables = append(variables, 255)
			i++
		}
		if lines[i][0] == "global" {
			global = append(global, []byte(lines[i][1])...)
			global = append(global, 0)
			i++
		}
	}

	return append(append(constants, variables...), global...)
}
