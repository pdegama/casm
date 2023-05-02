// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// instruction generation

package x86_64

import "fmt"

// generation instruction
func genInsrtuction(opcode archOpcode, inst instruction, bitMode int) error {

	instBinCode := []uint16{}

	fmt.Println(inst)
	fmt.Println(opcode)

	for _, i := range opcode.opcode {
		switch i {
		case modRM:
			return fmt.Errorf("todo modRM") // todo
		case modRM0:
			return fmt.Errorf("todo modRM0") // todo
		case modRM1:
			return fmt.Errorf("todo modRM1") // todo
		case modRM2:
			return fmt.Errorf("todo modRM2") // todo
		case modRM3:
			return fmt.Errorf("todo modRM3") // todo
		case modRM4:
			return fmt.Errorf("todo modRM4") // todo
		case modRM5:
			return fmt.Errorf("todo modRM5") // todo
		case modRM6:
			return fmt.Errorf("todo modRM6") // todo
		case modRM7:
			return fmt.Errorf("todo modRM7") // todo
		case plusRB:
			return fmt.Errorf("todo plusRB") // todo
		case plusRW:
			return fmt.Errorf("todo plusRW") // todo
		case plusRD:
			return fmt.Errorf("todo plusRD") // todo
		case plusRO:
			return fmt.Errorf("todo plusRO") // todo
		case valIB:
			return fmt.Errorf("todo valIB") // todo
		case valIW:
			return fmt.Errorf("todo valIW") // todo
		case valID:
			return fmt.Errorf("todo valID") // todo
		case valIO:
			return fmt.Errorf("todo valIO") // todo
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
			instBinCode = append(instBinCode, uint16(i))
		}
	}

	fmt.Printf("%x", instBinCode)

	return nil
}
