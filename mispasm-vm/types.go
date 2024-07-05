package main

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
)

var type_sizes = []byte{1, 2, 4, 8, 1, 2, 4, 8, 4, 8, 0, 2}
