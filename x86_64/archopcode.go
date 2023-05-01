// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// architecture opcode

package x86_64

// instruction architecture opcode structure
type archOpcode struct {
	name           string        // instruction mnemonic
	operands       []archOperand // instruction operand type by order
	opcode         []int         // instruction opcode
	valid32BitMode bool          // instruction is valid in 32-bit mode
	valid64BitMode bool          // instruction is valid in 64-bit mode
}

// operand structure for arch opcode list
type archOperand struct {
	t operandType // operand type
	v int         // operand value
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
	np = -23 // 

	todoOpcode = -100 // todo opcode

	//
)
