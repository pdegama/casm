// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// instruction generation

package x86_64

import "fmt"

// generation instruction
func genInsrtuction(opcode archOpcode, inst instruction, bitMode int) ([]uint8, error) {

	instBinCode := []uint8{}
	instPrefix := prefix{}

	fmt.Println(inst)
	fmt.Println(opcode)

	for _, i := range opcode.opcode {
		switch i {
		case modRM:

			// add modrm bytes
			modrmByte, err := addModRM(&opcode, &inst, bitMode, &instPrefix)
			if err != nil {
				// if error then return error
				return nil, err
			}
			instBinCode = append(instBinCode, modrmByte...)

		case modRM0, modRM1, modRM2, modRM3, modRM4, modRM5, modRM6, modRM7:

			// add modrm reg field 0
			modrmByte, err := addModRMfixRegField(&opcode, &inst, i, bitMode, &instPrefix)
			if err != nil {
				// if error then return error
				return nil, err
			}
			instBinCode = append(instBinCode, modrmByte...)

		case plusRB, plusRW, plusRD, plusRO:

			// add plus byte

			plusByte, err := plusRbyte(&opcode, &inst, bitMode, &instPrefix)
			if err != nil {
				return nil, err
			}

			instBinCode[len(instBinCode)-1] = instBinCode[len(instBinCode)-1] + plusByte

		case valIB, valIW, valID, valIO:

			// imm bytes

			immBytes, err := addImmIB(&opcode, &inst, i, bitMode, &instPrefix)
			if err != nil {
				// if error then return error
				return nil, err
			}
			instBinCode = append(instBinCode, immBytes...)

		case valCB:
			return nil, fmt.Errorf("todo valCB") // todo
		case valCW:
			return nil, fmt.Errorf("todo valCW") // todo
		case valCD:
			return nil, fmt.Errorf("todo valCD") // todo
		case valCP:
			return nil, fmt.Errorf("todo valCP") // todo
		case valCO:
			return nil, fmt.Errorf("todo valCO") // todo
		case valCT:
			return nil, fmt.Errorf("todo valCT") // todo
		case np:
			return nil, fmt.Errorf("todo np") // todo
		default:
			instBinCode = append(instBinCode, uint8(i))
		}
	}

	// check register operands
	err := checkRegisterOperand(&opcode, &inst, bitMode, &instPrefix)
	if err != nil {
		return nil, err
	}

	prefixByte := genPrefix(&instPrefix)
	instBinCode = append(prefixByte, instBinCode...)

	for _, b := range instBinCode {
		fmt.Printf("%x ", b)
	}

	return instBinCode, nil
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
