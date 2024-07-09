package main

var instructions = map[string]byte{
	"add":   0,
	"sub":   1,
	"fcal":  2,
	"mul":   3,
	"div":   4,
	"call":  5,
	"push":  6,
	"pop":   7,
	"mov":   8,
	"label": 9,
	"jmp":   10,
	"mod":   11,
	"cmp":   12,
	"je":    13,
	"jne":   14,
	"jg":    15,
	"jge":   16,
	"jl":    17,
	"jle":   18,
	"inc":   19,
	"dec":   20,
	"ret":   21,
	"def":   22,
	"set":   23,
	"and":   24,
	"or":    25,
	"xor":   26,
	"shr":   27,
	"shl":   28,
	"not":   29,
	"hlt":   30,
}

func instruction_to_bytes(line []string, var_types map[string][]byte, labels map[string]string) []byte {
	instruction := []byte{instructions[line[0]]}
	if instruction[0] == 5 && len(line[1]) >= 3 && line[1][0:3] != "u8-" {
		instruction = append(instruction, calls[line[1]]...)
		return instruction
	}
	if instruction[0] == 10 {
		if line[1][0:3] != "u8-" {
			line[1] = labels[line[1]]
		}
	}
	if instruction[0] >= 13 && instruction[0] <= 18 {
		if line[1][0:3] != "u8-" {
			line[1] = labels[line[1]]
		}
		if line[2][0:3] != "u8-" {
			line[2] = labels[line[2]]
		}
	}
	for i := 1; i < len(line); i++ {
		instruction = append(append(instruction, get_type(line[i], var_types)), value_to_byte(line[i], var_types)...)
	}
	return instruction
}
