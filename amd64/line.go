// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// line

package amd64

// line type
type lineType string

// line type
const ( // value is tmp

	//
	lineUnknow lineType = "Unknow" // unknow
	lineInst   lineType = "Inst"   // instruction
	lineData   lineType = "Data"   // data
	lineLabel  lineType = "Label"  // label
	lineModulo lineType = "Modulo" // modulo

	//
)

// line structure
type asmLine struct {
	tokens   []token  // tokens
	index    int      // line index
	filePath *string  // filename
	lineType lineType // line type
}
