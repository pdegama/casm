// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// instruction

package x86_64

import "fmt"

// instruction structure
type instruction struct {
	name     string    // instruction mnemonic name
	operands []operand // instruction operand by order
}

// instruction gen
func instructionGen(line asmLine, bitMode int) ([]uint8, error) {

	// parse instruction from tokens
	inst, err := parseInst(line.tokens)
	if err != nil {
		// inst parse err
		return nil, err
	}

	fmt.Println()
	fmt.Println(inst)
	fmt.Println("-----------------------------------------------------------------------")

	validInstOpcde := findArchOpcode(&inst, bitMode) // find valid opcode(s)
	if len(validInstOpcde) == 0 {
		// if opcode len is zero then return error
		return nil, fmt.Errorf("invalid instruction")
	}

	var instBinCode []uint8

	// print valid inst - tmp
	for _, opcode := range validInstOpcde {
		instBinCode, err = genInsrtuction(opcode, inst, bitMode)
		if err != nil {
			return nil, err
		}
		break
	}
	fmt.Println()

	return instBinCode, nil
}
