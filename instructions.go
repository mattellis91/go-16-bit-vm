package main

const (
	MOV_LIT_REG byte = 0x10 //Move literal value to register
	MOV_REG_REG byte = 0x11 //Move register value to another register
	MOV_REG_MEM byte = 0x12 //Move register value to memory address
	MOV_MEM_REG byte = 0x13 //Move from value from memory address into register
	ADD_REG_REG byte = 0x14 //Add two registers and store result in accumulator register
	JMP_NOT_EQU byte = 0x15 //Jump if not equal.
	PSH_LIT     byte = 0x17 //Push literal value to stack
	PSH_REG     byte = 0x18 //Push register to stack
	POP         byte = 0x1A //Pop stack to register
	CAL_LIT     byte = 0x5E //Call sub routine at literal address
	CAL_REG     byte = 0x5F //Call sub routine at address stored in register
	RET         byte = 0x60 //Return from sub routine
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
	SP
	FP
)
