// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// imm rel

package x86_64

// add imm byte
func addImmBytes(opcode *archOpcode, inst *instruction, immOperType int, bitMode int, pf *prefix) ([]uint8, []label, error) {

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

	// check override prefix
	err := checkOperandOverride(&operand{t: immType}, bitMode, pf)
	if err != nil {
		return nil, []label{}, err
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

	labels := []label{}

	if immOperand.l {
		// if token is lable

		l := label{
			labelPos:  0,
			labelName: immOperand.n,
			labelType: immType,
			value:     0,
		}

		labels = append(labels, l) // append to label

	}

	return leBytes, labels, nil
}

// add rel byte
func addRelBytes(opcode *archOpcode, inst *instruction, relOperType int, bitMode int, pf *prefix) ([]uint8, []label, error) {

	var relType operandType // rel type

	switch relOperType {
	case valCB:
		relType = rel8
	case valCW:
		relType = rel16
	case valCD:
		relType = rel32
	}

	var relOperand *operand // rel operand

	// loop of arch operands
	for aOperIndex, aOper := range opcode.operands {
		if aOper.t == relType {
			/*
				if arch operand type is match relType
				then assign inst operand to relOperand
			*/
			relOperand = &inst.operands[aOperIndex]
		}
	}

	// check override prefix
/* 	err := checkOperandOverride(&operand{t: relType}, bitMode, pf)
	if err != nil {
		return nil, []label{}, err
	}
 */
	var leBytes []uint8 // le rel bytes

	switch relType {
	case rel8:
		/*
			if operand type is rel8 then
			convert value to rel8 and le bytes
		*/
		leBytes = uint8le(uint8(0))
	case rel16:
		/*
			if operand type is rel16 ...
		*/
		leBytes = uint16le(uint16(0))
	case rel32:
		/*
			if operand type is rel32 ...
		*/
		leBytes = uint32le(uint32(0))
	}

	labels := []label{}

	if relOperand.l {
		// if token is lable

		l := label{
			labelPos:  0,
			labelName: relOperand.n,
			labelType: relType,
			value:     relOperand.v,
		}

		labels = append(labels, l) // append to label

	}

	return leBytes, labels, nil

}
