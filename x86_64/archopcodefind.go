// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// architecture opcode find

package x86_64

import "fmt"

// find arch opcode
func archOpcodeFind(inst *instruction) []archOpcode {

	validOpcodes := []archOpcode{}        // valid insts opcodes stack
	validPerfectOpcodes := []archOpcode{} // valid insts opcodes stack

	// loop of instruction opcode
	for _, opcode := range archOpcodeList {

		// match mnemonic
		if opcode.name == inst.name {

			operTypeMatch := true       // operand type is match
			operTypePerfecMatch := true // operand type is match

			// operands type
			if len(opcode.operands) == len(inst.operands) { // if size of operands match

				// loop of opcode operand
				for opeIndex, opeType := range opcode.operands {

					// check operand is valid or not valid
					switch operandIsValid(opeType.t, &inst.operands[opeIndex]) {
					case operandMatch:
						/*
							operand type match but not perfect then
							this inst opcode is not perfect
						*/
						operTypePerfecMatch = false
					case operandNotMatch:
						/*
							if any operand type is not match then
							this inst id invalid and operTypeMatch
							set false
						*/
						operTypeMatch = false
						operTypePerfecMatch = false
					}
				}

				// operand all type match then append to stack
				if operTypeMatch {
					// if operTypeMatch is true then all type match
					if operTypePerfecMatch {
						/*
							if operTypePerfecMatch is true then
							this opcode is perfect match, and
							append to validPerfectOpcodes
						*/
						validPerfectOpcodes = append(validPerfectOpcodes, opcode)
						fmt.Println(">>>", opcode)
					} else {
						/*
							if opecode is not perfect match then
							append to validOpcodes
						*/
						validOpcodes = append(validOpcodes, opcode)
						fmt.Println("---", opcode)
					}
				}

			}

		}
	}

	// append valid arch opcode to valid perfect arch opcode
	validPerfectOpcodes = append(validPerfectOpcodes, validOpcodes...)

	return validPerfectOpcodes // return opcodes
}

// operand match type
const (
	operandNotMatch     = iota // operand type not match
	operandPerfectMatch        // operand type perfectly match
	operandMatch               // operand type match
)

// check operand type is valid
func operandIsValid(withOperand operandType, thisOperand *operand) int {

	/*
		thisOperand is operand, to be check with withOperand
		withOperand is operand type
	*/

	if withOperand == thisOperand.operandType {
		// if withOperand and thisOperand type is same
		return operandPerfectMatch
	}

	/*
		check imm type
	*/

	if thisOperand.operandType == imm {

		//fmt.Println(parseImmType(thisOperand.operand))
		if withOperand == parseImmType(thisOperand.operandVal) {
			return operandPerfectMatch
		}

		switch withOperand {
		case imm8, imm16, imm32, imm64:
			/*
				if withOperand is imm8, imm16, imm32
				or imm64, then accept imm
			*/
			return operandMatch
		}

	}

	/*
		check regMem
	*/

	if thisOperand.operandType == regMem {
		switch withOperand {
		case regMem8, regMem16, regMem32, regMem64:
			/*
				if withOperand is regMem8, regMem16,
				regMem32 or regMem64, then accept regMem
			*/
			return operandPerfectMatch
		}
	}

	/*
		check register
	*/

	if withOperand == regMem8 && thisOperand.operandType == reg8 {
		// if withOperand is regMem8 then accept reg8
		return operandPerfectMatch
	}

	if withOperand == regMem16 && thisOperand.operandType == reg16 {
		// if withOperand is regMem16 then accept reg16
		return operandPerfectMatch
	}

	if withOperand == regMem32 && thisOperand.operandType == reg32 {
		// if withOperand is regMem32 then accept reg32
		return operandPerfectMatch
	}

	if withOperand == regMem64 && thisOperand.operandType == reg64 {
		// if withOperand is regMem64 then accept reg64
		return operandPerfectMatch
	}

	// operand not match
	return operandNotMatch
}
