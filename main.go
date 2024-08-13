package main

import (
	"fmt"
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
	}, CreateMemory(256*256))

	cpu.mem[incIp(&ip)] = MOV_MEM_REG
	cpu.mem[incIp(&ip)] = 0x01
	cpu.mem[incIp(&ip)] = 0x00 //0x0100
	cpu.mem[incIp(&ip)] = R1

	cpu.mem[incIp(&ip)] = MOV_LIT_REG
	cpu.mem[incIp(&ip)] = 0x00
	cpu.mem[incIp(&ip)] = 0x01 //0x0001
	cpu.mem[incIp(&ip)] = R2

	cpu.mem[incIp(&ip)] = ADD_REG_REG
	cpu.mem[incIp(&ip)] = R1
	cpu.mem[incIp(&ip)] = R2

	cpu.mem[incIp(&ip)] = MOV_REG_MEM
	cpu.mem[incIp(&ip)] = ACC
	cpu.mem[incIp(&ip)] = 0x01
	cpu.mem[incIp(&ip)] = 0x00 //0x0100

	cpu.mem[incIp(&ip)] = JMP_NOT_EQU
	cpu.mem[incIp(&ip)] = 0x00
	cpu.mem[incIp(&ip)] = 0x05 // 0x0003
	cpu.mem[incIp(&ip)] = 0x00
	cpu.mem[incIp(&ip)] = 0x00 // 0x0000

	stepProg(cpu)
}

func stepProg(cpu *CPU) {
	for {
		var inp int
		fmt.Scanf("%v", &inp)
		cpu.step()
		cpu.PrintMemoryAt(cpu.GetRegister("ip"))
		cpu.PrintRegisters()
	}
}
