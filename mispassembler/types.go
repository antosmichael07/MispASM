package main

import (
	"strconv"
	"strings"
)

/*
const (

	t_int8 byte = iota
	t_int16
	t_int32
	t_int64
	t_uint8
	t_uint16
	t_uint32
	t_uint64
	t_float32
	t_float64
	t_string
	t_reg
	t_const
	t_var

)
*/

var types = map[string]byte{
	"i8":  0,
	"i16": 1,
	"i32": 2,
	"i64": 3,
	"u8":  4,
	"u16": 5,
	"u32": 6,
	"u64": 7,
	"f32": 8,
	"f64": 9,
	"s":   10,
	"r":   11,
	"c":   12,
	"v":   13,
}

var type_sizes = [14]byte{1, 2, 4, 8, 1, 2, 4, 8, 4, 8, 0, 0, 0, 0}

func get_type(str string) byte {
	return types[str[:strings.Index(str, "-")]]
}

var value_to_byte = [13]func(string, map[string][]byte) []byte{
	func(str string, _ map[string][]byte) []byte {
		i, _ := strconv.ParseInt(str, 10, 8)
		return []byte{byte(i)}
	},
	func(str string, _ map[string][]byte) []byte {
		i, _ := strconv.ParseInt(str, 10, 16)
		return []byte{byte(i >> 8), byte(i)}
	},
	func(str string, _ map[string][]byte) []byte {
		i, _ := strconv.ParseInt(str, 10, 32)
		return []byte{byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
	},
	func(str string, _ map[string][]byte) []byte {
		i, _ := strconv.ParseInt(str, 10, 64)
		return []byte{byte(i >> 56), byte(i >> 48), byte(i >> 40), byte(i >> 32), byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
	},
	func(str string, _ map[string][]byte) []byte {
		i, _ := strconv.ParseUint(str, 10, 8)
		return []byte{byte(i)}
	},
	func(str string, _ map[string][]byte) []byte {
		i, _ := strconv.ParseUint(str, 10, 16)
		return []byte{byte(i >> 8), byte(i)}
	},
	func(str string, _ map[string][]byte) []byte {
		i, _ := strconv.ParseUint(str, 10, 32)
		return []byte{byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
	},
	func(str string, _ map[string][]byte) []byte {
		i, _ := strconv.ParseUint(str, 10, 64)
		return []byte{byte(i >> 56), byte(i >> 48), byte(i >> 40), byte(i >> 32), byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
	},
	func(str string, _ map[string][]byte) []byte {
		f, _ := strconv.ParseFloat(str, 32)
		return []byte{byte(f)}
	},
	func(str string, _ map[string][]byte) []byte {
		f, _ := strconv.ParseFloat(str, 64)
		return []byte{byte(f)}
	},
	func(str string, _ map[string][]byte) []byte {
		str = "\"" + str + "\""
		unquoted, _ := strconv.Unquote(str)
		return append([]byte(unquoted), 0)
	},
	func(str string, _ map[string][]byte) []byte {
		li := strings.LastIndex(str, "-")
		i, _ := strconv.ParseInt(str[li+1:], 10, 8)
		return []byte{registers[str[:li]], byte(i)}
	},
	func(str string, var_types map[string][]byte) []byte {
		return append(var_types[str], append([]byte(str), 0)...)
	},
}
