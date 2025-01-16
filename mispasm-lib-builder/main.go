package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <destination> <name>\n", os.Args[0])
		os.Exit(1)
	} else if strings.ToLower(os.Args[2]) == "lib" {
		fmt.Printf("Name cannot be '%s'\n", os.Args[2])
		os.Exit(1)
	}

	dest := os.Args[1]
	name := os.Args[2]

	os.Mkdir(dest+"/lib", 0755)
	os.Mkdir(dest+"/test", 0755)
	os.WriteFile(dest+"/build.bat", []byte("cd lib\ngo mod init "+name+"-lib\ngo build -buildmode=c-shared -o "+name+".misplib\ndel "+name+".h\n"), 0644)
	os.WriteFile(dest+"/lib/lib.go", []byte("package main\n/*\n#cgo CFLAGS: -I.\ntypedef struct {\n\tvoid* data;\n\tint val_len;\n\tint data_len;\n} response;\n*/\nimport \"C\"\nimport (\n\t\"encoding/binary\"\n\t\"math\"\n\t\"unsafe\"\n)\nconst (\n\tregister_bi byte = iota\n\tregister_si\n\tregister_li\n\tregister_lli\n\tregister_bui\n\tregister_sui\n\tregister_lui\n\tregister_llui\n\tregister_lf\n\tregister_llf\n\tregister_s\n\tregister_rbi\n\tregister_rsi\n\tregister_rli\n\tregister_rlli\n\tregister_rbui\n\tregister_rsui\n\tregister_rlui\n\tregister_rllui\n\tregister_rlf\n\tregister_rllf\n\tregister_rs\n)\ntype stack struct {\n\tname  byte\n\tindex byte\n\tdata  any\n}\n//export Call"+name+"\nfunc Call"+name+"(call *C.char, data unsafe.Pointer, val_len C.int, data_len C.int) C.response {\n\tstack_arr := decode_stack(data, val_len, data_len)\n\tres := calls[C.GoString(call)](stack_arr)\n\treturn encode_stack(res)\n}\nfunc stack_push(s *[]stack, reg byte, index byte, data any) {\n\t*s = append(*s, stack{reg, index, data})\n}\nfunc stack_pop(s *[]stack, reg byte, index byte) {\n\tfor i := range (*s) {\n\t\tif (*s)[i].name == reg && (*s)[i].index == index {\n\t\t\t*s = append((*s)[:i], (*s)[i+1:]...)\n\t\t\tbreak\n\t\t}\n\t}\n}\nfunc decode_stack(c_data unsafe.Pointer, value_length C.int, length C.int) []stack {\n\tdata := C.GoBytes(c_data, length)\n\tname_arr := data[:value_length]\n\tindex_arr := data[value_length : value_length*2]\n\tdata_arr := []any{}\n\toffset := int(value_length * 2)\n\tfor i := 0; i < int(value_length); i++ {\n\t\tsize := get_size(data[offset:])\n\t\tdata_arr = append(data_arr, convert_to_value[data[offset]](data[offset:offset+size+1]))\n\t\toffset += size + 1\n\t}\n\tstack_arr := []stack{}\n\tfor i := 0; i < int(value_length); i++ {\n\t\tstack_arr = append(stack_arr, stack{name_arr[i], index_arr[i], data_arr[i]})\n\t}\n\treturn stack_arr\n}\nfunc encode_stack(s []stack) C.response {\n\tdata := []byte{}\n\tvalue_length := C.int(0)\n\tfor i := 0; i < len(s); i++ {\n\t\tdata = append(data, s[i].name)\n\t\tvalue_length++\n\t}\n\tfor i := 0; i < len(s); i++ {\n\t\tdata = append(data, s[i].index)\n\t}\n\tfor i := 0; i < len(s); i++ {\n\t\tdata = append(data, convert_to_bytes[s[i].name%11](s[i].name%11, s[i].data)...)\n\t}\n\treturn C.response{C.CBytes(data), C.int(value_length), C.int(len(data))}\n}\nfunc get_size(data []byte) int {\n\tif data[0] <= 9 {\n\t\treturn type_sizes[data[0]]\n\t} else {\n\t\tsize := 0\n\t\tfor i := 0; i < len(data); i++ {\n\t\t\tif data[i] == 0 {\n\t\t\t\tbreak\n\t\t\t}\n\t\t\tsize++\n\t\t}\n\t\treturn size\n\t}\n}\nvar type_sizes = [11]int{1, 2, 4, 8, 1, 2, 4, 8, 4, 8, 0}\nvar convert_to_value = [11]func([]byte) any{\n\tfunc(data []byte) any { return byte_to_int8(data[1]) },\n\tfunc(data []byte) any { return bytes_to_int16(data[1:]) },\n\tfunc(data []byte) any { return bytes_to_int32(data[1:]) },\n\tfunc(data []byte) any { return bytes_to_int64(data[1:]) },\n\tfunc(data []byte) any { return byte_to_uint8(data[1]) },\n\tfunc(data []byte) any { return bytes_to_uint16(data[1:]) },\n\tfunc(data []byte) any { return bytes_to_uint32(data[1:]) },\n\tfunc(data []byte) any { return bytes_to_uint64(data[1:]) },\n\tfunc(data []byte) any { return bytes_to_float32(data[1:]) },\n\tfunc(data []byte) any { return bytes_to_float64(data[1:]) },\n\tfunc(data []byte) any { return bytes_to_string(data[1:]) },\n}\nfunc byte_to_int8(b byte) int8 {\n\treturn int8(b)\n}\nfunc bytes_to_int16(b []byte) int16 {\n\treturn int16(b[0])<<8 | int16(b[1])\n}\nfunc bytes_to_int32(b []byte) int32 {\n\treturn int32(b[0])<<24 | int32(b[1])<<16 | int32(b[2])<<8 | int32(b[3])\n}\nfunc bytes_to_int64(b []byte) int64 {\n\treturn int64(b[0])<<56 | int64(b[1])<<48 | int64(b[2])<<40 | int64(b[3])<<32 | int64(b[4])<<24 | int64(b[5])<<16 | int64(b[6])<<8 | int64(b[7])\n}\nfunc byte_to_uint8(b byte) uint8 {\n\treturn uint8(b)\n}\nfunc bytes_to_uint16(b []byte) uint16 {\n\treturn uint16(b[0])<<8 | uint16(b[1])\n}\nfunc bytes_to_uint32(b []byte) uint32 {\n\treturn uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])\n}\nfunc bytes_to_uint64(b []byte) uint64 {\n\treturn uint64(b[0])<<56 | uint64(b[1])<<48 | uint64(b[2])<<40 | uint64(b[3])<<32 | uint64(b[4])<<24 | uint64(b[5])<<16 | uint64(b[6])<<8 | uint64(b[7])\n}\nfunc bytes_to_float32(bytes []byte) float32 {\n\treturn math.Float32frombits(binary.LittleEndian.Uint32(bytes))\n}\nfunc bytes_to_float64(bytes []byte) float64 {\n\treturn math.Float64frombits(binary.LittleEndian.Uint64(bytes))\n}\nfunc bytes_to_string(b []byte) string {\n\treturn string(b[:len(b)-1])\n}\nvar convert_to_bytes = [11]func(byte, any) []byte{\n\tfunc(t byte, data any) []byte { return []byte{t, int8_to_byte(data.(int8))} },\n\tfunc(t byte, data any) []byte { return append([]byte{t}, int16_to_bytes(data.(int16))...) },\n\tfunc(t byte, data any) []byte { return append([]byte{t}, int32_to_bytes(data.(int32))...) },\n\tfunc(t byte, data any) []byte { return append([]byte{t}, int64_to_bytes(data.(int64))...) },\n\tfunc(t byte, data any) []byte { return []byte{t, uint8_to_byte(data.(uint8))} },\n\tfunc(t byte, data any) []byte { return append([]byte{t}, uint16_to_bytes(data.(uint16))...) },\n\tfunc(t byte, data any) []byte { return append([]byte{t}, uint32_to_bytes(data.(uint32))...) },\n\tfunc(t byte, data any) []byte { return append([]byte{t}, uint64_to_bytes(data.(uint64))...) },\n\tfunc(t byte, data any) []byte { return append([]byte{t}, float32_to_bytes(data.(float32))...) },\n\tfunc(t byte, data any) []byte { return append([]byte{t}, float64_to_bytes(data.(float64))...) },\n\tfunc(t byte, data any) []byte { return append([]byte{t}, string_to_bytes(data.(string))...) },\n}\nfunc int8_to_byte(i int8) byte {\n\treturn byte(i)\n}\nfunc int16_to_bytes(i int16) []byte {\n\treturn []byte{byte(i >> 8), byte(i)}\n}\nfunc int32_to_bytes(i int32) []byte {\n\treturn []byte{byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}\n}\nfunc int64_to_bytes(i int64) []byte {\n\treturn []byte{byte(i >> 56), byte(i >> 48), byte(i >> 40), byte(i >> 32), byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}\n}\nfunc uint8_to_byte(i uint8) byte {\n\treturn byte(i)\n}\nfunc uint16_to_bytes(i uint16) []byte {\n\treturn []byte{byte(i >> 8), byte(i)}\n}\nfunc uint32_to_bytes(i uint32) []byte {\n\treturn []byte{byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}\n}\nfunc uint64_to_bytes(i uint64) []byte {\n\treturn []byte{byte(i >> 56), byte(i >> 48), byte(i >> 40), byte(i >> 32), byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}\n}\nfunc float32_to_bytes(f float32) []byte {\n\treturn int32_to_bytes(int32(f))\n}\nfunc float64_to_bytes(f float64) []byte {\n\treturn int64_to_bytes(int64(f))\n}\nfunc string_to_bytes(s string) []byte {\n\treturn append([]byte(s), 0)\n}\nfunc main() {}"), 0644)
	os.WriteFile(dest+"/lib/"+name+".go", []byte("package main\n\nimport (\n\t\"fmt\"\n)\n\nvar calls = map[string]func([]stack) []stack{\n\t\"hello\": func(s []stack) (res []stack) {\n\t\tfmt.Println(\"Hello, World!\")\n\t\treturn res\n\t},\n}\n"), 0644)
	os.WriteFile(dest+"/test/test.mispasm", []byte("global _start\n\n_start\n\tcall \""+name+".hello\"\n"), 0644)
}
