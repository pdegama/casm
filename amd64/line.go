// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// line

package amd64

// line type
type lineType string

const ( // value is tmp
	lineUnknow lineType = "Unknow" // unknow
	lineInst            = "Inst"   // instruction
	lineData            = "Data"   // data
	lineLabel           = "Label"  // label
	lineModulo          = "Modulo" // modulo
)

// line structure
type asmLine struct {
	tokens   []token  // tokens
	index    int      // line index
	filePath *string  // filename
	lineType lineType // line type
}
