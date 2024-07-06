package main

import (
	"fmt"
)

const (
	print byte = iota
)

var calls = [1]func([]byte, []byte){}

func init_calls(p *Program) {
	calls[print] = func(arg1 []byte, arg2 []byte) {
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
