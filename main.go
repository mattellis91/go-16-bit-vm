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

	cpu.mem[incIp(&ip)] = MOV_LIT_REG
	cpu.mem[incIp(&ip)] = 0x51
	cpu.mem[incIp(&ip)] = 0x51
	cpu.mem[incIp(&ip)] = R1

	cpu.mem[incIp(&ip)] = MOV_LIT_REG
	cpu.mem[incIp(&ip)] = 0x42
	cpu.mem[incIp(&ip)] = 0x42
	cpu.mem[incIp(&ip)] = R2

	cpu.mem[incIp(&ip)] = PSH_REG
	cpu.mem[incIp(&ip)] = R1

	cpu.mem[incIp(&ip)] = PSH_REG
	cpu.mem[incIp(&ip)] = R2

	cpu.mem[incIp(&ip)] = POP
	cpu.mem[incIp(&ip)] = R1

	cpu.mem[incIp((&ip))] = POP
	cpu.mem[incIp((&ip))] = R2

	cpu.PrintRegisters()
	cpu.PrintMemoryAt(cpu.GetRegister("ip"))
	cpu.PrintMemoryAt(cpu.GetRegister("sp"))

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
