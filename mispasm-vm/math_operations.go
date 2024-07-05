package main

var add_math_operation_reg = [10]func(any, any){
	func(val1 any, val2 any) {
		register_set[rbi](0, val1.(int8)+val2.(int8))
	},
	func(val1 any, val2 any) {
		register_set[rsi](0, val1.(int16)+val2.(int16))
	},
	func(val1 any, val2 any) {
		register_set[rli](0, val1.(int32)+val2.(int32))
	},
	func(val1 any, val2 any) {
		register_set[rlli](0, val1.(int64)+val2.(int64))
	},
	func(val1 any, val2 any) {
		register_set[rbui](0, val1.(uint8)+val2.(uint8))
	},
	func(val1 any, val2 any) {
		register_set[rsui](0, val1.(uint16)+val2.(uint16))
	},
	func(val1 any, val2 any) {
		register_set[rlui](0, val1.(uint32)+val2.(uint32))
	},
	func(val1 any, val2 any) {
		register_set[rllui](0, val1.(uint64)+val2.(uint64))
	},
	func(val1 any, val2 any) {
		register_set[rlf](0, val1.(float32)+val2.(float32))
	},
	func(val1 any, val2 any) {
		register_set[rllf](0, val1.(float64)+val2.(float64))
	},
}
var add_math_operation = [13]func([]byte, []byte){
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rbi](0, byte_to_int8(arg1[1])+byte_to_int8(arg2[1]))
		} else {
			register_set[rbi](0, byte_to_int8(arg1[1])+register_get[arg2[1]](int(arg2[2])).(int8))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rsi](0, bytes_to_int16(arg1[1:])+bytes_to_int16(arg2[1:]))
		} else {
			register_set[rsi](0, bytes_to_int16(arg1[1:])+register_get[arg2[1]](int(arg2[2])).(int16))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rli](0, bytes_to_int32(arg1[1:])+bytes_to_int32(arg2[1:]))
		} else {
			register_set[rli](0, bytes_to_int32(arg1[1:])+register_get[arg2[1]](int(arg2[2])).(int32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rlli](0, bytes_to_int64(arg1[1:])+bytes_to_int64(arg2[1:]))
		} else {
			register_set[rlli](0, bytes_to_int64(arg1[1:])+register_get[arg2[1]](int(arg2[2])).(int64))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rbui](0, byte_to_uint8(arg1[1])+byte_to_uint8(arg2[1]))
		} else {
			register_set[rbui](0, byte_to_uint8(arg1[1])+register_get[arg2[1]](int(arg2[2])).(uint8))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])+bytes_to_uint16(arg2[1:]))
		} else {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])+register_get[arg2[1]](int(arg2[2])).(uint16))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])+bytes_to_uint32(arg2[1:]))
		} else {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])+register_get[arg2[1]](int(arg2[2])).(uint32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])+bytes_to_uint64(arg2[1:]))
		} else {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])+register_get[arg2[1]](int(arg2[2])).(uint64))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rlf](0, bytes_to_float32(arg1[1:])+bytes_to_float32(arg2[1:]))
		} else {
			register_set[rlf](0, bytes_to_float32(arg1[1:])+register_get[arg2[1]](int(arg2[2])).(float32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rllf](0, bytes_to_float64(arg1[1:])+bytes_to_float64(arg2[1:]))
		} else {
			register_set[rllf](0, bytes_to_float64(arg1[1:])+register_get[arg2[1]](int(arg2[2])).(float64))
		}
	},
	func(arg1 []byte, arg2 []byte) {},
	func(arg1 []byte, arg2 []byte) {},
	func(arg1 []byte, arg2 []byte) {
		add_math_operation_reg[arg1[1]%11](convert_to_value[arg1[0]](arg1), convert_to_value[arg2[0]](arg2))
	},
}

var sub_math_operation_reg = [10]func(any, any){
	func(val1 any, val2 any) {
		register_set[rbi](0, val1.(int8)-val2.(int8))
	},
	func(val1 any, val2 any) {
		register_set[rsi](0, val1.(int16)-val2.(int16))
	},
	func(val1 any, val2 any) {
		register_set[rli](0, val1.(int32)-val2.(int32))
	},
	func(val1 any, val2 any) {
		register_set[rlli](0, val1.(int64)-val2.(int64))
	},
	func(val1 any, val2 any) {
		register_set[rbui](0, val1.(uint8)-val2.(uint8))
	},
	func(val1 any, val2 any) {
		register_set[rsui](0, val1.(uint16)-val2.(uint16))
	},
	func(val1 any, val2 any) {
		register_set[rlui](0, val1.(uint32)-val2.(uint32))
	},
	func(val1 any, val2 any) {
		register_set[rllui](0, val1.(uint64)-val2.(uint64))
	},
	func(val1 any, val2 any) {
		register_set[rlf](0, val1.(float32)-val2.(float32))
	},
	func(val1 any, val2 any) {
		register_set[rllf](0, val1.(float64)-val2.(float64))
	},
}

var sub_math_operation = [13]func([]byte, []byte){
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rbi](0, byte_to_int8(arg1[1])-byte_to_int8(arg2[1]))
		} else {
			register_set[rbi](0, byte_to_int8(arg1[1])-register_get[arg2[1]](int(arg2[2])).(int8))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rsi](0, bytes_to_int16(arg1[1:])-bytes_to_int16(arg2[1:]))
		} else {
			register_set[rsi](0, bytes_to_int16(arg1[1:])-register_get[arg2[1]](int(arg2[2])).(int16))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rli](0, bytes_to_int32(arg1[1:])-bytes_to_int32(arg2[1:]))
		} else {
			register_set[rli](0, bytes_to_int32(arg1[1:])-register_get[arg2[1]](int(arg2[2])).(int32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rlli](0, bytes_to_int64(arg1[1:])-bytes_to_int64(arg2[1:]))
		} else {
			register_set[rlli](0, bytes_to_int64(arg1[1:])-register_get[arg2[1]](int(arg2[2])).(int64))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rbui](0, byte_to_uint8(arg1[1])-byte_to_uint8(arg2[1]))
		} else {
			register_set[rbui](0, byte_to_uint8(arg1[1])-register_get[arg2[1]](int(arg2[2])).(uint8))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])-bytes_to_uint16(arg2[1:]))
		} else {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])-register_get[arg2[1]](int(arg2[2])).(uint16))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])-bytes_to_uint32(arg2[1:]))
		} else {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])-register_get[arg2[1]](int(arg2[2])).(uint32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])-bytes_to_uint64(arg2[1:]))
		} else {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])-register_get[arg2[1]](int(arg2[2])).(uint64))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rlf](0, bytes_to_float32(arg1[1:])-bytes_to_float32(arg2[1:]))
		} else {
			register_set[rlf](0, bytes_to_float32(arg1[1:])-register_get[arg2[1]](int(arg2[2])).(float32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rllf](0, bytes_to_float64(arg1[1:])-bytes_to_float64(arg2[1:]))
		} else {
			register_set[rllf](0, bytes_to_float64(arg1[1:])-register_get[arg2[1]](int(arg2[2])).(float64))
		}
	},
	func(arg1 []byte, arg2 []byte) {},
	func(arg1 []byte, arg2 []byte) {},
	func(arg1 []byte, arg2 []byte) {
		sub_math_operation_reg[arg1[1]%11](convert_to_value[arg1[0]](arg1), convert_to_value[arg2[0]](arg2))
	},
}

var mul_math_operation_reg = [10]func(any, any){
	func(val1 any, val2 any) {
		register_set[rbi](0, val1.(int8)*val2.(int8))
	},
	func(val1 any, val2 any) {
		register_set[rsi](0, val1.(int16)*val2.(int16))
	},
	func(val1 any, val2 any) {
		register_set[rli](0, val1.(int32)*val2.(int32))
	},
	func(val1 any, val2 any) {
		register_set[rlli](0, val1.(int64)*val2.(int64))
	},
	func(val1 any, val2 any) {
		register_set[rbui](0, val1.(uint8)*val2.(uint8))
	},
	func(val1 any, val2 any) {
		register_set[rsui](0, val1.(uint16)*val2.(uint16))
	},
	func(val1 any, val2 any) {
		register_set[rlui](0, val1.(uint32)*val2.(uint32))
	},
	func(val1 any, val2 any) {
		register_set[rllui](0, val1.(uint64)*val2.(uint64))
	},
	func(val1 any, val2 any) {
		register_set[rlf](0, val1.(float32)*val2.(float32))
	},
	func(val1 any, val2 any) {
		register_set[rllf](0, val1.(float64)*val2.(float64))
	},
}

var mul_math_operation = [13]func([]byte, []byte){
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rbi](0, byte_to_int8(arg1[1])*byte_to_int8(arg2[1]))
		} else {
			register_set[rbi](0, byte_to_int8(arg1[1])*register_get[arg2[1]](int(arg2[2])).(int8))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rsi](0, bytes_to_int16(arg1[1:])*bytes_to_int16(arg2[1:]))
		} else {
			register_set[rsi](0, bytes_to_int16(arg1[1:])*register_get[arg2[1]](int(arg2[2])).(int16))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rli](0, bytes_to_int32(arg1[1:])*bytes_to_int32(arg2[1:]))
		} else {
			register_set[rli](0, bytes_to_int32(arg1[1:])*register_get[arg2[1]](int(arg2[2])).(int32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rlli](0, bytes_to_int64(arg1[1:])*bytes_to_int64(arg2[1:]))
		} else {
			register_set[rlli](0, bytes_to_int64(arg1[1:])*register_get[arg2[1]](int(arg2[2])).(int64))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rbui](0, byte_to_uint8(arg1[1])*byte_to_uint8(arg2[1]))
		} else {
			register_set[rbui](0, byte_to_uint8(arg1[1])*register_get[arg2[1]](int(arg2[2])).(uint8))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])*bytes_to_uint16(arg2[1:]))
		} else {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])*register_get[arg2[1]](int(arg2[2])).(uint16))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])*bytes_to_uint32(arg2[1:]))
		} else {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])*register_get[arg2[1]](int(arg2[2])).(uint32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])*bytes_to_uint64(arg2[1:]))
		} else {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])*register_get[arg2[1]](int(arg2[2])).(uint64))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rlf](0, bytes_to_float32(arg1[1:])*bytes_to_float32(arg2[1:]))
		} else {
			register_set[rlf](0, bytes_to_float32(arg1[1:])*register_get[arg2[1]](int(arg2[2])).(float32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rllf](0, bytes_to_float64(arg1[1:])*bytes_to_float64(arg2[1:]))
		} else {
			register_set[rllf](0, bytes_to_float64(arg1[1:])*register_get[arg2[1]](int(arg2[2])).(float64))
		}
	},
	func(arg1 []byte, arg2 []byte) {},
	func(arg1 []byte, arg2 []byte) {},
	func(arg1 []byte, arg2 []byte) {
		mul_math_operation_reg[arg1[1]%11](convert_to_value[arg1[0]](arg1), convert_to_value[arg2[0]](arg2))
	},
}

var div_math_operation_reg = [10]func(any, any){
	func(val1 any, val2 any) {
		register_set[rbi](0, val1.(int8)/val2.(int8))
	},
	func(val1 any, val2 any) {
		register_set[rsi](0, val1.(int16)/val2.(int16))
	},
	func(val1 any, val2 any) {
		register_set[rli](0, val1.(int32)/val2.(int32))
	},
	func(val1 any, val2 any) {
		register_set[rlli](0, val1.(int64)/val2.(int64))
	},
	func(val1 any, val2 any) {
		register_set[rbui](0, val1.(uint8)/val2.(uint8))
	},
	func(val1 any, val2 any) {
		register_set[rsui](0, val1.(uint16)/val2.(uint16))
	},
	func(val1 any, val2 any) {
		register_set[rlui](0, val1.(uint32)/val2.(uint32))
	},
	func(val1 any, val2 any) {
		register_set[rllui](0, val1.(uint64)/val2.(uint64))
	},
	func(val1 any, val2 any) {
		register_set[rlf](0, val1.(float32)/val2.(float32))
	},
	func(val1 any, val2 any) {
		register_set[rllf](0, val1.(float64)/val2.(float64))
	},
}

var div_math_operation = [13]func([]byte, []byte){
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rbi](0, byte_to_int8(arg1[1])/byte_to_int8(arg2[1]))
		} else {
			register_set[rbi](0, byte_to_int8(arg1[1])/register_get[arg2[1]](int(arg2[2])).(int8))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rsi](0, bytes_to_int16(arg1[1:])/bytes_to_int16(arg2[1:]))
		} else {
			register_set[rsi](0, bytes_to_int16(arg1[1:])/register_get[arg2[1]](int(arg2[2])).(int16))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rli](0, bytes_to_int32(arg1[1:])/bytes_to_int32(arg2[1:]))
		} else {
			register_set[rli](0, bytes_to_int32(arg1[1:])/register_get[arg2[1]](int(arg2[2])).(int32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rlli](0, bytes_to_int64(arg1[1:])/bytes_to_int64(arg2[1:]))
		} else {
			register_set[rlli](0, bytes_to_int64(arg1[1:])/register_get[arg2[1]](int(arg2[2])).(int64))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rbui](0, byte_to_uint8(arg1[1])/byte_to_uint8(arg2[1]))
		} else {
			register_set[rbui](0, byte_to_uint8(arg1[1])/register_get[arg2[1]](int(arg2[2])).(uint8))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])/bytes_to_uint16(arg2[1:]))
		} else {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])/register_get[arg2[1]](int(arg2[2])).(uint16))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])/bytes_to_uint32(arg2[1:]))
		} else {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])/register_get[arg2[1]](int(arg2[2])).(uint32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])/bytes_to_uint64(arg2[1:]))
		} else {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])/register_get[arg2[1]](int(arg2[2])).(uint64))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rlf](0, bytes_to_float32(arg1[1:])/bytes_to_float32(arg2[1:]))
		} else {
			register_set[rlf](0, bytes_to_float32(arg1[1:])/register_get[arg2[1]](int(arg2[2])).(float32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rllf](0, bytes_to_float64(arg1[1:])/bytes_to_float64(arg2[1:]))
		} else {
			register_set[rllf](0, bytes_to_float64(arg1[1:])/register_get[arg2[1]](int(arg2[2])).(float64))
		}
	},
	func(arg1 []byte, arg2 []byte) {},
	func(arg1 []byte, arg2 []byte) {},
	func(arg1 []byte, arg2 []byte) {
		div_math_operation_reg[arg1[1]%11](convert_to_value[arg1[0]](arg1), convert_to_value[arg2[0]](arg2))
	},
}

var mod_math_operation_reg = [10]func(any, any){
	func(val1 any, val2 any) {
		register_set[rbi](0, val1.(int8)%val2.(int8))
	},
	func(val1 any, val2 any) {
		register_set[rsi](0, val1.(int16)%val2.(int16))
	},
	func(val1 any, val2 any) {
		register_set[rli](0, val1.(int32)%val2.(int32))
	},
	func(val1 any, val2 any) {
		register_set[rlli](0, val1.(int64)%val2.(int64))
	},
	func(val1 any, val2 any) {
		register_set[rbui](0, val1.(uint8)%val2.(uint8))
	},
	func(val1 any, val2 any) {
		register_set[rsui](0, val1.(uint16)%val2.(uint16))
	},
	func(val1 any, val2 any) {
		register_set[rlui](0, val1.(uint32)%val2.(uint32))
	},
	func(val1 any, val2 any) {
		register_set[rllui](0, val1.(uint64)%val2.(uint64))
	},
	func(val1 any, val2 any) {
		register_set[rlf](0, float32(int(val1.(float32))%int(val2.(float32))))
	},
	func(val1 any, val2 any) {
		register_set[rllf](0, float64(int(val1.(float64))%int(val2.(float64))))
	},
}

var mod_math_operation = [13]func([]byte, []byte){
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rbi](0, byte_to_int8(arg1[1])%byte_to_int8(arg2[1]))
		} else {
			register_set[rbi](0, byte_to_int8(arg1[1])%register_get[arg2[1]](int(arg2[2])).(int8))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rsi](0, bytes_to_int16(arg1[1:])%bytes_to_int16(arg2[1:]))
		} else {
			register_set[rsi](0, bytes_to_int16(arg1[1:])%register_get[arg2[1]](int(arg2[2])).(int16))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rli](0, bytes_to_int32(arg1[1:])%bytes_to_int32(arg2[1:]))
		} else {
			register_set[rli](0, bytes_to_int32(arg1[1:])%register_get[arg2[1]](int(arg2[2])).(int32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rlli](0, bytes_to_int64(arg1[1:])%bytes_to_int64(arg2[1:]))
		} else {
			register_set[rlli](0, bytes_to_int64(arg1[1:])%register_get[arg2[1]](int(arg2[2])).(int64))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rbui](0, byte_to_uint8(arg1[1])%byte_to_uint8(arg2[1]))
		} else {
			register_set[rbui](0, byte_to_uint8(arg1[1])%register_get[arg2[1]](int(arg2[2])).(uint8))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])%bytes_to_uint16(arg2[1:]))
		} else {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])%register_get[arg2[1]](int(arg2[2])).(uint16))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])%bytes_to_uint32(arg2[1:]))
		} else {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])%register_get[arg2[1]](int(arg2[2])).(uint32))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])%bytes_to_uint64(arg2[1:]))
		} else {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])%register_get[arg2[1]](int(arg2[2])).(uint64))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rlf](0, float32(int(bytes_to_float32(arg1[1:]))%int(bytes_to_float32(arg2[1:]))))
		} else {
			register_set[rlf](0, float32(int(bytes_to_float32(arg1[1:]))%int(register_get[arg2[1]](int(arg2[2])).(float32))))
		}
	},
	func(arg1 []byte, arg2 []byte) {
		if arg2[0] != t_reg {
			register_set[rllf](0, float64(int(bytes_to_float64(arg1[1:]))%int(bytes_to_float64(arg2[1:]))))
		} else {
			register_set[rllf](0, float64(int(bytes_to_float64(arg1[1:]))%int(register_get[arg2[1]](int(arg2[2])).(float64))))
		}
	},
	func(arg1 []byte, arg2 []byte) {},
	func(arg1 []byte, arg2 []byte) {},
	func(arg1 []byte, arg2 []byte) {
		mod_math_operation_reg[arg1[1]%11](convert_to_value[arg1[0]](arg1), convert_to_value[arg2[0]](arg2))
	},
}

var inc_math_operation = [21]func(){
	func() {
		register_bi[0]++
	},
	func() {
		register_si[0]++
	},
	func() {
		register_li[0]++
	},
	func() {
		register_lli[0]++
	},
	func() {
		register_bui[0]++
	},
	func() {
		register_sui[0]++
	},
	func() {
		register_lui[0]++
	},
	func() {
		register_llui[0]++
	},
	func() {
		register_lf[0]++
	},
	func() {
		register_llf[0]++
	},
	func() {},
	func() {
		register_rbi[0]++
	},
	func() {
		register_rsi[0]++
	},
	func() {
		register_rli[0]++
	},
	func() {
		register_rlli[0]++
	},
	func() {
		register_rbui[0]++
	},
	func() {
		register_rsui[0]++
	},
	func() {
		register_rlui[0]++
	},
	func() {
		register_rllui[0]++
	},
	func() {
		register_rlf[0]++
	},
	func() {
		register_rllf[0]++
	},
}

var dec_math_operation = [21]func(){
	func() {
		register_bi[0]--
	},
	func() {
		register_si[0]--
	},
	func() {
		register_li[0]--
	},
	func() {
		register_lli[0]--
	},
	func() {
		register_bui[0]--
	},
	func() {
		register_sui[0]--
	},
	func() {
		register_lui[0]--
	},
	func() {
		register_llui[0]--
	},
	func() {
		register_lf[0]--
	},
	func() {
		register_llf[0]--
	},
	func() {},
	func() {
		register_rbi[0]--
	},
	func() {
		register_rsi[0]--
	},
	func() {
		register_rli[0]--
	},
	func() {
		register_rlli[0]--
	},
	func() {
		register_rbui[0]--
	},
	func() {
		register_rsui[0]--
	},
	func() {
		register_rlui[0]--
	},
	func() {
		register_rllui[0]--
	},
	func() {
		register_rlf[0]--
	},
	func() {
		register_rllf[0]--
	},
}
