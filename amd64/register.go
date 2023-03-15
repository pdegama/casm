// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// register

package amd64

import (
	"strings"
)

// register structure
type register struct {
	name       string // register name
	globIndex  int    // register globle index for this assembler
	bitSize    int    // register bit size 8, 16, 32, or 64-bit
	index      int    // register index
	baseOffset int    // register base offset +rb, +rw, +rd, or +rq value
}

// check or get register form register name
func isRegister(regName string) (bool, register) {

	regName = strings.ToUpper(regName) // convert uppercase

	// loop of register
	for _, reg := range registers {
		if reg.name == regName {
			// if register is match
			return true, reg
		}
	}

	// if register not found
	return false, register{}
}


