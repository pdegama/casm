// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// binary generation

package amd64

import "fmt"

// code generation
func codeGen(lines []asmLine) error {

	for _, line := range lines {
		if len(line.tokens) != 0 {
			switch line.lineType{
			case lineInst:
				// if line type is insrtuction
				genCodeInst(line) // gen code for instruction
			}
		}
	}
	
	return nil
}

// code generation for instruction
func genCodeInst(line asmLine) {

	toks := parseInst(line.tokens)
	fmt.Println(toks)

}
