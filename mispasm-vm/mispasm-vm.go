package main

type Program struct {
	should_close   bool
	data           []byte
	global         string
	funcs          map[string]function
	constants      []constant
	variables      []variable
	register_bi    [256]int8
	register_si    [256]int16
	register_li    [256]int32
	register_lli   [256]int64
	register_bui   [256]uint8
	register_sui   [256]uint16
	register_lui   [256]uint32
	register_llui  [256]uint64
	register_lf    [256]float32
	register_llf   [256]float64
	register_s     [256]string
	register_rbi   [256]int8
	register_rsi   [256]int16
	register_rli   [256]int32
	register_rlli  [256]int64
	register_rbui  [256]uint8
	register_rsui  [256]uint16
	register_rlui  [256]uint32
	register_rllui [256]uint64
	register_rlf   [256]float32
	register_rllf  [256]float64
	register_rs    [256]string
	register_cmp   [2][]byte
	stack          []stack
	instructions   [instruction_count]func([]byte, []byte, *function, *int, *int)
	calls          [1]func([]byte, []byte)
}

func NewProgram(data []byte) Program {
	return Program{data: data}
}

func (program *Program) Run() {
	program.load_program()
	program.init_instructions()
	program.init_calls()

	run_function(program.instructions, program.funcs[program.global], &program.should_close)
}
