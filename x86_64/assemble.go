// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// assemble

package x86_64

import (
	"fmt"
)

// assemble asm file
func assemble(arch *x86_64, oFile string, fFmt int) {

	// tokenization
	lines, err := lexer(arch.asmFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// tmp
	/* for _, line := range lines {
		fmt.Println(line)
	} */

	// parse asm lines
	lines, errs := parseLines(lines)
	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		return
	}

	binGen := binaryGen{}
	binGen.setBitMode(16)     // set bit mode
	binGen.setAsmLines(lines) // set asm lines

	errs = binGen.gen() // gen binary
	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		return
	}

	switch fFmt {
	case fFmtElf64:
		// if elf file
		elfGen := elf{}
		errs = elfGen.buildELF(&binGen)
		if len(errs) != 0 {
			for _, err := range errs {
				fmt.Printf("%s %v\n", errorStr, err)
			}
			return
		}
		elfGen.saveBinFile(oFile)

	case fFmtRawBin:
		// if raw bin
		rawBinGen := rawbin{}
		errs := rawBinGen.buildBin(&binGen)
		if len(errs) != 0 {
			for _, err := range errs {
				fmt.Printf("%s %v\n", errorStr, err)
			}
			return
		}
		rawBinGen.saveBinFile(oFile)

	}
	// tmp
	/* for _, line := range lines {
		fmt.Println(line)
	} */

}
