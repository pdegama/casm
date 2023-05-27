// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// architecture opcode find

package x86_64

import "fmt"

// find arch opcode
func findArchOpcode(inst *instruction, bitMode int) []archOpcode {

	validOpcodes := []archOpcode{}        // valid insts opcodes stack
	validPerfectOpcodes := []archOpcode{} // valid insts opcodes stack

	// loop of instruction opcode
	for _, opcode := range archOpcodeList {

		// match mnemonic
		if opcode.name == inst.name {

			if bitMode == 32 || bitMode == 16 {
				// if bit mode is 32 or 16
				if !opcode.valid32BitMode {
					// this opcode is not valid in 32-bit or 16-bit mode
					continue
				}
			}

			if bitMode == 64 {
				// if bit mode is 64
				if !opcode.valid64BitMode {
					// this opcode is not valid in 64-bit mode
					continue
				}
			}

			operTypeMatch := true       // operand type is match
			operTypePerfecMatch := true // operand type is match

			// operands type
			if len(opcode.operands) == len(inst.operands) { // if size of operands match

				// loop of opcode operand
				for opeIndex, opeType := range opcode.operands {

					// check operand is valid or not valid
					switch operandIsValid(&opeType, &inst.operands[opeIndex]) {
					case operandMatch:
						/*
							operand type match but not perfect then
							this inst opcode is not perfect
						*/
						operTypePerfecMatch = false
					case operandNotMatch:
						/*
							if any operand type is not match then
							this inst id invalid and operTypeMatch
							set false
						*/
						operTypeMatch = false
						operTypePerfecMatch = false
					}
				}

				// operand all type match then append to stack
				if operTypeMatch {
					// if operTypeMatch is true then all type match
					if operTypePerfecMatch {
						/*
							if operTypePerfecMatch is true then
							this opcode is perfect match, and
							append to validPerfectOpcodes
						*/
						validPerfectOpcodes = append(validPerfectOpcodes, opcode)
						fmt.Println(">>>", opcode)
					} else {
						/*
							if opecode is not perfect match then
							append to validOpcodes
						*/
						validOpcodes = append(validOpcodes, opcode)
						fmt.Println("---", opcode)
					}
				}

			}

		}
	}

	// filter opcodes
	filterOpcodeImm(inst, &validOpcodes)

	// append valid arch opcode to valid perfect arch opcode
	validPerfectOpcodes = append(validPerfectOpcodes, validOpcodes...)

	return validPerfectOpcodes // return opcodes
}

// operand match type
const (
	operandNotMatch     = iota // operand type not match
	operandPerfectMatch        // operand type perfectly match
	operandMatch               // operand type match
)

// check operand type is valid
func operandIsValid(withOperand *archOperand, thisOperand *operand) int {

	/*
		thisOperand is operand, to be check with withOperand
		withOperand is operand type and value
	*/

	if withOperand.v != anyValue {
		// if operand have fix value
		if withOperand.v != int(thisOperand.v) {
			// fix value not match return operandNotMatch
			return operandNotMatch
		}
	}

	if withOperand.t == thisOperand.t {
		// if withOperand and thisOperand type is same
		return operandPerfectMatch
	}

	/*
		check imm type
	*/

	if thisOperand.t == imm {

		//fmt.Println(parseImmType(thisOperand.operand))
		if withOperand.t == parseImmType(thisOperand.v) {
			if !thisOperand.l {
				// if this operand is not lable
				return operandPerfectMatch
			}
		}

		switch withOperand.t {
		case imm8, imm16, imm32, imm64:
			/*
				if withOperand is imm8, imm16, imm32
				or imm64, then accept imm
			*/
			return operandMatch
		}

	}

	/*
		check rel type
	*/

	if thisOperand.t == imm {

		switch parseImmType(thisOperand.v) {
		case imm8:
			// if value type is imm8
			if withOperand.t == rel8 {
				if !thisOperand.l {
					// if this operand is not lable
					return operandPerfectMatch
				}
			}

		case imm16:
			// if value type is imm16
			if withOperand.t == rel16 {
				if !thisOperand.l {
					// if this operand is not lable
					return operandPerfectMatch
				}
			}

		case imm32:
			// if value type is imm32
			if withOperand.t == rel32 {
				if !thisOperand.l {
					// if this operand is not lable
					return operandPerfectMatch
				}
			}
		}

		switch withOperand.t {
		case rel8, rel16, rel32:
			/*
				if withOperand is rel8, rel16, rel32
				or imm64, then accept imm
			*/
			return operandMatch
		}

	} else {

		switch thisOperand.t {
		case imm8:
			// if value type is imm8
			if withOperand.t == rel8 {
				return operandPerfectMatch
			}

		case imm16:
			// if value type is imm16
			if withOperand.t == rel16 {
				return operandPerfectMatch
			}

		case imm32:
			// if value type is imm32
			if withOperand.t == rel32 {
				return operandPerfectMatch
			}
		}

	}

	/*
		check mem
	*/

	if thisOperand.t == mem {
		if isMemOperand(withOperand.t) {
			/*
				if withOperand is mem operand then
				accept mem
			*/
			return operandPerfectMatch
		}

	}

	/*
		check reg
	*/

	if withOperand.t == regMem8 && thisOperand.t == reg8 {
		// if withOperand is regMem8 then accept reg8
		return operandPerfectMatch
	}

	if withOperand.t == regMem16 && thisOperand.t == reg16 {
		// if withOperand is regMem16 then accept reg16
		return operandPerfectMatch
	}

	if withOperand.t == regMem32 && thisOperand.t == reg32 {
		// if withOperand is regMem32 then accept reg32
		return operandPerfectMatch
	}

	if withOperand.t == regMem64 && thisOperand.t == reg64 {
		// if withOperand is regMem64 then accept reg64
		return operandPerfectMatch
	}

	// operand not match
	return operandNotMatch
}

// filter opcodes according imm
func filterOpcodeImm(inst *instruction, opcodes *[]archOpcode) {

	immPos := -1                               // imm operand position
	immIsLabel := false                        // imm operand is label
	for operPos, oper := range inst.operands { // loop of instruction operands
		if oper.t == imm {
			/*
				if oprand is imm set immPos
				is operand posation (index)
			*/
			immPos = operPos

			if oper.l {
				/*
					if operand is label then
					immIsLabel set true
				*/
				immIsLabel = true
			}
			break
		}
	}

	if immPos == -1 {
		/*
			if immPos is -1 then no any imm
			operand in this instruction
		*/
		return
	}

	for _, opcode := range *opcodes {
		fmt.Println(opcode)
	}
	fmt.Println("---------")

	// sorting opcodes
	for i := 0; i < len(*opcodes); i++ {
		for j := 0; j < i; j++ {
			if isLessThanImmRel((*opcodes)[j].operands[immPos].t, (*opcodes)[i].operands[immPos].t) {
				tmpOper := (*opcodes)[i]
				(*opcodes)[i] = (*opcodes)[j]
				(*opcodes)[j] = tmpOper

			}

		}

	}

	for _, opcode := range *opcodes {
		fmt.Println(opcode)
	}
	fmt.Println("---------")

	if immIsLabel {
		/*
			if imm operand is lable then
			no more sorting this opcodes
		*/
		return
	}

	greatThanOpcodes := []archOpcode{}
	lessThanOpcodes := []archOpcode{}

	for _, opcode := range *opcodes {
		// loop of opcodes

		/*
			parse imm value type and check
			instruction imm is less then
			opcode imm type
		*/
		if isLessThanImmRel(parseImmType(inst.operands[immPos].v), opcode.operands[immPos].t) {
			/*
				if instruction imm type is less
				then opcode imm then opcode
				append greatThanOpcode
			*/
			greatThanOpcodes = append(greatThanOpcodes, opcode)
		} else {
			/*
				if instruction imm type is great
				then opcode imm then opcode
				append lessThanOpcodes
			*/
			lessThanOpcodes = append(lessThanOpcodes, opcode)
		}
	}

	for _, opcode := range greatThanOpcodes {
		fmt.Println(opcode)
	}
	fmt.Println("g====")

	for i := 0; i < len(greatThanOpcodes)/2; i++ {
		// reverse greatThanOpcodes
		tmpOpcode := greatThanOpcodes[i]
		greatThanOpcodes[i] = greatThanOpcodes[len(greatThanOpcodes)-1-i]
		greatThanOpcodes[len(greatThanOpcodes)-1-i] = tmpOpcode
	}

	for _, opcode := range greatThanOpcodes {
		fmt.Println(opcode)
	}
	fmt.Println("g====")

	// make temp opcodes map after assign opcodes
	tmpOpcodes := append(greatThanOpcodes, lessThanOpcodes...)
	copy(*opcodes, tmpOpcodes) // copy temp opcodes to *opcodes

	for _, opcode := range *opcodes {
		fmt.Println(opcode)
	}
	fmt.Println("---------")

}

// is a imm or rel is smaller then b imm or rel
func isLessThanImmRel(a operandType, b operandType) bool {

	// a imm is less than b imm?

	switch b {
	case imm16:
		// if b is imm16

		if a == imm8 {
			/*
				and a is imm8 then return true
				beacuse...
			*/
			return true
		}

		if a == rel8 {
			/*
				and a is rel8 then return true
				beacuse...
			*/
			return true
		}

	case imm32:
		// if b is imm32

		if a == imm8 || a == imm16 {
			/*
				and a is imm8 or imm16 then return
				true beacuse...
			*/
			return true
		}

		if a == rel8 || a == rel16 {
			/*
				and a is rel8 or rel16 then return
				true beacuse...
			*/
			return true
		}

	case imm64:
		// if b is imm64

		if a == imm8 || a == imm16 || a == imm32 {
			/*
				and a is imm8, imm16 or imm32 then
				return true beacuse...
			*/
			return true
		}

		if a == rel8 || a == rel16 || a == rel32 {
			/*
				and a is rel8, rel16 or rel64 then
				return true beacuse...
			*/
			return true
		}

	case rel16:
		// if b is rel16

		if a == rel8 {
			/*
				and a is rel8 then return true
				beacuse...
			*/
			return true
		}

		if a == imm8 {
			/*
				and a is imm8 then return true
				beacuse...
			*/
			return true
		}

	case rel32:
		// if b is rel32

		if a == rel8 || a == rel16 {
			/*
				and a is rel8 or rel16 then return
				true beacuse...
			*/
			return true
		}

		if a == imm8 || a == imm16 {
			/*
				and a is imm8 or imm16 then return
				true beacuse...
			*/
			return true
		}

	}

	return false

}
