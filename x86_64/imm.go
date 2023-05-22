// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// imm

package x86_64

import "fmt"

// add imm byte
func addImmIB(opcode *archOpcode, inst *instruction, immOperType int, bitMode int, pf *prefix) ([]uint8, error) {

	var immType operandType // imm type

	switch immOperType {
	case valIB:
		immType = imm8
	case valIW:
		immType = imm16
	case valID:
		immType = imm32
	case valIO:
		immType = imm64
	}

	var immOperand *operand // imm operand

	// loop of arch operands
	for aOperIndex, aOper := range opcode.operands {
		if aOper.t == immType {
			/*
				if arch operand type is match immType
				then assign inst operand to immOperand
			*/
			immOperand = &inst.operands[aOperIndex]
		}
	}

	if immOperand.l {
		// todo: lable is not support
		return nil, fmt.Errorf("%v label not support", immType)
	}

	// check override prefix
	err := checkOperandOverride(&operand{t: immType}, bitMode, pf)
	if err != nil {
		return nil, err
	}

	var leBytes []uint8 // le imm bytes

	switch immType {
	case imm8:
		/*
			if operand type is imm8 then
			convert value to imm8 and le bytes
		*/
		leBytes = uint8le(uint8(immOperand.v))
	case imm16:
		/*
			if operand type is imm16 ...
		*/
		leBytes = uint16le(uint16(immOperand.v))
	case imm32:
		/*
			if operand type is imm32 ...
		*/
		leBytes = uint32le(uint32(immOperand.v))
	case imm64:
		/*
			if operand type is imm64 ...
		*/
		leBytes = uint64le(uint64(immOperand.v))
	}

	return leBytes, nil
}