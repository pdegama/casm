// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// instruction generation

package x86_64

import "fmt"

// generation instruction
func genInsrtuction(opcode archOpcode, inst instruction, bitMode int) (bytesStructure, error) {

	instBytes := []uint8{} // inst bytes
	instPrefix := prefix{} // inst prefix

	instBytesStruct := bytesStructure{} // inst bytes struct

	instLabels := []label{} // inst labels

	fmt.Println(inst)
	fmt.Println(opcode)

	for _, i := range opcode.opcode {
		switch i {
		case modRM:

			// add modrm bytes
			modrmByte, modrmLabels, err := addModRM(&opcode, &inst, bitMode, &instPrefix)
			if err != nil {
				// if error then return error
				return instBytesStruct, err
			}
			addCurrentPosLabel(&modrmLabels, len(instBytes)) // set label pos
			instBytes = append(instBytes, modrmByte...)      // append inst bytes
			instLabels = append(instLabels, modrmLabels...)  // append labels

		case modRM0, modRM1, modRM2, modRM3, modRM4, modRM5, modRM6, modRM7:

			// add modrm reg field 0
			modrmByte, modrmLabels, err := addModRMfixRegField(&opcode, &inst, i, bitMode, &instPrefix)
			if err != nil {
				// if error then return error
				return instBytesStruct, err
			}

			addCurrentPosLabel(&modrmLabels, len(instBytes)) // set label pos
			instBytes = append(instBytes, modrmByte...)      // append inst bytes
			instLabels = append(instLabels, modrmLabels...)  // append labels

		case plusRB, plusRW, plusRD, plusRO:

			// add plus byte

			plusByte, err := plusRbyte(&opcode, &inst, bitMode, &instPrefix)
			if err != nil {
				return instBytesStruct, err
			}

			instBytes[len(instBytes)-1] = instBytes[len(instBytes)-1] + plusByte

		case valIB, valIW, valID, valIO:

			// imm bytes

			immBytes, immLabels, err := addImmBytes(&opcode, &inst, i, bitMode, &instPrefix)
			if err != nil {
				// if error then return error
				return instBytesStruct, err
			}

			addCurrentPosLabel(&immLabels, len(instBytes)) // set label pos
			instBytes = append(instBytes, immBytes...)     // append inst bytes
			instLabels = append(instLabels, immLabels...)  // append labels

		case valCB, valCW, valCD:
			return instBytesStruct, fmt.Errorf("todo valCB") // todo

			// offset bytes

			offsetBytes, offsetLabels, err := addRelBytes(&opcode, &inst, i, bitMode, &instPrefix)
			if err != nil {
				// if error then return error
				return instBytesStruct, err
			}

			addCurrentPosLabel(&offsetLabels, len(instBytes)) // set label pos
			instBytes = append(instBytes, offsetBytes...)     // append inst bytes
			instLabels = append(instLabels, offsetLabels...)  // append labels

		case valCP:
			return instBytesStruct, fmt.Errorf("todo valCP") // todo
		case valCO:
			return instBytesStruct, fmt.Errorf("todo valCO") // todo
		case valCT:
			return instBytesStruct, fmt.Errorf("todo valCT") // todo
		case np:
			return instBytesStruct, fmt.Errorf("todo np") // todo
		default:
			instBytes = append(instBytes, uint8(i))
		}
	}

	// check register operands
	err := checkRegisterOperand(&opcode, &inst, bitMode, &instPrefix)
	if err != nil {
		return instBytesStruct, err
	}

	prefixByte := genPrefix(&instPrefix)         // prefix bytes
	instBytes = append(prefixByte, instBytes...) // append prefix

	addPosLabel(&instLabels, len(prefixByte)) // add prefix length in label pos

	for _, b := range instBytes {
		fmt.Printf("%x ", b)
	}

	instBytesStruct.bytes = instBytes
	instBytesStruct.len = len(instBytes)
	instBytesStruct.labels = instLabels

	return instBytesStruct, nil
}

// plus +r* byte
func plusRbyte(opcode *archOpcode, inst *instruction, bitMode int, pf *prefix) (uint8, error) {

	var regPos int // reg pos

	// loop off opcode operands
	for operPos, oper := range opcode.operands {
		if isRegOperand(oper.t) {
			/*
				if operans is register then
				set regPos is operPos
			*/
			regPos = operPos
		}
	}

	regOper := inst.operands[regPos] // register operand

	// register info
	regInfo, err := registerInfo(int(regOper.v))
	if err != nil {
		return 0, err
	}

	// check is valid or not
	err = registerIsValid(regInfo, bitMode)
	if err != nil {
		return 0, err
	}

	// check operand override prefix
	err = checkOperandOverride(&regOper, bitMode, pf)
	if err != nil {
		return 0, err
	}

	// check rex.b prefix
	err = checkREXbPrexif(&regOper, bitMode, pf)
	if err != nil {
		return 0, err
	}

	// plus byte is register base offset
	return uint8(regInfo.baseOffset), nil

}

// check register operand...
func checkRegisterOperand(opcode *archOpcode, inst *instruction, bitMode int, pf *prefix) error {

	r := 0 // register operand count
	m := 0 // memory operand count

	for _, oper := range opcode.operands {
		// loop of opcode operands

		if isRegOperand(oper.t) {
			r++ // if operand is register oper
		}

		if isMemOperand(oper.t) {
			m++ // if operand is memory oper
		}
	}

	if r != 0 {

		// register is not zero in this opcode

		for operIndex, oper := range opcode.operands {
			// loop of opcode operands

			if isRegOperand(oper.t) {
				// if operand is register

				regOperand := inst.operands[operIndex] // register operand

				// check operand override prefix
				err := checkOperandOverride(&regOperand, bitMode, pf)
				if err != nil {
					return err
				}

				if m == 0 {
					// if no memory operand
					err := checkREXbPrexif(&regOperand, bitMode, pf)
					if err != nil {
						return err
					}
				}

			}

		}

	}

	return nil

}

// add currennt pos in label
func addCurrentPosLabel(labels *[]label, cp int) {

	// loop of labels
	for i := range *labels {
		// add current position in label
		(*labels)[i].labelPos = (*labels)[i].labelPos + cp
	}

}

// add pos in label
func addPosLabel(labels *[]label, cp int) {

	// loop of labels
	for i := range *labels {
		// add current position in label
		(*labels)[i].labelPos = (*labels)[i].labelPos + cp
	}

}
