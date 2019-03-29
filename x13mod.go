package main

/*
void x13bcd_hash(const char* input, char* output);
#cgo LDFLAGS: -L. -Lx13bcd_hash -lx13bcd
 */
import "C"
import (
	"unsafe"
)

func X13BCD(data []byte) []byte {
	output := make([]byte, 32)
	C.x13bcd_hash((*C.char)(unsafe.Pointer(&data[0])), (*C.char)(unsafe.Pointer(&output[0])))
	return output
}
