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

var register_bi = [256]int8{}
var register_si = [256]int16{}
var register_li = [256]int32{}
var register_lli = [256]int64{}
var register_bui = [256]uint8{}
var register_sui = [256]uint16{}
var register_lui = [256]uint32{}
var register_llui = [256]uint64{}
var register_lf = [256]float32{}
var register_llf = [256]float64{}
var register_s = [256]string{}
var register_rbi = [256]int8{}
var register_rsi = [256]int16{}
var register_rli = [256]int32{}
var register_rlli = [256]int64{}
var register_rbui = [256]uint8{}
var register_rsui = [256]uint16{}
var register_rlui = [256]uint32{}
var register_rllui = [256]uint64{}
var register_rlf = [256]float32{}
var register_rllf = [256]float64{}
var register_rs = [256]string{}

var register_set = [22]func(byte, any){
	func(index byte, data any) { register_bi[index] = data.(int8) },
	func(index byte, data any) { register_si[index] = data.(int16) },
	func(index byte, data any) { register_li[index] = data.(int32) },
	func(index byte, data any) { register_lli[index] = data.(int64) },
	func(index byte, data any) { register_bui[index] = data.(uint8) },
	func(index byte, data any) { register_sui[index] = data.(uint16) },
	func(index byte, data any) { register_lui[index] = data.(uint32) },
	func(index byte, data any) { register_llui[index] = data.(uint64) },
	func(index byte, data any) { register_lf[index] = data.(float32) },
	func(index byte, data any) { register_llf[index] = data.(float64) },
	func(index byte, data any) { register_s[index] = data.(string) },
	func(index byte, data any) { register_rbi[index] = data.(int8) },
	func(index byte, data any) { register_rsi[index] = data.(int16) },
	func(index byte, data any) { register_rli[index] = data.(int32) },
	func(index byte, data any) { register_rlli[index] = data.(int64) },
	func(index byte, data any) { register_rbui[index] = data.(uint8) },
	func(index byte, data any) { register_rsui[index] = data.(uint16) },
	func(index byte, data any) { register_rlui[index] = data.(uint32) },
	func(index byte, data any) { register_rllui[index] = data.(uint64) },
	func(index byte, data any) { register_rlf[index] = data.(float32) },
	func(index byte, data any) { register_rllf[index] = data.(float64) },
	func(index byte, data any) { register_rs[index] = data.(string) },
}

var register_get = [22]func(int) any{
	func(index int) any { return register_bi[index] },
	func(index int) any { return register_si[index] },
	func(index int) any { return register_li[index] },
	func(index int) any { return register_lli[index] },
	func(index int) any { return register_bui[index] },
	func(index int) any { return register_sui[index] },
	func(index int) any { return register_lui[index] },
	func(index int) any { return register_llui[index] },
	func(index int) any { return register_lf[index] },
	func(index int) any { return register_llf[index] },
	func(index int) any { return register_s[index] },
	func(index int) any { return register_rbi[index] },
	func(index int) any { return register_rsi[index] },
	func(index int) any { return register_rli[index] },
	func(index int) any { return register_rlli[index] },
	func(index int) any { return register_rbui[index] },
	func(index int) any { return register_rsui[index] },
	func(index int) any { return register_rlui[index] },
	func(index int) any { return register_rllui[index] },
	func(index int) any { return register_rlf[index] },
	func(index int) any { return register_rllf[index] },
	func(index int) any { return register_rs[index] },
}

var register_force_set = [13][22]func(byte, any){
	{
		func(index byte, data any) { register_bi[index] = data.(int8) },
		func(index byte, data any) { register_si[index] = int16(data.(int8)) },
		func(index byte, data any) { register_li[index] = int32(data.(int8)) },
		func(index byte, data any) { register_lli[index] = int64(data.(int8)) },
		func(index byte, data any) { register_bui[index] = uint8(data.(int8)) },
		func(index byte, data any) { register_sui[index] = uint16(data.(int8)) },
		func(index byte, data any) { register_lui[index] = uint32(data.(int8)) },
		func(index byte, data any) { register_llui[index] = uint64(data.(int8)) },
		func(index byte, data any) { register_lf[index] = float32(data.(int8)) },
		func(index byte, data any) { register_llf[index] = float64(data.(int8)) },
		func(index byte, data any) { register_s[index] = string(rune(data.(int8))) },
		func(index byte, data any) { register_rbi[index] = data.(int8) },
		func(index byte, data any) { register_rsi[index] = int16(data.(int8)) },
		func(index byte, data any) { register_rli[index] = int32(data.(int8)) },
		func(index byte, data any) { register_rlli[index] = int64(data.(int8)) },
		func(index byte, data any) { register_rbui[index] = uint8(data.(int8)) },
		func(index byte, data any) { register_rsui[index] = uint16(data.(int8)) },
		func(index byte, data any) { register_rlui[index] = uint32(data.(int8)) },
		func(index byte, data any) { register_rllui[index] = uint64(data.(int8)) },
		func(index byte, data any) { register_rlf[index] = float32(data.(int8)) },
		func(index byte, data any) { register_rllf[index] = float64(data.(int8)) },
		func(index byte, data any) { register_rs[index] = string(rune(data.(int8))) },
	},
	{
		func(index byte, data any) { register_bi[index] = int8(data.(int16)) },
		func(index byte, data any) { register_si[index] = data.(int16) },
		func(index byte, data any) { register_li[index] = int32(data.(int16)) },
		func(index byte, data any) { register_lli[index] = int64(data.(int16)) },
		func(index byte, data any) { register_bui[index] = uint8(data.(int16)) },
		func(index byte, data any) { register_sui[index] = uint16(data.(int16)) },
		func(index byte, data any) { register_lui[index] = uint32(data.(int16)) },
		func(index byte, data any) { register_llui[index] = uint64(data.(int16)) },
		func(index byte, data any) { register_lf[index] = float32(data.(int16)) },
		func(index byte, data any) { register_llf[index] = float64(data.(int16)) },
		func(index byte, data any) { register_s[index] = string(rune(data.(int16))) },
		func(index byte, data any) { register_rbi[index] = int8(data.(int16)) },
		func(index byte, data any) { register_rsi[index] = data.(int16) },
		func(index byte, data any) { register_rli[index] = int32(data.(int16)) },
		func(index byte, data any) { register_rlli[index] = int64(data.(int16)) },
		func(index byte, data any) { register_rbui[index] = uint8(data.(int16)) },
		func(index byte, data any) { register_rsui[index] = uint16(data.(int16)) },
		func(index byte, data any) { register_rlui[index] = uint32(data.(int16)) },
		func(index byte, data any) { register_rllui[index] = uint64(data.(int16)) },
		func(index byte, data any) { register_rlf[index] = float32(data.(int16)) },
		func(index byte, data any) { register_rllf[index] = float64(data.(int16)) },
		func(index byte, data any) { register_rs[index] = string(rune(data.(int16))) },
	},
	{
		func(index byte, data any) { register_bi[index] = int8(data.(int32)) },
		func(index byte, data any) { register_si[index] = int16(data.(int32)) },
		func(index byte, data any) { register_li[index] = data.(int32) },
		func(index byte, data any) { register_lli[index] = int64(data.(int32)) },
		func(index byte, data any) { register_bui[index] = uint8(data.(int32)) },
		func(index byte, data any) { register_sui[index] = uint16(data.(int32)) },
		func(index byte, data any) { register_lui[index] = uint32(data.(int32)) },
		func(index byte, data any) { register_llui[index] = uint64(data.(int32)) },
		func(index byte, data any) { register_lf[index] = float32(data.(int32)) },
		func(index byte, data any) { register_llf[index] = float64(data.(int32)) },
		func(index byte, data any) { register_s[index] = string(rune(data.(int32))) },
		func(index byte, data any) { register_rbi[index] = int8(data.(int32)) },
		func(index byte, data any) { register_rsi[index] = int16(data.(int32)) },
		func(index byte, data any) { register_rli[index] = data.(int32) },
		func(index byte, data any) { register_rlli[index] = int64(data.(int32)) },
		func(index byte, data any) { register_rbui[index] = uint8(data.(int32)) },
		func(index byte, data any) { register_rsui[index] = uint16(data.(int32)) },
		func(index byte, data any) { register_rlui[index] = uint32(data.(int32)) },
		func(index byte, data any) { register_rllui[index] = uint64(data.(int32)) },
		func(index byte, data any) { register_rlf[index] = float32(data.(int32)) },
		func(index byte, data any) { register_rllf[index] = float64(data.(int32)) },
		func(index byte, data any) { register_rs[index] = string(rune(data.(int32))) },
	},
	{
		func(index byte, data any) { register_bi[index] = int8(data.(int64)) },
		func(index byte, data any) { register_si[index] = int16(data.(int64)) },
		func(index byte, data any) { register_li[index] = int32(data.(int64)) },
		func(index byte, data any) { register_lli[index] = data.(int64) },
		func(index byte, data any) { register_bui[index] = uint8(data.(int64)) },
		func(index byte, data any) { register_sui[index] = uint16(data.(int64)) },
		func(index byte, data any) { register_lui[index] = uint32(data.(int64)) },
		func(index byte, data any) { register_llui[index] = uint64(data.(int64)) },
		func(index byte, data any) { register_lf[index] = float32(data.(int64)) },
		func(index byte, data any) { register_llf[index] = float64(data.(int64)) },
		func(index byte, data any) { register_s[index] = string(rune(data.(int64))) },
		func(index byte, data any) { register_rbi[index] = int8(data.(int64)) },
		func(index byte, data any) { register_rsi[index] = int16(data.(int64)) },
		func(index byte, data any) { register_rli[index] = int32(data.(int64)) },
		func(index byte, data any) { register_rlli[index] = data.(int64) },
		func(index byte, data any) { register_rbui[index] = uint8(data.(int64)) },
		func(index byte, data any) { register_rsui[index] = uint16(data.(int64)) },
		func(index byte, data any) { register_rlui[index] = uint32(data.(int64)) },
		func(index byte, data any) { register_rllui[index] = uint64(data.(int64)) },
		func(index byte, data any) { register_rlf[index] = float32(data.(int64)) },
		func(index byte, data any) { register_rllf[index] = float64(data.(int64)) },
		func(index byte, data any) { register_rs[index] = string(rune(data.(int64))) },
	},
	{
		func(index byte, data any) { register_bi[index] = int8(data.(uint8)) },
		func(index byte, data any) { register_si[index] = int16(data.(uint8)) },
		func(index byte, data any) { register_li[index] = int32(data.(uint8)) },
		func(index byte, data any) { register_lli[index] = int64(data.(uint8)) },
		func(index byte, data any) { register_bui[index] = data.(uint8) },
		func(index byte, data any) { register_sui[index] = uint16(data.(uint8)) },
		func(index byte, data any) { register_lui[index] = uint32(data.(uint8)) },
		func(index byte, data any) { register_llui[index] = uint64(data.(uint8)) },
		func(index byte, data any) { register_lf[index] = float32(data.(uint8)) },
		func(index byte, data any) { register_llf[index] = float64(data.(uint8)) },
		func(index byte, data any) { register_s[index] = string(rune(data.(uint8))) },
		func(index byte, data any) { register_rbi[index] = int8(data.(uint8)) },
		func(index byte, data any) { register_rsi[index] = int16(data.(uint8)) },
		func(index byte, data any) { register_rli[index] = int32(data.(uint8)) },
		func(index byte, data any) { register_rlli[index] = int64(data.(uint8)) },
		func(index byte, data any) { register_rbui[index] = data.(uint8) },
		func(index byte, data any) { register_rsui[index] = uint16(data.(uint8)) },
		func(index byte, data any) { register_rlui[index] = uint32(data.(uint8)) },
		func(index byte, data any) { register_rllui[index] = uint64(data.(uint8)) },
		func(index byte, data any) { register_rlf[index] = float32(data.(uint8)) },
		func(index byte, data any) { register_rllf[index] = float64(data.(uint8)) },
		func(index byte, data any) { register_rs[index] = string(rune(data.(uint8))) },
	},
	{
		func(index byte, data any) { register_bi[index] = int8(data.(uint16)) },
		func(index byte, data any) { register_si[index] = int16(data.(uint16)) },
		func(index byte, data any) { register_li[index] = int32(data.(uint16)) },
		func(index byte, data any) { register_lli[index] = int64(data.(uint16)) },
		func(index byte, data any) { register_bui[index] = uint8(data.(uint16)) },
		func(index byte, data any) { register_sui[index] = data.(uint16) },
		func(index byte, data any) { register_lui[index] = uint32(data.(uint16)) },
		func(index byte, data any) { register_llui[index] = uint64(data.(uint16)) },
		func(index byte, data any) { register_lf[index] = float32(data.(uint16)) },
		func(index byte, data any) { register_llf[index] = float64(data.(uint16)) },
		func(index byte, data any) { register_s[index] = string(rune(data.(uint16))) },
		func(index byte, data any) { register_rbi[index] = int8(data.(uint16)) },
		func(index byte, data any) { register_rsi[index] = int16(data.(uint16)) },
		func(index byte, data any) { register_rli[index] = int32(data.(uint16)) },
		func(index byte, data any) { register_rlli[index] = int64(data.(uint16)) },
		func(index byte, data any) { register_rbui[index] = uint8(data.(uint16)) },
		func(index byte, data any) { register_rsui[index] = data.(uint16) },
		func(index byte, data any) { register_rlui[index] = uint32(data.(uint16)) },
		func(index byte, data any) { register_rllui[index] = uint64(data.(uint16)) },
		func(index byte, data any) { register_rlf[index] = float32(data.(uint16)) },
		func(index byte, data any) { register_rllf[index] = float64(data.(uint16)) },
		func(index byte, data any) { register_rs[index] = string(rune(data.(uint16))) },
	},
	{
		func(index byte, data any) { register_bi[index] = int8(data.(uint32)) },
		func(index byte, data any) { register_si[index] = int16(data.(uint32)) },
		func(index byte, data any) { register_li[index] = int32(data.(uint32)) },
		func(index byte, data any) { register_lli[index] = int64(data.(uint32)) },
		func(index byte, data any) { register_bui[index] = uint8(data.(uint32)) },
		func(index byte, data any) { register_sui[index] = uint16(data.(uint32)) },
		func(index byte, data any) { register_lui[index] = data.(uint32) },
		func(index byte, data any) { register_llui[index] = uint64(data.(uint32)) },
		func(index byte, data any) { register_lf[index] = float32(data.(uint32)) },
		func(index byte, data any) { register_llf[index] = float64(data.(uint32)) },
		func(index byte, data any) { register_s[index] = string(rune(data.(uint32))) },
		func(index byte, data any) { register_rbi[index] = int8(data.(uint32)) },
		func(index byte, data any) { register_rsi[index] = int16(data.(uint32)) },
		func(index byte, data any) { register_rli[index] = int32(data.(uint32)) },
		func(index byte, data any) { register_rlli[index] = int64(data.(uint32)) },
		func(index byte, data any) { register_rbui[index] = uint8(data.(uint32)) },
		func(index byte, data any) { register_rsui[index] = uint16(data.(uint32)) },
		func(index byte, data any) { register_rlui[index] = data.(uint32) },
		func(index byte, data any) { register_rllui[index] = uint64(data.(uint32)) },
		func(index byte, data any) { register_rlf[index] = float32(data.(uint32)) },
		func(index byte, data any) { register_rllf[index] = float64(data.(uint32)) },
		func(index byte, data any) { register_rs[index] = string(rune(data.(uint32))) },
	},
	{
		func(index byte, data any) { register_bi[index] = int8(data.(uint64)) },
		func(index byte, data any) { register_si[index] = int16(data.(uint64)) },
		func(index byte, data any) { register_li[index] = int32(data.(uint64)) },
		func(index byte, data any) { register_lli[index] = int64(data.(uint64)) },
		func(index byte, data any) { register_bui[index] = uint8(data.(uint64)) },
		func(index byte, data any) { register_sui[index] = uint16(data.(uint64)) },
		func(index byte, data any) { register_lui[index] = uint32(data.(uint64)) },
		func(index byte, data any) { register_llui[index] = data.(uint64) },
		func(index byte, data any) { register_lf[index] = float32(data.(uint64)) },
		func(index byte, data any) { register_llf[index] = float64(data.(uint64)) },
		func(index byte, data any) { register_s[index] = string(rune(data.(uint64))) },
		func(index byte, data any) { register_rbi[index] = int8(data.(uint64)) },
		func(index byte, data any) { register_rsi[index] = int16(data.(uint64)) },
		func(index byte, data any) { register_rli[index] = int32(data.(uint64)) },
		func(index byte, data any) { register_rlli[index] = int64(data.(uint64)) },
		func(index byte, data any) { register_rbui[index] = uint8(data.(uint64)) },
		func(index byte, data any) { register_rsui[index] = uint16(data.(uint64)) },
		func(index byte, data any) { register_rlui[index] = uint32(data.(uint64)) },
		func(index byte, data any) { register_rllui[index] = data.(uint64) },
		func(index byte, data any) { register_rlf[index] = float32(data.(uint64)) },
		func(index byte, data any) { register_rllf[index] = float64(data.(uint64)) },
		func(index byte, data any) { register_rs[index] = string(rune(data.(uint64))) },
	},
	{
		func(index byte, data any) { register_bi[index] = int8(data.(float32)) },
		func(index byte, data any) { register_si[index] = int16(data.(float32)) },
		func(index byte, data any) { register_li[index] = int32(data.(float32)) },
		func(index byte, data any) { register_lli[index] = int64(data.(float32)) },
		func(index byte, data any) { register_bui[index] = uint8(data.(float32)) },
		func(index byte, data any) { register_sui[index] = uint16(data.(float32)) },
		func(index byte, data any) { register_lui[index] = uint32(data.(float32)) },
		func(index byte, data any) { register_llui[index] = uint64(data.(float32)) },
		func(index byte, data any) { register_lf[index] = data.(float32) },
		func(index byte, data any) { register_llf[index] = float64(data.(float32)) },
		func(_ byte, _ any) {},
		func(index byte, data any) { register_rbi[index] = int8(data.(float32)) },
		func(index byte, data any) { register_rsi[index] = int16(data.(float32)) },
		func(index byte, data any) { register_rli[index] = int32(data.(float32)) },
		func(index byte, data any) { register_rlli[index] = int64(data.(float32)) },
		func(index byte, data any) { register_rbui[index] = uint8(data.(float32)) },
		func(index byte, data any) { register_rsui[index] = uint16(data.(float32)) },
		func(index byte, data any) { register_rlui[index] = uint32(data.(float32)) },
		func(index byte, data any) { register_rllui[index] = uint64(data.(float32)) },
		func(index byte, data any) { register_rlf[index] = data.(float32) },
		func(index byte, data any) { register_rllf[index] = float64(data.(float32)) },
		func(_ byte, _ any) {},
	},
	{
		func(index byte, data any) { register_bi[index] = int8(data.(float64)) },
		func(index byte, data any) { register_si[index] = int16(data.(float64)) },
		func(index byte, data any) { register_li[index] = int32(data.(float64)) },
		func(index byte, data any) { register_lli[index] = int64(data.(float64)) },
		func(index byte, data any) { register_bui[index] = uint8(data.(float64)) },
		func(index byte, data any) { register_sui[index] = uint16(data.(float64)) },
		func(index byte, data any) { register_lui[index] = uint32(data.(float64)) },
		func(index byte, data any) { register_llui[index] = uint64(data.(float64)) },
		func(index byte, data any) { register_lf[index] = float32(data.(float64)) },
		func(index byte, data any) { register_llf[index] = data.(float64) },
		func(_ byte, _ any) {},
		func(index byte, data any) { register_rbi[index] = int8(data.(float64)) },
		func(index byte, data any) { register_rsi[index] = int16(data.(float64)) },
		func(index byte, data any) { register_rli[index] = int32(data.(float64)) },
		func(index byte, data any) { register_rlli[index] = int64(data.(float64)) },
		func(index byte, data any) { register_rbui[index] = uint8(data.(float64)) },
		func(index byte, data any) { register_rsui[index] = uint16(data.(float64)) },
		func(index byte, data any) { register_rlui[index] = uint32(data.(float64)) },
		func(index byte, data any) { register_rllui[index] = uint64(data.(float64)) },
		func(index byte, data any) { register_rlf[index] = float32(data.(float64)) },
		func(index byte, data any) { register_rllf[index] = data.(float64) },
		func(_ byte, _ any) {},
	},
	{
		func(index byte, data any) { register_bi[index] = int8(data.(string)[0]) },
		func(index byte, data any) { register_si[index] = int16(data.(string)[0]) },
		func(index byte, data any) { register_li[index] = int32(data.(string)[0]) },
		func(index byte, data any) { register_lli[index] = int64(data.(string)[0]) },
		func(index byte, data any) { register_bui[index] = uint8(data.(string)[0]) },
		func(index byte, data any) { register_sui[index] = uint16(data.(string)[0]) },
		func(index byte, data any) { register_lui[index] = uint32(data.(string)[0]) },
		func(index byte, data any) { register_llui[index] = uint64(data.(string)[0]) },
		func(index byte, data any) { register_lf[index] = float32(data.(string)[0]) },
		func(index byte, data any) { register_llf[index] = float64(data.(string)[0]) },
		func(index byte, data any) { register_s[index] = data.(string) },
		func(index byte, data any) { register_rbi[index] = int8(data.(string)[0]) },
		func(index byte, data any) { register_rsi[index] = int16(data.(string)[0]) },
		func(index byte, data any) { register_rli[index] = int32(data.(string)[0]) },
		func(index byte, data any) { register_rlli[index] = int64(data.(string)[0]) },
		func(index byte, data any) { register_rbui[index] = uint8(data.(string)[0]) },
		func(index byte, data any) { register_rsui[index] = uint16(data.(string)[0]) },
		func(index byte, data any) { register_rlui[index] = uint32(data.(string)[0]) },
		func(index byte, data any) { register_rllui[index] = uint64(data.(string)[0]) },
		func(index byte, data any) { register_rlf[index] = float32(data.(string)[0]) },
		func(index byte, data any) { register_rllf[index] = float64(data.(string)[0]) },
		func(index byte, data any) { register_rs[index] = data.(string) },
	},
	{},
	{
		func(index byte, data any) {
			register_bi[index] = convert_to_value[data.([]byte)[0]](data.([]byte)).(int8)
		},
		func(index byte, data any) {
			register_si[index] = convert_to_value[data.([]byte)[0]](data.([]byte)).(int16)
		},
		func(index byte, data any) {
			register_li[index] = convert_to_value[data.([]byte)[0]](data.([]byte)).(int32)
		},
		func(index byte, data any) {
			register_lli[index] = convert_to_value[data.([]byte)[0]](data.([]byte)).(int64)
		},
		func(index byte, data any) {
			register_bui[index] = convert_to_value[data.([]byte)[0]](data.([]byte)).(uint8)
		},
		func(index byte, data any) {
			register_sui[index] = convert_to_value[data.([]byte)[0]](data.([]byte)).(uint16)
		},
		func(index byte, data any) {
			register_lui[index] = convert_to_value[data.([]byte)[0]](data.([]byte)).(uint32)
		},
		func(index byte, data any) {
			register_llui[index] = convert_to_value[data.([]byte)[0]](data.([]byte)).(uint64)
		},
		func(index byte, data any) {
			register_lf[index] = convert_to_value[data.([]byte)[0]](data.([]byte)).(float32)
		},
		func(index byte, data any) {
			register_llf[index] = convert_to_value[data.([]byte)[0]](data.([]byte)).(float64)
		},
		func(index byte, data any) {
			register_s[index] = convert_to_value[data.([]byte)[0]](data.([]byte)).(string)
		},
		func(index byte, data any) {
			register_rbi[index] = convert_to_value[data.([]byte)[0]](data.([]byte)).(int8)
		},
		func(index byte, data any) {
			register_rsi[index] = convert_to_value[data.([]byte)[0]](data.([]byte)).(int16)
		},
		func(index byte, data any) {
			register_rli[index] = convert_to_value[data.([]byte)[0]](data.([]byte)).(int32)
		},
		func(index byte, data any) {
			register_rlli[index] = convert_to_value[data.([]byte)[0]](data.([]byte)).(int64)
		},
		func(index byte, data any) {
			register_rbui[index] = convert_to_value[data.([]byte)[0]](data.([]byte)).(uint8)
		},
		func(index byte, data any) {
			register_rsui[index] = convert_to_value[data.([]byte)[0]](data.([]byte)).(uint16)
		},
		func(index byte, data any) {
			register_rlui[index] = convert_to_value[data.([]byte)[0]](data.([]byte)).(uint32)
		},
		func(index byte, data any) {
			register_rllui[index] = convert_to_value[data.([]byte)[0]](data.([]byte)).(uint64)
		},
		func(index byte, data any) {
			register_rlf[index] = convert_to_value[data.([]byte)[0]](data.([]byte)).(float32)
		},
		func(index byte, data any) {
			register_rllf[index] = convert_to_value[data.([]byte)[0]](data.([]byte)).(float64)
		},
		func(index byte, data any) {
			register_rs[index] = convert_to_value[data.([]byte)[0]](data.([]byte)).(string)
		},
	},
}
