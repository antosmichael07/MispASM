package main

type Stack struct {
	name  byte
	index int
	data  any
}

var stack = []Stack{}

func stack_push(reg byte, index int) {
	stack = append(stack, Stack{reg, index, register_get[reg](index)})
}

func stack_pop(reg byte, index int) {
	for i, v := range stack {
		if v.name == reg && v.index == index {
			stack = append(stack[:i], stack[i+1:]...)
			break
		}
	}
}
