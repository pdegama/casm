// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// Assembler CLI

package main

import (
	"hellocomputers/casm/x86_64"
	"os"
)

func main() {

	asmFilePath := os.Args[1]

	// mew x86_64 program
	asmProg := x86_64.NewX86_64()
	asmProg.SetAsmFile(asmFilePath) // set asm file
	asmProg.Assemble()              // assemble

}
