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
}

func instruction_to_bytes(line []string) []byte {
	instruction := []byte{instructions[line[0]]}
	for i := 1; i < len(line); i++ {
		instruction = append(append(instruction, types[line[i][:strings.Index(line[i], "-")]]), value_to_byte[types[line[i][:strings.Index(line[i], "-")]]](line[i][strings.Index(line[i], "-")+1:])...)
	}
	return instruction
}
