// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// assemble

package amd64

import "fmt"

// assemble asm file
func assemble(arch *amd64) {

	// tokenization
	lines, err := lexer(arch.asmFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(lines) // tmp

	// parse asm lines
	lines, errs := parseLines(lines)
	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		return
	}

	fmt.Println(lines) // tmp

}
