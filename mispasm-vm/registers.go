package main

const (
	bi byte = iota
	si
	li
	lli
	bui
	sui
	lui
	llui
	lf
	llf
	s
	rbi
	rsi
	rli
	rlli
	rbui
	rsui
	rlui
	rllui
	rlf
	rllf
	rs
)

var register_set = [22]func(byte, any, *Program){
	func(index byte, data any, p *Program) { p.register_bi[index] = data.(int8) },
	func(index byte, data any, p *Program) { p.register_si[index] = data.(int16) },
	func(index byte, data any, p *Program) { p.register_li[index] = data.(int32) },
	func(index byte, data any, p *Program) { p.register_lli[index] = data.(int64) },
	func(index byte, data any, p *Program) { p.register_bui[index] = data.(uint8) },
	func(index byte, data any, p *Program) { p.register_sui[index] = data.(uint16) },
	func(index byte, data any, p *Program) { p.register_lui[index] = data.(uint32) },
	func(index byte, data any, p *Program) { p.register_llui[index] = data.(uint64) },
	func(index byte, data any, p *Program) { p.register_lf[index] = data.(float32) },
	func(index byte, data any, p *Program) { p.register_llf[index] = data.(float64) },
	func(index byte, data any, p *Program) { p.register_s[index] = data.(string) },
	func(index byte, data any, p *Program) { p.register_rbi[index] = data.(int8) },
	func(index byte, data any, p *Program) { p.register_rsi[index] = data.(int16) },
	func(index byte, data any, p *Program) { p.register_rli[index] = data.(int32) },
	func(index byte, data any, p *Program) { p.register_rlli[index] = data.(int64) },
	func(index byte, data any, p *Program) { p.register_rbui[index] = data.(uint8) },
	func(index byte, data any, p *Program) { p.register_rsui[index] = data.(uint16) },
	func(index byte, data any, p *Program) { p.register_rlui[index] = data.(uint32) },
	func(index byte, data any, p *Program) { p.register_rllui[index] = data.(uint64) },
	func(index byte, data any, p *Program) { p.register_rlf[index] = data.(float32) },
	func(index byte, data any, p *Program) { p.register_rllf[index] = data.(float64) },
	func(index byte, data any, p *Program) { p.register_rs[index] = data.(string) },
}

var register_get = [22]func(byte, Program) any{
	func(index byte, p Program) any { return p.register_bi[index] },
	func(index byte, p Program) any { return p.register_si[index] },
	func(index byte, p Program) any { return p.register_li[index] },
	func(index byte, p Program) any { return p.register_lli[index] },
	func(index byte, p Program) any { return p.register_bui[index] },
	func(index byte, p Program) any { return p.register_sui[index] },
	func(index byte, p Program) any { return p.register_lui[index] },
	func(index byte, p Program) any { return p.register_llui[index] },
	func(index byte, p Program) any { return p.register_lf[index] },
	func(index byte, p Program) any { return p.register_llf[index] },
	func(index byte, p Program) any { return p.register_s[index] },
	func(index byte, p Program) any { return p.register_rbi[index] },
	func(index byte, p Program) any { return p.register_rsi[index] },
	func(index byte, p Program) any { return p.register_rli[index] },
	func(index byte, p Program) any { return p.register_rlli[index] },
	func(index byte, p Program) any { return p.register_rbui[index] },
	func(index byte, p Program) any { return p.register_rsui[index] },
	func(index byte, p Program) any { return p.register_rlui[index] },
	func(index byte, p Program) any { return p.register_rllui[index] },
	func(index byte, p Program) any { return p.register_rlf[index] },
	func(index byte, p Program) any { return p.register_rllf[index] },
	func(index byte, p Program) any { return p.register_rs[index] },
}

var register_force_set = [11][22]func(byte, any, *Program){
	{
		func(index byte, data any, p *Program) { p.register_bi[index] = data.(int8) },
		func(index byte, data any, p *Program) { p.register_si[index] = int16(data.(int8)) },
		func(index byte, data any, p *Program) { p.register_li[index] = int32(data.(int8)) },
		func(index byte, data any, p *Program) { p.register_lli[index] = int64(data.(int8)) },
		func(index byte, data any, p *Program) { p.register_bui[index] = uint8(data.(int8)) },
		func(index byte, data any, p *Program) { p.register_sui[index] = uint16(data.(int8)) },
		func(index byte, data any, p *Program) { p.register_lui[index] = uint32(data.(int8)) },
		func(index byte, data any, p *Program) { p.register_llui[index] = uint64(data.(int8)) },
		func(index byte, data any, p *Program) { p.register_lf[index] = float32(data.(int8)) },
		func(index byte, data any, p *Program) { p.register_llf[index] = float64(data.(int8)) },
		func(index byte, data any, p *Program) { p.register_s[index] = string(rune(data.(int8))) },
		func(index byte, data any, p *Program) { p.register_rbi[index] = data.(int8) },
		func(index byte, data any, p *Program) { p.register_rsi[index] = int16(data.(int8)) },
		func(index byte, data any, p *Program) { p.register_rli[index] = int32(data.(int8)) },
		func(index byte, data any, p *Program) { p.register_rlli[index] = int64(data.(int8)) },
		func(index byte, data any, p *Program) { p.register_rbui[index] = uint8(data.(int8)) },
		func(index byte, data any, p *Program) { p.register_rsui[index] = uint16(data.(int8)) },
		func(index byte, data any, p *Program) { p.register_rlui[index] = uint32(data.(int8)) },
		func(index byte, data any, p *Program) { p.register_rllui[index] = uint64(data.(int8)) },
		func(index byte, data any, p *Program) { p.register_rlf[index] = float32(data.(int8)) },
		func(index byte, data any, p *Program) { p.register_rllf[index] = float64(data.(int8)) },
		func(index byte, data any, p *Program) { p.register_rs[index] = string(rune(data.(int8))) },
	},
	{
		func(index byte, data any, p *Program) { p.register_bi[index] = int8(data.(int16)) },
		func(index byte, data any, p *Program) { p.register_si[index] = data.(int16) },
		func(index byte, data any, p *Program) { p.register_li[index] = int32(data.(int16)) },
		func(index byte, data any, p *Program) { p.register_lli[index] = int64(data.(int16)) },
		func(index byte, data any, p *Program) { p.register_bui[index] = uint8(data.(int16)) },
		func(index byte, data any, p *Program) { p.register_sui[index] = uint16(data.(int16)) },
		func(index byte, data any, p *Program) { p.register_lui[index] = uint32(data.(int16)) },
		func(index byte, data any, p *Program) { p.register_llui[index] = uint64(data.(int16)) },
		func(index byte, data any, p *Program) { p.register_lf[index] = float32(data.(int16)) },
		func(index byte, data any, p *Program) { p.register_llf[index] = float64(data.(int16)) },
		func(index byte, data any, p *Program) { p.register_s[index] = string(rune(data.(int16))) },
		func(index byte, data any, p *Program) { p.register_rbi[index] = int8(data.(int16)) },
		func(index byte, data any, p *Program) { p.register_rsi[index] = data.(int16) },
		func(index byte, data any, p *Program) { p.register_rli[index] = int32(data.(int16)) },
		func(index byte, data any, p *Program) { p.register_rlli[index] = int64(data.(int16)) },
		func(index byte, data any, p *Program) { p.register_rbui[index] = uint8(data.(int16)) },
		func(index byte, data any, p *Program) { p.register_rsui[index] = uint16(data.(int16)) },
		func(index byte, data any, p *Program) { p.register_rlui[index] = uint32(data.(int16)) },
		func(index byte, data any, p *Program) { p.register_rllui[index] = uint64(data.(int16)) },
		func(index byte, data any, p *Program) { p.register_rlf[index] = float32(data.(int16)) },
		func(index byte, data any, p *Program) { p.register_rllf[index] = float64(data.(int16)) },
		func(index byte, data any, p *Program) { p.register_rs[index] = string(rune(data.(int16))) },
	},
	{
		func(index byte, data any, p *Program) { p.register_bi[index] = int8(data.(int32)) },
		func(index byte, data any, p *Program) { p.register_si[index] = int16(data.(int32)) },
		func(index byte, data any, p *Program) { p.register_li[index] = data.(int32) },
		func(index byte, data any, p *Program) { p.register_lli[index] = int64(data.(int32)) },
		func(index byte, data any, p *Program) { p.register_bui[index] = uint8(data.(int32)) },
		func(index byte, data any, p *Program) { p.register_sui[index] = uint16(data.(int32)) },
		func(index byte, data any, p *Program) { p.register_lui[index] = uint32(data.(int32)) },
		func(index byte, data any, p *Program) { p.register_llui[index] = uint64(data.(int32)) },
		func(index byte, data any, p *Program) { p.register_lf[index] = float32(data.(int32)) },
		func(index byte, data any, p *Program) { p.register_llf[index] = float64(data.(int32)) },
		func(index byte, data any, p *Program) { p.register_s[index] = string(rune(data.(int32))) },
		func(index byte, data any, p *Program) { p.register_rbi[index] = int8(data.(int32)) },
		func(index byte, data any, p *Program) { p.register_rsi[index] = int16(data.(int32)) },
		func(index byte, data any, p *Program) { p.register_rli[index] = data.(int32) },
		func(index byte, data any, p *Program) { p.register_rlli[index] = int64(data.(int32)) },
		func(index byte, data any, p *Program) { p.register_rbui[index] = uint8(data.(int32)) },
		func(index byte, data any, p *Program) { p.register_rsui[index] = uint16(data.(int32)) },
		func(index byte, data any, p *Program) { p.register_rlui[index] = uint32(data.(int32)) },
		func(index byte, data any, p *Program) { p.register_rllui[index] = uint64(data.(int32)) },
		func(index byte, data any, p *Program) { p.register_rlf[index] = float32(data.(int32)) },
		func(index byte, data any, p *Program) { p.register_rllf[index] = float64(data.(int32)) },
		func(index byte, data any, p *Program) { p.register_rs[index] = string(rune(data.(int32))) },
	},
	{
		func(index byte, data any, p *Program) { p.register_bi[index] = int8(data.(int64)) },
		func(index byte, data any, p *Program) { p.register_si[index] = int16(data.(int64)) },
		func(index byte, data any, p *Program) { p.register_li[index] = int32(data.(int64)) },
		func(index byte, data any, p *Program) { p.register_lli[index] = data.(int64) },
		func(index byte, data any, p *Program) { p.register_bui[index] = uint8(data.(int64)) },
		func(index byte, data any, p *Program) { p.register_sui[index] = uint16(data.(int64)) },
		func(index byte, data any, p *Program) { p.register_lui[index] = uint32(data.(int64)) },
		func(index byte, data any, p *Program) { p.register_llui[index] = uint64(data.(int64)) },
		func(index byte, data any, p *Program) { p.register_lf[index] = float32(data.(int64)) },
		func(index byte, data any, p *Program) { p.register_llf[index] = float64(data.(int64)) },
		func(index byte, data any, p *Program) { p.register_s[index] = string(rune(data.(int64))) },
		func(index byte, data any, p *Program) { p.register_rbi[index] = int8(data.(int64)) },
		func(index byte, data any, p *Program) { p.register_rsi[index] = int16(data.(int64)) },
		func(index byte, data any, p *Program) { p.register_rli[index] = int32(data.(int64)) },
		func(index byte, data any, p *Program) { p.register_rlli[index] = data.(int64) },
		func(index byte, data any, p *Program) { p.register_rbui[index] = uint8(data.(int64)) },
		func(index byte, data any, p *Program) { p.register_rsui[index] = uint16(data.(int64)) },
		func(index byte, data any, p *Program) { p.register_rlui[index] = uint32(data.(int64)) },
		func(index byte, data any, p *Program) { p.register_rllui[index] = uint64(data.(int64)) },
		func(index byte, data any, p *Program) { p.register_rlf[index] = float32(data.(int64)) },
		func(index byte, data any, p *Program) { p.register_rllf[index] = float64(data.(int64)) },
		func(index byte, data any, p *Program) { p.register_rs[index] = string(rune(data.(int64))) },
	},
	{
		func(index byte, data any, p *Program) { p.register_bi[index] = int8(data.(uint8)) },
		func(index byte, data any, p *Program) { p.register_si[index] = int16(data.(uint8)) },
		func(index byte, data any, p *Program) { p.register_li[index] = int32(data.(uint8)) },
		func(index byte, data any, p *Program) { p.register_lli[index] = int64(data.(uint8)) },
		func(index byte, data any, p *Program) { p.register_bui[index] = data.(uint8) },
		func(index byte, data any, p *Program) { p.register_sui[index] = uint16(data.(uint8)) },
		func(index byte, data any, p *Program) { p.register_lui[index] = uint32(data.(uint8)) },
		func(index byte, data any, p *Program) { p.register_llui[index] = uint64(data.(uint8)) },
		func(index byte, data any, p *Program) { p.register_lf[index] = float32(data.(uint8)) },
		func(index byte, data any, p *Program) { p.register_llf[index] = float64(data.(uint8)) },
		func(index byte, data any, p *Program) { p.register_s[index] = string(rune(data.(uint8))) },
		func(index byte, data any, p *Program) { p.register_rbi[index] = int8(data.(uint8)) },
		func(index byte, data any, p *Program) { p.register_rsi[index] = int16(data.(uint8)) },
		func(index byte, data any, p *Program) { p.register_rli[index] = int32(data.(uint8)) },
		func(index byte, data any, p *Program) { p.register_rlli[index] = int64(data.(uint8)) },
		func(index byte, data any, p *Program) { p.register_rbui[index] = data.(uint8) },
		func(index byte, data any, p *Program) { p.register_rsui[index] = uint16(data.(uint8)) },
		func(index byte, data any, p *Program) { p.register_rlui[index] = uint32(data.(uint8)) },
		func(index byte, data any, p *Program) { p.register_rllui[index] = uint64(data.(uint8)) },
		func(index byte, data any, p *Program) { p.register_rlf[index] = float32(data.(uint8)) },
		func(index byte, data any, p *Program) { p.register_rllf[index] = float64(data.(uint8)) },
		func(index byte, data any, p *Program) { p.register_rs[index] = string(rune(data.(uint8))) },
	},
	{
		func(index byte, data any, p *Program) { p.register_bi[index] = int8(data.(uint16)) },
		func(index byte, data any, p *Program) { p.register_si[index] = int16(data.(uint16)) },
		func(index byte, data any, p *Program) { p.register_li[index] = int32(data.(uint16)) },
		func(index byte, data any, p *Program) { p.register_lli[index] = int64(data.(uint16)) },
		func(index byte, data any, p *Program) { p.register_bui[index] = uint8(data.(uint16)) },
		func(index byte, data any, p *Program) { p.register_sui[index] = data.(uint16) },
		func(index byte, data any, p *Program) { p.register_lui[index] = uint32(data.(uint16)) },
		func(index byte, data any, p *Program) { p.register_llui[index] = uint64(data.(uint16)) },
		func(index byte, data any, p *Program) { p.register_lf[index] = float32(data.(uint16)) },
		func(index byte, data any, p *Program) { p.register_llf[index] = float64(data.(uint16)) },
		func(index byte, data any, p *Program) { p.register_s[index] = string(rune(data.(uint16))) },
		func(index byte, data any, p *Program) { p.register_rbi[index] = int8(data.(uint16)) },
		func(index byte, data any, p *Program) { p.register_rsi[index] = int16(data.(uint16)) },
		func(index byte, data any, p *Program) { p.register_rli[index] = int32(data.(uint16)) },
		func(index byte, data any, p *Program) { p.register_rlli[index] = int64(data.(uint16)) },
		func(index byte, data any, p *Program) { p.register_rbui[index] = uint8(data.(uint16)) },
		func(index byte, data any, p *Program) { p.register_rsui[index] = data.(uint16) },
		func(index byte, data any, p *Program) { p.register_rlui[index] = uint32(data.(uint16)) },
		func(index byte, data any, p *Program) { p.register_rllui[index] = uint64(data.(uint16)) },
		func(index byte, data any, p *Program) { p.register_rlf[index] = float32(data.(uint16)) },
		func(index byte, data any, p *Program) { p.register_rllf[index] = float64(data.(uint16)) },
		func(index byte, data any, p *Program) { p.register_rs[index] = string(rune(data.(uint16))) },
	},
	{
		func(index byte, data any, p *Program) { p.register_bi[index] = int8(data.(uint32)) },
		func(index byte, data any, p *Program) { p.register_si[index] = int16(data.(uint32)) },
		func(index byte, data any, p *Program) { p.register_li[index] = int32(data.(uint32)) },
		func(index byte, data any, p *Program) { p.register_lli[index] = int64(data.(uint32)) },
		func(index byte, data any, p *Program) { p.register_bui[index] = uint8(data.(uint32)) },
		func(index byte, data any, p *Program) { p.register_sui[index] = uint16(data.(uint32)) },
		func(index byte, data any, p *Program) { p.register_lui[index] = data.(uint32) },
		func(index byte, data any, p *Program) { p.register_llui[index] = uint64(data.(uint32)) },
		func(index byte, data any, p *Program) { p.register_lf[index] = float32(data.(uint32)) },
		func(index byte, data any, p *Program) { p.register_llf[index] = float64(data.(uint32)) },
		func(index byte, data any, p *Program) { p.register_s[index] = string(rune(data.(uint32))) },
		func(index byte, data any, p *Program) { p.register_rbi[index] = int8(data.(uint32)) },
		func(index byte, data any, p *Program) { p.register_rsi[index] = int16(data.(uint32)) },
		func(index byte, data any, p *Program) { p.register_rli[index] = int32(data.(uint32)) },
		func(index byte, data any, p *Program) { p.register_rlli[index] = int64(data.(uint32)) },
		func(index byte, data any, p *Program) { p.register_rbui[index] = uint8(data.(uint32)) },
		func(index byte, data any, p *Program) { p.register_rsui[index] = uint16(data.(uint32)) },
		func(index byte, data any, p *Program) { p.register_rlui[index] = data.(uint32) },
		func(index byte, data any, p *Program) { p.register_rllui[index] = uint64(data.(uint32)) },
		func(index byte, data any, p *Program) { p.register_rlf[index] = float32(data.(uint32)) },
		func(index byte, data any, p *Program) { p.register_rllf[index] = float64(data.(uint32)) },
		func(index byte, data any, p *Program) { p.register_rs[index] = string(rune(data.(uint32))) },
	},
	{
		func(index byte, data any, p *Program) { p.register_bi[index] = int8(data.(uint64)) },
		func(index byte, data any, p *Program) { p.register_si[index] = int16(data.(uint64)) },
		func(index byte, data any, p *Program) { p.register_li[index] = int32(data.(uint64)) },
		func(index byte, data any, p *Program) { p.register_lli[index] = int64(data.(uint64)) },
		func(index byte, data any, p *Program) { p.register_bui[index] = uint8(data.(uint64)) },
		func(index byte, data any, p *Program) { p.register_sui[index] = uint16(data.(uint64)) },
		func(index byte, data any, p *Program) { p.register_lui[index] = uint32(data.(uint64)) },
		func(index byte, data any, p *Program) { p.register_llui[index] = data.(uint64) },
		func(index byte, data any, p *Program) { p.register_lf[index] = float32(data.(uint64)) },
		func(index byte, data any, p *Program) { p.register_llf[index] = float64(data.(uint64)) },
		func(index byte, data any, p *Program) { p.register_s[index] = string(rune(data.(uint64))) },
		func(index byte, data any, p *Program) { p.register_rbi[index] = int8(data.(uint64)) },
		func(index byte, data any, p *Program) { p.register_rsi[index] = int16(data.(uint64)) },
		func(index byte, data any, p *Program) { p.register_rli[index] = int32(data.(uint64)) },
		func(index byte, data any, p *Program) { p.register_rlli[index] = int64(data.(uint64)) },
		func(index byte, data any, p *Program) { p.register_rbui[index] = uint8(data.(uint64)) },
		func(index byte, data any, p *Program) { p.register_rsui[index] = uint16(data.(uint64)) },
		func(index byte, data any, p *Program) { p.register_rlui[index] = uint32(data.(uint64)) },
		func(index byte, data any, p *Program) { p.register_rllui[index] = data.(uint64) },
		func(index byte, data any, p *Program) { p.register_rlf[index] = float32(data.(uint64)) },
		func(index byte, data any, p *Program) { p.register_rllf[index] = float64(data.(uint64)) },
		func(index byte, data any, p *Program) { p.register_rs[index] = string(rune(data.(uint64))) },
	},
	{
		func(index byte, data any, p *Program) { p.register_bi[index] = int8(data.(float32)) },
		func(index byte, data any, p *Program) { p.register_si[index] = int16(data.(float32)) },
		func(index byte, data any, p *Program) { p.register_li[index] = int32(data.(float32)) },
		func(index byte, data any, p *Program) { p.register_lli[index] = int64(data.(float32)) },
		func(index byte, data any, p *Program) { p.register_bui[index] = uint8(data.(float32)) },
		func(index byte, data any, p *Program) { p.register_sui[index] = uint16(data.(float32)) },
		func(index byte, data any, p *Program) { p.register_lui[index] = uint32(data.(float32)) },
		func(index byte, data any, p *Program) { p.register_llui[index] = uint64(data.(float32)) },
		func(index byte, data any, p *Program) { p.register_lf[index] = data.(float32) },
		func(index byte, data any, p *Program) { p.register_llf[index] = float64(data.(float32)) },
		func(_ byte, _ any, _ *Program) {},
		func(index byte, data any, p *Program) { p.register_rbi[index] = int8(data.(float32)) },
		func(index byte, data any, p *Program) { p.register_rsi[index] = int16(data.(float32)) },
		func(index byte, data any, p *Program) { p.register_rli[index] = int32(data.(float32)) },
		func(index byte, data any, p *Program) { p.register_rlli[index] = int64(data.(float32)) },
		func(index byte, data any, p *Program) { p.register_rbui[index] = uint8(data.(float32)) },
		func(index byte, data any, p *Program) { p.register_rsui[index] = uint16(data.(float32)) },
		func(index byte, data any, p *Program) { p.register_rlui[index] = uint32(data.(float32)) },
		func(index byte, data any, p *Program) { p.register_rllui[index] = uint64(data.(float32)) },
		func(index byte, data any, p *Program) { p.register_rlf[index] = data.(float32) },
		func(index byte, data any, p *Program) { p.register_rllf[index] = float64(data.(float32)) },
		func(_ byte, _ any, _ *Program) {},
	},
	{
		func(index byte, data any, p *Program) { p.register_bi[index] = int8(data.(float64)) },
		func(index byte, data any, p *Program) { p.register_si[index] = int16(data.(float64)) },
		func(index byte, data any, p *Program) { p.register_li[index] = int32(data.(float64)) },
		func(index byte, data any, p *Program) { p.register_lli[index] = int64(data.(float64)) },
		func(index byte, data any, p *Program) { p.register_bui[index] = uint8(data.(float64)) },
		func(index byte, data any, p *Program) { p.register_sui[index] = uint16(data.(float64)) },
		func(index byte, data any, p *Program) { p.register_lui[index] = uint32(data.(float64)) },
		func(index byte, data any, p *Program) { p.register_llui[index] = uint64(data.(float64)) },
		func(index byte, data any, p *Program) { p.register_lf[index] = float32(data.(float64)) },
		func(index byte, data any, p *Program) { p.register_llf[index] = data.(float64) },
		func(_ byte, _ any, _ *Program) {},
		func(index byte, data any, p *Program) { p.register_rbi[index] = int8(data.(float64)) },
		func(index byte, data any, p *Program) { p.register_rsi[index] = int16(data.(float64)) },
		func(index byte, data any, p *Program) { p.register_rli[index] = int32(data.(float64)) },
		func(index byte, data any, p *Program) { p.register_rlli[index] = int64(data.(float64)) },
		func(index byte, data any, p *Program) { p.register_rbui[index] = uint8(data.(float64)) },
		func(index byte, data any, p *Program) { p.register_rsui[index] = uint16(data.(float64)) },
		func(index byte, data any, p *Program) { p.register_rlui[index] = uint32(data.(float64)) },
		func(index byte, data any, p *Program) { p.register_rllui[index] = uint64(data.(float64)) },
		func(index byte, data any, p *Program) { p.register_rlf[index] = float32(data.(float64)) },
		func(index byte, data any, p *Program) { p.register_rllf[index] = data.(float64) },
		func(_ byte, _ any, _ *Program) {},
	},
	{
		func(index byte, data any, p *Program) { p.register_bi[index] = int8(data.(string)[0]) },
		func(index byte, data any, p *Program) { p.register_si[index] = int16(data.(string)[0]) },
		func(index byte, data any, p *Program) { p.register_li[index] = int32(data.(string)[0]) },
		func(index byte, data any, p *Program) { p.register_lli[index] = int64(data.(string)[0]) },
		func(index byte, data any, p *Program) { p.register_bui[index] = uint8(data.(string)[0]) },
		func(index byte, data any, p *Program) { p.register_sui[index] = uint16(data.(string)[0]) },
		func(index byte, data any, p *Program) { p.register_lui[index] = uint32(data.(string)[0]) },
		func(index byte, data any, p *Program) { p.register_llui[index] = uint64(data.(string)[0]) },
		func(index byte, data any, p *Program) { p.register_lf[index] = float32(data.(string)[0]) },
		func(index byte, data any, p *Program) { p.register_llf[index] = float64(data.(string)[0]) },
		func(index byte, data any, p *Program) { p.register_s[index] = data.(string) },
		func(index byte, data any, p *Program) { p.register_rbi[index] = int8(data.(string)[0]) },
		func(index byte, data any, p *Program) { p.register_rsi[index] = int16(data.(string)[0]) },
		func(index byte, data any, p *Program) { p.register_rli[index] = int32(data.(string)[0]) },
		func(index byte, data any, p *Program) { p.register_rlli[index] = int64(data.(string)[0]) },
		func(index byte, data any, p *Program) { p.register_rbui[index] = uint8(data.(string)[0]) },
		func(index byte, data any, p *Program) { p.register_rsui[index] = uint16(data.(string)[0]) },
		func(index byte, data any, p *Program) { p.register_rlui[index] = uint32(data.(string)[0]) },
		func(index byte, data any, p *Program) { p.register_rllui[index] = uint64(data.(string)[0]) },
		func(index byte, data any, p *Program) { p.register_rlf[index] = float32(data.(string)[0]) },
		func(index byte, data any, p *Program) { p.register_rllf[index] = float64(data.(string)[0]) },
		func(index byte, data any, p *Program) { p.register_rs[index] = data.(string) },
	},
}
