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

typedef void (*callFunc)(const char*, void*, int, int);

void bridge_call(callFunc f, const char* call, void* data, int val_len, int data_len) {
	f(call, data, val_len, data_len);
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

	library := C.load_library(C.CString("./lib/" + lib + ".misplib"))
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

func (p Program) call(lib string, cal string) {
	data, val_len, data_len := encode_stack(p.stack)
	C.bridge_call(p.libs[lib], C.CString(cal), data, val_len, data_len)
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

	return C.CBytes(data), value_length, C.int(len(data))
}

func (p *Program) init_calls() {
	for _, f := range p.funcs {
		for j := 0; j < len(f.instructions); j++ {
			arg1, _, _, _, arg_size := get_args(f.instructions, j)
			if f.instructions[j] == call {
				lib := convert_to_value[t_string](arg1[:len(arg1)-1], *p).(string)[:strings.Index(convert_to_value[t_string](arg1[:len(arg1)-1], *p).(string), ".")]
				if arg1[1] != '#' && !p.is_lib_loaded(lib) {
					p.load_lib(lib)
				}
			}
			j += arg_size
		}
	}
}
