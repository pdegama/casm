// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// binary generation

package x86_64

import (
	"fmt"
)

// segments
const (
	textSegment = 0
	dataSegment = 1
)

// binary gen structure
type binaryGen struct {
	lines           []asmLine        // asm lines
	bitMode         int              // bit mode 16, 32 or 64
	bytesStruct     []bytesStructure // bytes
	textBytesStruct []bytesStructure // text bytes struct
	dataBytesStruct []bytesStructure // data bytes struct
	bianry          []uint8          // binary
	segment         int              // current segent (text, data)
}

// bytes structure
type bytesStructure struct {
	pos    uint    // postion
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

// binary generation
func (b *binaryGen) gen() []error {

	errs := []error{} // errors

	for _, line := range b.lines {

		if len(line.tokens) != 0 {

			switch line.lineType {
			case lineInst:

				// if line is insrtuction
				instBytesStruct, err := instructionGen(line, b.bitMode) // gen code for instruction
				if err != nil {
					// if error
					tErr := fmt.Errorf("%s %v:%v %v", errorStr, *line.filePath, line.index+1, err)
					errs = append(errs, tErr)
				} else {
					// if not error then append to bytes struct
					b.addByteStruct(instBytesStruct)
				}

			case lineLabel:

				// if line is label line

				labelStruct, err := parseLabel(line) // parse label
				if err != nil {                      // if error
					tErr := fmt.Errorf("%s %v:%v %v", errorStr, *line.filePath, line.index+1, err)
					errs = append(errs, tErr)
				} else {
					// if not error then append to bytes struct
					b.addByteStruct(labelStruct)
				}

			case lineData:

				// if line is data line
				dataBytesStruct, err := dataBytes(line) // get data byte
				if err != nil {                         // if error
					tErr := fmt.Errorf("%s %v:%v %v", errorStr, *line.filePath, line.index+1, err)
					errs = append(errs, tErr)
				} else {
					// if not error then append to bytes struct
					b.addByteStruct(dataBytesStruct)
				}

			case lineModulo:

				// if line is modulo
				err := parseModulo(line, b) // parse modulo
				if err != nil {
					tErr := fmt.Errorf("%s %v:%v %v", errorStr, *line.filePath, line.index+1, err)
					errs = append(errs, tErr)
				}

			}
		}
	}

	return errs
}

// add bytes struct
func (b *binaryGen) addByteStruct(bd bytesStructure) {

	if b.segment == textSegment {
		/*
			if current segment is text segment
			then append byte struct to text segment
		*/
		b.textBytesStruct = append(b.textBytesStruct, bd)
	} else if b.segment == dataSegment {
		/*
			if current segment is data segment
			then append byte struct to data segment
		*/
		b.dataBytesStruct = append(b.dataBytesStruct, bd)
	}

}

// merge segments
func (b *binaryGen) mergeSegments() {

	// first append text segments to bytes structure
	b.bytesStruct = append(b.bytesStruct, b.textBytesStruct...)

	// second append data segments to bytes structure
	b.bytesStruct = append(b.bytesStruct, b.dataBytesStruct...)

}

// set position
func (b *binaryGen) setPos() {

	var cPos uint // current pos
	cPos = 0

	// loop of bytes structure
	for bsIndex, bs := range b.bytesStruct {
		b.bytesStruct[bsIndex].pos = cPos // set curren position
		cPos += uint(bs.len)              // byte len add to current position
	}

}

// get insts
func (b *binaryGen) genBinary() {

	b.bianry = []uint8{} // clear binary

	for _, instBytes := range b.bytesStruct {
		// loop of bytes structure

		// append inst bytes to array
		b.bianry = append(b.bianry, instBytes.bytes...)

	}

}

// get row binary
func (s *binaryGen) getbinary() []uint8 {
	// return binary
	return s.bianry
}
