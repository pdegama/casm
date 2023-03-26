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

// operand match type
const (
	operandNotMatch     = iota // operand type not match
	operandPerfectMatch        // operand type perfectly match
	operandMatch               // operand type match
)

// find insrtuction
func findInstruction(inst *instruction) []instructionOpcode {

	validInstsOpcode := []instructionOpcode{}        // valid insts opcodes stack
	validPerfectInstsOpcode := []instructionOpcode{} // valid insts opcodes stack

	// loop of instruction opcode
	for _, instOpcode := range instOpcodes {

		if instOpcode.mnemonic == inst.mnemonic { // match mnemonic

			operTypeMatch := true       // operand type is match
			operTypePerfecMatch := true // operand type is match

			// operands type
			if len(instOpcode.operandsType) == len(inst.operands) { // if size of operands match
				// loop of operands type
				for operTypeIndex, operType := range instOpcode.operandsType {

					// check operand is valid
					switch validOperand(operType, &inst.operands[operTypeIndex]) {
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
			}

			// operand all type match then append to stack
			if operTypeMatch {
				// if operTypeMatch is true then all type match
				//
				if operTypePerfecMatch {
					// if operand type perfect match
					validPerfectInstsOpcode = append(validPerfectInstsOpcode, instOpcode)
					fmt.Println(">>>", instOpcode)
				} else {
					validInstsOpcode = append(validInstsOpcode, instOpcode)
					fmt.Println("---", instOpcode)
				}
			}

		}

	}

	// append validInstsOpcode to validPerfectInstsOpcode
	validPerfectInstsOpcode = append(validPerfectInstsOpcode, validInstsOpcode...)

	return validPerfectInstsOpcode // return
}

// check operand type is valid
func validOperand(withOperand operandType, thisOperand *operand) int {

	// if operand type not match then deep check
	if withOperand == thisOperand.operandType {
		return operandPerfectMatch
	}

	/*
		check imm type
	*/

	if thisOperand.operandType == imm {

		//fmt.Println(parseImmType(thisOperand.operand))
		if withOperand == parseImmType(thisOperand.operand) {
			return operandPerfectMatch
		}

		if withOperand == imm8 || withOperand == imm16 || withOperand == imm32 || withOperand == imm64 {
			return operandMatch
		}
	}

	/*
		check register side
	*/

	// regMem8 is accept reg8
	if withOperand == regMem8 && thisOperand.operandType == reg8 {
		return operandPerfectMatch
	}

	// regMem16 is accept reg16
	if withOperand == regMem16 && thisOperand.operandType == reg16 {
		return operandPerfectMatch
	}

	// regMem32 is accept reg32
	if withOperand == regMem32 && thisOperand.operandType == reg32 {
		return operandPerfectMatch
	}

	// regMem64 is accept reg64
	if withOperand == regMem64 && thisOperand.operandType == reg64 {
		return operandPerfectMatch
	}

	// operand not match
	return operandNotMatch
}
