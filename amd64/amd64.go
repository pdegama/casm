// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// amd64 architecture

package amd64

// amd64 architecture structure
type amd64 struct {
	asmFile string
}

// set asm file
func (arch *amd64) SetAsmFile(fileName string) {
	arch.asmFile = fileName
}

// assemble asm file
func (arch *amd64) Assemble() {
	// call assemble function
	assemble(arch)
}

// new amd64 architecture instant
func NewAMD64() amd64 {
	return amd64{}
}
