package main

import (
	_ "fmt"
)

func CreateMemory(size int) []byte {
	return make([]byte, size)
}

var ip = 0

func incIp(ip *int) int {
	x := *ip
	*ip += 1
	return x
}

func main() {
	cpu := NewCPU([]string{
		"ip", "acc", "r1", "r2", "r3", "r4", "r5", "r6", "r7", "r8",
	}, CreateMemory(256 * 256))

	cpu.mem[incIp(&ip)] = MOV_LIT_REG
	cpu.mem[incIp(&ip)] = 0x12
	cpu.mem[incIp(&ip)] = 0x34
	cpu.mem[incIp(&ip)] = R1

	cpu.mem[incIp(&ip)] = MOV_LIT_REG
	cpu.mem[incIp(&ip)] = 0xab
	cpu.mem[incIp(&ip)] = 0xcd
	cpu.mem[incIp(&ip)] = R2

	cpu.mem[incIp(&ip)] = ADD_REG_REG
	cpu.mem[incIp(&ip)] = R1 //index of r1 register
	cpu.mem[incIp(&ip)] = R2 //index of r2 register

	cpu.mem[incIp(&ip)] = MOV_REG_MEM
	cpu.mem[incIp(&ip)] = ACC
	cpu.mem[incIp(&ip)] = 0x01
	cpu.mem[incIp(&ip)] = 0x00

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