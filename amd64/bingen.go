// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// binary generation

package amd64

import "fmt"

// code generation
func codeGen(lines []asmLine) error {

	for _, line := range lines {
		if len(line.tokens) != 0 {
			fmt.Println(line)
		}
	}

	return nil
}
