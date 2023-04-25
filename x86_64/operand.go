// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// operand

package x86_64

// operand type
type operandType string

// any operandValue
const anyValue = -1

// operand type
const ( // value is tmp

	//
	noneOperand operandType = "None" // none

	rel8        operandType = "rel8"
	rel16       operandType = "rel16"
	rel32       operandType = "rel32"
	ptr16_16    operandType = "ptr16:16"
	ptr16_32    operandType = "ptr16:32"
	reg8        operandType = "reg8"     // 8-bit register
	reg16       operandType = "reg16"    // 16-bit register
	reg32       operandType = "reg32"    // 32-bit register
	reg64       operandType = "reg64"    // 64-bit register
	imm8        operandType = "imm8"     // 8-bit immediate value
	imm16       operandType = "imm16"    // 16-bit immediate value
	imm32       operandType = "imm32"    // 32-bit immediate value
	imm64       operandType = "imm64"    // 64-bit immediate value
	regMem8     operandType = "regMem8"  // 8-bit register or memory
	regMem16    operandType = "regMem16" // 16-bit register or memory
	regMem32    operandType = "regMem32" // 32-bit register or memory
	regMem64    operandType = "regMem64" // 64-bit register or memory
	mem8        operandType = "mem8"     // 8-bit memory
	mem16       operandType = "mem16"    // 16-bit memory
	mem32       operandType = "mem32"    // 32-bit memory
	mem64       operandType = "mem64"    // 64-bit memory
	mem128      operandType = "mem64"    // 128-bit memory
	mem16_16    operandType = "mem16:16"
	mem16_32    operandType = "mem16:32"
	mem16_64    operandType = "mem16:64"
	mem16n32    operandType = "mem16&32"
	mem16n16    operandType = "mem16&16"
	mem32n32    operandType = "mem32&32"
	mem16n64    operandType = "mem16&64"
	mem80bcd    operandType = "mem80bcd"
	memOffset8  operandType = "memOffset8"
	memOffset16 operandType = "memOffset16"
	memOffset32 operandType = "memOffset32"
	memOffset64 operandType = "memOffset64"
	sReg        operandType = "sReg"
	mem32fp     operandType = "mem32fp"
	mem64fp     operandType = "mem64fp"
	mem80fp     operandType = "mem80fp"
	mem16int    operandType = "mem16int"
	mem32int    operandType = "mem32int"
	mem64int    operandType = "mem64int"
	st          operandType = "st"
	st0         operandType = "st0"
	st1         operandType = "st1"
	st2         operandType = "st2"
	st3         operandType = "st3"
	st4         operandType = "st4"
	st5         operandType = "st5"
	st6         operandType = "st6"
	st7         operandType = "st7"
	mmx         operandType = "mmx"
	mmxMem32    operandType = "mmxMem32"
	mmxMem64    operandType = "mmxMem64"
	xmm         operandType = "xmm"
	xmmMem32    operandType = "xmmMem32"
	xmmMem64    operandType = "xmmMem64"
	xmmMem128   operandType = "xmmMem128"
	ymm         operandType = "ymm"
	mem256      operandType = "mem256"
	ymmMem256   operandType = "ymmMem256"
	bnd         operandType = "bnd"
	mib         operandType = "mib"
	mem512      operandType = "mem512"
	zmmMem512   operandType = "zmmMem512"
	// TODO: <XXM0>, <YMM0>, {k1}{z}, {k1}, k1, mv, vm32{x,y, z}, vm64{x,y, z}, zmm/m512/m32bcst, zmm/m512/m64bcst, <ZMM0>, {er}, {sae}, SRC1, SRC2, SRC3, SRC, DST

	reg    operandType = "reg"    // group of reg*
	mem    operandType = "mem"    // group of mem*
	regMem operandType = "regMem" // group of regMem*
	imm    operandType = "imm"    // group of imm*

	//
)

// operand structure
type operand struct {
	t operandType // operand type
	v uint        // operand value
	m []operand   // mem operands
	l bool        // is label
}
