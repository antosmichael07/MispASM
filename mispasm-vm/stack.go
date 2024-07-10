package main

type stack struct {
	name  byte
	index byte
	data  any
}

func stack_push(reg byte, index byte, p *Program) {
	p.stack = append(p.stack, stack{reg, index, register_get[reg](index, *p)})
}

func stack_pop(reg byte, index byte, p *Program) {
	for i, v := range p.stack {
		if v.name == reg && v.index == index {
			p.stack = append(p.stack[:i], p.stack[i+1:]...)
			break
		}
	}
}
