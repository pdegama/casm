// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// binary generation

package x86_64

import (
	"fmt"
)

// binary gen structure
type binaryGen struct {
	lines       []asmLine        // asm lines
	bitMode     int              // bit mode 16, 32 or 64
	bytesStruct []bytesStructure // bytes
	bianry      []uint8          // binary
}

// bytes structure
type bytesStructure struct {
	pos    int     // postion
	label  bool    // is lable
	name   string  // if label then name
	bytes  []uint8 // bytes
	len    int     // bytes length
	labels []label // label
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
func (s *binaryGen) genBinary() {

	s.bianry = []uint8{} // clear binary

	for _, instBytes := range s.bytesStruct {
		// loop of bytes structure

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

				// if line is insrtuction
				instBytes, err := instructionGen(line, s.bitMode) // gen code for instruction
				if err != nil {
					// if error
					tErr := fmt.Errorf("%s %v:%v %v", errorStr, *line.filePath, line.index+1, err)
					errs = append(errs, tErr)
				} else {
					// if not error then append to bytes struct
					s.bytesStruct = append(s.bytesStruct, instBytes)
				}

			case lineLabel:

				// if line is label line

				labelStruct, err := parseLabel(line) // parse label
				if err != nil {                      // if error
					tErr := fmt.Errorf("%s %v:%v %v", errorStr, *line.filePath, line.index+1, err)
					errs = append(errs, tErr)
				} else {
					// if not error then append to bytes struct
					s.bytesStruct = append(s.bytesStruct, labelStruct)
				}

			case lineData:

				// if line is data line
				dataBytes, err := dataBytes(line) // get data byte
				if err != nil {                   // if error
					tErr := fmt.Errorf("%s %v:%v %v", errorStr, *line.filePath, line.index+1, err)
					errs = append(errs, tErr)
				} else {
					// if not error then append to bytes struct
					s.bytesStruct = append(s.bytesStruct, dataBytes)
				}

			case lineModulo:

				// if line is modulo
				err := parseModulo(line, s) // parse modulo
				if err != nil {
					tErr := fmt.Errorf("%s %v:%v %v", errorStr, *line.filePath, line.index+1, err)
					errs = append(errs, tErr)
				}

			}
		}
	}

	return errs
}
