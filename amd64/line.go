// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// line

package amd64

// line type
const (
	lineUnknow = iota
	lineInst   // instruction line
	lineLabel  // label line
	lineModulo // modulo line
)

// line structure
type asmLine struct {
	tokens   []token // tokens
	index    int     // line index
	filePath *string // filename
	lineType int     // line type
}
