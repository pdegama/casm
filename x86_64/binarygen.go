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
				err := instructionGen(line) // gen code for instruction
				if err != nil {
					// if error
					tErr := fmt.Errorf("%s %v:%v %v", errorStr, *line.filePath, line.index+1, err)
					errs = append(errs, tErr)
				}
				
			}

		}
	}

	return errs
}

