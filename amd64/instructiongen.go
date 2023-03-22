// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// instruction generation

package amd64

import (
	"fmt"
)

// instruction gen
func instructionGen(inst instruction) error {
	fmt.Println(inst)
	validInstOpcde := findInstruction(&inst)
	if len(validInstOpcde) == 0 {
		return fmt.Errorf("invalid instruction")
	}

	// print valid inst - tmp
	for _, i := range validInstOpcde {
		fmt.Println(i)
	}
	fmt.Println()

	return nil
}

// find insrtuction
func findInstruction(inst *instruction) []instructionOpcode {

	validInstsOpcode := []instructionOpcode{} // valid insts opcodes

	// loop of instruction opcode
	for _, instOpcode := range instOpcodes {
		if instOpcode.mnemonic == inst.mnemonic { // match mnemonic
			// check operand
			if validOperand(instOpcode.operandFirstType, &inst.operandFirst) && validOperand(instOpcode.operandSecondType, &inst.operandSecond) {
				// inst is valid then append to stack
				validInstsOpcode = append(validInstsOpcode, instOpcode)
			}
		}
	}

	return validInstsOpcode // return
}

// check operand type is valid
func validOperand(withOperand operandType, thisOperand *operand) bool {

	/*
		check imm type
	*/

	if thisOperand.operandType == imm {

		fmt.Println(int64(thisOperand.operand))

		if withOperand == imm8 || withOperand == imm16 || withOperand == imm32 || withOperand == imm64 {
			return true
		}
	}

	/*
		check register side
	*/

	// regMem8 is accept reg8
	if withOperand == regMem8 && thisOperand.operandType == reg8 {
		return true
	}

	// regMem16 is accept reg16
	if withOperand == regMem16 && thisOperand.operandType == reg16 {
		return true
	}

	// regMem32 is accept reg32
	if withOperand == regMem32 && thisOperand.operandType == reg32 {
		return true
	}

	// regMem64 is accept reg64
	if withOperand == regMem64 && thisOperand.operandType == reg64 {
		return true
	}

	// if operand type not match then deep check
	return withOperand == thisOperand.operandType
}
