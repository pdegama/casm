// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// instruction generation

package x86_64

import "fmt"

// generation instruction
func genInsrtuction(opcode archOpcode, inst instruction, bitMode int) error {

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
				return err
			}
			instBinCode = append(instBinCode, modrmByte...)

		case modRM0, modRM1, modRM2, modRM3, modRM4, modRM5, modRM6, modRM7:

			// add modrm reg field 0
			modrmByte, err := addModRMfixRegField(&opcode, &inst, i, bitMode, &instPrefix)
			if err != nil {
				// if error then return error
				return err
			}
			instBinCode = append(instBinCode, modrmByte...)

		case plusRB:
			return fmt.Errorf("todo plusRB") // todo
		case plusRW:
			return fmt.Errorf("todo plusRW") // todo
		case plusRD:
			return fmt.Errorf("todo plusRD") // todo
		case plusRO:
			return fmt.Errorf("todo plusRO") // todo
		case valIB:

			immBytes, err := addImmIB(&opcode, &inst, imm8, bitMode, &instPrefix)
			if err != nil {
				// if error then return error
				return err
			}
			instBinCode = append(instBinCode, immBytes...)

		case valIW:

			immBytes, err := addImmIB(&opcode, &inst, imm16, bitMode, &instPrefix)
			if err != nil {
				// if error then return error
				return err
			}
			instBinCode = append(instBinCode, immBytes...)

		case valID:

			immBytes, err := addImmIB(&opcode, &inst, imm32, bitMode, &instPrefix)
			if err != nil {
				// if error then return error
				return err
			}
			instBinCode = append(instBinCode, immBytes...)

		case valIO:

			immBytes, err := addImmIB(&opcode, &inst, imm64, bitMode, &instPrefix)
			if err != nil {
				// if error then return error
				return err
			}
			instBinCode = append(instBinCode, immBytes...)

		case valCB:
			return fmt.Errorf("todo valCB") // todo
		case valCW:
			return fmt.Errorf("todo valCW") // todo
		case valCD:
			return fmt.Errorf("todo valCD") // todo
		case valCP:
			return fmt.Errorf("todo valCP") // todo
		case valCO:
			return fmt.Errorf("todo valCO") // todo
		case valCT:
			return fmt.Errorf("todo valCT") // todo
		case np:
			return fmt.Errorf("todo np") // todo
		default:
			instBinCode = append(instBinCode, uint8(i))
		}
	}

	prefixByte := genPrefix(&instPrefix)
	instBinCode = append(prefixByte, instBinCode...)

	for _, b := range instBinCode {
		fmt.Printf("%x ", b)
	}

	return nil
}
