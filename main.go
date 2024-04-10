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

func GetRegister(cpu *CPU, name string) uint {
	//TODO: get value (2 bytes) from memory based on the register name
	return 0
}


func main() {
	cpu := NewCPU([]string{
		"ip", "acc", "r1", "r2", "r3", "r4", "r5", "r6", "r7", "r8",
	}, CreateMemory(1024))
	fmt.Printf("%v\n", cpu)
}