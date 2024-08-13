package main

const (
	MOV_LIT_REG byte = 0x10
	MOV_REG_REG byte = 0x11
	MOV_REG_MEM byte = 0x12
	MOV_MEM_REG byte = 0x13
	ADD_REG_REG byte = 0x14
)

const (
	IP byte = iota
	ACC
	R1
	R2
	R3
	R4
	R5
	R6
	R7
	R8
)