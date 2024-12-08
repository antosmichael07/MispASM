package main

import (
	"math"
	"strconv"
	"strings"
)

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
}

func get_type(str string, var_types map[string][]byte) byte {
	if str[0] == '"' {
		return 10
	} else if strings.Contains(str, "-") {
		return types[str[:strings.Index(str, "-")]]
	} else if str[0] >= '0' && str[0] <= '9' {
		if strings.Contains(str, ".") {
			return 8
		} else {
			return 2
		}
	} else {
		if ok, _, _ := is_register(str); ok {
			return 11
		} else {
			return var_types[str][0]
		}
	}
}

func value_to_byte(str string, var_types map[string][]byte) []byte {
	if str[0] == '"' {
		unquoted, _ := strconv.Unquote(str)
		return append([]byte(unquoted), 0)
	} else if strings.Contains(str, "-") {
		return value_to_byte_arr[get_type(str, var_types)](str[strings.Index(str, "-")+1:])
	} else if str[0] >= '0' && str[0] <= '9' {
		if strings.Contains(str, ".") {
			return value_to_byte_arr[8](str)
		} else {
			return value_to_byte_arr[2](str)
		}
	} else {
		if ok, reg, i := is_register(str); ok {
			return []byte{registers[reg], i}
		} else {
			return append([]byte{var_types[str][1]}, append([]byte(str), 0)...)
		}
	}
}

var value_to_byte_arr = [10]func(string) []byte{
	func(str string) []byte {
		i, _ := strconv.ParseInt(str, 10, 8)
		return []byte{byte(i)}
	},
	func(str string) []byte {
		i, _ := strconv.ParseInt(str, 10, 16)
		return []byte{byte(i >> 8), byte(i)}
	},
	func(str string) []byte {
		i, _ := strconv.ParseInt(str, 10, 32)
		return []byte{byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
	},
	func(str string) []byte {
		i, _ := strconv.ParseInt(str, 10, 64)
		return []byte{byte(i >> 56), byte(i >> 48), byte(i >> 40), byte(i >> 32), byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
	},
	func(str string) []byte {
		i, _ := strconv.ParseUint(str, 10, 8)
		return []byte{byte(i)}
	},
	func(str string) []byte {
		i, _ := strconv.ParseUint(str, 10, 16)
		return []byte{byte(i >> 8), byte(i)}
	},
	func(str string) []byte {
		i, _ := strconv.ParseUint(str, 10, 32)
		return []byte{byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
	},
	func(str string) []byte {
		i, _ := strconv.ParseUint(str, 10, 64)
		return []byte{byte(i >> 56), byte(i >> 48), byte(i >> 40), byte(i >> 32), byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
	},
	func(str string) []byte {
		f, _ := strconv.ParseFloat(str, 32)
		return []byte{byte(math.Float32bits(float32(f))), byte(math.Float32bits(float32(f)) >> 8), byte(math.Float32bits(float32(f)) >> 16), byte(math.Float32bits(float32(f)) >> 24)}
	},
	func(str string) []byte {
		f, _ := strconv.ParseFloat(str, 64)
		return []byte{byte(math.Float64bits(f)), byte(math.Float64bits(f) >> 8), byte(math.Float64bits(f) >> 16), byte(math.Float64bits(f) >> 24), byte(math.Float64bits(f) >> 32), byte(math.Float64bits(f) >> 40), byte(math.Float64bits(f) >> 48), byte(math.Float64bits(f) >> 56)}
	},
}
