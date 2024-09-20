package utils

import  "C"

func Inportb(port uint16) uint8 {
	return uint8(C.inportb(C.uint16_t(port)))
}

func Outportb(port uint16, data uint8) {
	C.outportb(C.uint16_t(port), C.uint8_t(data))
}