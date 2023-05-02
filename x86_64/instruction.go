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
func instructionGen(line asmLine, bitMode int) error {

	// parse instruction from tokens
	inst, err := parseInst(line.tokens)
	if err != nil {
		// inst parse err
		return fmt.Errorf("%s %v:%v %v", errorStr, *line.filePath, line.index+1, err)
	}

	fmt.Println(inst)
	fmt.Println("-----------------------------------------------------------------------")

	validInstOpcde := findArchOpcode(&inst, bitMode) // find valid opcode(s)
	if len(validInstOpcde) == 0 {
		// if opcode len is zero then return error
		return fmt.Errorf("invalid instruction")
	}

	// print valid inst - tmp
	for _, opcode := range validInstOpcde {
		genInsrtuction(opcode, inst)
		break
	}
	fmt.Println()

	return nil
}
