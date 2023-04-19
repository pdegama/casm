// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// operand

package x86_64

// operand type
type operandType string

// operand type
const ( // value is tmp

	//
	noneOperand operandType = "None" // none

	reg8     operandType = "reg8"     // 8-bit register
	reg16    operandType = "reg16"    // 16-bit register
	reg32    operandType = "reg32"    // 32-bit register
	reg64    operandType = "reg64"    // 64-bit register
	mem8     operandType = "mem8"     // 8-bit memory
	mem16    operandType = "mem16"    // 16-bit memory
	mem32    operandType = "mem32"    // 32-bit memory
	mem64    operandType = "mem64"    // 64-bit memory
	regMem8  operandType = "regMem8"  // 8-bit register or memory
	regMem16 operandType = "regMem16" // 16-bit register or memory
	regMem32 operandType = "regMem32" // 32-bit register or memory
	regMem64 operandType = "regMem64" // 64-bit register or memory
	imm8     operandType = "imm8"     // 8-bit immediate value
	imm16    operandType = "imm16"    // 16-bit immediate value
	imm32    operandType = "imm32"    // 32-bit immediate value
	imm64    operandType = "imm64"    // 64-bit immediate value

	reg    operandType = "reg"    // group of Reg*
	regMem operandType = "regMem" // group of RegMem*
	imm    operandType = "imm"    // group of imm*

	//
)

// operand structure
type operand struct {
	operandType operandType // operand type
	operandVal  uint        // operand value
	operandMem  []operand   // mem operands
}
