package main

var and_bitwise_operation_reg = [8]func(any, any, *Program){
	func(val1 any, val2 any, p *Program) {
		register_set[rbi](0, val1.(int8)&val2.(int8), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rsi](0, val1.(int16)&val2.(int16), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rli](0, val1.(int32)&val2.(int32), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rlli](0, val1.(int64)&val2.(int64), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rbui](0, val1.(uint8)&val2.(uint8), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rsui](0, val1.(uint16)&val2.(uint16), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rlui](0, val1.(uint32)&val2.(uint32), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rllui](0, val1.(uint64)&val2.(uint64), p)
	},
}
var and_bitwise_operation = [12]func([]byte, []byte, *Program){
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rbi](0, byte_to_int8(arg1[1])&byte_to_int8(arg2[1]), p)
		} else {
			register_set[rbi](0, byte_to_int8(arg1[1])&register_get[arg2[1]](arg2[2], *p).(int8), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rsi](0, bytes_to_int16(arg1[1:])&bytes_to_int16(arg2[1:]), p)
		} else {
			register_set[rsi](0, bytes_to_int16(arg1[1:])&register_get[arg2[1]](arg2[2], *p).(int16), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rli](0, bytes_to_int32(arg1[1:])&bytes_to_int32(arg2[1:]), p)
		} else {
			register_set[rli](0, bytes_to_int32(arg1[1:])&register_get[arg2[1]](arg2[2], *p).(int32), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rlli](0, bytes_to_int64(arg1[1:])&bytes_to_int64(arg2[1:]), p)
		} else {
			register_set[rlli](0, bytes_to_int64(arg1[1:])&register_get[arg2[1]](arg2[2], *p).(int64), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rbui](0, byte_to_uint8(arg1[1])&byte_to_uint8(arg2[1]), p)
		} else {
			register_set[rbui](0, byte_to_uint8(arg1[1])&register_get[arg2[1]](arg2[2], *p).(uint8), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])&bytes_to_uint16(arg2[1:]), p)
		} else {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])&register_get[arg2[1]](arg2[2], *p).(uint16), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])&bytes_to_uint32(arg2[1:]), p)
		} else {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])&register_get[arg2[1]](arg2[2], *p).(uint32), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])&bytes_to_uint64(arg2[1:]), p)
		} else {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])&register_get[arg2[1]](arg2[2], *p).(uint64), p)
		}
	},
	func(_ []byte, _ []byte, _ *Program) {},
	func(_ []byte, _ []byte, _ *Program) {},
	func(_ []byte, _ []byte, _ *Program) {},
	func(arg1 []byte, arg2 []byte, p *Program) {
		and_bitwise_operation_reg[arg1[1]%11](convert_to_value[arg1[0]](arg1, *p), convert_to_value[arg2[0]](arg2, *p), p)
	},
}

var or_bitwise_operation_reg = [8]func(any, any, *Program){
	func(val1 any, val2 any, p *Program) {
		register_set[rbi](0, val1.(int8)|val2.(int8), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rsi](0, val1.(int16)|val2.(int16), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rli](0, val1.(int32)|val2.(int32), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rlli](0, val1.(int64)|val2.(int64), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rbui](0, val1.(uint8)|val2.(uint8), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rsui](0, val1.(uint16)|val2.(uint16), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rlui](0, val1.(uint32)|val2.(uint32), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rllui](0, val1.(uint64)|val2.(uint64), p)
	},
}
var or_bitwise_operation = [12]func([]byte, []byte, *Program){
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rbi](0, byte_to_int8(arg1[1])|byte_to_int8(arg2[1]), p)
		} else {
			register_set[rbi](0, byte_to_int8(arg1[1])|register_get[arg2[1]](arg2[2], *p).(int8), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rsi](0, bytes_to_int16(arg1[1:])|bytes_to_int16(arg2[1:]), p)
		} else {
			register_set[rsi](0, bytes_to_int16(arg1[1:])|register_get[arg2[1]](arg2[2], *p).(int16), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rli](0, bytes_to_int32(arg1[1:])|bytes_to_int32(arg2[1:]), p)
		} else {
			register_set[rli](0, bytes_to_int32(arg1[1:])|register_get[arg2[1]](arg2[2], *p).(int32), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rlli](0, bytes_to_int64(arg1[1:])|bytes_to_int64(arg2[1:]), p)
		} else {
			register_set[rlli](0, bytes_to_int64(arg1[1:])|register_get[arg2[1]](arg2[2], *p).(int64), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rbui](0, byte_to_uint8(arg1[1])|byte_to_uint8(arg2[1]), p)
		} else {
			register_set[rbui](0, byte_to_uint8(arg1[1])|register_get[arg2[1]](arg2[2], *p).(uint8), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])|bytes_to_uint16(arg2[1:]), p)
		} else {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])|register_get[arg2[1]](arg2[2], *p).(uint16), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])|bytes_to_uint32(arg2[1:]), p)
		} else {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])|register_get[arg2[1]](arg2[2], *p).(uint32), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])|bytes_to_uint64(arg2[1:]), p)
		} else {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])|register_get[arg2[1]](arg2[2], *p).(uint64), p)
		}
	},
	func(_ []byte, _ []byte, _ *Program) {},
	func(_ []byte, _ []byte, _ *Program) {},
	func(_ []byte, _ []byte, _ *Program) {},
	func(arg1 []byte, arg2 []byte, p *Program) {
		or_bitwise_operation_reg[arg1[1]%11](convert_to_value[arg1[0]](arg1, *p), convert_to_value[arg2[0]](arg2, *p), p)
	},
}

var xor_bitwise_operation_reg = [8]func(any, any, *Program){
	func(val1 any, val2 any, p *Program) {
		register_set[rbi](0, val1.(int8)^val2.(int8), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rsi](0, val1.(int16)^val2.(int16), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rli](0, val1.(int32)^val2.(int32), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rlli](0, val1.(int64)^val2.(int64), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rbui](0, val1.(uint8)^val2.(uint8), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rsui](0, val1.(uint16)^val2.(uint16), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rlui](0, val1.(uint32)^val2.(uint32), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rllui](0, val1.(uint64)^val2.(uint64), p)
	},
}
var xor_bitwise_operation = [12]func([]byte, []byte, *Program){
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rbi](0, byte_to_int8(arg1[1])^byte_to_int8(arg2[1]), p)
		} else {
			register_set[rbi](0, byte_to_int8(arg1[1])^register_get[arg2[1]](arg2[2], *p).(int8), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rsi](0, bytes_to_int16(arg1[1:])^bytes_to_int16(arg2[1:]), p)
		} else {
			register_set[rsi](0, bytes_to_int16(arg1[1:])^register_get[arg2[1]](arg2[2], *p).(int16), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rli](0, bytes_to_int32(arg1[1:])^bytes_to_int32(arg2[1:]), p)
		} else {
			register_set[rli](0, bytes_to_int32(arg1[1:])^register_get[arg2[1]](arg2[2], *p).(int32), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rlli](0, bytes_to_int64(arg1[1:])^bytes_to_int64(arg2[1:]), p)
		} else {
			register_set[rlli](0, bytes_to_int64(arg1[1:])^register_get[arg2[1]](arg2[2], *p).(int64), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rbui](0, byte_to_uint8(arg1[1])^byte_to_uint8(arg2[1]), p)
		} else {
			register_set[rbui](0, byte_to_uint8(arg1[1])^register_get[arg2[1]](arg2[2], *p).(uint8), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])^bytes_to_uint16(arg2[1:]), p)
		} else {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])^register_get[arg2[1]](arg2[2], *p).(uint16), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])^bytes_to_uint32(arg2[1:]), p)
		} else {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])^register_get[arg2[1]](arg2[2], *p).(uint32), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])^bytes_to_uint64(arg2[1:]), p)
		} else {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])^register_get[arg2[1]](arg2[2], *p).(uint64), p)
		}
	},
	func(_ []byte, _ []byte, _ *Program) {},
	func(_ []byte, _ []byte, _ *Program) {},
	func(_ []byte, _ []byte, _ *Program) {},
	func(arg1 []byte, arg2 []byte, p *Program) {
		xor_bitwise_operation_reg[arg1[1]%11](convert_to_value[arg1[0]](arg1, *p), convert_to_value[arg2[0]](arg2, *p), p)
	},
}

var shr_bitwise_operation_reg = [8]func(any, any, *Program){
	func(val1 any, val2 any, p *Program) {
		register_set[rbi](0, val1.(int8)>>val2.(int8), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rsi](0, val1.(int16)>>val2.(int16), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rli](0, val1.(int32)>>val2.(int32), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rlli](0, val1.(int64)>>val2.(int64), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rbui](0, val1.(uint8)>>val2.(uint8), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rsui](0, val1.(uint16)>>val2.(uint16), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rlui](0, val1.(uint32)>>val2.(uint32), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rllui](0, val1.(uint64)>>val2.(uint64), p)
	},
}
var shr_bitwise_operation = [12]func([]byte, []byte, *Program){
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rbi](0, byte_to_int8(arg1[1])>>byte_to_int8(arg2[1]), p)
		} else {
			register_set[rbi](0, byte_to_int8(arg1[1])>>register_get[arg2[1]](arg2[2], *p).(int8), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rsi](0, bytes_to_int16(arg1[1:])>>bytes_to_int16(arg2[1:]), p)
		} else {
			register_set[rsi](0, bytes_to_int16(arg1[1:])>>register_get[arg2[1]](arg2[2], *p).(int16), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rli](0, bytes_to_int32(arg1[1:])>>bytes_to_int32(arg2[1:]), p)
		} else {
			register_set[rli](0, bytes_to_int32(arg1[1:])>>register_get[arg2[1]](arg2[2], *p).(int32), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rlli](0, bytes_to_int64(arg1[1:])>>bytes_to_int64(arg2[1:]), p)
		} else {
			register_set[rlli](0, bytes_to_int64(arg1[1:])>>register_get[arg2[1]](arg2[2], *p).(int64), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rbui](0, byte_to_uint8(arg1[1])>>byte_to_uint8(arg2[1]), p)
		} else {
			register_set[rbui](0, byte_to_uint8(arg1[1])>>register_get[arg2[1]](arg2[2], *p).(uint8), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])>>bytes_to_uint16(arg2[1:]), p)
		} else {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])>>register_get[arg2[1]](arg2[2], *p).(uint16), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])>>bytes_to_uint32(arg2[1:]), p)
		} else {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])>>register_get[arg2[1]](arg2[2], *p).(uint32), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])>>bytes_to_uint64(arg2[1:]), p)
		} else {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])>>register_get[arg2[1]](arg2[2], *p).(uint64), p)
		}
	},
	func(_ []byte, _ []byte, _ *Program) {},
	func(_ []byte, _ []byte, _ *Program) {},
	func(_ []byte, _ []byte, _ *Program) {},
	func(arg1 []byte, arg2 []byte, p *Program) {
		shr_bitwise_operation_reg[arg1[1]%11](convert_to_value[arg1[0]](arg1, *p), convert_to_value[arg2[0]](arg2, *p), p)
	},
}

var shl_bitwise_operation_reg = [8]func(any, any, *Program){
	func(val1 any, val2 any, p *Program) {
		register_set[rbi](0, val1.(int8)<<val2.(int8), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rsi](0, val1.(int16)<<val2.(int16), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rli](0, val1.(int32)<<val2.(int32), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rlli](0, val1.(int64)<<val2.(int64), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rbui](0, val1.(uint8)<<val2.(uint8), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rsui](0, val1.(uint16)<<val2.(uint16), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rlui](0, val1.(uint32)<<val2.(uint32), p)
	},
	func(val1 any, val2 any, p *Program) {
		register_set[rllui](0, val1.(uint64)<<val2.(uint64), p)
	},
}
var shl_bitwise_operation = [12]func([]byte, []byte, *Program){
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rbi](0, byte_to_int8(arg1[1])<<byte_to_int8(arg2[1]), p)
		} else {
			register_set[rbi](0, byte_to_int8(arg1[1])<<register_get[arg2[1]](arg2[2], *p).(int8), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rsi](0, bytes_to_int16(arg1[1:])<<bytes_to_int16(arg2[1:]), p)
		} else {
			register_set[rsi](0, bytes_to_int16(arg1[1:])<<register_get[arg2[1]](arg2[2], *p).(int16), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rli](0, bytes_to_int32(arg1[1:])<<bytes_to_int32(arg2[1:]), p)
		} else {
			register_set[rli](0, bytes_to_int32(arg1[1:])<<register_get[arg2[1]](arg2[2], *p).(int32), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rlli](0, bytes_to_int64(arg1[1:])<<bytes_to_int64(arg2[1:]), p)
		} else {
			register_set[rlli](0, bytes_to_int64(arg1[1:])<<register_get[arg2[1]](arg2[2], *p).(int64), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rbui](0, byte_to_uint8(arg1[1])<<byte_to_uint8(arg2[1]), p)
		} else {
			register_set[rbui](0, byte_to_uint8(arg1[1])<<register_get[arg2[1]](arg2[2], *p).(uint8), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])<<bytes_to_uint16(arg2[1:]), p)
		} else {
			register_set[rsui](0, bytes_to_uint16(arg1[1:])<<register_get[arg2[1]](arg2[2], *p).(uint16), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])<<bytes_to_uint32(arg2[1:]), p)
		} else {
			register_set[rlui](0, bytes_to_uint32(arg1[1:])<<register_get[arg2[1]](arg2[2], *p).(uint32), p)
		}
	},
	func(arg1 []byte, arg2 []byte, p *Program) {
		if arg2[0] != t_reg {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])<<bytes_to_uint64(arg2[1:]), p)
		} else {
			register_set[rllui](0, bytes_to_uint64(arg1[1:])<<register_get[arg2[1]](arg2[2], *p).(uint64), p)
		}
	},
	func(_ []byte, _ []byte, _ *Program) {},
	func(_ []byte, _ []byte, _ *Program) {},
	func(_ []byte, _ []byte, _ *Program) {},
	func(arg1 []byte, arg2 []byte, p *Program) {
		shl_bitwise_operation_reg[arg1[1]%11](convert_to_value[arg1[0]](arg1, *p), convert_to_value[arg2[0]](arg2, *p), p)
	},
}

var not_bitwise_operation_reg = [8]func(any, *Program){
	func(val1 any, p *Program) {
		register_set[rbi](0, ^val1.(int8), p)
	},
	func(val1 any, p *Program) {
		register_set[rsi](0, ^val1.(int16), p)
	},
	func(val1 any, p *Program) {
		register_set[rli](0, ^val1.(int32), p)
	},
	func(val1 any, p *Program) {
		register_set[rlli](0, ^val1.(int64), p)
	},
	func(val1 any, p *Program) {
		register_set[rbui](0, ^val1.(uint8), p)
	},
	func(val1 any, p *Program) {
		register_set[rsui](0, ^val1.(uint16), p)
	},
	func(val1 any, p *Program) {
		register_set[rlui](0, ^val1.(uint32), p)
	},
	func(val1 any, p *Program) {
		register_set[rllui](0, ^val1.(uint64), p)
	},
}
var not_bitwise_operation = [12]func([]byte, *Program){
	func(arg1 []byte, p *Program) {
		register_set[rbi](0, ^byte_to_int8(arg1[1]), p)
	},
	func(arg1 []byte, p *Program) {
		register_set[rsi](0, ^bytes_to_int16(arg1[1:]), p)
	},
	func(arg1 []byte, p *Program) {
		register_set[rli](0, ^bytes_to_int32(arg1[1:]), p)
	},
	func(arg1 []byte, p *Program) {
		register_set[rlli](0, ^bytes_to_int64(arg1[1:]), p)
	},
	func(arg1 []byte, p *Program) {
		register_set[rbui](0, ^byte_to_uint8(arg1[1]), p)
	},
	func(arg1 []byte, p *Program) {
		register_set[rsui](0, ^bytes_to_uint16(arg1[1:]), p)
	},
	func(arg1 []byte, p *Program) {
		register_set[rlui](0, ^bytes_to_uint32(arg1[1:]), p)
	},
	func(arg1 []byte, p *Program) {
		register_set[rllui](0, ^bytes_to_uint64(arg1[1:]), p)
	},
	func(_ []byte, _ *Program) {},
	func(_ []byte, _ *Program) {},
	func(_ []byte, _ *Program) {},
	func(arg1 []byte, p *Program) {
		not_bitwise_operation_reg[arg1[1]%11](convert_to_value[arg1[0]](arg1, *p), p)
	},
}
