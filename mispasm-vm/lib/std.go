package main

/*
#cgo CFLAGS: -I.

typedef struct {
	void* data;
	int len;
} response;
*/
import "C"
import (
	"encoding/binary"
	"fmt"
	"math"
	"unsafe"
)

type stack struct {
	name  byte
	index byte
	data  any
}

var calls = map[string]func([]stack){
	"print": func(s []stack) {
		msg := ""
		for _, v := range s {
			msg = fmt.Sprint(msg, v.data)
		}
		fmt.Print(msg)
	},
}

//export Callstd
func Callstd(call *C.char, data unsafe.Pointer, val_len C.int, data_len C.int) C.response {
	stack_arr := decode_stack(data, val_len, data_len)

	calls[C.GoString(call)](stack_arr)

	return C.response{
		data: unsafe.Pointer(C.CBytes([]byte{1, 2, 3})),
		len:  C.int(3),
	}
}

func decode_stack(c_data unsafe.Pointer, value_length C.int, length C.int) []stack {
	data := C.GoBytes(c_data, length)

	name_arr := data[:value_length]
	index_arr := data[value_length : value_length*2]
	data_arr := []any{}

	offset := int(value_length * 2)
	for i := 0; i < int(value_length); i++ {
		size := get_size(data[offset:])
		data_arr = append(data_arr, convert_to_value[data[offset]](data[offset:offset+size]))
		offset += size
	}

	stack_arr := []stack{}
	for i := 0; i < int(value_length); i++ {
		stack_arr = append(stack_arr, stack{name_arr[i], index_arr[i], data_arr[i]})
	}

	return stack_arr
}

func get_size(data []byte) int {
	if data[0] <= 9 {
		return type_sizes[data[0]]
	} else {
		size := 0
		for i := 0; i < len(data); i++ {
			size++
			if data[i] == 0 {
				break
			}
		}
		return size
	}
}

var type_sizes = [11]int{1, 2, 4, 8, 1, 2, 4, 8, 4, 8, 0}

var convert_to_value = [11]func([]byte) any{
	func(data []byte) any { return byte_to_int8(data[1]) },
	func(data []byte) any { return bytes_to_int16(data[1:]) },
	func(data []byte) any { return bytes_to_int32(data[1:]) },
	func(data []byte) any { return bytes_to_int64(data[1:]) },
	func(data []byte) any { return byte_to_uint8(data[1]) },
	func(data []byte) any { return bytes_to_uint16(data[1:]) },
	func(data []byte) any { return bytes_to_uint32(data[1:]) },
	func(data []byte) any { return bytes_to_uint64(data[1:]) },
	func(data []byte) any { return bytes_to_float32(data[1:]) },
	func(data []byte) any { return bytes_to_float64(data[1:]) },
	func(data []byte) any { return bytes_to_string(data[1:]) },
}

func byte_to_int8(b byte) int8 {
	return int8(b)
}

func bytes_to_int16(b []byte) int16 {
	return int16(b[0])<<8 | int16(b[1])
}

func bytes_to_int32(b []byte) int32 {
	return int32(b[0])<<24 | int32(b[1])<<16 | int32(b[2])<<8 | int32(b[3])
}

func bytes_to_int64(b []byte) int64 {
	return int64(b[0])<<56 | int64(b[1])<<48 | int64(b[2])<<40 | int64(b[3])<<32 | int64(b[4])<<24 | int64(b[5])<<16 | int64(b[6])<<8 | int64(b[7])
}

func byte_to_uint8(b byte) uint8 {
	return uint8(b)
}

func bytes_to_uint16(b []byte) uint16 {
	return uint16(b[0])<<8 | uint16(b[1])
}

func bytes_to_uint32(b []byte) uint32 {
	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}

func bytes_to_uint64(b []byte) uint64 {
	return uint64(b[0])<<56 | uint64(b[1])<<48 | uint64(b[2])<<40 | uint64(b[3])<<32 | uint64(b[4])<<24 | uint64(b[5])<<16 | uint64(b[6])<<8 | uint64(b[7])
}

func bytes_to_float32(bytes []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(bytes))
}

func bytes_to_float64(bytes []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(bytes))
}

func bytes_to_string(b []byte) string {
	return string(b)
}

func main() {}
