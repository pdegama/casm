// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// modrm

package x86_64

import "fmt"

// add modRM byte
func addModRM(opcode *archOpcode, inst *instruction, bitMode int, pf *prefix) ([]uint8, error) {

	// calculate modRM and return
	if len(inst.operands) != 2 {
		return nil, fmt.Errorf("internal error: todo: operand size is not two")
	}

	modRMrmOper := &operand{t: undefinedOperand}  // modrm r/m operand
	modRMregOper := &operand{t: undefinedOperand} // modrm reg operand

	// loop of arch operands
	for aOperIndex, aOper := range opcode.operands {
		if isMemOperand(aOper.t) {
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

	// modRMreg info
	modRMregInfo, err := registerInfo(int(modRMregOper.v))
	if err != nil {
		return nil, err
	}

	// check modRMreg is valid or not
	err = registerIsValid(modRMregInfo, bitMode)
	if err != nil {
		return nil, err
	}

	// check operand override prefix
	err = checkOperandOverride(modRMregOper, bitMode, pf)
	if err != nil {
		return nil, err
	}

	// check rex.r prefix
	err = checkREXrPrexif(modRMregOper, bitMode, pf)
	if err != nil {
		return nil, err
	}

	// get reg field
	regField, err := modRMregField(*modRMregOper)
	if err != nil {
		return nil, err
	}

	fmt.Println(modRMrmOper, modRMregOper, regField)

	// calc modrm
	modRMByte, err := calcModRM(modRMrmOper, regField, bitMode, pf)
	if err != nil {
		return nil, err
	}

	return modRMByte, nil
}

// calc modrm
func calcModRM(rmOper *operand, regField int, bitMode int, pf *prefix) ([]uint8, error) {

	modrmBytes := []uint8{} // modrm bytes

	if rmOper.t == mem {
		// todo: modrm mem operand support
		//fmt.Println(rmOper.m)

		switch len(rmOper.m) {
		case 1:
			// if only one operand in memory operand
			memOper := rmOper.m[0]

			if isRegOperand(memOper.t) {

				// get mem register info
				regInfo, err := registerInfo(int(memOper.v))
				if err != nil {
					return nil, err
				}

				// check register is valid in mem or not
				err = registerIsValidInMem(regInfo, bitMode)
				if err != nil {
					return nil, err
				}

				// check address override prefix
				err = checkAddressOverride(&memOper, bitMode, pf)
				if err != nil {
					return nil, err
				}

				// check rex.b prefix
				err = checkREXbPrexif(&memOper, bitMode, pf)
				if err != nil {
					return nil, err
				}

				// get mem field
				modrmMemField, err := modRMmemRegField(regInfo)
				if err != nil {
					return nil, err
				}

				// if bitmode
				if regInfo.bitSize == 16 {
					if modrmMemField == 6 {
						return threeMemOperModRM(
							[]operand{
								memOper,
								{t: operPrePlus, v: 0, m: nil, l: false},
								{t: imm, v: 0, m: nil, l: false},
							},
							regField,
							bitMode,
							pf,
						)
					}
				} else if regInfo.bitSize == 32 || regInfo.bitSize == 64 {
					if modrmMemField == 5 {
						return threeMemOperModRM(
							[]operand{
								memOper,
								{t: operPrePlus, v: 0, m: nil, l: false},
								{t: imm, v: 0, m: nil, l: false},
							},
							regField,
							bitMode,
							pf,
						)
					} else if modrmMemField == 4 {
						return nil, fmt.Errorf("todo: %v register in address", regInfo.name)
					}
				}

				// gen modrm byte
				modRMByte := modRMbyte(0b00, regField, modrmMemField)
				modrmBytes = append(modrmBytes, uint8(modRMByte))

				return modrmBytes, nil // return modrm byte

			} else if isImmOperand(memOper.t) {

				if memOper.l {
					// if this is label then todo error
					return nil, fmt.Errorf("todo: modrm label disp mem opernad not support")
				}

				switch memOper.t {
				case imm:
					// this operand is not fix imm

					switch bitMode {
					case 16:
						// if bit mode is 16-bit

						immVal := uint16(memOper.v) // imm value

						immLeByte := uint16le(immVal) // convert imm value to little endian

						modrmMemField := 6 // modrm r/m field is 6 is 16-bit modrm

						// gen modrm byte
						modRMByte := modRMbyte(0b00, regField, modrmMemField)
						modrmBytes = append(modrmBytes, uint8(modRMByte))

						// and append disp imm bytes
						modrmBytes = append(modrmBytes, immLeByte...)

					case 32:
						// if bit mode is 32-bit

						immVal := uint32(memOper.v) // imm value

						immLeByte := uint32le(immVal) // convert imm value to little endian

						modrmMemField := 5 // modrm r/m field is 5 in 32-bit modrm

						// gen modrm byte
						modRMByte := modRMbyte(0b00, regField, modrmMemField)
						modrmBytes = append(modrmBytes, uint8(modRMByte))

						// and append disp imm bytes
						modrmBytes = append(modrmBytes, immLeByte...)

					case 64:
						// if bit mode is 64-bit

						/*
							64-bit mode support 32-bit (dword)
							disp imm address then convert
							imm value convert to 32-bit value
						*/

						immVal := uint32(memOper.v) // imm value

						immLeByte := uint32le(immVal) // convert imm value to little endian

						// 64-bit direct disp imm address is require SIB
						modrmMemField := 4 // SIB modrm r/m field is 4 in 64-bit modrm

						// gen modrm byte
						modRMByte := modRMbyte(0b00, regField, modrmMemField)
						modrmBytes = append(modrmBytes, uint8(modRMByte))

						// append 25 SIB byte because 0x25 is base disp
						modrmBytes = append(modrmBytes, 0x25)

						// and append disp imm bytes
						modrmBytes = append(modrmBytes, immLeByte...)

					}

				case imm8, imm16:

					if bitMode == 64 {
						/*
							if bit mode is 64 then return
							error because 64 bit is not support
							16-bit bisp imm value
						*/
						return nil, fmt.Errorf("disp imm16 is not support in 64-bit")
					}

					immVal := uint16(memOper.v) // convert imm value to 16-bit value

					immLeByte := uint16le(immVal) // convert imm value to little endian

					// check address override prefix
					err := checkAddressOverride(&operand{t: imm16}, bitMode, pf)
					if err != nil {
						return nil, err
					}

					modrmMemField := 6 // modrm r/m field is 6 is 16-bit modrm

					// gen modrm byte
					modRMByte := modRMbyte(0b00, regField, modrmMemField)
					modrmBytes = append(modrmBytes, uint8(modRMByte))

					// and append disp imm bytes
					modrmBytes = append(modrmBytes, immLeByte...)

				case imm32, imm64:

					if memOper.t == imm64 && bitMode != 64 {
						/*
							if bit mode is not 64 amnd imm is
							64-bit then return error because
							64-bit is not support 16-bit bisp
							imm value
						*/
						return nil, fmt.Errorf("disp imm64 is only support in 64-bit")
					}

					immVal := uint32(memOper.v) // convert imm value to 32-bit value

					immLeByte := uint32le(immVal) // convert imm value to little endian

					if bitMode == 64 {

						// 64-bit direct disp imm address is require SIB
						modrmMemField := 4 // SIB modrm r/m field is 4 in 64-bit modrm

						// gen modrm byte
						modRMByte := modRMbyte(0b00, regField, modrmMemField)
						modrmBytes = append(modrmBytes, uint8(modRMByte))

						// append 25 SIB byte because 0x25 is base disp
						modrmBytes = append(modrmBytes, 0x25)

					} else if bitMode == 32 || bitMode == 16 {

						// check address override prefix
						err := checkAddressOverride(&operand{t: imm32}, bitMode, pf)
						if err != nil {
							return nil, err
						}

						modrmMemField := 5 // modrm r/m field is 5 in 32-bit or 64-bit modrm

						// gen modrm byte
						modRMByte := modRMbyte(0b00, regField, modrmMemField)
						modrmBytes = append(modrmBytes, uint8(modRMByte))

					}

					// and append disp imm bytes
					modrmBytes = append(modrmBytes, immLeByte...)

				default:
					// if this is label then todo error
					return nil, fmt.Errorf("internal error: modrm")
				}

				return modrmBytes, nil
			}

			return nil, fmt.Errorf("internal error: modrm")

		case 3:

			modrmBytes, err := threeMemOperModRM(rmOper.m, regField, bitMode, pf)
			if err != nil {
				return nil, err
			}
			return modrmBytes, nil

		default:
			//
			return nil, fmt.Errorf("invalid syntax memory operand")
		}

	}

	if isRegOperand(rmOper.t) {
		// if rm operand is register

		// get r/m register info
		modRMrmRegInfo, err := registerInfo(int(rmOper.v))
		if err != nil {
			// if error then return error
			return nil, err
		}

		// check modRMreg is valid or not
		err = registerIsValid(modRMrmRegInfo, bitMode)
		if err != nil {
			return nil, err
		}

		// check rex.b prefix
		err = checkREXbPrexif(rmOper, bitMode, pf)
		if err != nil {
			return nil, err
		}

		// gen modrm byte
		modRMByte := modRMbyte(0b11, regField, modRMrmRegInfo.baseOffset)
		modrmBytes = append(modrmBytes, uint8(modRMByte))

		return modrmBytes, nil // return modrm byte

	}

	return nil, fmt.Errorf("internal error: modrm")
}

// 3 mem operand modrm
func threeMemOperModRM(opers []operand, regField int, bitMode int, pf *prefix) ([]uint8, error) {

	modrmBytes := []uint8{} // modrm bytes

	regOper := opers[0]        // reg operand firest operand
	operantionOper := opers[1] // operation opernad
	immOper := opers[2]        // imm operand

	if !isRegOperand(regOper.t) {
		// if first operand is not register
		return nil, fmt.Errorf("invalid mem operand")
	}

	switch operantionOper.t {
	case operPrePlus, operPreMinus:
		// pass
	default:
		// if seconde operand os not operation operand
		return nil, fmt.Errorf("invalid mem operand")
	}

	if !isImmOperand(immOper.t) {
		// if third operand is not imm operand
		return nil, fmt.Errorf("invalid mem operand")

	}

	regOperInfo, err := registerInfo(int(regOper.v))
	if err != nil {
		return nil, err
	}

	// check register is valid in mem or not
	err = registerIsValidInMem(regOperInfo, bitMode)
	if err != nil {
		return nil, err
	}

	// check address override prefix
	err = checkAddressOverride(&regOper, bitMode, pf)
	if err != nil {
		return nil, err
	}

	// check rex.b prefix
	err = checkREXbPrexif(&regOper, bitMode, pf)
	if err != nil {
		return nil, err
	}

	// get mem field
	modrmMemField, err := modRMmemRegField(regOperInfo)
	if err != nil {
		return nil, err
	}

	modRMmode := 0b10
	var immLeByte []uint8

	if immOper.l {
		// if this is label then todo error
		return nil, fmt.Errorf("todo: modrm label disp mem opernad not support")
	}

	if immOper.v <= 0xff {
		// if imm value is 8-bit
		modRMmode = 0b01
		immVal := uint8(immOper.v)

		if operantionOper.t == operPreMinus {
			// if operation is minus
			immVal = 0xff - immVal + 1
		}

		immLeByte = uint8le(immVal)

	} else {

		switch regOperInfo.bitSize {
		case 16:
			// if register is 16-bit register
			modRMmode = 0b10

			immVal := uint16(immOper.v)

			if operantionOper.t == operPreMinus {
				// if operation is minus
				immVal = 0xffff - immVal + 1
			}

			immLeByte = uint16le(immVal)

		case 32, 64:
			// if register is 16-bit or 64-bit register
			modRMmode = 0b10

			immVal := uint32(immOper.v)

			if operantionOper.t == operPreMinus {
				// if operation is minus
				immVal = 0xffffffff - immVal + 1
			}

			immLeByte = uint32le(immVal)
		default:
			return nil, fmt.Errorf("internal error: modrm")
		}

	}

	// gen modrm byte
	modRMByte := modRMbyte(modRMmode, regField, modrmMemField)
	modrmBytes = append(modrmBytes, uint8(modRMByte))
	modrmBytes = append(modrmBytes, immLeByte...)

	return modrmBytes, nil

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

// modRM mem field
func modRMmemRegField(r register) (int, error) {

	switch r.bitSize {
	case 16:
		// if register is 16 bit

		switch r.globleIndex {
		case si:
			return 4, nil
		case di:
			return 5, nil
		case bp:
			return 6, nil
		case bx:
			return 7, nil
		default:
			return 0, fmt.Errorf("%v register is not support in modrm address", r.name)
		}

	case 32, 64:
		// if register is 32 or 64 bit

		return r.baseOffset, nil

	}

	return 0, fmt.Errorf("invalid modrm reg mem")

}

// get modrm byte
func modRMbyte(mod int, reg int, rm int) int {

	modRMByte := mod<<6 | reg<<3 | rm<<0

	return modRMByte

}
