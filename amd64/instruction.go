// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// instruction

package amd64

// operand type
type operandType string

const ( // value is tmp
	noneOperand operandType = "None" // none

	reg8     = "reg8"     // 8-bit register
	reg16    = "reg16"    // 16-bit register
	reg32    = "reg32"    // 32-bit register
	reg64    = "reg64"    // 64-bit register
	regMem8  = "regMem8"  // 8-bit register or memory
	regMem16 = "regMem16" // 16-bit register or memory
	regMem32 = "regMem32" // 32-bit register or memory
	regMem64 = "regMem64" // 64-bit register or memory
	imm8     = "imm8"     // 8-bit immediate value
	imm16    = "imm16"    // 16-bit immediate value
	imm32    = "imm32"    // 32-bit immediate value
	imm64    = "imm64"    // 64-bit immediate value

	reg    = "regGroup"    // group of Reg*
	regMem = "regMemGroup" // group of RegMem*
	imm    = "imm"         // group of imm*
)

// operand structure
type operand struct {
	operandType operandType // operand type
	operand     int         // operand
}
