package main

import (
	"encoding/binary"
	"fmt"
)

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

func (cpu *CPU) GetRegister(name string) uint16 {
	if val, ok := cpu.registerMap[name]; ok {
		return uint16(cpu.registers[val]) << 8 | uint16(cpu.registers[val + 1])
	}
	panic(fmt.Sprintf("Register not found: %s", name));
}

func (cpu *CPU) SetRegister(name string, value uint16) {
	if val, ok := cpu.registerMap[name]; ok {
		b := Uint16ToBytes(value) 
		cpu.registers[val] = b[0]
		cpu.registers[val+1] = b[1]
	} else {
		panic(fmt.Sprintf("Register not found: %s", name));
	}
}

func (cpu *CPU) fetch() byte {
	nextInstructionAddress := cpu.GetRegister("ip") 
	cpu.SetRegister("ip", nextInstructionAddress + 1)
	return cpu.mem[nextInstructionAddress]
}

func (cpu *CPU) fetch16() uint16 {
	nextInstructionAddress := cpu.GetRegister("ip")
	cpu.SetRegister("ip", nextInstructionAddress + 2)
	return uint16(cpu.mem[nextInstructionAddress]) << 8 | uint16(cpu.mem[nextInstructionAddress + 1])
}

func (cpu *CPU) execute(instruction byte) {
	switch(instruction) {
		case MOV_LIT_R1: //MOVE literal value into r1 register
			literal := cpu.fetch16()
			cpu.SetRegister("r1", literal)
		case MOV_LIT_R2: //MOVE literal value into r2 register
			literal := cpu.fetch16()
			cpu.SetRegister("r2", literal)
		case ADD_REG_REG: //ADD register to register
			regA := cpu.fetch()
			regB := cpu.fetch() 
			regAValue := BytesToUint16(cpu.registers[regA * 2], cpu.registers[(regA * 2) + 1])
			regBValue := BytesToUint16(cpu.registers[regB * 2], cpu.registers[(regB * 2) + 1])
			cpu.SetRegister("acc", regAValue + regBValue)
	}
}

func (cpu *CPU) step() {
	instruction := cpu.fetch()
	cpu.execute(instruction)
}

func Uint16ToBytes(value uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, value)
	return b
}

func BytesToUint16(b1 byte, b2 byte) uint16 {
	return uint16(b1) << 8 | uint16(b2)
}

func (cpu *CPU) Debug() {
	for _, name := range cpu.registerNames {
		fmt.Printf("%s : %x\n", name, cpu.GetRegister(name))
	}
	fmt.Printf("mem: %v", cpu.mem)
	fmt.Printf("\n\n")
}