package main

/*
#cgo CFLAGS: -I.
#include <windows.h>
#include <stdio.h>

HINSTANCE load_library(const char* libname) {
    return LoadLibrary(libname);
}

FARPROC get_proc(HINSTANCE h, const char* procname) {
    return GetProcAddress(h, procname);
}

typedef struct {
	void* data;
	int val_len;
	int data_len;
} response;

typedef response (*callFunc)(const char*, void*, int, int);

response bridge_call(callFunc f, const char* call, void* data, int val_len, int data_len) {
	return f(call, data, val_len, data_len);
}
*/
import "C"
import (
	"strings"
	"unsafe"
)

func (p *Program) load_lib(lib string) {
	if p.libs == nil {
		p.libs = make(map[string]C.callFunc)
	}

	library := C.load_library(C.CString("../lib/" + lib + ".misplib"))
	proc := C.get_proc(library, C.CString("Call"+lib))
	fun := C.callFunc(proc)

	p.libs[lib] = fun
}

func (p Program) is_lib_loaded(lib string) bool {
	for k := range p.libs {
		if k == lib {
			return true
		}
	}
	return false
}

func (p *Program) call(lib string, cal string) {
	data, val_len, data_len := encode_stack(p.stack)
	c_res := C.bridge_call(p.libs[lib], C.CString(cal), data, val_len, data_len)
	res := decode_stack(c_res.data, c_res.val_len, c_res.data_len)

	for i := 0; i < len(res); i++ {
		register_force_set[res[i].name%11][res[i].name](res[i].index, res[i].data, p)
	}
}

func encode_stack(s []stack) (unsafe.Pointer, C.int, C.int) {
	data := []byte{}
	value_length := C.int(0)

	for i := 0; i < len(s); i++ {
		data = append(data, s[i].name)
		value_length++
	}
	for i := 0; i < len(s); i++ {
		data = append(data, s[i].index)
	}
	for i := 0; i < len(s); i++ {
		data = append(data, convert_to_bytes[s[i].name%11](s[i].name%11, s[i].data)...)
	}

	return C.CBytes(data), C.int(value_length), C.int(len(data))
}

func decode_stack(c_data unsafe.Pointer, value_length C.int, length C.int) []stack {
	data := C.GoBytes(c_data, length)

	name_arr := data[:value_length]
	index_arr := data[value_length : value_length*2]
	data_arr := []any{}

	offset := int(value_length * 2)
	for i := 0; i < int(value_length); i++ {
		size := get_size(data[offset:])
		data_arr = append(data_arr, convert_to_value[data[offset]](data[offset:offset+size+1], Program{}))
		offset += size + 1
	}

	stack_arr := []stack{}
	for i := 0; i < int(value_length); i++ {
		stack_arr = append(stack_arr, stack{name_arr[i], index_arr[i], data_arr[i]})
	}

	return stack_arr
}

func get_size(data []byte) int {
	if data[0] <= 9 {
		return int(type_sizes[data[0]])
	} else {
		size := 0
		for i := 0; i < len(data); i++ {
			if data[i] == 0 {
				break
			}
			size++
		}
		return size
	}
}

func (p *Program) init_calls() {
	for _, f := range p.funcs {
		for j := 0; j < len(f.instructions); j++ {
			arg1, _, _, _, arg_size := get_args(f.instructions, j)
			if f.instructions[j] == call {
				lib := convert_to_value[t_string](arg1[:len(arg1)-1], *p).(string)[:strings.Index(convert_to_value[t_string](arg1[:len(arg1)-1], *p).(string), ".")]
				if !p.is_lib_loaded(lib) {
					p.load_lib(lib)
				}
			}
			j += arg_size
		}
	}
}
