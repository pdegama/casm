// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// Assembler CLI

package main

import (
	"flag"
	"hellocomputers/casm/x86_64"
	"os"
)

func main() {

	asmFile := flag.String("ifile", "", "assembly file")
	binFile := flag.String("ofile", "./a.out", "output file")
	fmtType := flag.String("filef", "elf64", "format [ elf64 | bin ]")

	flag.Parse()
	// fmt.Println(*asmFile, *binFile, *fmtType)

	if *asmFile == "" {
		if len(flag.CommandLine.Args()) > 0 {
			*asmFile = flag.CommandLine.Args()[0]
		} else {
			flag.PrintDefaults()
			os.Exit(0)
		}
	}

	fType := 1

	switch *fmtType {
	case "elf64":
		fType = 1
	case "bin":
		fType = 0
	default:
		flag.PrintDefaults()
		os.Exit(0)
	}

	// new x86_64 program
	asmProg := x86_64.NewX86_64()
	asmProg.SetAsmFile(*asmFile)      // set asm file
	asmProg.Assemble(*binFile, fType) // assemble

}
