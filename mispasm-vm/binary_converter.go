package main

import (
	"encoding/binary"
	"math"
)

var convert_to_value = [14]func([]byte, Program) any{
	func(data []byte, _ Program) any { return byte_to_int8(data[1]) },
	func(data []byte, _ Program) any { return bytes_to_int16(data[1:]) },
	func(data []byte, _ Program) any { return bytes_to_int32(data[1:]) },
	func(data []byte, _ Program) any { return bytes_to_int64(data[1:]) },
	func(data []byte, _ Program) any { return byte_to_uint8(data[1]) },
	func(data []byte, _ Program) any { return bytes_to_uint16(data[1:]) },
	func(data []byte, _ Program) any { return bytes_to_uint32(data[1:]) },
	func(data []byte, _ Program) any { return bytes_to_uint64(data[1:]) },
	func(data []byte, _ Program) any { return bytes_to_float32(data[1:]) },
	func(data []byte, _ Program) any { return bytes_to_float64(data[1:]) },
	func(data []byte, _ Program) any { return bytes_to_string(data[1:]) },
	func(data []byte, p Program) any { return register_get[data[1]](data[2], p) },
	func(data []byte, p Program) any { return p.get_const_value(data) },
	func(data []byte, p Program) any { return p.get_var_value(data) },
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
	return string(b[:len(b)-1])
}

var convert_to_bytes = [11]func(byte, any) []byte{
	func(t byte, data any) []byte { return []byte{t, int8_to_byte(data.(int8))} },
	func(t byte, data any) []byte { return append([]byte{t}, int16_to_bytes(data.(int16))...) },
	func(t byte, data any) []byte { return append([]byte{t}, int32_to_bytes(data.(int32))...) },
	func(t byte, data any) []byte { return append([]byte{t}, int64_to_bytes(data.(int64))...) },
	func(t byte, data any) []byte { return []byte{t, uint8_to_byte(data.(uint8))} },
	func(t byte, data any) []byte { return append([]byte{t}, uint16_to_bytes(data.(uint16))...) },
	func(t byte, data any) []byte { return append([]byte{t}, uint32_to_bytes(data.(uint32))...) },
	func(t byte, data any) []byte { return append([]byte{t}, uint64_to_bytes(data.(uint64))...) },
	func(t byte, data any) []byte { return append([]byte{t}, float32_to_bytes(data.(float32))...) },
	func(t byte, data any) []byte { return append([]byte{t}, float64_to_bytes(data.(float64))...) },
	func(t byte, data any) []byte { return append([]byte{t}, string_to_bytes(data.(string))...) },
}

func int8_to_byte(i int8) byte {
	return byte(i)
}

func int16_to_bytes(i int16) []byte {
	return []byte{byte(i >> 8), byte(i)}
}

func int32_to_bytes(i int32) []byte {
	return []byte{byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
}

func int64_to_bytes(i int64) []byte {
	return []byte{byte(i >> 56), byte(i >> 48), byte(i >> 40), byte(i >> 32), byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
}

func uint8_to_byte(i uint8) byte {
	return byte(i)
}

func uint16_to_bytes(i uint16) []byte {
	return []byte{byte(i >> 8), byte(i)}
}

func uint32_to_bytes(i uint32) []byte {
	return []byte{byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
}

func uint64_to_bytes(i uint64) []byte {
	return []byte{byte(i >> 56), byte(i >> 48), byte(i >> 40), byte(i >> 32), byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}
}

func float32_to_bytes(f float32) []byte {
	return int32_to_bytes(int32(f))
}

func float64_to_bytes(f float64) []byte {
	return int64_to_bytes(int64(f))
}

func string_to_bytes(s string) []byte {
	return append([]byte(s), 0)
}
