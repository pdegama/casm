// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// architecture opcode find

package x86_64

import (
	"fmt"
)

// find arch opcode
func archOpcodeFind(inst *instruction, bitMode int) []archOpcode {

	validOpcodes := []archOpcode{}        // valid insts opcodes stack
	validPerfectOpcodes := []archOpcode{} // valid insts opcodes stack

	// loop of instruction opcode
	for _, opcode := range archOpcodeList {

		// match mnemonic
		if opcode.name == inst.name {

			if bitMode == 32 || bitMode == 16 {
				// if bit mode is 32 or 16
				if !opcode.valid32BitMode {
					// this opcode is not valid in 32-bit or 16-bit mode
					continue
				}
			}

			if bitMode == 64 {
				// if bit mode is 64
				if !opcode.valid64BitMode {
					// this opcode is not valid in 64-bit mode
					continue
				}
			}

			operTypeMatch := true       // operand type is match
			operTypePerfecMatch := true // operand type is match

			// operands type
			if len(opcode.operands) == len(inst.operands) { // if size of operands match

				// loop of opcode operand
				for opeIndex, opeType := range opcode.operands {

					// check operand is valid or not valid
					switch operandIsValid(&opeType, &inst.operands[opeIndex]) {
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

	// filter opcodes
	filterOpcodeImm(inst, &validOpcodes)

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
func operandIsValid(withOperand *archOperand, thisOperand *operand) int {

	/*
		thisOperand is operand, to be check with withOperand
		withOperand is operand type and value
	*/

	if withOperand.v != anyValue {
		// if operand have fix value
		if withOperand.v != int(thisOperand.v) {
			// fix value not match return operandNotMatch
			return operandNotMatch
		}
	}

	if withOperand.t == thisOperand.t {
		// if withOperand and thisOperand type is same
		return operandPerfectMatch
	}

	/*
		check imm type
	*/

	if thisOperand.t == imm {

		//fmt.Println(parseImmType(thisOperand.operand))
		if withOperand.t == parseImmType(thisOperand.v) {
			if !thisOperand.l {
				/*
					if this operand is not lable
				*/
				return operandPerfectMatch
			}
		}

		switch withOperand.t {
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

	if thisOperand.t == regMem {
		switch withOperand.t {
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

	if withOperand.t == regMem8 && thisOperand.t == reg8 {
		// if withOperand is regMem8 then accept reg8
		return operandPerfectMatch
	}

	if withOperand.t == regMem16 && thisOperand.t == reg16 {
		// if withOperand is regMem16 then accept reg16
		return operandPerfectMatch
	}

	if withOperand.t == regMem32 && thisOperand.t == reg32 {
		// if withOperand is regMem32 then accept reg32
		return operandPerfectMatch
	}

	if withOperand.t == regMem64 && thisOperand.t == reg64 {
		// if withOperand is regMem64 then accept reg64
		return operandPerfectMatch
	}

	// operand not match
	return operandNotMatch
}

// filter opcodes according imm
func filterOpcodeImm(inst *instruction, opcodes *[]archOpcode) {

	immPos := -1                               // imm operand position
	immIsLabel := false                        // imm operand is label
	for operPos, oper := range inst.operands { // loop of instruction operands
		if oper.t == imm {
			/*
				if oprand is imm set immPos
				is operand posation (index)
			*/
			immPos = operPos

			if oper.l {
				/*
					if operand is label then
					immIsLabel set true
				*/
				immIsLabel = true
			}
			break
		}
	}

	if immPos == -1 {
		/*
			if immPos is -1 then no any imm
			operand in this instruction
		*/
		return
	}

	// sorting opcodes
	for i := 0; i < len(*opcodes); i++ {

		for j := 0; j < i; j++ {

			if isSmallImm((*opcodes)[j].operands[immPos].t, (*opcodes)[i].operands[immPos].t) {

				tmpOper := (*opcodes)[i]
				(*opcodes)[i] = (*opcodes)[j]
				(*opcodes)[j] = tmpOper

			}

		}

	}

	if immIsLabel {
		/*
			if imm operand is lable then
			no more sorting this opcodes
		*/
		return
	}

	fmt.Println("Sorting...")
	_ = immIsLabel

}

// is smaller imm
func isSmallImm(a operandType, b operandType) bool {

	// a imm is smaller then b imm?

	switch b {
	case imm16:
		// if b is imm16
		if a == imm8 {
			/*
				and a is imm8 then return true
				beacuse imm8 (a) is smaller
				then imm16 (b)
			*/
			return true
		}
	case imm32:
		// if b is imm32
		if a == imm8 || a == imm16 {
			/*
				and a is imm8 or imm16 then return
				true beacuse imm8 or imm16 (a) is
				smaller then imm32 (b)
			*/
			return true
		}
	case imm64:
		// if b is imm64
		if a == imm8 || a == imm16 || a == imm32 {
			/*
				and a is imm8, imm16 or imm32 then
				return true beacuse imm8, imm16 or
				imm32 (a) is smaller then imm64 (b)
			*/
			return true
		}
	}

	return false

}
