package main

var compare = [6]func([]byte, []byte, *function, *int, Program){
	func(arg1 []byte, arg2 []byte, function *function, index *int, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p) == convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(arg1 []byte, arg2 []byte, function *function, index *int, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p) != convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(arg1 []byte, arg2 []byte, function *function, index *int, p Program) {
		if p.register_cmp[0][0] != t_reg {
			compare_jg[p.register_cmp[0][0]](function, index, arg1, arg2, p)
		} else {
			compare_jg[p.register_cmp[0][1]%11](function, index, arg1, arg2, p)
		}
	},
	func(arg1 []byte, arg2 []byte, function *function, index *int, p Program) {
		if p.register_cmp[0][0] != t_reg {
			compare_jge[p.register_cmp[0][0]](function, index, arg1, arg2, p)
		} else {
			compare_jge[p.register_cmp[0][1]%11](function, index, arg1, arg2, p)
		}
	},
	func(arg1 []byte, arg2 []byte, function *function, index *int, p Program) {
		if p.register_cmp[0][0] != t_reg {
			compare_jl[p.register_cmp[0][0]](function, index, arg1, arg2, p)
		} else {
			compare_jl[p.register_cmp[0][1]%11](function, index, arg1, arg2, p)
		}
	},
	func(arg1 []byte, arg2 []byte, function *function, index *int, p Program) {
		if p.register_cmp[0][0] != t_reg {
			compare_jle[p.register_cmp[0][0]](function, index, arg1, arg2, p)
		} else {
			compare_jle[p.register_cmp[0][1]%11](function, index, arg1, arg2, p)
		}
	},
}

var compare_jg = [10]func(*function, *int, []byte, []byte, Program){
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(int8) > convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(int8) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(int16) > convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(int16) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(int32) > convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(int32) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(int64) > convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(int64) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(uint8) > convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(uint8) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(uint16) > convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(uint16) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(uint32) > convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(uint32) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(uint64) > convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(uint64) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(float32) > convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(float32) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(float64) > convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(float64) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
}

var compare_jge = [10]func(*function, *int, []byte, []byte, Program){
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(int8) >= convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(int8) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(int16) >= convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(int16) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(int32) >= convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(int32) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(int64) >= convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(int64) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(uint8) >= convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(uint8) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(uint16) >= convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(uint16) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(uint32) >= convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(uint32) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(uint64) >= convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(uint64) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(float32) >= convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(float32) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(float64) >= convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(float64) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
}

var compare_jl = [10]func(*function, *int, []byte, []byte, Program){
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(int8) < convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(int8) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(int16) < convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(int16) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(int32) < convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(int32) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(int64) < convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(int64) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(uint8) < convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(uint8) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(uint16) < convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(uint16) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(uint32) < convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(uint32) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(uint64) < convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(uint64) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(float32) < convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(float32) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(float64) < convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(float64) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
}

var compare_jle = [10]func(*function, *int, []byte, []byte, Program){
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(int8) <= convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(int8) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(int16) <= convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(int16) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(int32) <= convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(int32) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(int64) <= convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(int64) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(uint8) <= convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(uint8) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(uint16) <= convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(uint16) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(uint32) <= convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(uint32) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(uint64) <= convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(uint64) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(float32) <= convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(float32) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *function, index *int, arg1 []byte, arg2 []byte, p Program) {
		if convert_to_value[p.register_cmp[0][0]](p.register_cmp[0], p).(float64) <= convert_to_value[p.register_cmp[1][0]](p.register_cmp[1], p).(float64) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
}
