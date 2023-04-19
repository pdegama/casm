// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

/*
	parse architecture opcode data
*/

/*

	# Each architecture code CSV line contains these fields

	1. The Intel manual instruction mnemonic. For example, "SHR r/m32, imm8".

	2. The Go assembler instruction mnemonic. For example, "SHRL imm8, r/m32".

	3. The GNU binutils instruction mnemonic. For example, "shrl imm8, r/m32".

	4. The instruction encoding. For example, "C1 /4 ib".

	5. The validity of the instruction in 32-bit (aka compatibility, legacy) mode.

	6. The validity of the instruction in 64-bit mode.

	7. The CPUID feature flags that signal support for the instruction.

	8. Additional comma-separated tags containing hints about the instruction.

	9. The read/write actions of the instruction on the arguments used in the Intel mnemonic. For example, "rw,r" to denote that "SHR r/m32, imm8" reads and writes its first argument but only reads its second argument.

	10. Whether the opcode used in the Intel mnemonic has encoding forms distinguished only by operand size, like most arithmetic instructions. The string "Y" indicates yes, the string "" indicates no.

	11. The data size of the operation in bits. In general this is the size corresponding to the Go and GNU assembler opcode suffix.


	The complete line used for the above examples is:
	"SHR r/m32, imm8","SHRL imm8, r/m32","shrl imm8, r/m32","C1 /5 ib","V","V","","operand32","","Y","32"


*/

package main

import (
	"fmt"
	"strings"
)

type operandType string // tmp

type archOpcode struct { // tmp
	mnemonicName     string
	mnemonicOperands []operandType
	opcode           []int
	valid32BitMode   bool
	valid64BitMode   bool
}

// parse architecture code data
func parseData(csvRow []string) {

	if len(csvRow) != 11 {
		// if csv row column is not equle 11 then return error
		panic("invalid csv row")
	}

	instMnemonic := csvRow[0]       // instruction mnemonic string
	instOpcode := csvRow[3]         // instruction opcode string
	instValid32bitMode := csvRow[4] // instruction is valid in 32 bit mode
	instValid64bitMode := csvRow[5] // instruction is valid in 64 bit mode

	_ = instOpcode
	_ = instValid32bitMode
	_ = instValid64bitMode

	parseMnemonic(instMnemonic)

}

// parse mnemonic
func parseMnemonic(mnemonicStr string) string {

	// split mnemonic string with " "
	mnemonicArray := strings.Split(mnemonicStr, " ")

	mnemonicName := mnemonicArray[0]      // mnemonic name
	mnemonicOperands := mnemonicArray[1:] // mnemonic operands

	for i, operand := range mnemonicOperands {
		// if operand string is end with `,` then remove it
		if operand[len(operand)-1] == ',' {
			mnemonicOperands[i] = operand[:len(operand)-1]
		}
	}

	fmt.Println(mnemonicName, mnemonicOperands)

	// return
	return mnemonicName
}
