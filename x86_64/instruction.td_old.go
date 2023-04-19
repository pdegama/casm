// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// instruction table

package x86_64

// instruction opcode
var instOpcodes []instructionOpcode = []instructionOpcode{

	// ADD - Signed or Unsigned Add
	{"ADD", []operandType{regMem8, imm8}, []int{0x80, modRM0, valIB}},
	{"ADD", []operandType{regMem16, imm16}, []int{0x81, modRM0, valIW}},
	{"ADD", []operandType{regMem32, imm32}, []int{0x81, modRM0, valID}},
	{"ADD", []operandType{regMem64, imm32}, []int{0x81, modRM0, valID}},
	{"ADD", []operandType{regMem16, imm8}, []int{0x83, modRM0, valIB}},
	{"ADD", []operandType{regMem32, imm8}, []int{0x83, modRM0, valIB}},
	{"ADD", []operandType{regMem64, imm8}, []int{0x83, modRM0, valIB}},
	{"ADD", []operandType{regMem8, reg8}, []int{0x00, modRM}},
	{"ADD", []operandType{regMem16, reg16}, []int{0x01, modRM}},
	{"ADD", []operandType{regMem32, reg32}, []int{0x01, modRM}},
	{"ADD", []operandType{regMem64, reg64}, []int{0x01, modRM}},
	{"ADD", []operandType{reg8, regMem8}, []int{0x02, modRM}},
	{"ADD", []operandType{reg16, regMem16}, []int{0x03, modRM}},
	{"ADD", []operandType{reg32, regMem32}, []int{0x03, modRM}},
	{"ADD", []operandType{reg64, regMem64}, []int{0x03, modRM}},

	// MOV - Move
	{"MOV", []operandType{regMem8, reg8}, []int{0x88, modRM}},
	{"MOV", []operandType{regMem16, reg16}, []int{0x89, modRM}},
	{"MOV", []operandType{regMem32, reg32}, []int{0x89, modRM}},
	{"MOV", []operandType{regMem64, reg64}, []int{0x89, modRM}},
	{"MOV", []operandType{reg8, regMem8}, []int{0x8a, modRM}},
	{"MOV", []operandType{reg16, regMem16}, []int{0x8b, modRM}},
	{"MOV", []operandType{reg32, regMem32}, []int{0x8b, modRM}},
	{"MOV", []operandType{reg64, regMem64}, []int{0x8b, modRM}},
	{"MOV", []operandType{reg8, imm8}, []int{0xb0, plusRB, valIB}},
	{"MOV", []operandType{reg16, imm16}, []int{0xb8, plusRW, valIW}},
	{"MOV", []operandType{reg32, imm32}, []int{0xb8, plusRD, valID}},
	{"MOV", []operandType{reg64, imm64}, []int{0xb8, plusRQ, valIQ}},
	{"MOV", []operandType{regMem8, imm8}, []int{0xc6, modRM0, valIB}},
	{"MOV", []operandType{regMem16, imm16}, []int{0xc7, modRM0, valIW}},
	{"MOV", []operandType{regMem32, imm32}, []int{0xc7, modRM0, valID}},
	{"MOV", []operandType{regMem64, imm32}, []int{0xc7, modRM0, valID}},

	// SYSCALL - Fast System Call
	{"SYSCALL", []operandType{}, []int{0x0f, 0x05}},

	// RET (near) - Near Return from Called Procedure
	{"RET", []operandType{}, []int{0xc3}},
	{"RET", []operandType{imm16}, []int{0xc2, valIW}},

	// todo other inst
}
