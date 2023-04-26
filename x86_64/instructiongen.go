// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// instruction generation

package x86_64

import (
	"fmt"
)

// instruction gen
func instructionGen(inst instruction) error {
	fmt.Println(inst)
	validInstOpcde := archOpcodeFind(&inst, 64)
	if len(validInstOpcde) == 0 {
		return fmt.Errorf("invalid instruction")
	}

	// print valid inst - tmp
	for _, i := range validInstOpcde {
		fmt.Println(i)
	}
	fmt.Println()

	return nil
}
