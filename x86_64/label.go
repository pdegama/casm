// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// label

package x86_64

import "fmt"

// label struct
type label struct {
	labelPos  int         // position of label value
	labelName string      // label name
	labelType operandType // valuse operand type imm*
	value     uint        // if value
	disp      bool        // displesment
}

// set label
func (b *binaryGen) setLabel(offset uint) {

	errs := []error{}
	_ = errs

	// loop of bytes structure
	for bsIndex, bs := range b.bytesStruct {

		fmt.Println(bs)

		for _, l := range bs.labels { // loop of labels

			lPos, err := b.getLabelPos(l.labelName)
			if err != nil {
				errs = append(errs, err)
				continue
			}

			lPos = lPos + offset

			switch l.labelType {
			case imm8:
				// label type is imm8

				leBytes := uint8le(uint8(lPos))
				for i, v := range leBytes {
					b.bytesStruct[bsIndex].bytes[l.labelPos+i] = v
				}

			case imm16:
				// label type is imm16

				leBytes := uint16le(uint16(lPos))
				for i, v := range leBytes {
					b.bytesStruct[bsIndex].bytes[l.labelPos+i] = v
				}

			case imm32:
				// label type is imm32
				leBytes := uint32le(uint32(lPos))
				for i, v := range leBytes {
					b.bytesStruct[bsIndex].bytes[l.labelPos+i] = v
				}

			case imm64:
				// label type is imm64

				leBytes := uint64le(uint64(lPos))
				for i, v := range leBytes {
					b.bytesStruct[bsIndex].bytes[l.labelPos+i] = v
				}

			}

		}

		fmt.Println(b.bytesStruct[bsIndex])

	}

}

// get label position
func (b *binaryGen) getLabelPos(l string) (uint, error) {

	// loop of bytes structure
	for _, bs := range b.bytesStruct {
		if bs.label {
			if bs.name == l {
				// if bs is label then return pos
				return bs.pos, nil
			}
		}
	}

	// if no any label found then return
	return 0x00, fmt.Errorf("%v is undeclared label", l)

}
