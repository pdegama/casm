// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// parser

package amd64

import (
	"fmt"

	"hellocomputers/casm/utils"
)

// parse lines
func parseLines(lines []asmLine) ([]asmLine, []error) {

	linesX := []asmLine{} // parsed line
	errs := []error{}     // errors

	for _, line := range lines {
		line, err := parseLine(line)
		if err != nil {
			errs = append(errs, err)
		}
		linesX = append(linesX, line)
	}

	return linesX, errs
}

// parse one line
func parseLine(line asmLine) (asmLine, error) {

	///	fmt.Printf("%v:%v %v %v\n", *line.filePath, line.index, line.lineType, line.tokens)

	if len(line.tokens) > 0 {
		// check first token

		if line.tokens[0].tokenType == tokenModulo {
			// if token is modulo then line is modulo
			line.lineType = lineModulo
			return line, nil
		}

		if line.tokens[0].tokenType == tokenUnknow {
			switch line.tokens[0].token {
			case "db", "dw", "dd", "dq": // if token is d* then line is token
				line.lineType = lineData
				return line, nil
			}
		}
	}

	for tokIndex, tok := range line.tokens {
		switch tok.tokenType {
		case tokenColon:
			// token is colon and token index is 1 then is valid
			if tokIndex == 1 {
				line.lineType = lineLabel
				return line, nil
			} else {
				return line, fmt.Errorf("%s %v:%v %v", utils.Error, *line.filePath, line.index+1, "invalid syntax")
			}
		}
	}

	//fmt.Printf("%v:%v %v %v\n", *line.filePath, line.index, line.lineType, line.tokens)

	// if first token is unknow then assume line is Instruction
	if len(line.tokens) > 0 {
		if line.tokens[0].tokenType == tokenUnknow {
			line.lineType = lineInst
		} else {
			return line, fmt.Errorf("%s %v:%v %v", utils.Error, *line.filePath, line.index+1, "invalid syntax")
		}
	}

	return line, nil
}
