// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// binary generation

package x86_64

import (
	"fmt"
)

// binary gen structure
type binaryGen struct {
	lines   []asmLine          // asm lines
	bitMode int                // bit mode 16, 32 or 64
	insts   []instructionBytes // insts
	bianry  []uint8            // binary
}

// set asm lines
func (s *binaryGen) setAsmLines(lines []asmLine) {
	s.lines = lines // assign asm lines to structure
}

// set bit mode
func (s *binaryGen) setBitMode(bitMode int) {
	s.bitMode = bitMode // assign bit mode  to structure
}

// get insts
func (s *binaryGen) getInstBytes() []instructionBytes {
	return s.insts
}

// get insts
func (s *binaryGen) genBinary() {

	s.bianry = []uint8{} // clear binary

	for _, instBytes := range s.insts {

		// loop of inst bytes
		// append inst bytes to array
		s.bianry = append(s.bianry, instBytes.bytes...)
	
	}

}

// get row binary
func (s *binaryGen) getbinary() []uint8 {
	// return binary
	return s.bianry
}

// binary generation
func (s *binaryGen) gen() []error {

	errs := []error{} // errors

	for _, line := range s.lines {

		if len(line.tokens) != 0 {

			switch line.lineType {
			case lineInst:

				// if line type is insrtuction
				instByteCode, err := instructionGen(line, s.bitMode) // gen code for instruction
				if err != nil {
					// if error
					tErr := fmt.Errorf("%s %v:%v %v", errorStr, *line.filePath, line.index+1, err)
					errs = append(errs, tErr)
				} else {
					s.insts = append(s.insts, instByteCode)
				}

			case lineData:
				_, err := dataBytes(line)
				if err != nil {
					tErr := fmt.Errorf("%s %v:%v %v", errorStr, *line.filePath, line.index+1, err)
					errs = append(errs, tErr)
				}

			case lineModulo:

				// if line is modulo
				err := parseModulo(line, s)
				if err != nil {
					tErr := fmt.Errorf("%s %v:%v %v", errorStr, *line.filePath, line.index+1, err)
					errs = append(errs, tErr)
				}

			}
		}
	}

	return errs
}
