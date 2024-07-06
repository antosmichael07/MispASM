package main

import (
	"fmt"
)

const (
	print byte = iota
)

func (p *Program) init_calls() {
	p.calls[print] = func(arg1 []byte, arg2 []byte) {
		fmt.Print(p.sprint_stack())
	}
}

func (p *Program) sprint_stack() string {
	msg := ""
	for _, v := range p.stack {
		msg = fmt.Sprint(msg, v.data)
	}
	return msg
}
