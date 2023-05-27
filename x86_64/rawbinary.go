// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// raw binary

package x86_64

import "os"

// rawbin structure
type rawbin struct {
	rawbinBytes []uint8
}

// build raw bin file
func (e *rawbin) buildBin(b *binaryGen, offset uint) []error {

	b.mergeSegments() // merge segment
	b.setPos()        // set bytes pos

	errs := b.setLabel(offset) // set labels
	if len(errs) != 0 {
		return errs
	}

	b.genBinary()        // genret binary
	bin := b.getbinary() // get binary

	e.addBytes(bin...)

	return []error{}

}

// add bytes
func (e *rawbin) addBytes(b ...byte) {
	e.rawbinBytes = append(e.rawbinBytes, b...)
}

// save bin file
func (e *rawbin) saveBinFile(fName string) {
	err := os.WriteFile(fName, e.rawbinBytes, 0777)
	if err != nil {
		panic(err)
	}
}
