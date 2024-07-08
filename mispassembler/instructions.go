package main

import "strings"

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

func instruction_to_bytes(line []string, var_types map[string][]byte) []byte {
	instruction := []byte{instructions[line[0]]}
	for i := 1; i < len(line); i++ {
		index := strings.Index(line[i], "-")
		if index == -1 {
			instruction = append(instruction, value_to_byte[12](line[i], var_types)...)
			continue
		}
		instruction = append(append(instruction, types[line[i][:index]]), value_to_byte[types[line[i][:index]]](line[i][index+1:], var_types)...)
	}
	return instruction
}
