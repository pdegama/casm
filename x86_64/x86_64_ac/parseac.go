// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

/*
	parse architecture code data
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
	"SHR r/m32, imm8","SHRL imm8, r/m32","shrl imm8, r/m32","C1 /5 ib","V","V","","operand32","rw,r","Y","32"


*/

package main

// parse architecture code data
func parseData() {

}
