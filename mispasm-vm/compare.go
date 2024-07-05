package main

var register_cmp = [2][]byte{}

var compare = [6]func(arg1 []byte, arg2 []byte, function *Function, index *int){
	func(arg1 []byte, arg2 []byte, function *Function, index *int) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]) == convert_to_value[register_cmp[1][0]](register_cmp[1]) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(arg1 []byte, arg2 []byte, function *Function, index *int) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]) != convert_to_value[register_cmp[1][0]](register_cmp[1]) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(arg1 []byte, arg2 []byte, function *Function, index *int) {
		if register_cmp[0][0] != t_reg {
			compare_jg[register_cmp[0][0]](function, index, arg1, arg2)
		} else {
			compare_jg[register_cmp[0][1]%11](function, index, arg1, arg2)
		}
	},
	func(arg1 []byte, arg2 []byte, function *Function, index *int) {
		if register_cmp[0][0] != t_reg {
			compare_jge[register_cmp[0][0]](function, index, arg1, arg2)
		} else {
			compare_jge[register_cmp[0][1]%11](function, index, arg1, arg2)
		}
	},
	func(arg1 []byte, arg2 []byte, function *Function, index *int) {
		if register_cmp[0][0] != t_reg {
			compare_jl[register_cmp[0][0]](function, index, arg1, arg2)
		} else {
			compare_jl[register_cmp[0][1]%11](function, index, arg1, arg2)
		}
	},
	func(arg1 []byte, arg2 []byte, function *Function, index *int) {
		if register_cmp[0][0] != t_reg {
			compare_jle[register_cmp[0][0]](function, index, arg1, arg2)
		} else {
			compare_jle[register_cmp[0][1]%11](function, index, arg1, arg2)
		}
	},
}

var compare_jg = [10]func(function *Function, index *int, arg1 []byte, arg2 []byte){
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int8) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(int8) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int16) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(int16) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int32) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(int32) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int64) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(int64) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint8) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint8) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint16) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint16) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint32) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint32) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint64) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint64) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(float32) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(float32) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(float64) > convert_to_value[register_cmp[1][0]](register_cmp[1]).(float64) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
}

var compare_jge = [10]func(function *Function, index *int, arg1 []byte, arg2 []byte){
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int8) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(int8) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int16) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(int16) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int32) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(int32) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int64) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(int64) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint8) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint8) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint16) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint16) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint32) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint32) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint64) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint64) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(float32) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(float32) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(float64) >= convert_to_value[register_cmp[1][0]](register_cmp[1]).(float64) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
}

var compare_jl = [10]func(function *Function, index *int, arg1 []byte, arg2 []byte){
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int8) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(int8) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int16) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(int16) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int32) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(int32) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int64) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(int64) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint8) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint8) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint16) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint16) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint32) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint32) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint64) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint64) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(float32) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(float32) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(float64) < convert_to_value[register_cmp[1][0]](register_cmp[1]).(float64) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
}

var compare_jle = [10]func(function *Function, index *int, arg1 []byte, arg2 []byte){
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int8) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(int8) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int16) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(int16) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int32) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(int32) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(int64) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(int64) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint8) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint8) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint16) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint16) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint32) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint32) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(uint64) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(uint64) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(float32) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(float32) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
	func(function *Function, index *int, arg1 []byte, arg2 []byte) {
		if convert_to_value[register_cmp[0][0]](register_cmp[0]).(float64) <= convert_to_value[register_cmp[1][0]](register_cmp[1]).(float64) {
			*index = function.labels[arg1[1]]
		} else {
			*index = function.labels[arg2[1]]
		}
	},
}
