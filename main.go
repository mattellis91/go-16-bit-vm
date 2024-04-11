package main

import (
	"fmt"
)

func CreateMemory(size int) []byte {
	return make([]byte, size)
}

type CPU struct {
	mem []byte 
	registerNames []string
	registers []byte
	registerMap map[string]uint
}

func NewCPU(registerNames []string, mem []byte) *CPU {

	rm := make(map[string]uint)
	for i, name := range registerNames {
		rm[name] = uint(i) * 2
	}

	return &CPU{
		mem: mem,
		registerNames: registerNames,
		registers: make([]byte, len(registerNames) * 2),
		registerMap: rm,
	}
}

func GetRegister(cpu *CPU, name string) uint16 {
	if val, ok := cpu.registerMap[name]; ok {
		return uint16(cpu.registers[val]) << 8 | uint16(cpu.registers[val + 1])
	}
	panic("Register not found");
}


func main() {
	cpu := NewCPU([]string{
		"ip", "acc", "r1", "r2", "r3", "r4", "r5", "r6", "r7", "r8",
	}, CreateMemory(1024))
	fmt.Printf("%v\n", cpu)
}