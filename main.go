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
	}, CreateMemory(256))

	cpu.mem[0] = MOV_LIT_R1
	cpu.mem[1] = 0x12
	cpu.mem[2] = 0x34

	cpu.mem[3] = MOV_LIT_R2
	cpu.mem[4] = 0xab
	cpu.mem[5] = 0xcd

	cpu.mem[6] = ADD_REG_REG
	cpu.mem[7] = 2 //index of r1 register
	cpu.mem[8] = 3 //index of r2 register

	// fmt.Printf("%x", cpu.GetRegister("ip"))

	cpu.Debug()

	cpu.step()

	cpu.step()

	cpu.Debug()

	cpu.step()

	cpu.Debug()

	cpu.step()

	cpu.Debug()
}