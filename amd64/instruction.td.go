// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// instruction table

package amd64

// instruction opcode
var instOpcodes []instructionOpcode = []instructionOpcode{

	// ADD - Signed or Unsigned Add
	{"ADD", regMem8, imm8, []int{0x80, modRM0, valIB}},
	{"ADD", regMem16, imm16, []int{0x81, modRM0, valIW}},
	{"ADD", regMem32, imm32, []int{0x81, modRM0, valID}},
	{"ADD", regMem64, imm32, []int{0x81, modRM0, valID}},
	{"ADD", regMem16, imm8, []int{0x83, modRM0, valIB}},
	{"ADD", regMem32, imm8, []int{0x83, modRM0, valIB}},
	{"ADD", regMem64, imm8, []int{0x83, modRM0, valIB}},
	{"ADD", regMem8, reg8, []int{0x00, modRM}},
	{"ADD", regMem16, reg16, []int{0x01, modRM}},
	{"ADD", regMem32, reg32, []int{0x01, modRM}},
	{"ADD", regMem64, reg64, []int{0x01, modRM}},
	{"ADD", reg8, regMem8, []int{0x02, modRM}},
	{"ADD", reg16, regMem16, []int{0x03, modRM}},
	{"ADD", reg32, regMem32, []int{0x03, modRM}},
	{"ADD", reg64, regMem64, []int{0x03, modRM}},

	// MOV - Move
	{"MOV", regMem8, reg8, []int{0x88, modRM}},
	{"MOV", regMem16, reg16, []int{0x89, modRM}},
	{"MOV", regMem32, reg32, []int{0x89, modRM}},
	{"MOV", regMem64, reg64, []int{0x89, modRM}},
	{"MOV", reg8, regMem8, []int{0x8a, modRM}},
	{"MOV", reg16, regMem16, []int{0x8b, modRM}},
	{"MOV", reg32, regMem32, []int{0x8b, modRM}},
	{"MOV", reg64, regMem64, []int{0x8b, modRM}},
	{"MOV", reg8, imm8, []int{0xb0, plusRB, valIB}},
	{"MOV", reg16, imm16, []int{0xb8, plusRW, valIW}},
	{"MOV", reg32, imm32, []int{0xb8, plusRD, valID}},
	{"MOV", reg64, imm64, []int{0xb8, plusRQ, valIQ}},
	{"MOV", regMem8, imm8, []int{0xc6, modRM0, valIB}},
	{"MOV", regMem16, imm16, []int{0xc7, modRM0, valIW}},
	{"MOV", regMem32, imm32, []int{0xc7, modRM0, valID}},
	{"MOV", regMem64, imm32, []int{0xc7, modRM0, valID}},

	// SYSCALL - Fast System Call
	{"SYSCALL", noneOperand, noneOperand, []int{0x0f, 0x05}},

	// RET (near) - Near Return from Called Procedure
	{"RET", noneOperand, noneOperand, []int{0xc3}},
	{"RET", imm16, noneOperand, []int{0xc2, valIW}},

	// todo other inst
}
