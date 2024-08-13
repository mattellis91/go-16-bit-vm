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
	if reg, ok := cpu.registerMap[name]; ok {
		return uint16(cpu.registers[reg]) << 8 | uint16(cpu.registers[reg + 1])
	}
	panic(fmt.Sprintf("Register not found: %s", name))
}

func (cpu *CPU) SetRegisterByName(name string, value uint16) {
	if reg, ok := cpu.registerMap[name]; ok {
		b := Uint16ToBytes(value) 
		cpu.registers[reg] = b[0]
		cpu.registers[reg+1] = b[1]
	} else {
		panic(fmt.Sprintf("Register not found: %s", name))
	}
}

func (cpu *CPU) SetRegisterByOffset(offset int, value uint16) {
	if offset < 0 || offset >= len(cpu.registers) {
		panic(fmt.Sprintf("Register offset out of bounds %d", offset))
	} else {
		b := Uint16ToBytes(value)
		cpu.registers[offset] = b[0]
		cpu.registers[offset + 1] = b[1]

	}
}

func (cpu *CPU) SetMemoryAtAddress(offset int, value uint16) {
	if offset < 0 || offset >= len(cpu.mem) {
		panic(fmt.Sprintf("Memory offset out of bounds %d", offset))
	} else {
		b := Uint16ToBytes(value)
		cpu.mem[offset] = b[0]
		cpu.mem[offset + 1] = b[1]
	}
}

func (cpu *CPU) fetch() byte {
	nextInstructionAddress := cpu.GetRegister("ip") 
	cpu.SetRegisterByName("ip", nextInstructionAddress + 1)
	return cpu.mem[nextInstructionAddress]
}

func (cpu *CPU) fetch16() uint16 {
	nextInstructionAddress := cpu.GetRegister("ip")
	cpu.SetRegisterByName("ip", nextInstructionAddress + 2)
	return uint16(cpu.mem[nextInstructionAddress]) << 8 | uint16(cpu.mem[nextInstructionAddress + 1])
}

func (cpu *CPU) readRegisterOffset() int {
	return (int(cpu.fetch()) % len(cpu.registerNames)) * 2
}

func (cpu *CPU) execute(instruction byte) {
	switch(instruction) {
		case MOV_LIT_REG: //MOVE literal value into register
			literal := cpu.fetch16()		
			regOffset := cpu.readRegisterOffset()
			cpu.SetRegisterByOffset(regOffset, literal)
		case MOV_REG_REG: //MOVE value from register to register
			regOffsetFrom := cpu.readRegisterOffset()
			regOffSetTo := cpu.readRegisterOffset()
			value := BytesToUint16(cpu.registers[regOffsetFrom], cpu.registers[regOffsetFrom + 1])
			cpu.SetRegisterByOffset(regOffSetTo, value)
		case MOV_REG_MEM: //MOVE register to memory
			regOffsetFrom := cpu.readRegisterOffset()
			address := cpu.fetch16()
			value := BytesToUint16(cpu.registers[regOffsetFrom], cpu.registers[regOffsetFrom + 1])
			cpu.SetMemoryAtAddress(int(address), value)
		case MOV_MEM_REG: //MOVE memory to register
			address := cpu.fetch16()
			regOffsetTo := cpu.readRegisterOffset()
			value := BytesToUint16(cpu.mem[address], cpu.mem[address + 1])
			cpu.SetRegisterByOffset(regOffsetTo, value)
		case ADD_REG_REG: //ADD register to register
			regA := cpu.fetch()
			regB := cpu.fetch() 
			regAValue := BytesToUint16(cpu.registers[regA * 2], cpu.registers[(regA * 2) + 1])
			regBValue := BytesToUint16(cpu.registers[regB * 2], cpu.registers[(regB * 2) + 1])
			cpu.SetRegisterByName("acc", regAValue + regBValue)
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

func (cpu *CPU) PrintRegisters() {
	for _, name := range cpu.registerNames {
		fmt.Printf("%s : %x\n", name, cpu.GetRegister(name))
	}
	fmt.Printf("\n\n")
}

func (cpu *CPU) PrintMemoryAt(address uint16) {
	nextEightBytes := cpu.mem[address:address + 8] 
	valString := fmt.Sprintf("%x: ", address)
	for _, value := range nextEightBytes {
		valString += fmt.Sprintf("%x ", value)
	}
	valString += "\n"
	fmt.Print(valString)
}