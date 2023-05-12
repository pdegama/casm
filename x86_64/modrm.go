// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// modrm

package x86_64

import "fmt"

// add modRM byte
func addModRM(opcode *archOpcode, inst *instruction, bitMode int) ([]uint8, error) {

	// calculate modRM and return
	if len(inst.operands) != 2 {
		return nil, fmt.Errorf("internal error: todo: operand size is not two")
	}

	modRMrmOper := &operand{t: undefinedOperand}  // modrm r/m operand
	modRMregOper := &operand{t: undefinedOperand} // modrm reg operand

	// loop of arch operands
	for aOperIndex, aOper := range opcode.operands {
		if isMemoryOperand(aOper.t) {
			/*
				if arch operand is memory operand then
				assign inst operand to modRMrmOper
			*/
			modRMrmOper = &inst.operands[aOperIndex]
		}
		if isRegOperand(aOper.t) {
			/*
				if arch operand is register operand then
				assign inst operand to modRMregOper
			*/
			modRMregOper = &inst.operands[aOperIndex]
		}
	}

	// get reg field
	regField, err := modRMregField(*modRMregOper)
	if err != nil {
		return nil, err
	}

	// calc modrm
	modRMByte, err := calcModRM(modRMrmOper, regField, bitMode)
	if err != nil {
		return nil, err
	}

	fmt.Println(modRMrmOper, modRMregOper, regField)

	return modRMByte, nil
}

// calc modrm
func calcModRM(rmOper *operand, regField int, bitMode int) ([]uint8, error) {

	modrmBytes := []uint8{} // modrm bytes

	if rmOper.t == mem {
		// todo: modrm mem operand support
		fmt.Println(rmOper.m)
		return nil, fmt.Errorf("todo: modrm mem opernad not support")
	}

	if isRegOperand(rmOper.t) {
		// if rm operand is register

		// get r/m register info
		regInfo, err := registerInfo(int(rmOper.v))
		if err != nil {
			// if error then return error
			return nil, err
		}

		/*
			register base offset is r
		*/

		modField := 0b11              // modrm mod field
		rmField := regInfo.baseOffset // modrm r/m field
		regField := regField          // modrm reg field

		modRMByte := modField<<6 | regField<<3 | rmField<<0

		modrmBytes = append(modrmBytes, uint8(modRMByte))

		return modrmBytes, nil

	}

	return nil, nil
}

// modrm reg field
func modRMregField(oper operand) (int, error) {

	if !isRegOperand(oper.t) {
		// if opernad is not register then return error
		return 0, fmt.Errorf("is not modrm reg type")
	}

	// get register info
	regInfo, err := registerInfo(int(oper.v))
	if err != nil {
		// if error then return error
		return 0, nil
	}

	/*
		return register base offset because
		register base offset is modrm reg field
	*/

	return regInfo.baseOffset, nil
}

// is memory operand type
func isMemoryOperand(operType operandType) bool {

	switch operType {
	case mem, mem8, mem16, mem32, mem64, mem128:
		/*
			retrun true because mem, mem8, mem16,
			mem32, mem64, mem128:
		*/
		return true
	case regMem, regMem8, regMem16, regMem32, regMem64:
		/*
			return true bacause regMem, regMem8,
			regMem16, regMem32, regMem64
		*/
		return true
	}

	/*
		other return false
	*/
	return false
}

// is reg operand type
func isRegOperand(operType operandType) bool {

	switch operType {
	case reg, reg8, reg16, reg32, reg64:
		/*
			retrun true because reg, reg8, reg16,
			reg32, reg64
		*/
		return true
	}

	/*
		other return false
	*/
	return false
}
