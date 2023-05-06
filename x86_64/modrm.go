// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// modrm

package x86_64

import "fmt"

// calculate modRM
func calcModRM(opers ...operand) (uint16, error) {

	// calculate modRM and return
	if len(opers) != 2 {
		return 0x00, fmt.Errorf("internal error: todo: operand size is not two")
	}

	return 0x00, nil
}
