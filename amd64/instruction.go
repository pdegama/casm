// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// instruction

package amd64

// operand type
type operandType string

const ( // value is tmp
	noneOperand operandType = "None" // none

	reg8     operandType = "reg8"     // 8-bit register
	reg16    operandType = "reg16"    // 16-bit register
	reg32    operandType = "reg32"    // 32-bit register
	reg64    operandType = "reg64"    // 64-bit register
	regMem8  operandType = "regMem8"  // 8-bit register or memory
	regMem16 operandType = "regMem16" // 16-bit register or memory
	regMem32 operandType = "regMem32" // 32-bit register or memory
	regMem64 operandType = "regMem64" // 64-bit register or memory
	imm8     operandType = "imm8"     // 8-bit immediate value
	imm16    operandType = "imm16"    // 16-bit immediate value
	imm32    operandType = "imm32"    // 32-bit immediate value
	imm64    operandType = "imm64"    // 64-bit immediate value

	reg    operandType = "regGroup"    // group of Reg*
	regMem operandType = "regMemGroup" // group of RegMem*
	imm    operandType = "imm"         // group of imm*
)

// operand structure
type operand struct {
	operandType operandType // operand type
	operand     int         // operand
}

// instruction structure
type instruction struct {
	mnemonic      string  // instruction mnemonic
	operandFirst  operand // first operand
	operandSecond operand // second operand
}
