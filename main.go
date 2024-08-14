package main

import (
	"fmt"
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
		"ip", "acc", "r1", "r2", "r3", "r4", "r5", "r6", "r7", "r8", "sp", "fp",
	}, CreateMemory(256*256))

	cpu.Init()

	subRoutineAddress := 0x3000

	cpu.mem[incIp(&ip)] = PSH_LIT
	cpu.mem[incIp(&ip)] = 0x33
	cpu.mem[incIp(&ip)] = 0x33

	cpu.mem[incIp(&ip)] = PSH_LIT
	cpu.mem[incIp(&ip)] = 0x22
	cpu.mem[incIp(&ip)] = 0x22

	cpu.mem[incIp(&ip)] = MOV_LIT_REG
	cpu.mem[incIp(&ip)] = 0x12
	cpu.mem[incIp(&ip)] = 0x34
	cpu.mem[incIp(&ip)] = R1

	cpu.mem[incIp(&ip)] = MOV_LIT_REG
	cpu.mem[incIp(&ip)] = 0x56
	cpu.mem[incIp(&ip)] = 0x78
	cpu.mem[incIp(&ip)] = R2

	cpu.mem[incIp(&ip)] = PSH_LIT
	cpu.mem[incIp(&ip)] = 0x00
	cpu.mem[incIp(&ip)] = 0x00

	cpu.mem[incIp(&ip)] = CAL_LIT
	cpu.mem[incIp(&ip)] = byte(subRoutineAddress & 0xff00 >> 8) 
	cpu.mem[incIp(&ip)] = byte(subRoutineAddress & 0x00ff)

	cpu.mem[incIp(&ip)] = PSH_LIT
	cpu.mem[incIp(&ip)] = 0x44
	cpu.mem[incIp(&ip)] = 0x44

	//sub routine
	ip = subRoutineAddress

	cpu.mem[incIp(&ip)] = PSH_LIT
	cpu.mem[incIp(&ip)] = 0x01
	cpu.mem[incIp(&ip)] = 0x02

	cpu.mem[incIp(&ip)] = PSH_LIT
	cpu.mem[incIp(&ip)] = 0x03
	cpu.mem[incIp(&ip)] = 0x04

	cpu.mem[incIp(&ip)] = PSH_LIT
	cpu.mem[incIp(&ip)] = 0x05
	cpu.mem[incIp(&ip)] = 0x06

	cpu.mem[incIp(&ip)] = MOV_LIT_REG
	cpu.mem[incIp(&ip)] = 0x07
	cpu.mem[incIp(&ip)] = 0x08
	cpu.mem[incIp(&ip)] = R1

	cpu.mem[incIp(&ip)] = MOV_LIT_REG
	cpu.mem[incIp(&ip)] = 0x09
	cpu.mem[incIp(&ip)] = 0x0A
	cpu.mem[incIp(&ip)] = R8

	cpu.mem[incIp(&ip)] = RET

	cpu.PrintRegisters()
	cpu.PrintMemoryAt(cpu.GetRegister("ip"), 8)
	cpu.PrintMemoryAt(cpu.GetRegister("sp") - 42, 44)

	stepProg(cpu)
}

func stepProg(cpu *CPU) {
	for {
		var inp int
		fmt.Scanf("%v", &inp)
		cpu.step()
		cpu.PrintMemoryAt(cpu.GetRegister("ip"), 8)
		cpu.PrintRegisters()
		cpu.PrintMemoryAt(cpu.GetRegister("sp"), 8)
	}
}
