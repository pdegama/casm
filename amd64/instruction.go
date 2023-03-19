// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// instruction

package amd64

// operand type
type operandType string

const ( // value is tmp

	//
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

	//
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

// instruction opcode structure
type instructionOpcode struct {
	mnemonic          string      // instruction name
	operandFirstType  operandType // first instruction operand type
	operandSecondType operandType // second instruction operand type
	opCode            []int       // instruction opcode
}

// opcode syntax code tmp will replace with original code
const (

	//
	NoneOpcode = -1

	ModRM  = -2  // ModRM byte
	ModRM0 = -3  // ModRM reg field 0
	ModRM1 = -4  // ModRM reg field 1
	ModRM2 = -5  // ModRM reg field 2
	ModRM3 = -6  // ModRM reg field 3
	ModRM4 = -7  // ModRM reg field 4
	ModRM5 = -8  // ModRM reg field 5
	ModRM6 = -9  // ModRM reg field 6
	ModRM7 = -10 // ModRM reg field 7

	PlusRB = -11 // 8-bit register value (base offset)
	PlusRW = -12 // 16-bit register value (base offset)
	PlusRD = -13 // 32-bit register value (base offset)
	PlusRQ = -14 // 64-bit register value (base offset)

	ValIB = -15 // immediate-operand value 8-bit
	ValIW = -16 // immediate-operand value 16-bit
	ValID = -17 // immediate-operand value 32-bit
	ValIQ = -18 // immediate-operand value 64-bit

	ValCB = -19 // code-offset value 8-bit
	ValCW = -20 // code-offset value 16-bit
	ValCD = -21 // code-offset value 32-bit
	ValCQ = -22 // code-offset value 64-bit

	//
)
