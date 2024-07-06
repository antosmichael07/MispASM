package main

import (
	"fmt"
)

const (
	print byte = iota
)

func (p *Program) init_calls() {
	p.calls[print] = func(arg1 []byte, arg2 []byte) {
		fmt.Print(sprint_stack(p))
	}
}

func sprint_stack(p *Program) string {
	msg := ""
	for _, v := range p.stack {
		msg = fmt.Sprint(msg, v.data)
	}
	return msg
}
