// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// assemble

package x86_64

import (
	"fmt"
)

// assemble asm file
func assemble(arch *x86_64) {

	// tokenization
	lines, err := lexer(arch.asmFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// tmp
	/* for _, line := range lines {
		fmt.Println(line)
	} */

	// parse asm lines
	lines, errs := parseLines(lines)
	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		return
	}

	binGen := binaryGen{}
	binGen.setBitMode(16)     // set bit mode
	binGen.setAsmLines(lines) // set asm lines
	
	errs = binGen.gen()       // gen binary
	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		return
	}

	// tmp
	/* for _, line := range lines {
		fmt.Println(line)
	} */

}
