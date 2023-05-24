// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// label

package x86_64

// label struct
type label struct {
	labelPos  int         // position of label value
	labelType operandType // valuse operand type imm*
}
