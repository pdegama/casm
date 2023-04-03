// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// Assembler CLI

package main

import (
	"hellocomputers/casm/amd64"
	"os"
)

func main() {

	asmFilePath := os.Args[1]

	// mew amd64 program
	asmProg := amd64.NewAMD64()
	asmProg.SetAsmFile(asmFilePath) // set asm file
	asmProg.Assemble()                         // assemble

}
