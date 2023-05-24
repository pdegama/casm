// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// label

package x86_64

import "fmt"

// label struct
type label struct {
	labelPos  int         // position of label value
	labelName string      // label name
	labelType operandType // valuse operand type imm*
	value     uint        // if value
	disp      bool        // displesment
}

// set label
func (b *binaryGen) setLabel() {

	// loop of bytes structure
	for _, bs := range b.bytesStruct {
		fmt.Println(bs)
	}

}
