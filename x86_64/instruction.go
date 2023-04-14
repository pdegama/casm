// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// instruction

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

// instruction structure
type instruction struct {
	mnemonic string    // instruction mnemonic
	operands []operand // nstruction operand by order
}

// instruction opcode structure
type instructionOpcode struct {
	mnemonic     string        // instruction name
	operandsType []operandType // instruction operand type by order
	opCode       []int         // instruction opcode
}

// opcode syntax code
const (

	// code is tmp will replace with original code

	//
	noneOpcode = -1

	modRM  = -2  // ModRM byte
	modRM0 = -3  // ModRM reg field 0
	modRM1 = -4  // ModRM reg field 1
	modRM2 = -5  // ModRM reg field 2
	modRM3 = -6  // ModRM reg field 3
	modRM4 = -7  // ModRM reg field 4
	modRM5 = -8  // ModRM reg field 5
	modRM6 = -9  // ModRM reg field 6
	modRM7 = -10 // ModRM reg field 7
	plusRB = -11 // 8-bit register value (base offset)
	plusRW = -12 // 16-bit register value (base offset)
	plusRD = -13 // 32-bit register value (base offset)
	plusRQ = -14 // 64-bit register value (base offset)
	valIB  = -15 // immediate-operand value 8-bit
	valIW  = -16 // immediate-operand value 16-bit
	valID  = -17 // immediate-operand value 32-bit
	valIQ  = -18 // immediate-operand value 64-bit
	valCB  = -19 // code-offset value 8-bit
	valCW  = -20 // code-offset value 16-bit
	valCD  = -21 // code-offset value 32-bit
	valCQ  = -22 // code-offset value 64-bit

	//
)
