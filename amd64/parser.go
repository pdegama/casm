// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// parser

package amd64

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"

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

// parse instruction
func parseInst(tokens []token) ([]token, error) {

	operandTokens := []token{} // operand token queue

	for tokIndex, tok := range tokens {

		if tokIndex == 0 {
			// if first token mnemonic
			// first token is
			if tok.tokenType == tokenUnknow {
				fmt.Println(tok.token, "Mnemonic")
				continue
			} else {
				fmt.Println("error: invalid token")
			}
		}

		switch tok.tokenType {
		case tokenUnknow, tokenLabel:
			// if token is unknow, or label
			operandTokens = append(operandTokens, tok)
		case tokenComma:
			// token is comma then current operand is over
			opr, err := parseOperand(operandTokens)
			if err != nil {
				return tokens, err
			}
			fmt.Println(opr)
			operandTokens = []token{} // clear operand
		default:
			// invalid token
			fmt.Println("error: invalid token")
		}
	}

	if len(operandTokens) > 0 {
		opr, err := parseOperand(operandTokens)
		if err != nil {
			return tokens, err
		}
		fmt.Println(opr)
	}

	return tokens, nil
}

// parse operand
func parseOperand(tokens []token) (operand, error) {

	if len(tokens) == 1 {

		tok := tokens[0] // first token

		switch tok.tokenType {
		case tokenUnknow:
			// if token is register
			isReg, reg := isRegister(tok.token)
			if isReg {
				return operand{
					operandType: getRegisterOperandType(reg),
					operand:     reg.globleIndex,
				}, nil
			}

			// check token is imm
			tokenVal, err := strconv.Atoi(tok.token) // try to convert imm
			if err == nil {
				// if err is nil then token is imm
				return operand{
					operandType: imm,
					operand:     tokenVal,
				}, nil
			}

			// check token is hex
			if strings.HasPrefix(tok.token, "0x") || strings.HasPrefix(tok.token, "0X") {

				hexString := tok.token[2:] // hex string
				if len(hexString)%2 != 0 {
					// if hex string is Odd len then add 0 prefix
					hexString = "0" + hexString
				}

				// token is start with 0x or 0X then hex
				tokenHex, err := hex.DecodeString(hexString)
				if err != nil {
					// if err not nil then hex is not valid
					return operand{}, fmt.Errorf("hex not valid")
				}

				tokenVal := big.NewInt(0).SetBytes(tokenHex).Uint64() // hex to int

				return operand{
					operandType: imm,
					operand:     int(tokenVal),
				}, nil

			}

		case tokenLabel:
			// token is label
			return operand{
				operandType: imm,
				operand:     0x00,
			}, nil
		}

	} else {
		return operand{}, fmt.Errorf("to do")
	}

	return operand{}, fmt.Errorf("operand is not valid")

}
