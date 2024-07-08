package main

import "strconv"

var registers = map[string]byte{
	"bi":    0,
	"si":    1,
	"li":    2,
	"lli":   3,
	"bui":   4,
	"sui":   5,
	"lui":   6,
	"llui":  7,
	"lf":    8,
	"llf":   9,
	"s":     10,
	"rbi":   11,
	"rsi":   12,
	"rli":   13,
	"rlli":  14,
	"rbui":  15,
	"rsui":  16,
	"rlui":  17,
	"rllui": 18,
	"rlf":   19,
	"rllf":  20,
	"rs":    21,
}

func is_register(str string) (bool, string, byte) {
	reg := ""
	i := 0
	for ; i < len(str) && str[i] >= 'a' && str[i] <= 'z'; i++ {
		reg += string(str[i])
	}
	index, _ := strconv.ParseUint(str[i:], 10, 8)

	_, ok := registers[reg]
	return ok, reg, byte(index)
}
