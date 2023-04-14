// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// x86_64 architecture

package x86_64

// x86_64 architecture structure
type x86_64 struct {
	asmFile string
}

// set asm file
func (arch *x86_64) SetAsmFile(fileName string) {
	arch.asmFile = fileName
}

// assemble asm file
func (arch *x86_64) Assemble() {
	// call assemble function
	assemble(arch)
}

// new x86_64 architecture instant
func NewX86_64() x86_64 {
	return x86_64{}
}
