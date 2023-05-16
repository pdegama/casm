// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// register

package x86_64

import (
	"fmt"
	"strings"
)

// register structure
type register struct {
	name             string // register name
	globleIndex      int    // register globle index for this assembler
	bitSize          int    // register bit size 8, 16, 32, or 64-bit
	index            int    // register index
	baseOffset       int    // register base offset according to modrm reg field, modrm r/m field and +rb, +rw, +rd, or +rq value
	onlyValidIn64Bit bool   // register is only valid in 64-bit mode
}

// check or get register form register name
func isRegister(regName string) (bool, register) {

	regName = strings.ToUpper(regName) // convert uppercase

	// loop of registers
	for _, reg := range registers {
		if reg.name == regName {
			// if register is match
			return true, reg
		}
	}

	// if register not found
	return false, register{}
}

// get register info
func registerInfo(reg int) (register, error) {

	// loop of register
	for _, r := range registers {
		if r.globleIndex == reg {
			// if register is match
			return r, nil
		}
	}

	return register{}, fmt.Errorf("modrm is not register")
}

// get register operand type
func getRegisterOperandType(reg register) operandType {
	switch reg.bitSize {
	case 8:
		// if register bit size is 8
		return reg8
	case 16:
		// if register bit size is 16
		return reg16
	case 32:
		// if register bit size is 32
		return reg32
	case 64:
		// if register bit size is 64
		return reg64
	default:
		// if register bit size is invalid
		panic("invalid register")
	}
}

// register is valid...
func registerIsValid(r register, bitMode int) error {

	if r.onlyValidIn64Bit {
		/*
			register is only valid
			in 64-bit mode
		*/
		if bitMode != 64 {
			/*
				and bit mode is not 64
				then return error
			*/
			return fmt.Errorf("%v register is only support in 64-bit mode", r.name)
		}
	}

	return nil
}

// register is valide in mem...
func registerIsValidInMem(r register, bitMode int) error {

	if r.onlyValidIn64Bit {
		// if register is only valid in 64-bit mode
		if bitMode != 64 {
			/*
				and current bit mode is not
				64 then return error
			*/
			return fmt.Errorf("%v register is only valid in 64-bit mode", r.name)
		}
	}

	switch r.bitSize {
	case 8:
		/*
			if register is 8-bit register then
			8-bit register is not addressable
			in any bit mode then return error
		*/
		return fmt.Errorf("%v register is not addressable in any bit mode", r.name)
	case 16:
		/*
			if register is 16-bit register then
			this register is addressable in
			16 and 32 bit mode, is not addressable
			in 64-bit mode
		*/
		if bitMode == 64 {
			// if bit mode is 64 then return error
			return fmt.Errorf("%v register is not addressable in 64-bit mode", r.name)
		}

	case 32:
		/*
			if register is 32-bit register then
			this register is addressable in
			any (16, 32 and 64) bit mode
		*/
	case 64:
		/*
			if register is 64-bit register then
			this register is addressable in
			only 64 bit mode
		*/
		if bitMode != 64 {
			// if bit mode is not 64 then return error
			return fmt.Errorf("%v register is only addressable in 64-bit mode", r.name)
		}
	}

	return nil
}
