// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// binary generation

package x86_64

import (
	"fmt"
)

// code generation
func codeGen(lines []asmLine) []error {

	errs := []error{} // errors

	for _, line := range lines {
		if len(line.tokens) != 0 {
			switch line.lineType {
			case lineInst:
				// if line type is insrtuction
				err := genCodeInst(line) // gen code for instruction
				if err != nil {
					// if error
					errs = append(errs, err)
				}
			}
		}
	}

	return errs
}

// code generation for instruction
func genCodeInst(line asmLine) error {

	inst, err := parseInst(line.tokens)
	if err != nil {
		// inst parse err
		return fmt.Errorf("%s %v:%v %v", errorStr, *line.filePath, line.index+1, err)
	}

	//  instruction gen
	err = instructionGen(inst)
	if err != nil {
		// inst gen error
		return fmt.Errorf("%s %v:%v %v", errorStr, *line.filePath, line.index+1, err)
	}

	return nil
}
