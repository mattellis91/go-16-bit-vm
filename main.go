package main

import (
	_ "fmt"
)

func CreateMemory(size int) []byte {
	return make([]byte, size)
}

func main() {
	cpu := NewCPU([]string{
		"ip", "acc", "r1", "r2", "r3", "r4", "r5", "r6", "r7", "r8",
	}, CreateMemory(256 * 256))

	cpu.mem[0] = MOV_LIT_REG
	cpu.mem[1] = 0x12
	cpu.mem[2] = 0x34
	cpu.mem[3] = R1

	cpu.mem[4] = MOV_LIT_REG
	cpu.mem[5] = 0xab
	cpu.mem[6] = 0xcd
	cpu.mem[7] = R2

	cpu.mem[8] = ADD_REG_REG
	cpu.mem[9] = R1 //index of r1 register
	cpu.mem[10] = R2 //index of r2 register

	cpu.mem[11] = MOV_REG_MEM
	cpu.mem[12] = ACC
	cpu.mem[13] = 0x01
	cpu.mem[14] = 0x00

	cpu.PrintMemoryAt(cpu.GetRegister("ip"))
	cpu.PrintRegisters()

	cpu.step()

	cpu.PrintMemoryAt(cpu.GetRegister("ip"))
	cpu.PrintRegisters()

	cpu.step()

	cpu.PrintMemoryAt(cpu.GetRegister("ip"))
	cpu.PrintRegisters()

	cpu.step()

	cpu.PrintMemoryAt(cpu.GetRegister("ip"))
	cpu.PrintRegisters()

	cpu.step()

	cpu.PrintMemoryAt(0x0100)
	cpu.PrintRegisters()

 }