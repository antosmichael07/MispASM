package main

import (
	"fmt"
)

const (
	printf byte = iota
)

var calls = [1]func([]byte, []byte){}

func init_calls(p *Program) {
	calls[printf] = func(arg1 []byte, arg2 []byte) {
		fmt.Print(sprintf_stack(p))
	}
}

func sprintf_stack(p *Program) string {
	msg := ""
	continu := false
	for _, v := range p.stack {
		if !continu {
			msg = v.data.(string)
			continu = true
			continue
		}
		msg = fmt.Sprintf(msg, v.data)
	}
	return msg
}
