// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// instruction

package x86_64

// instruction structure
type instruction struct {
	name     string    // instruction mnemonic name
	operands []operand // instruction operand by order
}
