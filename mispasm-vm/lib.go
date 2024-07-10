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
import "fmt"

func (p *Program) load_lib(lib string) {
	library := C.load_library(C.CString(lib + ".mispcal"))
	if library == nil {
		fmt.Println("Error loading library")
		return
	}
	f := C.callFunc(C.get_proc(library, C.CString("Call"+lib)))
	p.libs[lib] = f
}

func (p Program) is_lib_loaded(lib string) bool {
	for k := range p.libs {
		if k == lib {
			return true
		}
	}
	return false
}
