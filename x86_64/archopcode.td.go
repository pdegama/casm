// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// architecture opcode table

package x86_64

// arch opcode list
var archOpcodeList []archOpcode = []archOpcode{
	{name: "AAA", operands: []archOperand{}, opcode: []int{0x37}, valid32BitMode: true, valid64BitMode: false},
	{name: "AAD", operands: []archOperand{}, opcode: []int{0xD5, 0x0A}, valid32BitMode: true, valid64BitMode: false},
	{name: "AAD", operands: []archOperand{{imm8, anyValue}}, opcode: []int{0xD5, valIB}, valid32BitMode: true, valid64BitMode: false},
	{name: "AAM", operands: []archOperand{}, opcode: []int{0xD4, 0x0A}, valid32BitMode: true, valid64BitMode: false},
	{name: "AAM", operands: []archOperand{{imm8, anyValue}}, opcode: []int{0xD4, valIB}, valid32BitMode: true, valid64BitMode: false},
	{name: "AAS", operands: []archOperand{}, opcode: []int{0x3F}, valid32BitMode: true, valid64BitMode: false},
	{name: "ADC", operands: []archOperand{{reg8, al}, {imm8, anyValue}}, opcode: []int{0x14, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADC", operands: []archOperand{{reg16, ax}, {imm16, anyValue}}, opcode: []int{0x15, valIW}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADC", operands: []archOperand{{reg32, eax}, {imm32, anyValue}}, opcode: []int{0x15, valID}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADC", operands: []archOperand{{reg64, rax}, {imm32, anyValue}}, opcode: []int{0x15, valID}, valid32BitMode: false, valid64BitMode: true},
	{name: "ADC", operands: []archOperand{{regMem16, anyValue}, {imm16, anyValue}}, opcode: []int{0x81, modRM2, valIW}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADC", operands: []archOperand{{regMem16, anyValue}, {imm8, anyValue}}, opcode: []int{0x83, modRM2, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADC", operands: []archOperand{{regMem16, anyValue}, {reg16, anyValue}}, opcode: []int{0x11, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADC", operands: []archOperand{{regMem32, anyValue}, {imm32, anyValue}}, opcode: []int{0x81, modRM2, valID}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADC", operands: []archOperand{{regMem32, anyValue}, {imm8, anyValue}}, opcode: []int{0x83, modRM2, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADC", operands: []archOperand{{regMem32, anyValue}, {reg32, anyValue}}, opcode: []int{0x11, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADC", operands: []archOperand{{regMem64, anyValue}, {imm32, anyValue}}, opcode: []int{0x81, modRM2, valID}, valid32BitMode: false, valid64BitMode: true},
	{name: "ADC", operands: []archOperand{{regMem64, anyValue}, {imm8, anyValue}}, opcode: []int{0x83, modRM2, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "ADC", operands: []archOperand{{regMem64, anyValue}, {reg64, anyValue}}, opcode: []int{0x11, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "ADC", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0x80, modRM2, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADC", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0x80, modRM2, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "ADC", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x10, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADC", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x10, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "ADC", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0x13, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADC", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0x13, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADC", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{0x13, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "ADC", operands: []archOperand{{reg8, anyValue}, {regMem8, anyValue}}, opcode: []int{0x12, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADC", operands: []archOperand{{reg8, anyValue}, {regMem8, anyValue}}, opcode: []int{0x12, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "ADCX", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0x66, 0x0F, 0x38, 0xF6, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADCX", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{0x66, 0x0F, 0x38, 0xF6, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "ADD", operands: []archOperand{{reg8, al}, {imm8, anyValue}}, opcode: []int{0x04, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADD", operands: []archOperand{{reg16, ax}, {imm16, anyValue}}, opcode: []int{0x05, valIW}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADD", operands: []archOperand{{reg32, eax}, {imm32, anyValue}}, opcode: []int{0x05, valID}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADD", operands: []archOperand{{reg64, rax}, {imm32, anyValue}}, opcode: []int{0x05, valID}, valid32BitMode: false, valid64BitMode: true},
	{name: "ADD", operands: []archOperand{{regMem16, anyValue}, {imm16, anyValue}}, opcode: []int{0x81, modRM0, valIW}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADD", operands: []archOperand{{regMem16, anyValue}, {imm8, anyValue}}, opcode: []int{0x83, modRM0, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADD", operands: []archOperand{{regMem16, anyValue}, {reg16, anyValue}}, opcode: []int{0x01, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADD", operands: []archOperand{{regMem32, anyValue}, {imm32, anyValue}}, opcode: []int{0x81, modRM0, valID}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADD", operands: []archOperand{{regMem32, anyValue}, {imm8, anyValue}}, opcode: []int{0x83, modRM0, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADD", operands: []archOperand{{regMem32, anyValue}, {reg32, anyValue}}, opcode: []int{0x01, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADD", operands: []archOperand{{regMem64, anyValue}, {imm32, anyValue}}, opcode: []int{0x81, modRM0, valID}, valid32BitMode: false, valid64BitMode: true},
	{name: "ADD", operands: []archOperand{{regMem64, anyValue}, {imm8, anyValue}}, opcode: []int{0x83, modRM0, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "ADD", operands: []archOperand{{regMem64, anyValue}, {reg64, anyValue}}, opcode: []int{0x01, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "ADD", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0x80, modRM0, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADD", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0x80, modRM0, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "ADD", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x00, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADD", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x00, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "ADD", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0x03, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADD", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0x03, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADD", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{0x03, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "ADD", operands: []archOperand{{reg8, anyValue}, {regMem8, anyValue}}, opcode: []int{0x02, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADD", operands: []archOperand{{reg8, anyValue}, {regMem8, anyValue}}, opcode: []int{0x02, modRM}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	{name: "ADOX", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0xF3, 0x0F, 0x38, 0xF6, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "ADOX", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{0xF3, 0x0F, 0x38, 0xF6, modRM}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand xmm1
	// todo: operand xmm
	// todo: operand xmm
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm
	// todo: operand xmm
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	{name: "AND", operands: []archOperand{{reg8, al}, {imm8, anyValue}}, opcode: []int{0x24, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "AND", operands: []archOperand{{reg16, ax}, {imm16, anyValue}}, opcode: []int{0x25, valIW}, valid32BitMode: true, valid64BitMode: true},
	{name: "AND", operands: []archOperand{{reg32, eax}, {imm32, anyValue}}, opcode: []int{0x25, valID}, valid32BitMode: true, valid64BitMode: true},
	{name: "AND", operands: []archOperand{{reg64, rax}, {imm32, anyValue}}, opcode: []int{0x25, valID}, valid32BitMode: false, valid64BitMode: true},
	{name: "AND", operands: []archOperand{{regMem16, anyValue}, {imm16, anyValue}}, opcode: []int{0x81, modRM4, valIW}, valid32BitMode: true, valid64BitMode: true},
	{name: "AND", operands: []archOperand{{regMem16, anyValue}, {imm8, anyValue}}, opcode: []int{0x83, modRM4, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "AND", operands: []archOperand{{regMem16, anyValue}, {reg16, anyValue}}, opcode: []int{0x21, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "AND", operands: []archOperand{{regMem32, anyValue}, {imm32, anyValue}}, opcode: []int{0x81, modRM4, valID}, valid32BitMode: true, valid64BitMode: true},
	{name: "AND", operands: []archOperand{{regMem32, anyValue}, {imm8, anyValue}}, opcode: []int{0x83, modRM4, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "AND", operands: []archOperand{{regMem32, anyValue}, {reg32, anyValue}}, opcode: []int{0x21, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "AND", operands: []archOperand{{regMem64, anyValue}, {imm32, anyValue}}, opcode: []int{0x81, modRM4, valID}, valid32BitMode: false, valid64BitMode: true},
	{name: "AND", operands: []archOperand{{regMem64, anyValue}, {imm8, anyValue}}, opcode: []int{0x83, modRM4, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "AND", operands: []archOperand{{regMem64, anyValue}, {reg64, anyValue}}, opcode: []int{0x21, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "AND", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0x80, modRM4, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "AND", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0x80, modRM4, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "AND", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x20, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "AND", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x20, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "AND", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0x23, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "AND", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0x23, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "AND", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{0x23, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "AND", operands: []archOperand{{reg8, anyValue}, {regMem8, anyValue}}, opcode: []int{0x22, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "AND", operands: []archOperand{{reg8, anyValue}, {regMem8, anyValue}}, opcode: []int{0x22, modRM}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand r32a
	// todo: operand r64a
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	{name: "ARPL", operands: []archOperand{{regMem16, anyValue}, {reg16, anyValue}}, opcode: []int{0x63, modRM}, valid32BitMode: true, valid64BitMode: false},
	// todo: operand r32a
	// todo: operand r64a
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	{name: "BLSI", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{todoOpcode}, valid32BitMode: true, valid64BitMode: true},
	{name: "BLSI", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{todoOpcode}, valid32BitMode: false, valid64BitMode: true},
	{name: "BLSMSK", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{todoOpcode}, valid32BitMode: true, valid64BitMode: true},
	{name: "BLSMSK", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{todoOpcode}, valid32BitMode: false, valid64BitMode: true},
	{name: "BLSR", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{todoOpcode}, valid32BitMode: true, valid64BitMode: true},
	{name: "BLSR", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{todoOpcode}, valid32BitMode: false, valid64BitMode: true},
	{name: "BNDCL", operands: []archOperand{{bnd, anyValue}, {regMem32, anyValue}}, opcode: []int{0xF3, 0x0F, 0x1A, modRM}, valid32BitMode: true, valid64BitMode: false},
	{name: "BNDCL", operands: []archOperand{{bnd, anyValue}, {regMem64, anyValue}}, opcode: []int{0xF3, 0x0F, 0x1A, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "BNDCN", operands: []archOperand{{bnd, anyValue}, {regMem32, anyValue}}, opcode: []int{0xF2, 0x0F, 0x1B, modRM}, valid32BitMode: true, valid64BitMode: false},
	{name: "BNDCN", operands: []archOperand{{bnd, anyValue}, {regMem64, anyValue}}, opcode: []int{0xF2, 0x0F, 0x1B, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "BNDCU", operands: []archOperand{{bnd, anyValue}, {regMem32, anyValue}}, opcode: []int{0xF2, 0x0F, 0x1A, modRM}, valid32BitMode: true, valid64BitMode: false},
	{name: "BNDCU", operands: []archOperand{{bnd, anyValue}, {regMem64, anyValue}}, opcode: []int{0xF2, 0x0F, 0x1A, modRM}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand bnd1
	{name: "BNDMK", operands: []archOperand{{bnd, anyValue}, {mem32, anyValue}}, opcode: []int{0xF3, 0x0F, 0x1B, modRM}, valid32BitMode: true, valid64BitMode: false},
	{name: "BNDMK", operands: []archOperand{{bnd, anyValue}, {mem64, anyValue}}, opcode: []int{0xF3, 0x0F, 0x1B, modRM}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand bnd1
	// todo: operand bnd1
	// todo: operand bnd1/m128
	// todo: operand bnd1/m64
	// todo: operand bnd1
	{name: "BOUND", operands: []archOperand{{reg16, anyValue}, {mem16n16, anyValue}}, opcode: []int{0x62, modRM}, valid32BitMode: true, valid64BitMode: false},
	{name: "BOUND", operands: []archOperand{{reg32, anyValue}, {mem32n32, anyValue}}, opcode: []int{0x62, modRM}, valid32BitMode: true, valid64BitMode: false},
	{name: "BSF", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0x0F, 0xBC, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "BSF", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0x0F, 0xBC, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "BSF", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{0x0F, 0xBC, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "BSR", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0x0F, 0xBD, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "BSR", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0x0F, 0xBD, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "BSR", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{0x0F, 0xBD, modRM}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand r16op
	{name: "BSWAP", operands: []archOperand{{reg32, anyValue}}, opcode: []int{0x0F, 0xC8, plusRD}, valid32BitMode: true, valid64BitMode: true},
	{name: "BSWAP", operands: []archOperand{{reg64, anyValue}}, opcode: []int{0x0F, 0xC8, plusRD}, valid32BitMode: false, valid64BitMode: true},
	{name: "BT", operands: []archOperand{{regMem16, anyValue}, {imm8, anyValue}}, opcode: []int{0x0F, 0xBA, modRM4, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "BT", operands: []archOperand{{regMem16, anyValue}, {reg16, anyValue}}, opcode: []int{0x0F, 0xA3, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "BT", operands: []archOperand{{regMem32, anyValue}, {imm8, anyValue}}, opcode: []int{0x0F, 0xBA, modRM4, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "BT", operands: []archOperand{{regMem32, anyValue}, {reg32, anyValue}}, opcode: []int{0x0F, 0xA3, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "BT", operands: []archOperand{{regMem64, anyValue}, {imm8, anyValue}}, opcode: []int{0x0F, 0xBA, modRM4, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "BT", operands: []archOperand{{regMem64, anyValue}, {reg64, anyValue}}, opcode: []int{0x0F, 0xA3, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "BTC", operands: []archOperand{{regMem16, anyValue}, {imm8, anyValue}}, opcode: []int{0x0F, 0xBA, modRM7, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "BTC", operands: []archOperand{{regMem16, anyValue}, {reg16, anyValue}}, opcode: []int{0x0F, 0xBB, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "BTC", operands: []archOperand{{regMem32, anyValue}, {imm8, anyValue}}, opcode: []int{0x0F, 0xBA, modRM7, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "BTC", operands: []archOperand{{regMem32, anyValue}, {reg32, anyValue}}, opcode: []int{0x0F, 0xBB, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "BTC", operands: []archOperand{{regMem64, anyValue}, {imm8, anyValue}}, opcode: []int{0x0F, 0xBA, modRM7, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "BTC", operands: []archOperand{{regMem64, anyValue}, {reg64, anyValue}}, opcode: []int{0x0F, 0xBB, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "BTR", operands: []archOperand{{regMem16, anyValue}, {imm8, anyValue}}, opcode: []int{0x0F, 0xBA, modRM6, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "BTR", operands: []archOperand{{regMem16, anyValue}, {reg16, anyValue}}, opcode: []int{0x0F, 0xB3, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "BTR", operands: []archOperand{{regMem32, anyValue}, {imm8, anyValue}}, opcode: []int{0x0F, 0xBA, modRM6, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "BTR", operands: []archOperand{{regMem32, anyValue}, {reg32, anyValue}}, opcode: []int{0x0F, 0xB3, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "BTR", operands: []archOperand{{regMem64, anyValue}, {imm8, anyValue}}, opcode: []int{0x0F, 0xBA, modRM6, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "BTR", operands: []archOperand{{regMem64, anyValue}, {reg64, anyValue}}, opcode: []int{0x0F, 0xB3, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "BTS", operands: []archOperand{{regMem16, anyValue}, {imm8, anyValue}}, opcode: []int{0x0F, 0xBA, modRM5, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "BTS", operands: []archOperand{{regMem16, anyValue}, {reg16, anyValue}}, opcode: []int{0x0F, 0xAB, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "BTS", operands: []archOperand{{regMem32, anyValue}, {imm8, anyValue}}, opcode: []int{0x0F, 0xBA, modRM5, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "BTS", operands: []archOperand{{regMem32, anyValue}, {reg32, anyValue}}, opcode: []int{0x0F, 0xAB, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "BTS", operands: []archOperand{{regMem64, anyValue}, {imm8, anyValue}}, opcode: []int{0x0F, 0xBA, modRM5, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "BTS", operands: []archOperand{{regMem64, anyValue}, {reg64, anyValue}}, opcode: []int{0x0F, 0xAB, modRM}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand r32a
	// todo: operand r64a
	{name: "CALL", operands: []archOperand{{regMem16, anyValue}}, opcode: []int{0xFF, modRM2}, valid32BitMode: true, valid64BitMode: false},
	{name: "CALL", operands: []archOperand{{regMem32, anyValue}}, opcode: []int{0xFF, modRM2}, valid32BitMode: true, valid64BitMode: false},
	{name: "CALL", operands: []archOperand{{regMem64, anyValue}}, opcode: []int{0xFF, modRM2}, valid32BitMode: false, valid64BitMode: true},
	{name: "CALL", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0xE8, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "CALL", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0xE8, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "CALL", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0xE8, valCD}, valid32BitMode: false, valid64BitMode: true},
	{name: "CALL_FAR", operands: []archOperand{{mem16_16, anyValue}}, opcode: []int{0xFF, modRM3}, valid32BitMode: true, valid64BitMode: true},
	{name: "CALL_FAR", operands: []archOperand{{mem16_32, anyValue}}, opcode: []int{0xFF, modRM3}, valid32BitMode: true, valid64BitMode: true},
	{name: "CALL_FAR", operands: []archOperand{{mem16_64, anyValue}}, opcode: []int{0xFF, modRM3}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand ptr16:16
	// todo: operand ptr16:32
	{name: "CBW", operands: []archOperand{}, opcode: []int{0x98}, valid32BitMode: true, valid64BitMode: true},
	{name: "CDQ", operands: []archOperand{}, opcode: []int{0x99}, valid32BitMode: true, valid64BitMode: true},
	{name: "CDQE", operands: []archOperand{}, opcode: []int{0x98}, valid32BitMode: false, valid64BitMode: true},
	{name: "CLAC", operands: []archOperand{}, opcode: []int{0x0F, 0x01, 0xCA}, valid32BitMode: true, valid64BitMode: true},
	{name: "CLC", operands: []archOperand{}, opcode: []int{0xF8}, valid32BitMode: true, valid64BitMode: true},
	{name: "CLD", operands: []archOperand{}, opcode: []int{0xFC}, valid32BitMode: true, valid64BitMode: true},
	{name: "CLI", operands: []archOperand{}, opcode: []int{0xFA}, valid32BitMode: true, valid64BitMode: true},
	{name: "CLRSSBSY", operands: []archOperand{{mem64, anyValue}}, opcode: []int{0xF3, 0x0F, 0xAE, modRM6}, valid32BitMode: true, valid64BitMode: true},
	{name: "CLTS", operands: []archOperand{}, opcode: []int{0x0F, 0x06}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMC", operands: []archOperand{}, opcode: []int{0xF5}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMOVA", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0x0F, 0x47, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMOVA", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0x0F, 0x47, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMOVA", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{0x0F, 0x47, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "CMOVAE", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0x0F, 0x43, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMOVAE", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0x0F, 0x43, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMOVAE", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{0x0F, 0x43, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "CMOVB", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0x0F, 0x42, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMOVB", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0x0F, 0x42, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMOVB", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{0x0F, 0x42, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "CMOVBE", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0x0F, 0x46, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMOVBE", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0x0F, 0x46, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMOVBE", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{0x0F, 0x46, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "CMOVC", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0x0F, 0x42, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMOVC", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0x0F, 0x42, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMOVC", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{0x0F, 0x42, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "CMOVE", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0x0F, 0x44, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMOVE", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0x0F, 0x44, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMOVE", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{0x0F, 0x44, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "CMOVG", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0x0F, 0x4F, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMOVG", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0x0F, 0x4F, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMP", operands: []archOperand{{reg8, al}, {imm8, anyValue}}, opcode: []int{0x3C, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMP", operands: []archOperand{{reg16, ax}, {imm16, anyValue}}, opcode: []int{0x3D, valIW}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMP", operands: []archOperand{{reg32, eax}, {imm32, anyValue}}, opcode: []int{0x3D, valID}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMP", operands: []archOperand{{reg64, rax}, {imm32, anyValue}}, opcode: []int{0x3D, valID}, valid32BitMode: false, valid64BitMode: true},
	{name: "CMP", operands: []archOperand{{regMem16, anyValue}, {imm16, anyValue}}, opcode: []int{0x81, modRM7, valIW}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMP", operands: []archOperand{{regMem16, anyValue}, {imm8, anyValue}}, opcode: []int{0x83, modRM7, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMP", operands: []archOperand{{regMem16, anyValue}, {reg16, anyValue}}, opcode: []int{0x39, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMP", operands: []archOperand{{regMem32, anyValue}, {imm32, anyValue}}, opcode: []int{0x81, modRM7, valID}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMP", operands: []archOperand{{regMem32, anyValue}, {imm8, anyValue}}, opcode: []int{0x83, modRM7, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMP", operands: []archOperand{{regMem32, anyValue}, {reg32, anyValue}}, opcode: []int{0x39, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMP", operands: []archOperand{{regMem64, anyValue}, {imm32, anyValue}}, opcode: []int{0x81, modRM7, valID}, valid32BitMode: false, valid64BitMode: true},
	{name: "CMP", operands: []archOperand{{regMem64, anyValue}, {imm8, anyValue}}, opcode: []int{0x83, modRM7, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "CMP", operands: []archOperand{{regMem64, anyValue}, {reg64, anyValue}}, opcode: []int{0x39, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "CMP", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0x80, modRM7, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMP", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0x80, modRM7, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "CMP", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x38, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMP", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x38, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "CMP", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0x3B, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMP", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0x3B, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMP", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{0x3B, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "CMP", operands: []archOperand{{reg8, anyValue}, {regMem8, anyValue}}, opcode: []int{0x3A, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMP", operands: []archOperand{{reg8, anyValue}, {regMem8, anyValue}}, opcode: []int{0x3A, modRM}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand xmm1
	// todo: operand xmm1
	{name: "CMPSB", operands: []archOperand{}, opcode: []int{0xA6}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMPSD", operands: []archOperand{}, opcode: []int{0xA7}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand xmm1
	{name: "CMPSQ", operands: []archOperand{}, opcode: []int{0xA7}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand xmm1
	{name: "CMPSW", operands: []archOperand{}, opcode: []int{0xA7}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMPXCHG", operands: []archOperand{{regMem16, anyValue}, {reg16, anyValue}}, opcode: []int{0x0F, 0xB1, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMPXCHG", operands: []archOperand{{regMem32, anyValue}, {reg32, anyValue}}, opcode: []int{0x0F, 0xB1, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMPXCHG", operands: []archOperand{{regMem64, anyValue}, {reg64, anyValue}}, opcode: []int{0x0F, 0xB1, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "CMPXCHG", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x0F, 0xB0, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CMPXCHG", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x0F, 0xB0, modRM}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand m128
	{name: "CMPXCHG8B", operands: []archOperand{{mem64, anyValue}}, opcode: []int{0x0F, 0xC7, modRM1}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand xmm1
	// todo: operand xmm1
	{name: "CPUID", operands: []archOperand{}, opcode: []int{0x0F, 0xA2}, valid32BitMode: true, valid64BitMode: true},
	{name: "CQO", operands: []archOperand{}, opcode: []int{0x99}, valid32BitMode: false, valid64BitMode: true},
	{name: "CRC32", operands: []archOperand{{reg32, anyValue}, {regMem16, anyValue}}, opcode: []int{0xF2, 0x0F, 0x38, 0xF1, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CRC32", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0xF2, 0x0F, 0x38, 0xF1, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CRC32", operands: []archOperand{{reg32, anyValue}, {regMem8, anyValue}}, opcode: []int{0xF2, 0x0F, 0x38, 0xF0, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "CRC32", operands: []archOperand{{reg32, anyValue}, {regMem8, anyValue}}, opcode: []int{0xF2, 0x0F, 0x38, 0xF0, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "CRC32", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{0xF2, 0x0F, 0x38, 0xF1, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "CRC32", operands: []archOperand{{reg64, anyValue}, {regMem8, anyValue}}, opcode: []int{0xF2, 0x0F, 0x38, 0xF0, modRM}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand xmm
	// todo: operand xmm
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1/m64
	// todo: operand xmm1/m64
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1/m32
	// todo: operand xmm1/m32
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1/m64
	// todo: operand xmm1/m64
	// todo: operand xmm1/m32
	// todo: operand xmm1/m32
	{name: "CWD", operands: []archOperand{}, opcode: []int{0x99}, valid32BitMode: true, valid64BitMode: true},
	{name: "CWDE", operands: []archOperand{}, opcode: []int{0x98}, valid32BitMode: true, valid64BitMode: true},
	{name: "DAA", operands: []archOperand{}, opcode: []int{0x27}, valid32BitMode: true, valid64BitMode: false},
	{name: "DAS", operands: []archOperand{}, opcode: []int{0x2F}, valid32BitMode: true, valid64BitMode: false},
	{name: "DEC", operands: []archOperand{{regMem16, anyValue}}, opcode: []int{0xFF, modRM1}, valid32BitMode: true, valid64BitMode: true},
	{name: "DEC", operands: []archOperand{{regMem32, anyValue}}, opcode: []int{0xFF, modRM1}, valid32BitMode: true, valid64BitMode: true},
	{name: "DEC", operands: []archOperand{{regMem64, anyValue}}, opcode: []int{0xFF, modRM1}, valid32BitMode: false, valid64BitMode: true},
	{name: "DEC", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0xFE, modRM1}, valid32BitMode: true, valid64BitMode: true},
	{name: "DEC", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0xFE, modRM1}, valid32BitMode: false, valid64BitMode: true},
	{name: "DEC", operands: []archOperand{{reg16, anyValue}}, opcode: []int{0x48, plusRW}, valid32BitMode: true, valid64BitMode: false},
	{name: "DEC", operands: []archOperand{{reg32, anyValue}}, opcode: []int{0x48, plusRD}, valid32BitMode: true, valid64BitMode: false},
	{name: "DIV", operands: []archOperand{{regMem16, anyValue}}, opcode: []int{0xF7, modRM6}, valid32BitMode: true, valid64BitMode: true},
	{name: "DIV", operands: []archOperand{{regMem32, anyValue}}, opcode: []int{0xF7, modRM6}, valid32BitMode: true, valid64BitMode: true},
	{name: "DIV", operands: []archOperand{{regMem64, anyValue}}, opcode: []int{0xF7, modRM6}, valid32BitMode: false, valid64BitMode: true},
	{name: "DIV", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0xF6, modRM6}, valid32BitMode: true, valid64BitMode: true},
	{name: "DIV", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0xF6, modRM6}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	{name: "EMMS", operands: []archOperand{}, opcode: []int{0x0F, 0x77}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand <XMM0-2>
	// todo: operand r32 <XMM0-6>
	{name: "ENDBR32", operands: []archOperand{}, opcode: []int{0xF3, 0x0F, 0x1E, 0xFB}, valid32BitMode: true, valid64BitMode: true},
	{name: "ENDBR64", operands: []archOperand{}, opcode: []int{0xF3, 0x0F, 0x1E, 0xFA}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand 0
	// todo: operand 1
	// todo: operand imm8b
	{name: "F2XM1", operands: []archOperand{}, opcode: []int{0xD9, 0xF0}, valid32BitMode: true, valid64BitMode: true},
	{name: "FABS", operands: []archOperand{}, opcode: []int{0xD9, 0xE1}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand ST(0)
	// todo: operand ST(i)
	// todo: operand m32fp
	// todo: operand m64fp
	{name: "FADDP", operands: []archOperand{}, opcode: []int{0xDE, 0xC1}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand ST(i)
	// todo: operand m80bcd
	// todo: operand m80bcd
	{name: "FCHS", operands: []archOperand{}, opcode: []int{0xD9, 0xE0}, valid32BitMode: true, valid64BitMode: true},
	{name: "FCLEX", operands: []archOperand{}, opcode: []int{0x9B, 0xDB, 0xE2}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand ST(0)
	// todo: operand ST(0)
	// todo: operand ST(0)
	// todo: operand ST(0)
	// todo: operand ST(0)
	// todo: operand ST(0)
	// todo: operand ST(0)
	// todo: operand ST(0)
	{name: "FCOM", operands: []archOperand{}, opcode: []int{0xD8, 0xD1}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand ST(i)
	// todo: operand m32fp
	// todo: operand m64fp
	// todo: operand ST(0)
	// todo: operand ST(0)
	{name: "FCOMP", operands: []archOperand{}, opcode: []int{0xD8, 0xD9}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand ST(i)
	// todo: operand m32fp
	// todo: operand m64fp
	{name: "FCOMPP", operands: []archOperand{}, opcode: []int{0xDE, 0xD9}, valid32BitMode: true, valid64BitMode: true},
	{name: "FCOS", operands: []archOperand{}, opcode: []int{0xD9, 0xFF}, valid32BitMode: true, valid64BitMode: true},
	{name: "FDECSTP", operands: []archOperand{}, opcode: []int{0xD9, 0xF6}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand ST(0)
	// todo: operand ST(i)
	// todo: operand m32fp
	// todo: operand m64fp
	{name: "FDIVP", operands: []archOperand{}, opcode: []int{0xDE, 0xF9}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand ST(i)
	// todo: operand ST(0)
	// todo: operand ST(i)
	// todo: operand m32fp
	// todo: operand m64fp
	{name: "FDIVRP", operands: []archOperand{}, opcode: []int{0xDE, 0xF1}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand ST(i)
	// todo: operand ST(i)
	// todo: operand ST(i)
	// todo: operand m16int
	// todo: operand m32int
	// todo: operand m16int
	// todo: operand m32int
	// todo: operand m16int
	// todo: operand m32int
	// todo: operand m16int
	// todo: operand m32int
	// todo: operand m16int
	// todo: operand m32int
	// todo: operand m16int
	// todo: operand m32int
	// todo: operand m64int
	// todo: operand m16int
	// todo: operand m32int
	{name: "FINCSTP", operands: []archOperand{}, opcode: []int{0xD9, 0xF7}, valid32BitMode: true, valid64BitMode: true},
	{name: "FINIT", operands: []archOperand{}, opcode: []int{0x9B, 0xDB, 0xE3}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand m16int
	// todo: operand m32int
	// todo: operand m16int
	// todo: operand m32int
	// todo: operand m64int
	// todo: operand m16int
	// todo: operand m32int
	// todo: operand m64int
	// todo: operand m16int
	// todo: operand m32int
	// todo: operand m16int
	// todo: operand m32int
	// todo: operand ST(i)
	// todo: operand m32fp
	// todo: operand m64fp
	// todo: operand m80fp
	{name: "FLD1", operands: []archOperand{}, opcode: []int{0xD9, 0xE8}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand m2byte
	// todo: operand m14/28byte
	{name: "FLDL2E", operands: []archOperand{}, opcode: []int{0xD9, 0xEA}, valid32BitMode: true, valid64BitMode: true},
	{name: "FLDL2T", operands: []archOperand{}, opcode: []int{0xD9, 0xE9}, valid32BitMode: true, valid64BitMode: true},
	{name: "FLDLG2", operands: []archOperand{}, opcode: []int{0xD9, 0xEC}, valid32BitMode: true, valid64BitMode: true},
	{name: "FLDPI", operands: []archOperand{}, opcode: []int{0xD9, 0xEB}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand ST(0)
	// todo: operand ST(i)
	// todo: operand m32fp
	// todo: operand m64fp
	{name: "FMULP", operands: []archOperand{}, opcode: []int{0xDE, 0xC9}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand ST(i)
	{name: "FNCLEX", operands: []archOperand{}, opcode: []int{0xDB, 0xE2}, valid32BitMode: true, valid64BitMode: true},
	{name: "FNINIT", operands: []archOperand{}, opcode: []int{0xDB, 0xE3}, valid32BitMode: true, valid64BitMode: true},
	{name: "FNOP", operands: []archOperand{}, opcode: []int{0xD9, 0xD0}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand m94/108byte
	// todo: operand m2byte
	// todo: operand m14/28byte
	{name: "FNSTSW", operands: []archOperand{{reg16, ax}}, opcode: []int{0xDF, 0xE0}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand m2byte
	{name: "FPATAN", operands: []archOperand{}, opcode: []int{0xD9, 0xF3}, valid32BitMode: true, valid64BitMode: true},
	{name: "FPREM", operands: []archOperand{}, opcode: []int{0xD9, 0xF8}, valid32BitMode: true, valid64BitMode: true},
	{name: "FPREM1", operands: []archOperand{}, opcode: []int{0xD9, 0xF5}, valid32BitMode: true, valid64BitMode: true},
	{name: "FPTAN", operands: []archOperand{}, opcode: []int{0xD9, 0xF2}, valid32BitMode: true, valid64BitMode: true},
	{name: "FRNDINT", operands: []archOperand{}, opcode: []int{0xD9, 0xFC}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand m94/108byte
	// todo: operand m94/108byte
	{name: "FSCALE", operands: []archOperand{}, opcode: []int{0xD9, 0xFD}, valid32BitMode: true, valid64BitMode: true},
	{name: "FSIN", operands: []archOperand{}, opcode: []int{0xD9, 0xFE}, valid32BitMode: true, valid64BitMode: true},
	{name: "FSINCOS", operands: []archOperand{}, opcode: []int{0xD9, 0xFB}, valid32BitMode: true, valid64BitMode: true},
	{name: "FSQRT", operands: []archOperand{}, opcode: []int{0xD9, 0xFA}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand ST(i)
	// todo: operand m32fp
	// todo: operand m64fp
	// todo: operand m2byte
	// todo: operand m14/28byte
	// todo: operand ST(i)
	// todo: operand m32fp
	// todo: operand m64fp
	// todo: operand m80fp
	{name: "FSTSW", operands: []archOperand{{reg16, ax}}, opcode: []int{0x9B, 0xDF, 0xE0}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand m2byte
	// todo: operand ST(0)
	// todo: operand ST(i)
	// todo: operand m32fp
	// todo: operand m64fp
	{name: "FSUBP", operands: []archOperand{}, opcode: []int{0xDE, 0xE9}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand ST(i)
	// todo: operand ST(0)
	// todo: operand ST(i)
	// todo: operand m32fp
	// todo: operand m64fp
	{name: "FSUBRP", operands: []archOperand{}, opcode: []int{0xDE, 0xE1}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand ST(i)
	{name: "FTST", operands: []archOperand{}, opcode: []int{0xD9, 0xE4}, valid32BitMode: true, valid64BitMode: true},
	{name: "FUCOM", operands: []archOperand{}, opcode: []int{0xDD, 0xE1}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand ST(i)
	// todo: operand ST(0)
	// todo: operand ST(0)
	{name: "FUCOMP", operands: []archOperand{}, opcode: []int{0xDD, 0xE9}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand ST(i)
	{name: "FUCOMPP", operands: []archOperand{}, opcode: []int{0xDA, 0xE9}, valid32BitMode: true, valid64BitMode: true},
	{name: "FWAIT", operands: []archOperand{}, opcode: []int{0x9B}, valid32BitMode: true, valid64BitMode: true},
	{name: "FXAM", operands: []archOperand{}, opcode: []int{0xD9, 0xE5}, valid32BitMode: true, valid64BitMode: true},
	{name: "FXCH", operands: []archOperand{}, opcode: []int{0xD9, 0xC9}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand ST(i)
	// todo: operand m512byte
	// todo: operand m512byte
	// todo: operand m512byte
	// todo: operand m512byte
	{name: "FXTRACT", operands: []archOperand{}, opcode: []int{0xD9, 0xF4}, valid32BitMode: true, valid64BitMode: true},
	{name: "FYL2X", operands: []archOperand{}, opcode: []int{0xD9, 0xF1}, valid32BitMode: true, valid64BitMode: true},
	{name: "FYL2XP1", operands: []archOperand{}, opcode: []int{0xD9, 0xF9}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand xmm1
	// todo: operand xmm1
	{name: "HLT", operands: []archOperand{}, opcode: []int{0xF4}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand xmm1
	// todo: operand xmm1
	{name: "ICEBP", operands: []archOperand{}, opcode: []int{0xF1}, valid32BitMode: true, valid64BitMode: true},
	{name: "IDIV", operands: []archOperand{{regMem16, anyValue}}, opcode: []int{0xF7, modRM7}, valid32BitMode: true, valid64BitMode: true},
	{name: "IDIV", operands: []archOperand{{regMem32, anyValue}}, opcode: []int{0xF7, modRM7}, valid32BitMode: true, valid64BitMode: true},
	{name: "IDIV", operands: []archOperand{{regMem64, anyValue}}, opcode: []int{0xF7, modRM7}, valid32BitMode: false, valid64BitMode: true},
	{name: "IDIV", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0xF6, modRM7}, valid32BitMode: true, valid64BitMode: true},
	{name: "IDIV", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0xF6, modRM7}, valid32BitMode: false, valid64BitMode: true},
	{name: "IMUL", operands: []archOperand{{regMem16, anyValue}}, opcode: []int{0xF7, modRM5}, valid32BitMode: true, valid64BitMode: true},
	{name: "IMUL", operands: []archOperand{{regMem32, anyValue}}, opcode: []int{0xF7, modRM5}, valid32BitMode: true, valid64BitMode: true},
	{name: "IMUL", operands: []archOperand{{regMem64, anyValue}}, opcode: []int{0xF7, modRM5}, valid32BitMode: false, valid64BitMode: true},
	{name: "IMUL", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0xF6, modRM5}, valid32BitMode: true, valid64BitMode: true},
	{name: "IMUL", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0x0F, 0xAF, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "IMUL", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}, {imm16, anyValue}}, opcode: []int{0x69, modRM, valIW}, valid32BitMode: true, valid64BitMode: true},
	{name: "IMUL", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}, {imm8, anyValue}}, opcode: []int{0x6B, modRM, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "IMUL", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0x0F, 0xAF, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "IMUL", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}, {imm32, anyValue}}, opcode: []int{0x69, modRM, valID}, valid32BitMode: true, valid64BitMode: true},
	{name: "IMUL", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}, {imm8, anyValue}}, opcode: []int{0x6B, modRM, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "IMUL", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{0x0F, 0xAF, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "IMUL", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}, {imm32, anyValue}}, opcode: []int{0x69, modRM, valID}, valid32BitMode: false, valid64BitMode: true},
	{name: "IMUL", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}, {imm8, anyValue}}, opcode: []int{0x6B, modRM, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "IN", operands: []archOperand{{reg8, al}, {reg16, dx}}, opcode: []int{0xEC}, valid32BitMode: true, valid64BitMode: true},
	{name: "IN", operands: []archOperand{{reg8, al}, {imm8, anyValue}}, opcode: []int{0xE4, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "IN", operands: []archOperand{{reg16, ax}, {reg16, dx}}, opcode: []int{0xED}, valid32BitMode: true, valid64BitMode: true},
	{name: "IN", operands: []archOperand{{reg16, ax}, {imm8, anyValue}}, opcode: []int{0xE5, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "IN", operands: []archOperand{{reg32, eax}, {reg16, dx}}, opcode: []int{0xED}, valid32BitMode: true, valid64BitMode: true},
	{name: "IN", operands: []archOperand{{reg32, eax}, {imm8, anyValue}}, opcode: []int{0xE5, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "INC", operands: []archOperand{{regMem16, anyValue}}, opcode: []int{0xFF, modRM0}, valid32BitMode: true, valid64BitMode: true},
	{name: "INC", operands: []archOperand{{regMem32, anyValue}}, opcode: []int{0xFF, modRM0}, valid32BitMode: true, valid64BitMode: true},
	{name: "INC", operands: []archOperand{{regMem64, anyValue}}, opcode: []int{0xFF, modRM0}, valid32BitMode: false, valid64BitMode: true},
	{name: "INC", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0xFE, modRM0}, valid32BitMode: true, valid64BitMode: true},
	{name: "INC", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0xFE, modRM0}, valid32BitMode: false, valid64BitMode: true},
	{name: "INC", operands: []archOperand{{reg16, anyValue}}, opcode: []int{0x40, plusRW}, valid32BitMode: true, valid64BitMode: false},
	{name: "INC", operands: []archOperand{{reg32, anyValue}}, opcode: []int{0x40, plusRD}, valid32BitMode: true, valid64BitMode: false},
	{name: "INCSSPD", operands: []archOperand{{reg32, anyValue}}, opcode: []int{0xF3, 0x0F, 0xAE, todoOpcode}, valid32BitMode: true, valid64BitMode: true},
	{name: "INCSSPQ", operands: []archOperand{{reg64, anyValue}}, opcode: []int{0xF3, 0x0F, 0xAE, todoOpcode}, valid32BitMode: false, valid64BitMode: true},
	{name: "INSB", operands: []archOperand{}, opcode: []int{0x6C}, valid32BitMode: true, valid64BitMode: true},
	{name: "INSD", operands: []archOperand{}, opcode: []int{0x6D}, valid32BitMode: true, valid64BitMode: true},
	{name: "INSW", operands: []archOperand{}, opcode: []int{0x6D}, valid32BitMode: true, valid64BitMode: true},
	{name: "INT", operands: []archOperand{{imm8, anyValue}}, opcode: []int{0xCD, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "INT1", operands: []archOperand{}, opcode: []int{0xF1}, valid32BitMode: true, valid64BitMode: true},
	{name: "INT3", operands: []archOperand{}, opcode: []int{0xCC}, valid32BitMode: true, valid64BitMode: true},
	{name: "INTO", operands: []archOperand{}, opcode: []int{0xCE}, valid32BitMode: true, valid64BitMode: false},
	{name: "INVD", operands: []archOperand{}, opcode: []int{0x0F, 0x08}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand m
	// todo: operand m128
	// todo: operand m128
	{name: "IRET", operands: []archOperand{}, opcode: []int{0xCF}, valid32BitMode: true, valid64BitMode: true},
	{name: "IRETD", operands: []archOperand{}, opcode: []int{0xCF}, valid32BitMode: true, valid64BitMode: true},
	{name: "IRETQ", operands: []archOperand{}, opcode: []int{0xCF}, valid32BitMode: false, valid64BitMode: true},
	{name: "JA", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x87, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JA", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x87, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JA", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x87, valCD}, valid32BitMode: false, valid64BitMode: true},
	{name: "JA", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x77, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JAE", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x83, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JAE", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x83, valCD}, valid32BitMode: false, valid64BitMode: true},
	{name: "JAE", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x83, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JAE", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x73, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JB", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x82, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JB", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x82, valCD}, valid32BitMode: false, valid64BitMode: true},
	{name: "JB", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x82, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JB", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x72, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JBE", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x86, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JBE", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x86, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JBE", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x86, valCD}, valid32BitMode: false, valid64BitMode: true},
	{name: "JBE", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x76, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JC", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x82, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JC", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x82, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JC", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x72, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JCXZ", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0xE3, valCB}, valid32BitMode: true, valid64BitMode: false},
	{name: "JE", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x84, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JE", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x84, valCD}, valid32BitMode: false, valid64BitMode: true},
	{name: "JE", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x84, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JE", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x74, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JECXZ", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0xE3, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JG", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x8F, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JG", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x8F, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JG", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x8F, valCD}, valid32BitMode: false, valid64BitMode: true},
	{name: "JG", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x7F, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JGE", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x8D, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JGE", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x8D, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JGE", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x8D, valCD}, valid32BitMode: false, valid64BitMode: true},
	{name: "JGE", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x7D, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JL", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x8C, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JL", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x8C, valCD}, valid32BitMode: false, valid64BitMode: true},
	{name: "JL", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x8C, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JL", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x7C, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JLE", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x8E, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JLE", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x8E, valCD}, valid32BitMode: false, valid64BitMode: true},
	{name: "JLE", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x8E, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JLE", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x7E, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JMP", operands: []archOperand{{regMem16, anyValue}}, opcode: []int{0xFF, modRM4}, valid32BitMode: true, valid64BitMode: false},
	{name: "JMP", operands: []archOperand{{regMem32, anyValue}}, opcode: []int{0xFF, modRM4}, valid32BitMode: true, valid64BitMode: false},
	{name: "JMP", operands: []archOperand{{regMem64, anyValue}}, opcode: []int{0xFF, modRM4}, valid32BitMode: false, valid64BitMode: true},
	{name: "JMP", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0xE9, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JMP", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0xE9, valCD}, valid32BitMode: false, valid64BitMode: true},
	{name: "JMP", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0xE9, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JMP", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0xEB, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JMP_FAR", operands: []archOperand{{mem16_16, anyValue}}, opcode: []int{0xFF, modRM5}, valid32BitMode: true, valid64BitMode: true},
	{name: "JMP_FAR", operands: []archOperand{{mem16_32, anyValue}}, opcode: []int{0xFF, modRM5}, valid32BitMode: true, valid64BitMode: true},
	{name: "JMP_FAR", operands: []archOperand{{mem16_64, anyValue}}, opcode: []int{0xFF, modRM5}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand ptr16:16
	// todo: operand ptr16:32
	{name: "JNA", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x86, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JNA", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x86, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNA", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x76, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNAE", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x82, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JNAE", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x82, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNAE", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x72, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNB", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x83, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JNB", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x83, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNB", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x73, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNBE", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x87, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JNBE", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x87, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNBE", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x77, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNC", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x83, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JNC", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x83, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNC", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x73, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNE", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x85, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JNE", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x85, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNE", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x85, valCD}, valid32BitMode: false, valid64BitMode: true},
	{name: "JNE", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x75, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNG", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x8E, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JNG", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x8E, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNG", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x7E, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNGE", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x8C, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JNGE", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x8C, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNGE", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x7C, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNL", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x8D, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JNL", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x8D, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNL", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x7D, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNLE", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x8F, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JNLE", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x8F, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNLE", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x7F, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNO", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x81, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JNO", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x81, valCD}, valid32BitMode: false, valid64BitMode: true},
	{name: "JNO", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x81, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNO", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x71, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNP", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x8B, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JNP", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x8B, valCD}, valid32BitMode: false, valid64BitMode: true},
	{name: "JNP", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x8B, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNP", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x7B, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNS", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x89, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JNS", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x89, valCD}, valid32BitMode: false, valid64BitMode: true},
	{name: "JNS", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x89, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNS", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x79, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNZ", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x85, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JNZ", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x85, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JNZ", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x75, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JO", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x80, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JO", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x80, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JO", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x80, valCD}, valid32BitMode: false, valid64BitMode: true},
	{name: "JO", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x70, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JP", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x8A, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JP", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x8A, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JP", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x8A, valCD}, valid32BitMode: false, valid64BitMode: true},
	{name: "JP", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x7A, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JPE", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x8A, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JPE", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x8A, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JPE", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x7A, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JPO", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x8B, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JPO", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x8B, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JPO", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x7B, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JRCXZ", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0xE3, valCB}, valid32BitMode: false, valid64BitMode: true},
	{name: "JS", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x88, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JS", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x88, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JS", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x88, valCD}, valid32BitMode: false, valid64BitMode: true},
	{name: "JS", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x78, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "JZ", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0x0F, 0x84, valCW}, valid32BitMode: true, valid64BitMode: false},
	{name: "JZ", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0x0F, 0x84, valCD}, valid32BitMode: true, valid64BitMode: true},
	{name: "JZ", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0x74, valCB}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	// todo: operand k1
	{name: "LAHF", operands: []archOperand{}, opcode: []int{0x9F}, valid32BitMode: true, valid64BitMode: true},
	{name: "LAR", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0x0F, 0x02, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "LAR", operands: []archOperand{{reg64, anyValue}, {regMem16, anyValue}}, opcode: []int{0x0F, 0x02, modRM}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand r32/m16
	// todo: operand xmm1
	{name: "LDMXCSR", operands: []archOperand{{mem32, anyValue}}, opcode: []int{0x0F, 0xAE, modRM2}, valid32BitMode: true, valid64BitMode: true},
	{name: "LDS", operands: []archOperand{{reg16, anyValue}, {mem16_16, anyValue}}, opcode: []int{0xC5, modRM}, valid32BitMode: true, valid64BitMode: false},
	{name: "LDS", operands: []archOperand{{reg32, anyValue}, {mem16_32, anyValue}}, opcode: []int{0xC5, modRM}, valid32BitMode: true, valid64BitMode: false},
	// todo: operand m
	// todo: operand m
	// todo: operand m
	{name: "LEAVE", operands: []archOperand{}, opcode: []int{0xC9}, valid32BitMode: true, valid64BitMode: true},
	{name: "LEAVE", operands: []archOperand{}, opcode: []int{0xC9}, valid32BitMode: false, valid64BitMode: true},
	{name: "LEAVE", operands: []archOperand{}, opcode: []int{0xC9}, valid32BitMode: true, valid64BitMode: false},
	{name: "LES", operands: []archOperand{{reg16, anyValue}, {mem16_16, anyValue}}, opcode: []int{0xC4, modRM}, valid32BitMode: true, valid64BitMode: false},
	{name: "LES", operands: []archOperand{{reg32, anyValue}, {mem16_32, anyValue}}, opcode: []int{0xC4, modRM}, valid32BitMode: true, valid64BitMode: false},
	{name: "LFS", operands: []archOperand{{reg16, anyValue}, {mem16_16, anyValue}}, opcode: []int{0x0F, 0xB4, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "LFS", operands: []archOperand{{reg32, anyValue}, {mem16_32, anyValue}}, opcode: []int{0x0F, 0xB4, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "LFS", operands: []archOperand{{reg64, anyValue}, {mem16_64, anyValue}}, opcode: []int{0x0F, 0xB4, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "LGDT", operands: []archOperand{{mem16n32, anyValue}}, opcode: []int{0x0F, 0x01, modRM2}, valid32BitMode: true, valid64BitMode: false},
	{name: "LGDT", operands: []archOperand{{mem16n64, anyValue}}, opcode: []int{0x0F, 0x01, modRM2}, valid32BitMode: false, valid64BitMode: true},
	{name: "LGS", operands: []archOperand{{reg16, anyValue}, {mem16_16, anyValue}}, opcode: []int{0x0F, 0xB5, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "LGS", operands: []archOperand{{reg32, anyValue}, {mem16_32, anyValue}}, opcode: []int{0x0F, 0xB5, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "LGS", operands: []archOperand{{reg64, anyValue}, {mem16_64, anyValue}}, opcode: []int{0x0F, 0xB5, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "LIDT", operands: []archOperand{{mem16n32, anyValue}}, opcode: []int{0x0F, 0x01, modRM3}, valid32BitMode: true, valid64BitMode: false},
	{name: "LIDT", operands: []archOperand{{mem16n64, anyValue}}, opcode: []int{0x0F, 0x01, modRM3}, valid32BitMode: false, valid64BitMode: true},
	{name: "LLDT", operands: []archOperand{{regMem16, anyValue}}, opcode: []int{0x0F, 0x00, modRM2}, valid32BitMode: true, valid64BitMode: true},
	{name: "LMSW", operands: []archOperand{{regMem16, anyValue}}, opcode: []int{0x0F, 0x01, modRM6}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand xmm1
	{name: "LOCK", operands: []archOperand{}, opcode: []int{0xF0}, valid32BitMode: true, valid64BitMode: true},
	{name: "LODSB", operands: []archOperand{}, opcode: []int{0xAC}, valid32BitMode: true, valid64BitMode: true},
	{name: "LODSD", operands: []archOperand{}, opcode: []int{0xAD}, valid32BitMode: true, valid64BitMode: true},
	{name: "LODSQ", operands: []archOperand{}, opcode: []int{0xAD}, valid32BitMode: false, valid64BitMode: true},
	{name: "LODSW", operands: []archOperand{}, opcode: []int{0xAD}, valid32BitMode: true, valid64BitMode: true},
	{name: "LOOP", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0xE2, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "LOOPE", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0xE1, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "LOOPNE", operands: []archOperand{{rel8, anyValue}}, opcode: []int{0xE0, valCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "LSL", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0x0F, 0x03, modRM}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand r32/m16
	// todo: operand r32/m16
	{name: "LSS", operands: []archOperand{{reg16, anyValue}, {mem16_16, anyValue}}, opcode: []int{0x0F, 0xB2, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "LSS", operands: []archOperand{{reg32, anyValue}, {mem16_32, anyValue}}, opcode: []int{0x0F, 0xB2, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "LSS", operands: []archOperand{{reg64, anyValue}, {mem16_64, anyValue}}, opcode: []int{0x0F, 0xB2, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "LTR", operands: []archOperand{{regMem16, anyValue}}, opcode: []int{0x0F, 0x00, modRM3}, valid32BitMode: true, valid64BitMode: true},
	{name: "LZCNT", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0xF3, 0x0F, 0xBD, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "LZCNT", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0xF3, 0x0F, 0xBD, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "LZCNT", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{0xF3, 0x0F, 0xBD, modRM}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand xmm1
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	{name: "MONITOR", operands: []archOperand{}, opcode: []int{0x0F, 0x01, 0xC8}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{reg8, al}, {memOffset8, anyValue}}, opcode: []int{0xA0}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{reg8, al}, {memOffset8, anyValue}}, opcode: []int{0xA0}, valid32BitMode: false, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{reg16, ax}, {memOffset16, anyValue}}, opcode: []int{0xA1}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand CR0-CR7
	// todo: operand CR0-CR7
	// todo: operand CR8
	// todo: operand DR0-DR7
	// todo: operand DR0-DR7
	{name: "MOV", operands: []archOperand{{reg32, eax}, {memOffset32, anyValue}}, opcode: []int{0xA1}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{reg64, rax}, {memOffset64, anyValue}}, opcode: []int{0xA1}, valid32BitMode: false, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{sReg, anyValue}, {regMem16, anyValue}}, opcode: []int{0x8E, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{sReg, anyValue}, {regMem16, anyValue}}, opcode: []int{0x8E, modRM}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand r32/m16
	{name: "MOV", operands: []archOperand{{memOffset16, anyValue}, {reg16, ax}}, opcode: []int{0xA3}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{memOffset32, anyValue}, {reg32, eax}}, opcode: []int{0xA3}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{memOffset64, anyValue}, {reg64, rax}}, opcode: []int{0xA3}, valid32BitMode: false, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{memOffset8, anyValue}, {reg8, al}}, opcode: []int{0xA2}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{memOffset8, anyValue}, {reg8, al}}, opcode: []int{0xA2}, valid32BitMode: false, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{regMem16, anyValue}, {sReg, anyValue}}, opcode: []int{0x8C, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{regMem16, anyValue}, {imm16, anyValue}}, opcode: []int{0xC7, modRM0, valIW}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{regMem16, anyValue}, {reg16, anyValue}}, opcode: []int{0x89, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{regMem32, anyValue}, {sReg, anyValue}}, opcode: []int{0x8C, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{regMem32, anyValue}, {imm32, anyValue}}, opcode: []int{0xC7, modRM0, valID}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{regMem32, anyValue}, {reg32, anyValue}}, opcode: []int{0x89, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{regMem64, anyValue}, {imm32, anyValue}}, opcode: []int{0xC7, modRM0, valID}, valid32BitMode: false, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{regMem64, anyValue}, {reg64, anyValue}}, opcode: []int{0x89, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0xC6, modRM0, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0xC6, modRM0, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x88, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x88, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{reg16, anyValue}, {imm16, anyValue}}, opcode: []int{0xB8, plusRW, valIW}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0x8B, modRM}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand r16/r32/m16
	// todo: operand CR0-CR7
	// todo: operand DR0-DR7
	{name: "MOV", operands: []archOperand{{reg32, anyValue}, {imm32, anyValue}}, opcode: []int{0xB8, plusRD, valID}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0x8B, modRM}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand CR0-CR7
	// todo: operand CR8
	// todo: operand DR0-DR7
	{name: "MOV", operands: []archOperand{{reg64, anyValue}, {imm64, anyValue}}, opcode: []int{0xB8, plusRD, valIO}, valid32BitMode: false, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{0x8B, modRM}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand r64/m16
	{name: "MOV", operands: []archOperand{{reg8, anyValue}, {imm8, anyValue}}, opcode: []int{0xB0, plusRB, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{reg8, anyValue}, {imm8, anyValue}}, opcode: []int{0xB0, plusRB, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{reg8, anyValue}, {regMem8, anyValue}}, opcode: []int{0x8A, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOV", operands: []archOperand{{reg8, anyValue}, {regMem8, anyValue}}, opcode: []int{0x8A, modRM}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand xmm1
	// todo: operand xmm2/m128
	// todo: operand xmm1
	// todo: operand xmm2/m128
	{name: "MOVBE", operands: []archOperand{{mem16, anyValue}, {reg16, anyValue}}, opcode: []int{0x0F, 0x38, 0xF1, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOVBE", operands: []archOperand{{mem32, anyValue}, {reg32, anyValue}}, opcode: []int{0x0F, 0x38, 0xF1, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOVBE", operands: []archOperand{{mem64, anyValue}, {reg64, anyValue}}, opcode: []int{0x0F, 0x38, 0xF1, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "MOVBE", operands: []archOperand{{reg16, anyValue}, {mem16, anyValue}}, opcode: []int{0x0F, 0x38, 0xF0, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOVBE", operands: []archOperand{{reg32, anyValue}, {mem32, anyValue}}, opcode: []int{0x0F, 0x38, 0xF0, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOVBE", operands: []archOperand{{reg64, anyValue}, {mem64, anyValue}}, opcode: []int{0x0F, 0x38, 0xF0, modRM}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand mm
	// todo: operand mm
	// todo: operand xmm
	// todo: operand xmm
	// todo: operand m512
	{name: "MOVDIRI", operands: []archOperand{{mem32, anyValue}, {reg32, anyValue}}, opcode: []int{0x0F, 0x38, 0xF9, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOVDIRI", operands: []archOperand{{mem64, anyValue}, {reg64, anyValue}}, opcode: []int{0x0F, 0x38, 0xF9, modRM}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand xmm2/m128
	// todo: operand xmm1
	// todo: operand xmm2/m128
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm
	// todo: operand xmm
	// todo: operand m128
	// todo: operand xmm1
	// todo: operand m128
	// todo: operand m128
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand mm
	// todo: operand mm
	// todo: operand mm
	// todo: operand xmm
	// todo: operand xmm
	// todo: operand xmm1
	// todo: operand xmm2/m64
	{name: "MOVSB", operands: []archOperand{}, opcode: []int{0xA4}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOVSD", operands: []archOperand{}, opcode: []int{0xA5}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand xmm1
	// todo: operand xmm1
	{name: "MOVSQ", operands: []archOperand{}, opcode: []int{0xA5}, valid32BitMode: false, valid64BitMode: true},
	{name: "MOVSW", operands: []archOperand{}, opcode: []int{0xA5}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOVSX", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0x0F, 0xBF, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOVSX", operands: []archOperand{{reg16, anyValue}, {regMem8, anyValue}}, opcode: []int{0x0F, 0xBE, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOVSX", operands: []archOperand{{reg32, anyValue}, {regMem16, anyValue}}, opcode: []int{0x0F, 0xBF, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOVSX", operands: []archOperand{{reg32, anyValue}, {regMem8, anyValue}}, opcode: []int{0x0F, 0xBE, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOVSX", operands: []archOperand{{reg64, anyValue}, {regMem16, anyValue}}, opcode: []int{0x0F, 0xBF, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "MOVSX", operands: []archOperand{{reg64, anyValue}, {regMem8, anyValue}}, opcode: []int{0x0F, 0xBE, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "MOVSXD", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0x63, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "MOVSXD", operands: []archOperand{{reg16, anyValue}, {regMem32, anyValue}}, opcode: []int{0x63, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "MOVSXD", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0x63, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "MOVSXD", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0x63, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "MOVSXD", operands: []archOperand{{reg64, anyValue}, {regMem32, anyValue}}, opcode: []int{0x63, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "MOVZX", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0x0F, 0xB7, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOVZX", operands: []archOperand{{reg16, anyValue}, {regMem8, anyValue}}, opcode: []int{0x0F, 0xB6, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOVZX", operands: []archOperand{{reg32, anyValue}, {regMem16, anyValue}}, opcode: []int{0x0F, 0xB7, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOVZX", operands: []archOperand{{reg32, anyValue}, {regMem8, anyValue}}, opcode: []int{0x0F, 0xB6, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "MOVZX", operands: []archOperand{{reg64, anyValue}, {regMem16, anyValue}}, opcode: []int{0x0F, 0xB7, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "MOVZX", operands: []archOperand{{reg64, anyValue}, {regMem8, anyValue}}, opcode: []int{0x0F, 0xB6, modRM}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand xmm1
	{name: "MUL", operands: []archOperand{{regMem16, anyValue}}, opcode: []int{0xF7, modRM4}, valid32BitMode: true, valid64BitMode: true},
	{name: "MUL", operands: []archOperand{{regMem32, anyValue}}, opcode: []int{0xF7, modRM4}, valid32BitMode: true, valid64BitMode: true},
	{name: "MUL", operands: []archOperand{{regMem64, anyValue}}, opcode: []int{0xF7, modRM4}, valid32BitMode: false, valid64BitMode: true},
	{name: "MUL", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0xF6, modRM4}, valid32BitMode: true, valid64BitMode: true},
	{name: "MUL", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0xF6, modRM4}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand r32a
	// todo: operand r64a
	{name: "MWAIT", operands: []archOperand{}, opcode: []int{0x0F, 0x01, 0xC9}, valid32BitMode: true, valid64BitMode: true},
	{name: "NEG", operands: []archOperand{{regMem16, anyValue}}, opcode: []int{0xF7, modRM3}, valid32BitMode: true, valid64BitMode: true},
	{name: "NEG", operands: []archOperand{{regMem32, anyValue}}, opcode: []int{0xF7, modRM3}, valid32BitMode: true, valid64BitMode: true},
	{name: "NEG", operands: []archOperand{{regMem64, anyValue}}, opcode: []int{0xF7, modRM3}, valid32BitMode: false, valid64BitMode: true},
	{name: "NEG", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0xF6, modRM3}, valid32BitMode: true, valid64BitMode: true},
	{name: "NEG", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0xF6, modRM3}, valid32BitMode: false, valid64BitMode: true},
	{name: "NOP", operands: []archOperand{}, opcode: []int{0x90}, valid32BitMode: true, valid64BitMode: true},
	{name: "NOP", operands: []archOperand{{regMem16, anyValue}}, opcode: []int{0x0F, 0x1F, modRM0}, valid32BitMode: true, valid64BitMode: true},
	{name: "NOP", operands: []archOperand{{regMem32, anyValue}}, opcode: []int{0x0F, 0x1F, modRM0}, valid32BitMode: true, valid64BitMode: true},
	{name: "NOT", operands: []archOperand{{regMem16, anyValue}}, opcode: []int{0xF7, modRM2}, valid32BitMode: true, valid64BitMode: true},
	{name: "NOT", operands: []archOperand{{regMem32, anyValue}}, opcode: []int{0xF7, modRM2}, valid32BitMode: true, valid64BitMode: true},
	{name: "NOT", operands: []archOperand{{regMem64, anyValue}}, opcode: []int{0xF7, modRM2}, valid32BitMode: false, valid64BitMode: true},
	{name: "NOT", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0xF6, modRM2}, valid32BitMode: true, valid64BitMode: true},
	{name: "NOT", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0xF6, modRM2}, valid32BitMode: false, valid64BitMode: true},
	{name: "OR", operands: []archOperand{{reg8, al}, {imm8, anyValue}}, opcode: []int{0x0C, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "OR", operands: []archOperand{{reg16, ax}, {imm16, anyValue}}, opcode: []int{0x0D, valIW}, valid32BitMode: true, valid64BitMode: true},
	{name: "OR", operands: []archOperand{{reg32, eax}, {imm32, anyValue}}, opcode: []int{0x0D, valID}, valid32BitMode: true, valid64BitMode: true},
	{name: "OR", operands: []archOperand{{reg64, rax}, {imm32, anyValue}}, opcode: []int{0x0D, valID}, valid32BitMode: false, valid64BitMode: true},
	{name: "OR", operands: []archOperand{{regMem16, anyValue}, {imm16, anyValue}}, opcode: []int{0x81, modRM1, valIW}, valid32BitMode: true, valid64BitMode: true},
	{name: "OR", operands: []archOperand{{regMem16, anyValue}, {imm8, anyValue}}, opcode: []int{0x83, modRM1, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "OR", operands: []archOperand{{regMem16, anyValue}, {reg16, anyValue}}, opcode: []int{0x09, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "OR", operands: []archOperand{{regMem32, anyValue}, {imm32, anyValue}}, opcode: []int{0x81, modRM1, valID}, valid32BitMode: true, valid64BitMode: true},
	{name: "OR", operands: []archOperand{{regMem32, anyValue}, {imm8, anyValue}}, opcode: []int{0x83, modRM1, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "OR", operands: []archOperand{{regMem32, anyValue}, {reg32, anyValue}}, opcode: []int{0x09, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "OR", operands: []archOperand{{regMem64, anyValue}, {imm32, anyValue}}, opcode: []int{0x81, modRM1, valID}, valid32BitMode: false, valid64BitMode: true},
	{name: "OR", operands: []archOperand{{regMem64, anyValue}, {imm8, anyValue}}, opcode: []int{0x83, modRM1, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "OR", operands: []archOperand{{regMem64, anyValue}, {reg64, anyValue}}, opcode: []int{0x09, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "OR", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0x80, modRM1, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "OR", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0x80, modRM1, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "OR", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x08, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "OR", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x08, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "OR", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0x0B, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "OR", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0x0B, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "OR", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{0x0B, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "OR", operands: []archOperand{{reg8, anyValue}, {regMem8, anyValue}}, opcode: []int{0x0A, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "OR", operands: []archOperand{{reg8, anyValue}, {regMem8, anyValue}}, opcode: []int{0x0A, modRM}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand xmm1
	// todo: operand xmm1
	{name: "OUT", operands: []archOperand{{reg16, dx}, {reg8, al}}, opcode: []int{0xEE}, valid32BitMode: true, valid64BitMode: true},
	{name: "OUT", operands: []archOperand{{reg16, dx}, {reg16, ax}}, opcode: []int{0xEF}, valid32BitMode: true, valid64BitMode: true},
	{name: "OUT", operands: []archOperand{{reg16, dx}, {reg32, eax}}, opcode: []int{0xEF}, valid32BitMode: true, valid64BitMode: true},
	{name: "OUT", operands: []archOperand{{imm8, anyValue}, {reg8, al}}, opcode: []int{0xE6, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "OUT", operands: []archOperand{{imm8, anyValue}, {reg16, ax}}, opcode: []int{0xE7, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "OUT", operands: []archOperand{{imm8, anyValue}, {reg32, eax}}, opcode: []int{0xE7, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "OUTSB", operands: []archOperand{}, opcode: []int{0x6E}, valid32BitMode: true, valid64BitMode: true},
	{name: "OUTSD", operands: []archOperand{}, opcode: []int{0x6F}, valid32BitMode: true, valid64BitMode: true},
	{name: "OUTSW", operands: []archOperand{}, opcode: []int{0x6F}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	{name: "PAUSE", operands: []archOperand{}, opcode: []int{0xF3, 0x90}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand r32a
	// todo: operand r64a
	// todo: operand r32a
	// todo: operand r64a
	// todo: operand xmm2
	// todo: operand xmm2
	// todo: operand xmm2
	// todo: operand mm
	// todo: operand xmm
	// todo: operand xmm
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand DS
	// todo: operand ES
	// todo: operand FS
	// todo: operand FS
	// todo: operand FS
	// todo: operand GS
	// todo: operand GS
	// todo: operand GS
	// todo: operand SS
	{name: "POP", operands: []archOperand{{regMem16, anyValue}}, opcode: []int{0x8F, modRM0}, valid32BitMode: true, valid64BitMode: true},
	{name: "POP", operands: []archOperand{{regMem32, anyValue}}, opcode: []int{0x8F, modRM0}, valid32BitMode: true, valid64BitMode: false},
	{name: "POP", operands: []archOperand{{regMem64, anyValue}}, opcode: []int{0x8F, modRM0}, valid32BitMode: false, valid64BitMode: true},
	{name: "POP", operands: []archOperand{{reg16, anyValue}}, opcode: []int{0x58, plusRW}, valid32BitMode: true, valid64BitMode: true},
	{name: "POP", operands: []archOperand{{reg32, anyValue}}, opcode: []int{0x58, plusRD}, valid32BitMode: true, valid64BitMode: false},
	{name: "POP", operands: []archOperand{{reg64, anyValue}}, opcode: []int{0x58, plusRD}, valid32BitMode: false, valid64BitMode: true},
	{name: "POPA", operands: []archOperand{}, opcode: []int{0x61}, valid32BitMode: true, valid64BitMode: false},
	{name: "POPAD", operands: []archOperand{}, opcode: []int{0x61}, valid32BitMode: true, valid64BitMode: false},
	{name: "POPCNT", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0xF3, 0x0F, 0xB8, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "POPCNT", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0xF3, 0x0F, 0xB8, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "POPCNT", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{0xF3, 0x0F, 0xB8, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "POPF", operands: []archOperand{}, opcode: []int{0x9D}, valid32BitMode: true, valid64BitMode: true},
	{name: "POPFD", operands: []archOperand{}, opcode: []int{0x9D}, valid32BitMode: true, valid64BitMode: false},
	{name: "POPFQ", operands: []archOperand{}, opcode: []int{0x9D}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand mm
	// todo: operand xmm1
	{name: "PREFETCHNTA", operands: []archOperand{{mem8, anyValue}}, opcode: []int{0x0F, 0x18, modRM0}, valid32BitMode: true, valid64BitMode: true},
	{name: "PREFETCHT0", operands: []archOperand{{mem8, anyValue}}, opcode: []int{0x0F, 0x18, modRM1}, valid32BitMode: true, valid64BitMode: true},
	{name: "PREFETCHT1", operands: []archOperand{{mem8, anyValue}}, opcode: []int{0x0F, 0x18, modRM2}, valid32BitMode: true, valid64BitMode: true},
	{name: "PREFETCHT2", operands: []archOperand{{mem8, anyValue}}, opcode: []int{0x0F, 0x18, modRM3}, valid32BitMode: true, valid64BitMode: true},
	{name: "PREFETCHW", operands: []archOperand{{mem8, anyValue}}, opcode: []int{0x0F, 0x0D, modRM1}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm1
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand mm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand CS
	// todo: operand DS
	// todo: operand ES
	// todo: operand FS
	// todo: operand GS
	// todo: operand SS
	{name: "PUSH", operands: []archOperand{{imm16, anyValue}}, opcode: []int{0x68, valIW}, valid32BitMode: true, valid64BitMode: true},
	{name: "PUSH", operands: []archOperand{{imm32, anyValue}}, opcode: []int{0x68, valID}, valid32BitMode: true, valid64BitMode: true},
	{name: "PUSH", operands: []archOperand{{imm8, anyValue}}, opcode: []int{0x6A, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "PUSH", operands: []archOperand{{regMem16, anyValue}}, opcode: []int{0xFF, modRM6}, valid32BitMode: true, valid64BitMode: true},
	{name: "PUSH", operands: []archOperand{{regMem32, anyValue}}, opcode: []int{0xFF, modRM6}, valid32BitMode: true, valid64BitMode: false},
	{name: "PUSH", operands: []archOperand{{regMem64, anyValue}}, opcode: []int{0xFF, modRM6}, valid32BitMode: false, valid64BitMode: true},
	{name: "PUSH", operands: []archOperand{{reg16, anyValue}}, opcode: []int{0x50, plusRW}, valid32BitMode: true, valid64BitMode: true},
	{name: "PUSH", operands: []archOperand{{reg32, anyValue}}, opcode: []int{0x50, plusRD}, valid32BitMode: true, valid64BitMode: false},
	{name: "PUSH", operands: []archOperand{{reg64, anyValue}}, opcode: []int{0x50, plusRD}, valid32BitMode: false, valid64BitMode: true},
	{name: "PUSHA", operands: []archOperand{}, opcode: []int{0x60}, valid32BitMode: true, valid64BitMode: false},
	{name: "PUSHAD", operands: []archOperand{}, opcode: []int{0x60}, valid32BitMode: true, valid64BitMode: false},
	{name: "PUSHF", operands: []archOperand{}, opcode: []int{0x9C}, valid32BitMode: true, valid64BitMode: true},
	{name: "PUSHFD", operands: []archOperand{}, opcode: []int{0x9C}, valid32BitMode: true, valid64BitMode: false},
	{name: "PUSHFQ", operands: []archOperand{}, opcode: []int{0x9C}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand mm
	// todo: operand xmm1
	// todo: operand 1
	{name: "RCL", operands: []archOperand{{regMem16, anyValue}, {reg8, cl}}, opcode: []int{0xD3, modRM2}, valid32BitMode: true, valid64BitMode: true},
	{name: "RCL", operands: []archOperand{{regMem16, anyValue}, {imm8, anyValue}}, opcode: []int{0xC1, modRM2, valIB}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand 1
	{name: "RCL", operands: []archOperand{{regMem32, anyValue}, {reg8, cl}}, opcode: []int{0xD3, modRM2}, valid32BitMode: true, valid64BitMode: true},
	{name: "RCL", operands: []archOperand{{regMem32, anyValue}, {imm8, anyValue}}, opcode: []int{0xC1, modRM2, valIB}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand 1
	{name: "RCL", operands: []archOperand{{regMem64, anyValue}, {reg8, cl}}, opcode: []int{0xD3, modRM2}, valid32BitMode: false, valid64BitMode: true},
	{name: "RCL", operands: []archOperand{{regMem64, anyValue}, {imm8, anyValue}}, opcode: []int{0xC1, modRM2, valIB}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand 1
	// todo: operand 1
	{name: "RCL", operands: []archOperand{{regMem8, anyValue}, {reg8, cl}}, opcode: []int{0xD2, modRM2}, valid32BitMode: true, valid64BitMode: true},
	{name: "RCL", operands: []archOperand{{regMem8, anyValue}, {reg8, cl}}, opcode: []int{0xD2, modRM2}, valid32BitMode: false, valid64BitMode: true},
	{name: "RCL", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0xC0, modRM2, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "RCL", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0xC0, modRM2, valIB}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand 1
	{name: "RCR", operands: []archOperand{{regMem16, anyValue}, {reg8, cl}}, opcode: []int{0xD3, modRM3}, valid32BitMode: true, valid64BitMode: true},
	{name: "RCR", operands: []archOperand{{regMem16, anyValue}, {imm8, anyValue}}, opcode: []int{0xC1, modRM3, valIB}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand 1
	{name: "RCR", operands: []archOperand{{regMem32, anyValue}, {reg8, cl}}, opcode: []int{0xD3, modRM3}, valid32BitMode: true, valid64BitMode: true},
	{name: "RCR", operands: []archOperand{{regMem32, anyValue}, {imm8, anyValue}}, opcode: []int{0xC1, modRM3, valIB}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand 1
	{name: "RCR", operands: []archOperand{{regMem64, anyValue}, {reg8, cl}}, opcode: []int{0xD3, modRM3}, valid32BitMode: false, valid64BitMode: true},
	{name: "RCR", operands: []archOperand{{regMem64, anyValue}, {imm8, anyValue}}, opcode: []int{0xC1, modRM3, valIB}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand 1
	// todo: operand 1
	{name: "RCR", operands: []archOperand{{regMem8, anyValue}, {reg8, cl}}, opcode: []int{0xD2, modRM3}, valid32BitMode: true, valid64BitMode: true},
	{name: "RCR", operands: []archOperand{{regMem8, anyValue}, {reg8, cl}}, opcode: []int{0xD2, modRM3}, valid32BitMode: false, valid64BitMode: true},
	{name: "RCR", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0xC0, modRM3, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "RCR", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0xC0, modRM3, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "RDFSBASE", operands: []archOperand{{reg32, anyValue}}, opcode: []int{0xF3, 0x0F, 0xAE, modRM0}, valid32BitMode: false, valid64BitMode: true},
	{name: "RDFSBASE", operands: []archOperand{{reg64, anyValue}}, opcode: []int{0xF3, 0x0F, 0xAE, modRM0}, valid32BitMode: false, valid64BitMode: true},
	{name: "RDGSBASE", operands: []archOperand{{reg32, anyValue}}, opcode: []int{0xF3, 0x0F, 0xAE, modRM1}, valid32BitMode: false, valid64BitMode: true},
	{name: "RDGSBASE", operands: []archOperand{{reg64, anyValue}}, opcode: []int{0xF3, 0x0F, 0xAE, modRM1}, valid32BitMode: false, valid64BitMode: true},
	{name: "RDMSR", operands: []archOperand{}, opcode: []int{0x0F, 0x32}, valid32BitMode: true, valid64BitMode: true},
	{name: "RDPID", operands: []archOperand{{reg32, anyValue}}, opcode: []int{0xF3, 0x0F, 0xC7, modRM7}, valid32BitMode: true, valid64BitMode: false},
	{name: "RDPID", operands: []archOperand{{reg64, anyValue}}, opcode: []int{0xF3, 0x0F, 0xC7, modRM7}, valid32BitMode: false, valid64BitMode: true},
	{name: "RDPKRU", operands: []archOperand{}, opcode: []int{0x0F, 0x01, 0xEE}, valid32BitMode: true, valid64BitMode: true},
	{name: "RDPMC", operands: []archOperand{}, opcode: []int{0x0F, 0x33}, valid32BitMode: true, valid64BitMode: true},
	{name: "RDRAND", operands: []archOperand{{reg16, anyValue}}, opcode: []int{0x0F, 0xC7, modRM6}, valid32BitMode: true, valid64BitMode: true},
	{name: "RDRAND", operands: []archOperand{{reg32, anyValue}}, opcode: []int{0x0F, 0xC7, modRM6}, valid32BitMode: true, valid64BitMode: true},
	{name: "RDRAND", operands: []archOperand{{reg64, anyValue}}, opcode: []int{0x0F, 0xC7, modRM6}, valid32BitMode: false, valid64BitMode: true},
	{name: "RDSEED", operands: []archOperand{{reg16, anyValue}}, opcode: []int{0x0F, 0xC7, modRM7}, valid32BitMode: true, valid64BitMode: true},
	{name: "RDSEED", operands: []archOperand{{reg32, anyValue}}, opcode: []int{0x0F, 0xC7, modRM7}, valid32BitMode: true, valid64BitMode: true},
	{name: "RDSEED", operands: []archOperand{{reg64, anyValue}}, opcode: []int{0x0F, 0xC7, modRM7}, valid32BitMode: false, valid64BitMode: true},
	{name: "RDSSPD", operands: []archOperand{{reg32, anyValue}}, opcode: []int{0xF3, 0x0F, 0x1E, modRM1, todoOpcode}, valid32BitMode: true, valid64BitMode: true},
	{name: "RDSSPQ", operands: []archOperand{{reg64, anyValue}}, opcode: []int{0xF3, 0x0F, 0x1E, modRM1, todoOpcode}, valid32BitMode: false, valid64BitMode: true},
	{name: "RDTSC", operands: []archOperand{}, opcode: []int{0x0F, 0x31}, valid32BitMode: true, valid64BitMode: true},
	{name: "RDTSCP", operands: []archOperand{}, opcode: []int{0x0F, 0x01, 0xF9}, valid32BitMode: true, valid64BitMode: true},
	{name: "RET", operands: []archOperand{}, opcode: []int{0xC3}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand imm16u
	{name: "RET_FAR", operands: []archOperand{}, opcode: []int{0xCB}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand imm16u
	// todo: operand 1
	{name: "ROL", operands: []archOperand{{regMem16, anyValue}, {reg8, cl}}, opcode: []int{0xD3, modRM0}, valid32BitMode: true, valid64BitMode: true},
	{name: "ROL", operands: []archOperand{{regMem16, anyValue}, {imm8, anyValue}}, opcode: []int{0xC1, modRM0, valIB}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand 1
	{name: "ROL", operands: []archOperand{{regMem32, anyValue}, {reg8, cl}}, opcode: []int{0xD3, modRM0}, valid32BitMode: true, valid64BitMode: true},
	{name: "ROL", operands: []archOperand{{regMem32, anyValue}, {imm8, anyValue}}, opcode: []int{0xC1, modRM0, valIB}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand 1
	{name: "ROL", operands: []archOperand{{regMem64, anyValue}, {reg8, cl}}, opcode: []int{0xD3, modRM0}, valid32BitMode: false, valid64BitMode: true},
	{name: "ROL", operands: []archOperand{{regMem64, anyValue}, {imm8, anyValue}}, opcode: []int{0xC1, modRM0, valIB}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand 1
	// todo: operand 1
	{name: "ROL", operands: []archOperand{{regMem8, anyValue}, {reg8, cl}}, opcode: []int{0xD2, modRM0}, valid32BitMode: true, valid64BitMode: true},
	{name: "ROL", operands: []archOperand{{regMem8, anyValue}, {reg8, cl}}, opcode: []int{0xD2, modRM0}, valid32BitMode: false, valid64BitMode: true},
	{name: "ROL", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0xC0, modRM0, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "ROL", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0xC0, modRM0, valIB}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand 1
	{name: "ROR", operands: []archOperand{{regMem16, anyValue}, {reg8, cl}}, opcode: []int{0xD3, modRM1}, valid32BitMode: true, valid64BitMode: true},
	{name: "ROR", operands: []archOperand{{regMem16, anyValue}, {imm8, anyValue}}, opcode: []int{0xC1, modRM1, valIB}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand 1
	{name: "ROR", operands: []archOperand{{regMem32, anyValue}, {reg8, cl}}, opcode: []int{0xD3, modRM1}, valid32BitMode: true, valid64BitMode: true},
	{name: "ROR", operands: []archOperand{{regMem32, anyValue}, {imm8, anyValue}}, opcode: []int{0xC1, modRM1, valIB}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand 1
	{name: "ROR", operands: []archOperand{{regMem64, anyValue}, {reg8, cl}}, opcode: []int{0xD3, modRM1}, valid32BitMode: false, valid64BitMode: true},
	{name: "ROR", operands: []archOperand{{regMem64, anyValue}, {imm8, anyValue}}, opcode: []int{0xC1, modRM1, valIB}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand 1
	// todo: operand 1
	{name: "ROR", operands: []archOperand{{regMem8, anyValue}, {reg8, cl}}, opcode: []int{0xD2, modRM1}, valid32BitMode: true, valid64BitMode: true},
	{name: "ROR", operands: []archOperand{{regMem8, anyValue}, {reg8, cl}}, opcode: []int{0xD2, modRM1}, valid32BitMode: false, valid64BitMode: true},
	{name: "ROR", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0xC0, modRM1, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "ROR", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0xC0, modRM1, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "RORX", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}, {imm8, anyValue}}, opcode: []int{todoOpcode}, valid32BitMode: true, valid64BitMode: true},
	{name: "RORX", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}, {imm8, anyValue}}, opcode: []int{todoOpcode}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	{name: "RSM", operands: []archOperand{}, opcode: []int{0x0F, 0xAA}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand xmm1
	// todo: operand xmm1
	{name: "RSTORSSP", operands: []archOperand{{mem64, anyValue}}, opcode: []int{0xF3, 0x0F, 0x01, modRM5, todoOpcode}, valid32BitMode: true, valid64BitMode: true},
	{name: "SAHF", operands: []archOperand{}, opcode: []int{0x9E}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand 1
	{name: "SAL", operands: []archOperand{{regMem16, anyValue}, {reg8, cl}}, opcode: []int{0xD3, modRM4}, valid32BitMode: true, valid64BitMode: true},
	{name: "SAL", operands: []archOperand{{regMem16, anyValue}, {imm8, anyValue}}, opcode: []int{0xC1, modRM4, valIB}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand 1
	{name: "SAL", operands: []archOperand{{regMem32, anyValue}, {reg8, cl}}, opcode: []int{0xD3, modRM4}, valid32BitMode: true, valid64BitMode: true},
	{name: "SAL", operands: []archOperand{{regMem32, anyValue}, {imm8, anyValue}}, opcode: []int{0xC1, modRM4, valIB}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand 1
	{name: "SAL", operands: []archOperand{{regMem64, anyValue}, {reg8, cl}}, opcode: []int{0xD3, modRM4}, valid32BitMode: false, valid64BitMode: true},
	{name: "SAL", operands: []archOperand{{regMem64, anyValue}, {imm8, anyValue}}, opcode: []int{0xC1, modRM4, valIB}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand 1
	// todo: operand 1
	{name: "SAL", operands: []archOperand{{regMem8, anyValue}, {reg8, cl}}, opcode: []int{0xD2, modRM4}, valid32BitMode: true, valid64BitMode: true},
	{name: "SAL", operands: []archOperand{{regMem8, anyValue}, {reg8, cl}}, opcode: []int{0xD2, modRM4}, valid32BitMode: false, valid64BitMode: true},
	{name: "SAL", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0xC0, modRM4, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "SAL", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0xC0, modRM4, valIB}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand 1
	{name: "SAR", operands: []archOperand{{regMem16, anyValue}, {reg8, cl}}, opcode: []int{0xD3, modRM7}, valid32BitMode: true, valid64BitMode: true},
	{name: "SAR", operands: []archOperand{{regMem16, anyValue}, {imm8, anyValue}}, opcode: []int{0xC1, modRM7, valIB}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand 1
	{name: "SAR", operands: []archOperand{{regMem32, anyValue}, {reg8, cl}}, opcode: []int{0xD3, modRM7}, valid32BitMode: true, valid64BitMode: true},
	{name: "SAR", operands: []archOperand{{regMem32, anyValue}, {imm8, anyValue}}, opcode: []int{0xC1, modRM7, valIB}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand 1
	{name: "SAR", operands: []archOperand{{regMem64, anyValue}, {reg8, cl}}, opcode: []int{0xD3, modRM7}, valid32BitMode: false, valid64BitMode: true},
	{name: "SAR", operands: []archOperand{{regMem64, anyValue}, {imm8, anyValue}}, opcode: []int{0xC1, modRM7, valIB}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand 1
	// todo: operand 1
	{name: "SAR", operands: []archOperand{{regMem8, anyValue}, {reg8, cl}}, opcode: []int{0xD2, modRM7}, valid32BitMode: true, valid64BitMode: true},
	{name: "SAR", operands: []archOperand{{regMem8, anyValue}, {reg8, cl}}, opcode: []int{0xD2, modRM7}, valid32BitMode: false, valid64BitMode: true},
	{name: "SAR", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0xC0, modRM7, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "SAR", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0xC0, modRM7, valIB}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand r32a
	// todo: operand r64a
	{name: "SAVEPREVSSP", operands: []archOperand{}, opcode: []int{0xF3, 0x0F, 0x01, 0xEA, todoOpcode}, valid32BitMode: true, valid64BitMode: true},
	{name: "SBB", operands: []archOperand{{reg8, al}, {imm8, anyValue}}, opcode: []int{0x1C, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "SBB", operands: []archOperand{{reg16, ax}, {imm16, anyValue}}, opcode: []int{0x1D, valIW}, valid32BitMode: true, valid64BitMode: true},
	{name: "SBB", operands: []archOperand{{reg32, eax}, {imm32, anyValue}}, opcode: []int{0x1D, valID}, valid32BitMode: true, valid64BitMode: true},
	{name: "SBB", operands: []archOperand{{reg64, rax}, {imm32, anyValue}}, opcode: []int{0x1D, valID}, valid32BitMode: false, valid64BitMode: true},
	{name: "SBB", operands: []archOperand{{regMem16, anyValue}, {imm16, anyValue}}, opcode: []int{0x81, modRM3, valIW}, valid32BitMode: true, valid64BitMode: true},
	{name: "SBB", operands: []archOperand{{regMem16, anyValue}, {imm8, anyValue}}, opcode: []int{0x83, modRM3, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "SBB", operands: []archOperand{{regMem16, anyValue}, {reg16, anyValue}}, opcode: []int{0x19, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "SBB", operands: []archOperand{{regMem32, anyValue}, {imm32, anyValue}}, opcode: []int{0x81, modRM3, valID}, valid32BitMode: true, valid64BitMode: true},
	{name: "SBB", operands: []archOperand{{regMem32, anyValue}, {imm8, anyValue}}, opcode: []int{0x83, modRM3, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "SBB", operands: []archOperand{{regMem32, anyValue}, {reg32, anyValue}}, opcode: []int{0x19, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "SBB", operands: []archOperand{{regMem64, anyValue}, {imm32, anyValue}}, opcode: []int{0x81, modRM3, valID}, valid32BitMode: false, valid64BitMode: true},
	{name: "SBB", operands: []archOperand{{regMem64, anyValue}, {imm8, anyValue}}, opcode: []int{0x83, modRM3, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "SBB", operands: []archOperand{{regMem64, anyValue}, {reg64, anyValue}}, opcode: []int{0x19, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "SBB", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0x80, modRM3, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "SBB", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0x80, modRM3, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "SBB", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x18, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "SBB", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x18, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "SBB", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0x1B, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "SBB", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0x1B, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "SBB", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{0x1B, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "SBB", operands: []archOperand{{reg8, anyValue}, {regMem8, anyValue}}, opcode: []int{0x1A, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "SBB", operands: []archOperand{{reg8, anyValue}, {regMem8, anyValue}}, opcode: []int{0x1A, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "SCASB", operands: []archOperand{}, opcode: []int{0xAE}, valid32BitMode: true, valid64BitMode: true},
	{name: "SCASD", operands: []archOperand{}, opcode: []int{0xAF}, valid32BitMode: true, valid64BitMode: true},
	{name: "SCASQ", operands: []archOperand{}, opcode: []int{0xAF}, valid32BitMode: false, valid64BitMode: true},
	{name: "SCASW", operands: []archOperand{}, opcode: []int{0xAF}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETA", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x97}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETA", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x97}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETAE", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x93}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETAE", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x93}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETB", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x92}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETB", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x92}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETBE", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x96}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETBE", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x96}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETC", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x92}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETC", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x92}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETE", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x94}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETE", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x94}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETG", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x9F}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETG", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x9F}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETGE", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x9D}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETGE", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x9D}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETL", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x9C}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETL", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x9C}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETLE", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x9E}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETLE", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x9E}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETNA", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x96}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETNA", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x96}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETNAE", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x92}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETNAE", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x92}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETNB", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x93}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETNB", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x93}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETNBE", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x97}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETNBE", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x97}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETNC", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x93}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETNC", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x93}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETNE", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x95}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETNE", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x95}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETNG", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x9E}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETNG", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x9E}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETNGE", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x9C}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETNGE", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x9C}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETNL", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x9D}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETNL", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x9D}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETNLE", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x9F}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETNLE", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x9F}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETNO", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x91}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETNO", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x91}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETNP", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x9B}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETNP", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x9B}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETNS", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x99}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETNS", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x99}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETNZ", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x95}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETNZ", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x95}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETO", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x90}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETO", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x90}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETP", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x9A}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETP", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x9A}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETPE", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x9A}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETPE", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x9A}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETPO", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x9B}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETPO", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x9B}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETS", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x98}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETS", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x98}, valid32BitMode: false, valid64BitMode: true},
	{name: "SETSSBSY", operands: []archOperand{}, opcode: []int{0xF3, 0x0F, 0x01, 0xE8}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETZ", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x94}, valid32BitMode: true, valid64BitMode: true},
	{name: "SETZ", operands: []archOperand{{regMem8, anyValue}}, opcode: []int{0x0F, 0x94}, valid32BitMode: false, valid64BitMode: true},
	{name: "SFENCE", operands: []archOperand{}, opcode: []int{0x0F, 0xAE, 0xF8}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand m
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand 1
	{name: "SHL", operands: []archOperand{{regMem16, anyValue}, {reg8, cl}}, opcode: []int{0xD3, modRM4}, valid32BitMode: true, valid64BitMode: true},
	{name: "SHL", operands: []archOperand{{regMem16, anyValue}, {imm8, anyValue}}, opcode: []int{0xC1, modRM4, valIB}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand 1
	{name: "SHL", operands: []archOperand{{regMem32, anyValue}, {reg8, cl}}, opcode: []int{0xD3, modRM4}, valid32BitMode: true, valid64BitMode: true},
	{name: "SHL", operands: []archOperand{{regMem32, anyValue}, {imm8, anyValue}}, opcode: []int{0xC1, modRM4, valIB}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand 1
	{name: "SHL", operands: []archOperand{{regMem64, anyValue}, {reg8, cl}}, opcode: []int{0xD3, modRM4}, valid32BitMode: false, valid64BitMode: true},
	{name: "SHL", operands: []archOperand{{regMem64, anyValue}, {imm8, anyValue}}, opcode: []int{0xC1, modRM4, valIB}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand 1
	// todo: operand 1
	{name: "SHL", operands: []archOperand{{regMem8, anyValue}, {reg8, cl}}, opcode: []int{0xD2, modRM4}, valid32BitMode: true, valid64BitMode: true},
	{name: "SHL", operands: []archOperand{{regMem8, anyValue}, {reg8, cl}}, opcode: []int{0xD2, modRM4}, valid32BitMode: false, valid64BitMode: true},
	{name: "SHL", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0xC0, modRM4, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "SHL", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0xC0, modRM4, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "SHLD", operands: []archOperand{{regMem16, anyValue}, {reg16, anyValue}, {reg8, cl}}, opcode: []int{0x0F, 0xA5, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "SHLD", operands: []archOperand{{regMem16, anyValue}, {reg16, anyValue}, {imm8, anyValue}}, opcode: []int{0x0F, 0xA4, modRM, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "SHLD", operands: []archOperand{{regMem32, anyValue}, {reg32, anyValue}, {reg8, cl}}, opcode: []int{0x0F, 0xA5, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "SHLD", operands: []archOperand{{regMem32, anyValue}, {reg32, anyValue}, {imm8, anyValue}}, opcode: []int{0x0F, 0xA4, modRM, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "SHLD", operands: []archOperand{{regMem64, anyValue}, {reg64, anyValue}, {reg8, cl}}, opcode: []int{0x0F, 0xA5, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "SHLD", operands: []archOperand{{regMem64, anyValue}, {reg64, anyValue}, {imm8, anyValue}}, opcode: []int{0x0F, 0xA4, modRM, valIB}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand r32a
	// todo: operand r64a
	// todo: operand 1
	{name: "SHR", operands: []archOperand{{regMem16, anyValue}, {reg8, cl}}, opcode: []int{0xD3, modRM5}, valid32BitMode: true, valid64BitMode: true},
	{name: "SHR", operands: []archOperand{{regMem16, anyValue}, {imm8, anyValue}}, opcode: []int{0xC1, modRM5, valIB}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand 1
	{name: "SHR", operands: []archOperand{{regMem32, anyValue}, {reg8, cl}}, opcode: []int{0xD3, modRM5}, valid32BitMode: true, valid64BitMode: true},
	{name: "SHR", operands: []archOperand{{regMem32, anyValue}, {imm8, anyValue}}, opcode: []int{0xC1, modRM5, valIB}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand 1
	{name: "SHR", operands: []archOperand{{regMem64, anyValue}, {reg8, cl}}, opcode: []int{0xD3, modRM5}, valid32BitMode: false, valid64BitMode: true},
	{name: "SHR", operands: []archOperand{{regMem64, anyValue}, {imm8, anyValue}}, opcode: []int{0xC1, modRM5, valIB}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand 1
	// todo: operand 1
	{name: "SHR", operands: []archOperand{{regMem8, anyValue}, {reg8, cl}}, opcode: []int{0xD2, modRM5}, valid32BitMode: true, valid64BitMode: true},
	{name: "SHR", operands: []archOperand{{regMem8, anyValue}, {reg8, cl}}, opcode: []int{0xD2, modRM5}, valid32BitMode: false, valid64BitMode: true},
	{name: "SHR", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0xC0, modRM5, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "SHR", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0xC0, modRM5, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "SHRD", operands: []archOperand{{regMem16, anyValue}, {reg16, anyValue}, {reg8, cl}}, opcode: []int{0x0F, 0xAD, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "SHRD", operands: []archOperand{{regMem16, anyValue}, {reg16, anyValue}, {imm8, anyValue}}, opcode: []int{0x0F, 0xAC, modRM, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "SHRD", operands: []archOperand{{regMem32, anyValue}, {reg32, anyValue}, {reg8, cl}}, opcode: []int{0x0F, 0xAD, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "SHRD", operands: []archOperand{{regMem32, anyValue}, {reg32, anyValue}, {imm8, anyValue}}, opcode: []int{0x0F, 0xAC, modRM, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "SHRD", operands: []archOperand{{regMem64, anyValue}, {reg64, anyValue}, {reg8, cl}}, opcode: []int{0x0F, 0xAD, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "SHRD", operands: []archOperand{{regMem64, anyValue}, {reg64, anyValue}, {imm8, anyValue}}, opcode: []int{0x0F, 0xAC, modRM, valIB}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand r32a
	// todo: operand r64a
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand m
	{name: "SLDT", operands: []archOperand{{regMem16, anyValue}}, opcode: []int{0x0F, 0x00, modRM0}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand r32/m16
	{name: "SMSW", operands: []archOperand{{regMem16, anyValue}}, opcode: []int{0x0F, 0x01, modRM4}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand r32/m16
	// todo: operand r64/m16
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	{name: "STAC", operands: []archOperand{}, opcode: []int{0x0F, 0x01, 0xCB}, valid32BitMode: true, valid64BitMode: true},
	{name: "STC", operands: []archOperand{}, opcode: []int{0xF9}, valid32BitMode: true, valid64BitMode: true},
	{name: "STD", operands: []archOperand{}, opcode: []int{0xFD}, valid32BitMode: true, valid64BitMode: true},
	{name: "STI", operands: []archOperand{}, opcode: []int{0xFB}, valid32BitMode: true, valid64BitMode: true},
	{name: "STMXCSR", operands: []archOperand{{mem32, anyValue}}, opcode: []int{0x0F, 0xAE, modRM3}, valid32BitMode: true, valid64BitMode: true},
	{name: "STOSB", operands: []archOperand{}, opcode: []int{0xAA}, valid32BitMode: true, valid64BitMode: true},
	{name: "STOSD", operands: []archOperand{}, opcode: []int{0xAB}, valid32BitMode: true, valid64BitMode: true},
	{name: "STOSQ", operands: []archOperand{}, opcode: []int{0xAB}, valid32BitMode: false, valid64BitMode: true},
	{name: "STOSW", operands: []archOperand{}, opcode: []int{0xAB}, valid32BitMode: true, valid64BitMode: true},
	{name: "STR", operands: []archOperand{{regMem16, anyValue}}, opcode: []int{0x0F, 0x00, modRM1}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand r32/m16
	// todo: operand r64/m16
	{name: "SUB", operands: []archOperand{{reg8, al}, {imm8, anyValue}}, opcode: []int{0x2C, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "SUB", operands: []archOperand{{reg16, ax}, {imm16, anyValue}}, opcode: []int{0x2D, valIW}, valid32BitMode: true, valid64BitMode: true},
	{name: "SUB", operands: []archOperand{{reg32, eax}, {imm32, anyValue}}, opcode: []int{0x2D, valID}, valid32BitMode: true, valid64BitMode: true},
	{name: "SUB", operands: []archOperand{{reg64, rax}, {imm32, anyValue}}, opcode: []int{0x2D, valID}, valid32BitMode: false, valid64BitMode: true},
	{name: "SUB", operands: []archOperand{{regMem16, anyValue}, {imm16, anyValue}}, opcode: []int{0x81, modRM5, valIW}, valid32BitMode: true, valid64BitMode: true},
	{name: "SUB", operands: []archOperand{{regMem16, anyValue}, {imm8, anyValue}}, opcode: []int{0x83, modRM5, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "SUB", operands: []archOperand{{regMem16, anyValue}, {reg16, anyValue}}, opcode: []int{0x29, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "SUB", operands: []archOperand{{regMem32, anyValue}, {imm32, anyValue}}, opcode: []int{0x81, modRM5, valID}, valid32BitMode: true, valid64BitMode: true},
	{name: "SUB", operands: []archOperand{{regMem32, anyValue}, {imm8, anyValue}}, opcode: []int{0x83, modRM5, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "SUB", operands: []archOperand{{regMem32, anyValue}, {reg32, anyValue}}, opcode: []int{0x29, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "SUB", operands: []archOperand{{regMem64, anyValue}, {imm32, anyValue}}, opcode: []int{0x81, modRM5, valID}, valid32BitMode: false, valid64BitMode: true},
	{name: "SUB", operands: []archOperand{{regMem64, anyValue}, {imm8, anyValue}}, opcode: []int{0x83, modRM5, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "SUB", operands: []archOperand{{regMem64, anyValue}, {reg64, anyValue}}, opcode: []int{0x29, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "SUB", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0x80, modRM5, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "SUB", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0x80, modRM5, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "SUB", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x28, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "SUB", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x28, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "SUB", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0x2B, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "SUB", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0x2B, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "SUB", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{0x2B, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "SUB", operands: []archOperand{{reg8, anyValue}, {regMem8, anyValue}}, opcode: []int{0x2A, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "SUB", operands: []archOperand{{reg8, anyValue}, {regMem8, anyValue}}, opcode: []int{0x2A, modRM}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand xmm1
	// todo: operand xmm1
	{name: "SWAPGS", operands: []archOperand{}, opcode: []int{0x0F, 0x01, 0xF8}, valid32BitMode: false, valid64BitMode: true},
	{name: "SYSCALL", operands: []archOperand{}, opcode: []int{0x0F, 0x05}, valid32BitMode: false, valid64BitMode: true},
	{name: "SYSENTER", operands: []archOperand{}, opcode: []int{0x0F, 0x34}, valid32BitMode: true, valid64BitMode: true},
	{name: "SYSEXIT", operands: []archOperand{}, opcode: []int{0x0F, 0x35}, valid32BitMode: true, valid64BitMode: true},
	{name: "SYSEXIT", operands: []archOperand{}, opcode: []int{0x0F, 0x35}, valid32BitMode: false, valid64BitMode: true},
	{name: "SYSRET", operands: []archOperand{}, opcode: []int{0x0F, 0x07}, valid32BitMode: false, valid64BitMode: true},
	{name: "SYSRET", operands: []archOperand{}, opcode: []int{0x0F, 0x07}, valid32BitMode: false, valid64BitMode: true},
	{name: "TEST", operands: []archOperand{{reg8, al}, {imm8, anyValue}}, opcode: []int{0xA8, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "TEST", operands: []archOperand{{reg16, ax}, {imm16, anyValue}}, opcode: []int{0xA9, valIW}, valid32BitMode: true, valid64BitMode: true},
	{name: "TEST", operands: []archOperand{{reg32, eax}, {imm32, anyValue}}, opcode: []int{0xA9, valID}, valid32BitMode: true, valid64BitMode: true},
	{name: "TEST", operands: []archOperand{{reg64, rax}, {imm32, anyValue}}, opcode: []int{0xA9, valID}, valid32BitMode: false, valid64BitMode: true},
	{name: "TEST", operands: []archOperand{{regMem16, anyValue}, {imm16, anyValue}}, opcode: []int{0xF7, modRM0, valIW}, valid32BitMode: true, valid64BitMode: true},
	{name: "TEST", operands: []archOperand{{regMem16, anyValue}, {reg16, anyValue}}, opcode: []int{0x85, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "TEST", operands: []archOperand{{regMem32, anyValue}, {imm32, anyValue}}, opcode: []int{0xF7, modRM0, valID}, valid32BitMode: true, valid64BitMode: true},
	{name: "TEST", operands: []archOperand{{regMem32, anyValue}, {reg32, anyValue}}, opcode: []int{0x85, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "TEST", operands: []archOperand{{regMem64, anyValue}, {imm32, anyValue}}, opcode: []int{0xF7, modRM0, valID}, valid32BitMode: false, valid64BitMode: true},
	{name: "TEST", operands: []archOperand{{regMem64, anyValue}, {reg64, anyValue}}, opcode: []int{0x85, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "TEST", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0xF6, modRM0, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "TEST", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0xF6, modRM0, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "TEST", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x84, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "TEST", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x84, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "TZCNT", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0xF3, 0x0F, 0xBC, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "TZCNT", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0xF3, 0x0F, 0xBC, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "TZCNT", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{0xF3, 0x0F, 0xBC, modRM}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand xmm1
	// todo: operand xmm1
	{name: "UD0", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0x0F, 0xFF, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "UD1", operands: []archOperand{}, opcode: []int{0x0F, 0xB9}, valid32BitMode: true, valid64BitMode: true},
	{name: "UD1", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0x0F, 0xB9, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "UD2", operands: []archOperand{}, opcode: []int{0x0F, 0x0B}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1{k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand ymm1
	// todo: operand zmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand ymm1
	// todo: operand zmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand ymm1
	// todo: operand zmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand ymm1
	// todo: operand zmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand k1 {k2}
	// todo: operand k1 {k2}
	// todo: operand k1 {k2}
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand k1 {k2}
	// todo: operand k1 {k2}
	// todo: operand k1 {k2}
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand k1 {k2}
	{name: "VCMPSDxmm1,xmm2,xmm3/m64,", operands: []archOperand{{imm8, anyValue}}, opcode: []int{todoOpcode}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand k1 {k2}
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1/m64
	// todo: operand xmm1/m64{er}
	// todo: operand xmm1/m64
	// todo: operand xmm1/m64{er}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1/m32
	// todo: operand xmm1/m32{er}
	// todo: operand xmm1/m32
	// todo: operand xmm1/m32{er}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1/m64
	// todo: operand xmm1/m64{sae}
	// todo: operand xmm1/m64
	// todo: operand xmm1/m64{sae}
	// todo: operand xmm1/m32
	// todo: operand xmm1/m32{sae}
	// todo: operand xmm1/m32
	// todo: operand xmm1/m32{sae}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand ymm1
	{name: "VLDMXCSR", operands: []archOperand{{mem32, anyValue}}, opcode: []int{todoOpcode}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand xmm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm2/m128 {k1}{z}
	// todo: operand xmm2/m128
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand ymm2/m256 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand zmm2/m512 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm2/m128 {k1}{z}
	// todo: operand xmm2/m128
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand ymm2/m256 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand zmm2/m512 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm2/m128
	// todo: operand ymm1
	// todo: operand ymm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm2/m128 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm2/m256 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand zmm2/m512 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm2/m128 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm2/m256 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand zmm2/m512 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm2/m128
	// todo: operand ymm1
	// todo: operand ymm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm2/m128 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm2/m256 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand zmm2/m512 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm2/m128 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm2/m256 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand zmm2/m512 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm2/m128 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm2/m256 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand zmm2/m512 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm2/m128 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm2/m256 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand zmm2/m512 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm2
	// todo: operand ymm2
	// todo: operand xmm2
	// todo: operand ymm2
	// todo: operand m128
	// todo: operand m128
	// todo: operand ymm1
	// todo: operand ymm1
	// todo: operand m512
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand ymm1
	// todo: operand zmm1
	// todo: operand m128
	// todo: operand m128
	// todo: operand ymm1
	// todo: operand ymm1
	// todo: operand m512
	// todo: operand m128
	// todo: operand m128
	// todo: operand ymm1
	// todo: operand ymm1
	// todo: operand m512
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1/m64
	// todo: operand xmm1/m64
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1{k1}{z}
	// todo: operand ymm1
	// todo: operand ymm1{k1}{z}
	// todo: operand zmm1{k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1{k1}{z}
	// todo: operand ymm1
	// todo: operand ymm1{k1}{z}
	// todo: operand zmm1{k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand ymm1
	// todo: operand zmm1
	// todo: operand k1 {k2}
	// todo: operand k1 {k2}
	// todo: operand k1 {k2}
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand k1 {k2}
	// todo: operand k1 {k2}
	// todo: operand k1 {k2}
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand k1 {k2}
	// todo: operand k1 {k2}
	// todo: operand k1 {k2}
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand k1 {k2}
	// todo: operand k1 {k2}
	// todo: operand k1 {k2}
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand k1 {k2}
	// todo: operand k1 {k2}
	// todo: operand k1 {k2}
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand k1 {k2}
	// todo: operand k1 {k2}
	// todo: operand k1 {k2}
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand k1 {k2}
	// todo: operand k1 {k2}
	// todo: operand k1 {k2}
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand k1 {k2}
	// todo: operand k1 {k2}
	// todo: operand k1 {k2}
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm2
	// todo: operand xmm2
	// todo: operand xmm2
	// todo: operand xmm2
	// todo: operand xmm2
	// todo: operand xmm2
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm2
	// todo: operand xmm2
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1{k1}{z}
	// todo: operand ymm1
	// todo: operand ymm1{k1}{z}
	// todo: operand zmm1{k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1{k1}{z}
	// todo: operand ymm1
	// todo: operand ymm1{k1}{z}
	// todo: operand zmm1{k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1{k1}{z}
	// todo: operand ymm1
	// todo: operand ymm1{k1}{z}
	// todo: operand zmm1{k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1{k1}{z}
	// todo: operand ymm1
	// todo: operand ymm1{k1}{z}
	// todo: operand zmm1{k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1{k1}{z}
	// todo: operand ymm1
	// todo: operand ymm1{k1}{z}
	// todo: operand zmm1{k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1{k1}{z}
	// todo: operand ymm1
	// todo: operand ymm1{k1}{z}
	// todo: operand zmm1{k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1{k1}{z}
	// todo: operand ymm1
	// todo: operand ymm1{k1}{z}
	// todo: operand zmm1{k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand ymm1
	// todo: operand zmm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand ymm1
	// todo: operand zmm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand ymm1 {k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand ymm1{k1}{z}
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand ymm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1{k1}{z}
	// todo: operand ymm1
	// todo: operand ymm1{k1}{z}
	// todo: operand zmm1{k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1{k1}{z}
	// todo: operand ymm1
	// todo: operand ymm1{k1}{z}
	// todo: operand zmm1{k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	{name: "VSTMXCSR", operands: []archOperand{{mem32, anyValue}}, opcode: []int{todoOpcode}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	// todo: operand xmm1 {k1}{z}
	// todo: operand xmm1
	// todo: operand ymm1 {k1}{z}
	// todo: operand ymm1
	// todo: operand zmm1 {k1}{z}
	{name: "WAIT", operands: []archOperand{}, opcode: []int{0x9B}, valid32BitMode: true, valid64BitMode: true},
	{name: "WBINVD", operands: []archOperand{}, opcode: []int{0x0F, 0x09}, valid32BitMode: true, valid64BitMode: true},
	{name: "WRFSBASE", operands: []archOperand{{reg32, anyValue}}, opcode: []int{0xF3, 0x0F, 0xAE, modRM2}, valid32BitMode: false, valid64BitMode: true},
	{name: "WRFSBASE", operands: []archOperand{{reg64, anyValue}}, opcode: []int{0xF3, 0x0F, 0xAE, modRM2}, valid32BitMode: false, valid64BitMode: true},
	{name: "WRGSBASE", operands: []archOperand{{reg32, anyValue}}, opcode: []int{0xF3, 0x0F, 0xAE, modRM3}, valid32BitMode: false, valid64BitMode: true},
	{name: "WRGSBASE", operands: []archOperand{{reg64, anyValue}}, opcode: []int{0xF3, 0x0F, 0xAE, modRM3}, valid32BitMode: false, valid64BitMode: true},
	{name: "WRMSR", operands: []archOperand{}, opcode: []int{0x0F, 0x30}, valid32BitMode: true, valid64BitMode: true},
	{name: "WRPKRU", operands: []archOperand{}, opcode: []int{0x0F, 0x01, 0xEF}, valid32BitMode: true, valid64BitMode: true},
	{name: "WRSSD", operands: []archOperand{{mem32, anyValue}, {reg32, anyValue}}, opcode: []int{0x0F, 0x38, 0xF6, todoOpcode}, valid32BitMode: true, valid64BitMode: true},
	{name: "WRSSQ", operands: []archOperand{{mem64, anyValue}, {reg64, anyValue}}, opcode: []int{0x0F, 0x38, 0xF6, todoOpcode}, valid32BitMode: false, valid64BitMode: true},
	{name: "WRUSSD", operands: []archOperand{{mem32, anyValue}, {reg32, anyValue}}, opcode: []int{0x66, 0x0F, 0x38, 0xF5, todoOpcode}, valid32BitMode: true, valid64BitMode: true},
	{name: "WRUSSQ", operands: []archOperand{{mem64, anyValue}, {reg64, anyValue}}, opcode: []int{0x66, 0x0F, 0x38, 0xF5, todoOpcode}, valid32BitMode: false, valid64BitMode: true},
	{name: "XABORT", operands: []archOperand{{imm8, anyValue}}, opcode: []int{0xC6, 0xF8, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "XACQUIRE", operands: []archOperand{}, opcode: []int{0xF2}, valid32BitMode: true, valid64BitMode: true},
	{name: "XADD", operands: []archOperand{{regMem16, anyValue}, {reg16, anyValue}}, opcode: []int{0x0F, 0xC1, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "XADD", operands: []archOperand{{regMem32, anyValue}, {reg32, anyValue}}, opcode: []int{0x0F, 0xC1, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "XADD", operands: []archOperand{{regMem64, anyValue}, {reg64, anyValue}}, opcode: []int{0x0F, 0xC1, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "XADD", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x0F, 0xC0, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "XADD", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x0F, 0xC0, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "XBEGIN", operands: []archOperand{{rel16, anyValue}}, opcode: []int{0xC7, 0xF8}, valid32BitMode: true, valid64BitMode: true},
	{name: "XBEGIN", operands: []archOperand{{rel32, anyValue}}, opcode: []int{0xC7, 0xF8}, valid32BitMode: true, valid64BitMode: true},
	// todo: operand r16op
	// todo: operand r32op
	// todo: operand r64op
	{name: "XCHG", operands: []archOperand{{regMem16, anyValue}, {reg16, anyValue}}, opcode: []int{0x87, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "XCHG", operands: []archOperand{{regMem32, anyValue}, {reg32, anyValue}}, opcode: []int{0x87, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "XCHG", operands: []archOperand{{regMem64, anyValue}, {reg64, anyValue}}, opcode: []int{0x87, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "XCHG", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x86, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "XCHG", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x86, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "XCHG", operands: []archOperand{{reg16, anyValue}, {reg16, ax}}, opcode: []int{0x90, plusRW}, valid32BitMode: true, valid64BitMode: true},
	{name: "XCHG", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0x87, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "XCHG", operands: []archOperand{{reg32, anyValue}, {reg32, eax}}, opcode: []int{0x90, plusRD}, valid32BitMode: true, valid64BitMode: true},
	{name: "XCHG", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0x87, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "XCHG", operands: []archOperand{{reg64, anyValue}, {reg64, rax}}, opcode: []int{0x90, plusRD}, valid32BitMode: false, valid64BitMode: true},
	{name: "XCHG", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{0x87, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "XCHG", operands: []archOperand{{reg8, anyValue}, {regMem8, anyValue}}, opcode: []int{0x86, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "XCHG", operands: []archOperand{{reg8, anyValue}, {regMem8, anyValue}}, opcode: []int{0x86, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "XEND", operands: []archOperand{}, opcode: []int{0x0F, 0x01, 0xD5}, valid32BitMode: true, valid64BitMode: true},
	{name: "XGETBV", operands: []archOperand{}, opcode: []int{0x0F, 0x01, 0xD0}, valid32BitMode: true, valid64BitMode: true},
	{name: "XLATB", operands: []archOperand{}, opcode: []int{0xD7}, valid32BitMode: true, valid64BitMode: true},
	{name: "XLATB", operands: []archOperand{}, opcode: []int{0xD7}, valid32BitMode: false, valid64BitMode: true},
	{name: "XOR", operands: []archOperand{{reg8, al}, {imm8, anyValue}}, opcode: []int{0x34, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "XOR", operands: []archOperand{{reg16, ax}, {imm16, anyValue}}, opcode: []int{0x35, valIW}, valid32BitMode: true, valid64BitMode: true},
	{name: "XOR", operands: []archOperand{{reg32, eax}, {imm32, anyValue}}, opcode: []int{0x35, valID}, valid32BitMode: true, valid64BitMode: true},
	{name: "XOR", operands: []archOperand{{reg64, rax}, {imm32, anyValue}}, opcode: []int{0x35, valID}, valid32BitMode: false, valid64BitMode: true},
	{name: "XOR", operands: []archOperand{{regMem16, anyValue}, {imm16, anyValue}}, opcode: []int{0x81, modRM6, valIW}, valid32BitMode: true, valid64BitMode: true},
	{name: "XOR", operands: []archOperand{{regMem16, anyValue}, {imm8, anyValue}}, opcode: []int{0x83, modRM6, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "XOR", operands: []archOperand{{regMem16, anyValue}, {reg16, anyValue}}, opcode: []int{0x31, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "XOR", operands: []archOperand{{regMem32, anyValue}, {imm32, anyValue}}, opcode: []int{0x81, modRM6, valID}, valid32BitMode: true, valid64BitMode: true},
	{name: "XOR", operands: []archOperand{{regMem32, anyValue}, {imm8, anyValue}}, opcode: []int{0x83, modRM6, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "XOR", operands: []archOperand{{regMem32, anyValue}, {reg32, anyValue}}, opcode: []int{0x31, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "XOR", operands: []archOperand{{regMem64, anyValue}, {imm32, anyValue}}, opcode: []int{0x81, modRM6, valID}, valid32BitMode: false, valid64BitMode: true},
	{name: "XOR", operands: []archOperand{{regMem64, anyValue}, {imm8, anyValue}}, opcode: []int{0x83, modRM6, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "XOR", operands: []archOperand{{regMem64, anyValue}, {reg64, anyValue}}, opcode: []int{0x31, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "XOR", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0x80, modRM6, valIB}, valid32BitMode: true, valid64BitMode: true},
	{name: "XOR", operands: []archOperand{{regMem8, anyValue}, {imm8, anyValue}}, opcode: []int{0x80, modRM6, valIB}, valid32BitMode: false, valid64BitMode: true},
	{name: "XOR", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x30, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "XOR", operands: []archOperand{{regMem8, anyValue}, {reg8, anyValue}}, opcode: []int{0x30, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "XOR", operands: []archOperand{{reg16, anyValue}, {regMem16, anyValue}}, opcode: []int{0x33, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "XOR", operands: []archOperand{{reg32, anyValue}, {regMem32, anyValue}}, opcode: []int{0x33, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "XOR", operands: []archOperand{{reg64, anyValue}, {regMem64, anyValue}}, opcode: []int{0x33, modRM}, valid32BitMode: false, valid64BitMode: true},
	{name: "XOR", operands: []archOperand{{reg8, anyValue}, {regMem8, anyValue}}, opcode: []int{0x32, modRM}, valid32BitMode: true, valid64BitMode: true},
	{name: "XOR", operands: []archOperand{{reg8, anyValue}, {regMem8, anyValue}}, opcode: []int{0x32, modRM}, valid32BitMode: false, valid64BitMode: true},
	// todo: operand xmm1
	// todo: operand xmm1
	{name: "XRELEASE", operands: []archOperand{}, opcode: []int{0xF3}, valid32BitMode: true, valid64BitMode: true},
	{name: "XSAVEOPT", operands: []archOperand{{mem, anyValue}}, opcode: []int{0x0F, 0xAE, modRM6}, valid32BitMode: true, valid64BitMode: true},
	{name: "XSAVEOPT64", operands: []archOperand{{mem, anyValue}}, opcode: []int{0x0F, 0xAE, modRM6}, valid32BitMode: true, valid64BitMode: true},
	{name: "XSETBV", operands: []archOperand{}, opcode: []int{0x0F, 0x01, 0xD1}, valid32BitMode: true, valid64BitMode: true},
	{name: "XTEST", operands: []archOperand{}, opcode: []int{0x0F, 0x01, 0xD6}, valid32BitMode: true, valid64BitMode: true},
}
