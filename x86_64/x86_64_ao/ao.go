// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// x86_64.csv (x86_64 architecture opcode) to go file

package main

import (
	"os"
)

// main function
func main() {

	// read csv file
	row := readCSV()

	// make and open table file
	tableFile := makeArchOpcodeTableFile()

	// set table file header
	setArchOpcodeFileHeader(tableFile)

	for _, v := range row {

		mName, mOperands, valid32Bit, valid64Bit, err := parseData(v)
		if err != nil {
			tableFile.WriteString("\t// ")
			tableFile.WriteString(err.Error())
			tableFile.WriteString("\n")
			continue
		}

		archOpcodeStr := makeArchOpcodeStruct(mName, mOperands, valid32Bit, valid64Bit)
		tableFile.WriteString(archOpcodeStr)
		tableFile.WriteString("\n")
		//fmt.Println(mName, mOperands, mValid32, mValid64)

	}

	// set table file footer
	setArchOpcodeFileFooter(tableFile)

	tableFile.Close()

	//fmt.Println(row[0])
}

// make arch opcode table go file
func makeArchOpcodeTableFile() *os.File {

	tableFilePath := "./x86_64/archopcode.td.go"

	// check state of table file
	_, err := os.Stat(tableFilePath)
	if err == nil {
		// if error is nil, file is already exist
		err := os.Remove(tableFilePath) // remove file
		if err != nil {
			panic(err)
		}
	}

	tableFile, err := os.OpenFile(tableFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	return tableFile
}

// set arch opcode file header
func setArchOpcodeFileHeader(tableFile *os.File) {
	fileHeader := `// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// architecture opcode table

package x86_64

// arch opcode list
var archOpcodeList []archOpcode = []archOpcode{
`
	tableFile.WriteString(fileHeader)
}

// set arch opcode file footer
func setArchOpcodeFileFooter(tableFile *os.File) {
	fileHeader := `}
`
	tableFile.WriteString(fileHeader)
}
