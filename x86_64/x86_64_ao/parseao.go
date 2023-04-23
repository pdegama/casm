// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// parse architecture opcode data

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

type archOpcode struct {
	name           string
	operands       []archOperand
	opcode         []int
	valid32BitMode bool
	valid64BitMode bool
}

type archOperand struct {
	t operandType
	v int
}

// parse architecture code data
func parseData(csvRow []string) (string, string, bool, bool, error) {

	if len(csvRow) != 11 {
		// if csv row column is not equle 11 then return error
		panic("invalid csv row")
	}

	instMnemonic := csvRow[0] // instruction mnemonic string
	// instOpcode := csvRow[3]         // instruction opcode string
	instValid32bitMode := csvRow[4] // instruction is valid in 32 bit mode
	instValid64bitMode := csvRow[5] // instruction is valid in 64 bit mode

	// parse inst mnemonic name and operands
	mnemonicName, mOperands, err := parseMnemonic(instMnemonic)
	if err != nil {
		panic(err)
	}

	// parse mnemonic operands
	mnemonicOperands, err := parseMnemonicOperand(mOperands)
	if err != nil {
		return "", "", false, false, err
	}

	// parse bit mode
	valid32bitMode := parseValidMode(instValid32bitMode)
	valid64bitMode := parseValidMode(instValid64bitMode)

	// return
	return mnemonicName, mnemonicOperands, valid32bitMode, valid64bitMode, nil

}

// parse mnemonic
func parseMnemonic(mnemonicStr string) (string, []string, error) {

	// split mnemonic string with " "
	mnemonicArray := strings.Split(mnemonicStr, " ")

	mnemonicName := mnemonicArray[0]                                             // mnemonic name
	mnemonicOperands := strings.Split(strings.Join(mnemonicArray[1:], " "), ",") // mnemonic operands

	if len(mnemonicOperands) == 1 && mnemonicOperands[0] == "" {
		return mnemonicName, []string{}, nil
	}
	for i, operand := range mnemonicOperands {
		// remove left and right space
		mnemonicOperands[i] = strings.Trim(operand, " ")
	}

	// return
	return mnemonicName, mnemonicOperands, nil
}

// parse valid on 32-bit or 64 bit mode
func parseValidMode(validStr string) bool {

	switch validStr {
	case "V": // valid
		return true
	case "N.E.", "I": // invalid
		return false
	case "N.S.": // ?
		return false
	}

	// str is not match ni switch-case then panic program
	panic(validStr)
}

// parse mnemonic operands
func parseMnemonicOperand(mOperands []string) (string, error) {

	operandString := ""

	for opeIndex, ope := range mOperands {

		opeType, opeVal, err := checkOperand(ope)
		if err != nil {
			return "", err
		}

		operandString += makeArchOperandStruct(opeType, opeVal)
		if len(mOperands)-1 != opeIndex {
			operandString += ", "
		}
	}

	return "[]archOperand{" + operandString + "}", nil
}

// check operand
func checkOperand(opr string) (string, string, error) {
	todoOper := hasList(todoOperand, opr)
	if todoOper {
		return "", "", fmt.Errorf("todo: operand %v", opr)
	}

	// if operand is register
	switch opr {
	case "AL":
		return "reg8", "al", nil
	case "AX":
		return "reg16", "ax", nil
	case "EAX":
		return "reg32", "eax", nil
	case "RAX":
		return "reg64", "rax", nil
	case "DX":
		return "reg16", "dx", nil
	case "CL":
		return "reg8", "cl", nil
	}

	// replace string
	opr = replaceOperand(opr)

	return opr, "1", nil
}

// replave operand string
func replaceOperand(opr string) string {

	// loop of operand replace list
	for k, v := range replaceOperandList {
		if opr == k {
			// if replace string is exist then return replace value
			return v
		}
	}

	// other replace opr
	return opr
}

// replace operand list
var replaceOperandList map[string]string = map[string]string{
	"ptr16:16": "ptr16_16",
	"ptr16:32": "ptr16_32",
	"r8":       "reg8",
	"r16":      "reg16",
	"r32":      "reg32",
	"r64":      "reg64",
	"m8":       "mem8",
	"m16":      "mem16",
	"m32":      "mem32",
	"m64":      "mem64",
	"m128":     "mem128",
	"r/m8":     "regMem8",
	"r/m16":    "regMem16",
	"r/m32":    "regMem32",
	"r/m64":    "regMem64",
	"m16&32":   "mem16n32",
	"m16&16":   "mem16n16",
	"m32&32":   "mem32n32",
	"m16&64":   "mem16n64",
	"m16:16":   "mem16_16",
	"m16:32":   "mem16_32",
	"m16:64":   "mem16_64",
	"mm":       "mmx",
	"xmm/m128": "xmmMem128",
	"m80bcd":   "mem80bcd",
	"m16int":   "mem16int",
	"m32int":   "mem32int",
	"m64int":   "mem64int",
	"m256":     "mem256",
	"Sreg":     "sReg",
	"moffs8":   "memOffset8",
	"moffs16":  "memOffset16",
	"moffs32":  "memOffset32",
	"moffs64":  "memOffset64",
}

// todo operands
var todoOperand = []string{
	"<XMM0>",
	"<XMM0-2>",
	"<XMM4-6>",
	"r32 <XMM0-6>",
	"m14/28byte",
	"m94/108byte",
	"{k1}{z}",
	"xmm1{k1}{z}", "xmm1 {k1}{z}",
	"ymm1{k1}{z}", "ymm1 {k1}{z}",
	"zmm1{k1}{z}", "zmm1 {k1}{z}",
	"xmm2/m128 {k1}{z}",
	"ymm2/m256 {k1}{z}",
	"zmm2/m512 {k1}{z}",
	"xmm",
	"xmm1", "xmm2",
	"xmm2/m128", "xmm2/m64", "xmm2/m32",
	"ymm1", "ymm2",
	"zmm1",
	"m384", "m512",
	"r32a", "r32b",
	"r64a", "r64b",
	"bnd1", "bnd2",
	"bnd1/m128", "bnd1/m64",
	"m",
	"r16op", // maybe invalid operand
	"r32op",
	"r64op",
	"ptr16:16", "ptr16:32",
	"m128",
	"mm",
	"mm1", "mm2",
	"xmm/m128",
	"k1", "k2", "k3",
	"xmm1/m64", "xmm1/m32",
	"imm8b",
	"ST(0)", "ST(i)",
	"0", "1",
	"m512byte",
	"m32fp", "m64fp", "m80fp",
	"m80bcd",
	"m16int", "m32int", "m64int",
	"r32/m16",
	"m2byte",
	"imm8u", "imm16u",
	"r16/r32/m16",
	"CR0-CR7", "CR8",
	"DR0-DR7",
	"r64/m16",
	"CS", "DS", "ES", "FS", "GS", "SS",
	"k1 {k2}",
	"xmm1/m64{er}",
	"xmm1/m32{er}",
	"xmm1/m64{sae}",
	"xmm1/m32{sae}",
}

// check is member in the list
func hasList[T comparable](list []T, m T) bool {
	for _, lm := range list {
		if lm == m {
			return true
		}
	}
	return false
}
