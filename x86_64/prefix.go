// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// opcode prefix

package x86_64

import "fmt"

// rex prefix
type rexPrefix struct {
	a bool // rex prefix is required
	w bool // 64-bit operand size
	r bool // modrm reg field, permitting access to 16 registers
	x bool // SIB index field, permitting access to 16 registers
	b bool // modrm r/m field, SIB base field , or opcode reg field, permitting access to 16 registers
}

// override prefix
type overridePrefix struct {
	operand bool // operand size override
	address bool // address size override
}

// prefix sructure
type prefix struct {
	override overridePrefix // override prefix
	rex      rexPrefix      // rex prefix
}

// check operand override prefix
func checkOperandOverride(ope *operand, bitMode int, pf *prefix) error {

	effectiveBits, err := effectiveBitSize(ope) // effective bit size
	if err != nil {
		return err
	}

	switch bitMode {
	case 16:
		// if bit mode is 16

		if effectiveBits == 32 {
			// if effective bit size is 32

			pf.override.operand = true
		}

	case 32:
		// if bit mode is 32

		if effectiveBits == 16 {
			// if effective bit size is 16

			pf.override.operand = true
		}

	case 64:
		// if bit mode is 64

		if effectiveBits == 16 {
			// if effective bit size is 16

			pf.override.operand = true
		} else if effectiveBits == 64 {
			// if effective bit size is 16

			pf.rex.w = true // set rex.w true
		}

	}

	return nil

}

// check address  override prefix
func checkAddressOverride(ope *operand, bitMode int, pf *prefix) error {

	effectiveBits, err := effectiveBitSize(ope) // effective bit size
	if err != nil {
		return err
	}

	switch bitMode {
	case 16:
		// if bit mode is 16

		if effectiveBits == 32 {
			// if effective bit size is 32

			pf.override.address = true
		}

	case 32:
		// if bit mode is 32

		if effectiveBits == 16 {
			// if effective bit size is 16

			pf.override.address = true
		}

	case 64:
		// if bit mode is 64

		if effectiveBits == 32 {
			// if effective bit size is 32

			pf.override.address = true
		}

	}

	return nil

}

// get effective bit size
func effectiveBitSize(ope *operand) (int, error) {
	if isRegOperand(ope.t) {
		// if operand is register
		regInfo, err := registerInfo(int(ope.v))
		if err != nil {
			return 0, err
		}

		return regInfo.bitSize, nil
	}

	return 0, fmt.Errorf("todo: imm")
}

// gen prefix
func genPrefix(p *prefix) []uint8 {

	prefixByte := []uint8{}

	if p.override.operand {
		// if operand size override
		prefixByte = append(prefixByte, 0x66)
	}

	if p.override.address {
		// if operand size override
		prefixByte = append(prefixByte, 0x67)
	}

	return prefixByte
}
