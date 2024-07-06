package main

var add_math_operation_reg = [10]func(any, any, *Program){
	func(val1 any, val2 any, p *Program) {
		register_set[rbi](0, val1.(int8)+val2.(int8), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rsi](0, val1.(int16)+val2.(int16), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rli](0, val1.(int32)+val2.(int32), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rlli](0, val1.(int64)+val2.(int64), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rbui](0, val1.(uint8)+val2.(uint8), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rsui](0, val1.(uint16)+val2.(uint16), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rlui](0, val1.(uint32)+val2.(uint32), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rllui](0, val1.(uint64)+val2.(uint64), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rlf](0, val1.(float32)+val2.(float32), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rllf](0, val1.(float64)+val2.(float64), p)
	},
}
var add_math_operation = [12]func([]byte, []byte, *Program){
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rbi](0, byte_to_int8(arg1[1])+byte_to_int8(arg2[1]), p)
		} else {
			register_set[rbi](0, byte_to_int8(arg1[1])+register_get[arg2[1]](int(arg2[2]), *p).(int8), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rsi](0, bytes_to_int16(arg1[1:])+bytes_to_int16(arg2[1:]), p)
		} else {
			register_set[rsi](0, bytes_to_int16(arg1[1:])+register_get[arg2[1]](int(arg2[2]), *p).(int16), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rli](0, bytes_to_int32(arg1[1:])+bytes_to_int32(arg2[1:]), p)
		} else {
			register_set[rli](0, bytes_to_int32(arg1[1:])+register_get[arg2[1]](int(arg2[2]), *p).(int32), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rlli](0, bytes_to_int64(arg1[1:])+bytes_to_int64(arg2[1:]), p)
		} else {
			register_set[rlli](0, bytes_to_int64(arg1[1:])+register_get[arg2[1]](int(arg2[2]), *p).(int64), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rbui](0, byte_to_uint8(arg1[1])+byte_to_uint8(arg2[1]), p)
		} else {
			register_set[rbui](0, byte_to_uint8(arg1[1])+register_get[arg2[1]](int(arg2[2]), *p).(uint8), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])+bytes_to_uint16(arg2[1:]), p)
		} else {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])+register_get[arg2[1]](int(arg2[2]), *p).(uint16), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])+bytes_to_uint32(arg2[1:]), p)
		} else {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])+register_get[arg2[1]](int(arg2[2]), *p).(uint32), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])+bytes_to_uint64(arg2[1:]), p)
		} else {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])+register_get[arg2[1]](int(arg2[2]), *p).(uint64), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rlf](0, bytes_to_float32(arg1[1:])+bytes_to_float32(arg2[1:]), p)
		} else {
			register_set[rlf](0, bytes_to_float32(arg1[1:])+register_get[arg2[1]](int(arg2[2]), *p).(float32), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rllf](0, bytes_to_float64(arg1[1:])+bytes_to_float64(arg2[1:]), p)
		} else {
			register_set[rllf](0, bytes_to_float64(arg1[1:])+register_get[arg2[1]](int(arg2[2]), *p).(float64), p)
		}
	},
	func(_ []byte, _ []byte, _ *Program) {},
	func(arg1 []byte, arg2 []byte, p *Program) {
		add_math_operation_reg[arg1[1]%11](convert_to_value[arg1[0]](arg1, *p), convert_to_value[arg2[0]](arg2, *p), p)
	},
}

var sub_math_operation_reg = [10]func(any, any, *Program){
	func(val1 any, val2 any, p *Program) {
		register_set[rbi](0, val1.(int8)-val2.(int8), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rsi](0, val1.(int16)-val2.(int16), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rli](0, val1.(int32)-val2.(int32), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rlli](0, val1.(int64)-val2.(int64), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rbui](0, val1.(uint8)-val2.(uint8), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rsui](0, val1.(uint16)-val2.(uint16), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rlui](0, val1.(uint32)-val2.(uint32), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rllui](0, val1.(uint64)-val2.(uint64), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rlf](0, val1.(float32)-val2.(float32), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rllf](0, val1.(float64)-val2.(float64), p)
	},
}

var sub_math_operation = [12]func([]byte, []byte, *Program){
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rbi](0, byte_to_int8(arg1[1])-byte_to_int8(arg2[1]), p)
		} else {
			register_set[rbi](0, byte_to_int8(arg1[1])-register_get[arg2[1]](int(arg2[2]), *p).(int8), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rsi](0, bytes_to_int16(arg1[1:])-bytes_to_int16(arg2[1:]), p)
		} else {
			register_set[rsi](0, bytes_to_int16(arg1[1:])-register_get[arg2[1]](int(arg2[2]), *p).(int16), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rli](0, bytes_to_int32(arg1[1:])-bytes_to_int32(arg2[1:]), p)
		} else {
			register_set[rli](0, bytes_to_int32(arg1[1:])-register_get[arg2[1]](int(arg2[2]), *p).(int32), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rlli](0, bytes_to_int64(arg1[1:])-bytes_to_int64(arg2[1:]), p)
		} else {
			register_set[rlli](0, bytes_to_int64(arg1[1:])-register_get[arg2[1]](int(arg2[2]), *p).(int64), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rbui](0, byte_to_uint8(arg1[1])-byte_to_uint8(arg2[1]), p)
		} else {
			register_set[rbui](0, byte_to_uint8(arg1[1])-register_get[arg2[1]](int(arg2[2]), *p).(uint8), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])-bytes_to_uint16(arg2[1:]), p)
		} else {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])-register_get[arg2[1]](int(arg2[2]), *p).(uint16), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])-bytes_to_uint32(arg2[1:]), p)
		} else {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])-register_get[arg2[1]](int(arg2[2]), *p).(uint32), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])-bytes_to_uint64(arg2[1:]), p)
		} else {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])-register_get[arg2[1]](int(arg2[2]), *p).(uint64), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rlf](0, bytes_to_float32(arg1[1:])-bytes_to_float32(arg2[1:]), p)
		} else {
			register_set[rlf](0, bytes_to_float32(arg1[1:])-register_get[arg2[1]](int(arg2[2]), *p).(float32), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rllf](0, bytes_to_float64(arg1[1:])-bytes_to_float64(arg2[1:]), p)
		} else {
			register_set[rllf](0, bytes_to_float64(arg1[1:])-register_get[arg2[1]](int(arg2[2]), *p).(float64), p)
		}
	},
	func(_ []byte, _ []byte, _ *Program) {},
	func(arg1 []byte, arg2 []byte, p *Program) {
		sub_math_operation_reg[arg1[1]%11](convert_to_value[arg1[0]](arg1, *p), convert_to_value[arg2[0]](arg2, *p), p)
	},
}

var mul_math_operation_reg = [10]func(any, any, *Program){
	func(val1 any, val2 any, p *Program) {
		register_set[rbi](0, val1.(int8)*val2.(int8), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rsi](0, val1.(int16)*val2.(int16), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rli](0, val1.(int32)*val2.(int32), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rlli](0, val1.(int64)*val2.(int64), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rbui](0, val1.(uint8)*val2.(uint8), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rsui](0, val1.(uint16)*val2.(uint16), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rlui](0, val1.(uint32)*val2.(uint32), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rllui](0, val1.(uint64)*val2.(uint64), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rlf](0, val1.(float32)*val2.(float32), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rllf](0, val1.(float64)*val2.(float64), p)
	},
}

var mul_math_operation = [12]func([]byte, []byte, *Program){
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rbi](0, byte_to_int8(arg1[1])*byte_to_int8(arg2[1]), p)
		} else {
			register_set[rbi](0, byte_to_int8(arg1[1])*register_get[arg2[1]](int(arg2[2]), *p).(int8), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rsi](0, bytes_to_int16(arg1[1:])*bytes_to_int16(arg2[1:]), p)
		} else {
			register_set[rsi](0, bytes_to_int16(arg1[1:])*register_get[arg2[1]](int(arg2[2]), *p).(int16), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rli](0, bytes_to_int32(arg1[1:])*bytes_to_int32(arg2[1:]), p)
		} else {
			register_set[rli](0, bytes_to_int32(arg1[1:])*register_get[arg2[1]](int(arg2[2]), *p).(int32), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rlli](0, bytes_to_int64(arg1[1:])*bytes_to_int64(arg2[1:]), p)
		} else {
			register_set[rlli](0, bytes_to_int64(arg1[1:])*register_get[arg2[1]](int(arg2[2]), *p).(int64), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rbui](0, byte_to_uint8(arg1[1])*byte_to_uint8(arg2[1]), p)
		} else {
			register_set[rbui](0, byte_to_uint8(arg1[1])*register_get[arg2[1]](int(arg2[2]), *p).(uint8), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])*bytes_to_uint16(arg2[1:]), p)
		} else {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])*register_get[arg2[1]](int(arg2[2]), *p).(uint16), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])*bytes_to_uint32(arg2[1:]), p)
		} else {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])*register_get[arg2[1]](int(arg2[2]), *p).(uint32), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])*bytes_to_uint64(arg2[1:]), p)
		} else {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])*register_get[arg2[1]](int(arg2[2]), *p).(uint64), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rlf](0, bytes_to_float32(arg1[1:])*bytes_to_float32(arg2[1:]), p)
		} else {
			register_set[rlf](0, bytes_to_float32(arg1[1:])*register_get[arg2[1]](int(arg2[2]), *p).(float32), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rllf](0, bytes_to_float64(arg1[1:])*bytes_to_float64(arg2[1:]), p)
		} else {
			register_set[rllf](0, bytes_to_float64(arg1[1:])*register_get[arg2[1]](int(arg2[2]), *p).(float64), p)
		}
	},
	func(_ []byte, _ []byte, _ *Program) {},
	func(arg1 []byte, arg2 []byte, p *Program) {
		mul_math_operation_reg[arg1[1]%11](convert_to_value[arg1[0]](arg1, *p), convert_to_value[arg2[0]](arg2, *p), p)
	},
}

var div_math_operation_reg = [10]func(any, any, *Program){
	func(val1 any, val2 any, p *Program) {
		register_set[rbi](0, val1.(int8)/val2.(int8), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rsi](0, val1.(int16)/val2.(int16), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rli](0, val1.(int32)/val2.(int32), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rlli](0, val1.(int64)/val2.(int64), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rbui](0, val1.(uint8)/val2.(uint8), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rsui](0, val1.(uint16)/val2.(uint16), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rlui](0, val1.(uint32)/val2.(uint32), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rllui](0, val1.(uint64)/val2.(uint64), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rlf](0, val1.(float32)/val2.(float32), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rllf](0, val1.(float64)/val2.(float64), p)
	},
}

var div_math_operation = [12]func([]byte, []byte, *Program){
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rbi](0, byte_to_int8(arg1[1])/byte_to_int8(arg2[1]), p)
		} else {
			register_set[rbi](0, byte_to_int8(arg1[1])/register_get[arg2[1]](int(arg2[2]), *p).(int8), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rsi](0, bytes_to_int16(arg1[1:])/bytes_to_int16(arg2[1:]), p)
		} else {
			register_set[rsi](0, bytes_to_int16(arg1[1:])/register_get[arg2[1]](int(arg2[2]), *p).(int16), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rli](0, bytes_to_int32(arg1[1:])/bytes_to_int32(arg2[1:]), p)
		} else {
			register_set[rli](0, bytes_to_int32(arg1[1:])/register_get[arg2[1]](int(arg2[2]), *p).(int32), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rlli](0, bytes_to_int64(arg1[1:])/bytes_to_int64(arg2[1:]), p)
		} else {
			register_set[rlli](0, bytes_to_int64(arg1[1:])/register_get[arg2[1]](int(arg2[2]), *p).(int64), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rbui](0, byte_to_uint8(arg1[1])/byte_to_uint8(arg2[1]), p)
		} else {
			register_set[rbui](0, byte_to_uint8(arg1[1])/register_get[arg2[1]](int(arg2[2]), *p).(uint8), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])/bytes_to_uint16(arg2[1:]), p)
		} else {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])/register_get[arg2[1]](int(arg2[2]), *p).(uint16), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])/bytes_to_uint32(arg2[1:]), p)
		} else {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])/register_get[arg2[1]](int(arg2[2]), *p).(uint32), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])/bytes_to_uint64(arg2[1:]), p)
		} else {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])/register_get[arg2[1]](int(arg2[2]), *p).(uint64), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rlf](0, bytes_to_float32(arg1[1:])/bytes_to_float32(arg2[1:]), p)
		} else {
			register_set[rlf](0, bytes_to_float32(arg1[1:])/register_get[arg2[1]](int(arg2[2]), *p).(float32), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rllf](0, bytes_to_float64(arg1[1:])/bytes_to_float64(arg2[1:]), p)
		} else {
			register_set[rllf](0, bytes_to_float64(arg1[1:])/register_get[arg2[1]](int(arg2[2]), *p).(float64), p)
		}
	},
	func(_ []byte, _ []byte, _ *Program) {},
	func(arg1 []byte, arg2 []byte, p *Program) {
		div_math_operation_reg[arg1[1]%11](convert_to_value[arg1[0]](arg1, *p), convert_to_value[arg2[0]](arg2, *p), p)
	},
}

var mod_math_operation_reg = [10]func(any, any, *Program){
	func(val1 any, val2 any, p *Program) {
		register_set[rbi](0, val1.(int8)%val2.(int8), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rsi](0, val1.(int16)%val2.(int16), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rli](0, val1.(int32)%val2.(int32), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rlli](0, val1.(int64)%val2.(int64), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rbui](0, val1.(uint8)%val2.(uint8), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rsui](0, val1.(uint16)%val2.(uint16), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rlui](0, val1.(uint32)%val2.(uint32), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rllui](0, val1.(uint64)%val2.(uint64), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rlf](0, float32(int(val1.(float32))%int(val2.(float32))), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rllf](0, float64(int(val1.(float64))%int(val2.(float64))), p)
	},
}

var mod_math_operation = [12]func([]byte, []byte, *Program){
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rbi](0, byte_to_int8(arg1[1])%byte_to_int8(arg2[1]), p)
		} else {
			register_set[rbi](0, byte_to_int8(arg1[1])%register_get[arg2[1]](int(arg2[2]), *p).(int8), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rsi](0, bytes_to_int16(arg1[1:])%bytes_to_int16(arg2[1:]), p)
		} else {
			register_set[rsi](0, bytes_to_int16(arg1[1:])%register_get[arg2[1]](int(arg2[2]), *p).(int16), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rli](0, bytes_to_int32(arg1[1:])%bytes_to_int32(arg2[1:]), p)
		} else {
			register_set[rli](0, bytes_to_int32(arg1[1:])%register_get[arg2[1]](int(arg2[2]), *p).(int32), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rlli](0, bytes_to_int64(arg1[1:])%bytes_to_int64(arg2[1:]), p)
		} else {
			register_set[rlli](0, bytes_to_int64(arg1[1:])%register_get[arg2[1]](int(arg2[2]), *p).(int64), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rbui](0, byte_to_uint8(arg1[1])%byte_to_uint8(arg2[1]), p)
		} else {
			register_set[rbui](0, byte_to_uint8(arg1[1])%register_get[arg2[1]](int(arg2[2]), *p).(uint8), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])%bytes_to_uint16(arg2[1:]), p)
		} else {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])%register_get[arg2[1]](int(arg2[2]), *p).(uint16), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])%bytes_to_uint32(arg2[1:]), p)
		} else {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])%register_get[arg2[1]](int(arg2[2]), *p).(uint32), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])%bytes_to_uint64(arg2[1:]), p)
		} else {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])%register_get[arg2[1]](int(arg2[2]), *p).(uint64), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rlf](0, float32(int(bytes_to_float32(arg1[1:]))%int(bytes_to_float32(arg2[1:]))), p)
		} else {
			register_set[rlf](0, float32(int(bytes_to_float32(arg1[1:]))%int(register_get[arg2[1]](int(arg2[2]), *p).(float32))), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rllf](0, float64(int(bytes_to_float64(arg1[1:]))%int(bytes_to_float64(arg2[1:]))), p)
		} else {
			register_set[rllf](0, float64(int(bytes_to_float64(arg1[1:]))%int(register_get[arg2[1]](int(arg2[2]), *p).(float64))), p)
		}
	},
	func(_ []byte, _ []byte, _ *Program) {},
	func(arg1 []byte, arg2 []byte, p *Program) {
		mod_math_operation_reg[arg1[1]%11](convert_to_value[arg1[0]](arg1, *p), convert_to_value[arg2[0]](arg2, *p), p)
	},
}

var inc_math_operation = [21]func(*Program){
	func(p *Program) {
		p.register_bi[0]++
	},
	func(p *Program) {
		p.register_si[0]++
	},
	func(p *Program) {
		p.register_li[0]++
	},
	func(p *Program) {
		p.register_lli[0]++
	},
	func(p *Program) {
		p.register_bui[0]++
	},
	func(p *Program) {
		p.register_sui[0]++
	},
	func(p *Program) {
		p.register_lui[0]++
	},
	func(p *Program) {
		p.register_llui[0]++
	},
	func(p *Program) {
		p.register_lf[0]++
	},
	func(p *Program) {
		p.register_llf[0]++
	},
	func(_ *Program) {},
	func(p *Program) {
		p.register_rbi[0]++
	},
	func(p *Program) {
		p.register_rsi[0]++
	},
	func(p *Program) {
		p.register_rli[0]++
	},
	func(p *Program) {
		p.register_rlli[0]++
	},
	func(p *Program) {
		p.register_rbui[0]++
	},
	func(p *Program) {
		p.register_rsui[0]++
	},
	func(p *Program) {
		p.register_rlui[0]++
	},
	func(p *Program) {
		p.register_rllui[0]++
	},
	func(p *Program) {
		p.register_rlf[0]++
	},
	func(p *Program) {
		p.register_rllf[0]++
	},
}

var dec_math_operation = [21]func(*Program){
	func(p *Program) {
		p.register_bi[0]--
	},
	func(p *Program) {
		p.register_si[0]--
	},
	func(p *Program) {
		p.register_li[0]--
	},
	func(p *Program) {
		p.register_lli[0]--
	},
	func(p *Program) {
		p.register_bui[0]--
	},
	func(p *Program) {
		p.register_sui[0]--
	},
	func(p *Program) {
		p.register_lui[0]--
	},
	func(p *Program) {
		p.register_llui[0]--
	},
	func(p *Program) {
		p.register_lf[0]--
	},
	func(p *Program) {
		p.register_llf[0]--
	},
	func(_ *Program) {},
	func(p *Program) {
		p.register_rbi[0]--
	},
	func(p *Program) {
		p.register_rsi[0]--
	},
	func(p *Program) {
		p.register_rli[0]--
	},
	func(p *Program) {
		p.register_rlli[0]--
	},
	func(p *Program) {
		p.register_rbui[0]--
	},
	func(p *Program) {
		p.register_rsui[0]--
	},
	func(p *Program) {
		p.register_rlui[0]--
	},
	func(p *Program) {
		p.register_rllui[0]--
	},
	func(p *Program) {
		p.register_rlf[0]--
	},
	func(p *Program) {
		p.register_rllf[0]--
	},
}
