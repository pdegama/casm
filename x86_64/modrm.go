// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// modrm

package x86_64

import "fmt"

// calculate modRM
func calcModRM(opcode *archOpcode, inst *instruction) (uint8, error) {

	// calculate modRM and return
	if len(inst.operands) != 2 {
		return 0x00, fmt.Errorf("internal error: todo: operand size is not two")
	}

	modRMrmOper := &operand{t: undefinedOperand}  // modrm r/m operand
	modRMregOper := &operand{t: undefinedOperand} // modrm reg operand

	for aOperIndex, aOper := range opcode.operands {
		if isMemoryOperand(aOper.t) {
			modRMrmOper = &inst.operands[aOperIndex]
		}
		if isRegOperand(aOper.t) {
			modRMregOper = &inst.operands[aOperIndex]
		}
	}

	if !(isRegOperand(modRMrmOper.t) && isRegOperand(modRMregOper.t)) {
		// todo: modrm mem support
		return 0x00, fmt.Errorf("internal error: modrm support only register not memory todo: memory support")
	}

	

	fmt.Println(modRMrmOper, modRMregOper)

	return 0x00, nil
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
