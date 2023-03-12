// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// Assembler CLI

package main

import (
	"hellocomputers/casm/amd64"
)

func main() {

	// mew amd64 program
	asmProg := amd64.NewAMD64()
	asmProg.SetAsmFile("./test_files/hello.asm") // set asm file
	asmProg.Assemble()                         // assemble

}
