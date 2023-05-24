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

	fmt.Println(inst)
	fmt.Println(opcode)

	for _, i := range opcode.opcode {
		switch i {
		case modRM:

			// add modrm bytes
			modrmByte, err := addModRM(&opcode, &inst, bitMode, &instPrefix)
			if err != nil {
				// if error then return error
				return instBytesStruct, err
			}
			instBytes = append(instBytes, modrmByte...)

		case modRM0, modRM1, modRM2, modRM3, modRM4, modRM5, modRM6, modRM7:

			// add modrm reg field 0
			modrmByte, err := addModRMfixRegField(&opcode, &inst, i, bitMode, &instPrefix)
			if err != nil {
				// if error then return error
				return instBytesStruct, err
			}
			instBytes = append(instBytes, modrmByte...)

		case plusRB, plusRW, plusRD, plusRO:

			// add plus byte

			plusByte, err := plusRbyte(&opcode, &inst, bitMode, &instPrefix)
			if err != nil {
				return instBytesStruct, err
			}

			instBytes[len(instBytes)-1] = instBytes[len(instBytes)-1] + plusByte

		case valIB, valIW, valID, valIO:

			// imm bytes

			immBytes, err := addImmIB(&opcode, &inst, i, bitMode, &instPrefix)
			if err != nil {
				// if error then return error
				return instBytesStruct, err
			}
			instBytes = append(instBytes, immBytes...)

		case valCB:
			return instBytesStruct, fmt.Errorf("todo valCB") // todo
		case valCW:
			return instBytesStruct, fmt.Errorf("todo valCW") // todo
		case valCD:
			return instBytesStruct, fmt.Errorf("todo valCD") // todo
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

	prefixByte := genPrefix(&instPrefix)
	instBytes = append(prefixByte, instBytes...)

	for _, b := range instBytes {
		fmt.Printf("%x ", b)
	}

	instBytesStruct.bytes = instBytes
	instBytesStruct.len = len(instBytes)

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
