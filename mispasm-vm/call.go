package main

import (
	"fmt"
)

const (
	print byte = iota
	scanf
)

func (p *Program) init_calls() {
	p.calls[0][print] = func(arg1 []byte, arg2 []byte) {
		fmt.Print(p.sprint_stack())
	}
	p.calls[0][scanf] = func(arg1 []byte, arg2 []byte) {
		scanf_func[p.stack[1].name](arg1, arg2, p)
	}
}

func (p *Program) sprint_stack() string {
	msg := ""
	for _, v := range p.stack {
		msg = fmt.Sprint(msg, v.data)
	}
	return msg
}

var scanf_func = [11]func([]byte, []byte, *Program){
	func(arg1 []byte, arg2 []byte, p *Program) {
		var data int8
		fmt.Scanf(p.stack[0].data.(string), &data)
		register_set[p.stack[1].name](0, data, p)
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		var data int16
		fmt.Scanf(p.stack[0].data.(string), &data)
		register_set[p.stack[1].name](0, data, p)
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		var data int32
		fmt.Scanf(p.stack[0].data.(string), &data)
		register_set[p.stack[1].name](0, data, p)
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		var data int64
		fmt.Scanf(p.stack[0].data.(string), &data)
		register_set[p.stack[1].name](0, data, p)
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		var data uint8
		fmt.Scanf(p.stack[0].data.(string), &data)
		register_set[p.stack[1].name](0, data, p)
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		var data uint16
		fmt.Scanf(p.stack[0].data.(string), &data)
		register_set[p.stack[1].name](0, data, p)
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		var data uint32
		fmt.Scanf(p.stack[0].data.(string), &data)
		register_set[p.stack[1].name](0, data, p)
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		var data uint64
		fmt.Scanf(p.stack[0].data.(string), &data)
		register_set[p.stack[1].name](0, data, p)
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		var data float32
		fmt.Scanf(p.stack[0].data.(string), &data)
		register_set[p.stack[1].name](0, data, p)
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		var data float64
		fmt.Scanf(p.stack[0].data.(string), &data)
		register_set[p.stack[1].name](0, data, p)
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		var data string
		fmt.Scanf(p.stack[0].data.(string), &data)
		register_set[p.stack[1].name](0, data, p)
	},
}
