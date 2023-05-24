// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// parser

package x86_64

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// lines parser
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

// one line parser
func parseLine(line asmLine) (asmLine, error) {

	if len(line.tokens) > 0 {
		// check first token

		if line.tokens[0].tokenType == tokenModulo {
			// if token is modulo then line is modulo
			line.lineType = lineModulo
			return line, nil
		}

		if line.tokens[0].tokenType == tokenUnknow {
			switch line.tokens[0].token {
			case "db", "dw", "dd", "dq", "str": // if token is d* then line is token
				line.lineType = lineData
				return line, nil
			}
		}
	}

	// check other token
	for tokIndex, tok := range line.tokens {
		switch tok.tokenType {
		case tokenColon:
			// token is colon and token index is 1 then is valid
			if tokIndex == 1 {
				line.lineType = lineLabel
				return line, nil
			} else {
				return line, fmt.Errorf("%s %v:%v %v", errorStr, *line.filePath, line.index+1, "invalid syntax")
			}
		}
	}

	//fmt.Printf("%v:%v %v %v\n", *line.filePath, line.index, line.lineType, line.tokens)

	// if first token is unknow then assume line is Instruction
	if len(line.tokens) > 0 {
		if line.tokens[0].tokenType == tokenUnknow {
			line.lineType = lineInst
		} else {
			return line, fmt.Errorf("%s %v:%v %v", errorStr, *line.filePath, line.index+1, "invalid syntax")
		}
	}

	return line, nil
}

// instruction parser
func parseInst(tokens []token) (instruction, error) {

	mnemonicName := "" // instruction mnemonic name

	operandTokens := []token{} // operand token queue
	operands := []operand{}    // operands queue

	mNameToken := tokens[0]      // mnemonic token
	mOperandTokens := tokens[1:] // mnemonic operand

	if mNameToken.tokenType == tokenUnknow {
		// if mnemonic token type is unknow
		mnemonicName = strings.ToUpper(mNameToken.token) // set inst mnemonic name
	} else {
		// invalid token
		return instruction{}, fmt.Errorf("invalid token")
	}

	// loop of opernad tokens
	for _, tok := range mOperandTokens {
		switch tok.tokenType {
		case tokenUnknow, tokenLabel, tokenBracketLeft, tokenBracketRight, tokenDoubleQuote, tokenPlus, tokenMinus:
			// if token is unknow, or label
			operandTokens = append(operandTokens, tok)
		case tokenComma:
			// token is comma then current operand is over
			opr, err := parseOperand(operandTokens)
			if err != nil {
				// operand parse error
				return instruction{}, err
			}
			operands = append(operands, opr) // add to operands
			operandTokens = []token{}        // clear operand
		default:
			// invalid token
			return instruction{}, fmt.Errorf("invalid token")
		}
	}

	if len(operandTokens) > 0 {
		// if operand tokens is more then zero then parse the tokens
		opr, err := parseOperand(operandTokens)
		if err != nil {
			// operand parse error
			return instruction{}, err
		}
		operands = append(operands, opr) // add to operands
	}

	// return instruction
	return instruction{
		name:     mnemonicName,
		operands: operands,
	}, nil
}

// operand parser
func parseOperand(tokens []token) (operand, error) {

	if len(tokens) == 0 {
		/* tokens len is zero then error tokens is invalid */
		return operand{}, fmt.Errorf("invalid syntax")
	}

	// fmt.Println(tokens)

	nextImm := false           // next token chance to imm
	nextImmType := noneOperand // next operand type
	opers := []operand{}       // operand
	isMem := false             // this operand is chance to mem

	bracketStart := false // bracket is strat or not

	// loop of token
	for _, tok := range tokens {

		switch tok.tokenType {
		case tokenUnknow:

			// if token is register
			isReg, reg := isRegister(tok.token)
			if isReg {
				if nextImm {
					/*
						if register is not imm value and this
						is chance the imm then return error
					*/
					return operand{}, fmt.Errorf("invalid operand register is not imm")
				}
				opers = append(opers, operand{
					t: getRegisterOperandType(reg),
					v: uint(reg.globleIndex),
					m: nil,
					l: false,
				})
				continue
			}

			// check token is imm
			tokenVal, err := strconv.Atoi(tok.token) // try to convert imm
			if err == nil {
				// if err is nil then token is imm

				/*
					parse and assign to oper, operand
					over set true and continue
				*/
				oprType := imm
				if nextImm {
					oprType = nextImmType
					nextImm = false
				}
				opers = append(opers, operand{
					t: oprType,
					v: uint(tokenVal),
					m: nil,
					l: false,
				})
				continue
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

				/*
					parse hex and assign to oper,
					operand over se true and continue
				*/
				oprType := imm
				if nextImm {
					oprType = nextImmType
					nextImm = false
				}
				opers = append(opers, operand{
					t: oprType,
					v: uint(tokenVal),
					m: nil,
					l: false,
				})
				continue

			}

			// check this is type
			switch tok.token {
			case "byte", "word", "dword", "qword":
				/*
					set type true and set type
				*/
				nextImm = true
				switch tok.token {
				case "byte":
					nextImmType = imm8
				case "word":
					nextImmType = imm16
				case "dword":
					nextImmType = imm32
				case "qword":
					nextImmType = imm64
				}
			default:
				return operand{}, fmt.Errorf("invalid token '%v'", tok.token)
			}

		case tokenBracketLeft:
			if isMem {
				/*
					if operand is aready start before
					operand start then return error
				*/
				return operand{}, fmt.Errorf("invalid operand / syntax")
			}

			if !bracketStart {
				// if bracket is not start
				bracketStart = true
			} else {
				return operand{}, fmt.Errorf("invalid bracket syntax")
			}

		case tokenBracketRight:
			if bracketStart {
				/*
					Mem operand have a brackets,
					if brackets is over then
					confirm this operand is mem
				*/
				isMem = true

				// if bracket is start
				bracketStart = false
			} else {
				return operand{}, fmt.Errorf("invalid bracket syntax")
			}

		case tokenPlus:
			opers = append(opers, operand{
				t: operPrePlus,
				v: 0x00,
				m: nil,
				l: false,
			})

		case tokenMinus:
			opers = append(opers, operand{
				t: operPreMinus,
				v: 0x00,
				m: nil,
				l: false,
			})

		case tokenLabel:
			// token is label
			oprType := imm
			if nextImm {
				oprType = nextImmType
				nextImm = false
			}
			opers = append(opers, operand{
				t: oprType,
				v: 0x00,
				m: nil,
				l: true,
				n: tok.token,
			})

		}
	}

	if bracketStart {
		// if bracket is not over
		return operand{}, fmt.Errorf("invalid bracket syntax")
	}

	if isMem {
		// operand is mem

		if len(opers) == 0 {
			/*
				if memory operand length is
				zero then return error
			*/
			return operand{}, fmt.Errorf("invalid memory operand/syntax")
		}

		oper, err := parseMem(&opers)
		if err != nil {
			// if error
			return operand{}, err
		}
		return oper, nil
	}

	if len(opers) == 0 {
		// if opers len is 0 then return error
		return operand{}, fmt.Errorf("invalid operand/syntax")
	}

	if len(opers) > 1 {
		// if oper len is more then 1 return error
		return operand{}, fmt.Errorf("invalid operand / todo")
	}

	return opers[0], nil
}

// mem parser
func parseMem(opers *[]operand) (operand, error) {

	isOperation := true // if this operand is operation operand (init value is true)

	memOpers := []operand{} // memory operands

	// fmt.Println("memory operands:", opers)

	// operands
	for _, oper := range *opers {
		switch oper.t {
		case reg8, reg16, reg32, reg64:
			// if operand is reg
			if isOperation {
				memOpers = append(memOpers, oper) // append to memory operands
				isOperation = false
			} else {
				return operand{}, fmt.Errorf("invalid memory operand/syntax")
			}
		case imm, imm8, imm16, imm32, imm64:
			// if operand id imm
			if isOperation {
				memOpers = append(memOpers, oper) // append to memory operands
				isOperation = false
			} else {
				return operand{}, fmt.Errorf("invalid memory operand/syntax")
			}
		case operPrePlus, operPreMinus:
			// is pre pluse
			memOpers = append(memOpers, oper) // append to memory operands
			isOperation = true                // this operand is operation operand then set true
		default:
			// invalid or todo
			return operand{}, fmt.Errorf("invalid operand / todo")
		}
	}

	// return mem operand
	return operand{
		t: mem,
		v: 0,
		m: memOpers,
		l: false,
	}, nil
}

// imm type parser
func parseImmType(i uint) operandType {

	/*
		get operand imm type (imm8, imm16, imm32, imm64)
	*/

	if i <= 0xff {
		// operand is equal or smaller than 0xff
		return imm8
	} else if i <= 0xffff {
		// operand is equal or smaller than 0xffff
		return imm16
	} else if i <= 0xffffffff {
		// operand is equal or smaller than 0xffffffff
		return imm32
	} else {
		// operand is equal or smaller than 0xffffffffffffffff
		return imm64
	}
}

// modulo parse
func parseModulo(line asmLine, b *binaryGen) error {

	if len(line.tokens) < 2 {
		return fmt.Errorf("invalid modulo instraction")
	}

	// loop of token
	for tokIndex, tok := range line.tokens {

		switch tokIndex {
		case 0:
			// if token index is zero (first token) and
			if tok.tokenType != tokenModulo {
				// this token is not modulo token then return error
				return fmt.Errorf("invalid modulo instraction")
			}

		case 1:
			// second token
			switch tok.token {
			case "bits", "bitmode":
				/*
					if modulo token i bits and bitmode
					then require three token
					% bitmod <bit>
				*/
				if len(line.tokens) == 3 {

					// third token is bits token
					bitsToken := line.tokens[2]

					switch bitsToken.token {
					case "16":
						/*
							if bitsToken value is 16 then
							set 16 in binaryGen
						*/
						b.bitMode = 16
					case "32":
						/*
							if bitsToken value is 32 then
							set 32 in binaryGen
						*/
						b.bitMode = 32
					case "64":
						/*
							if bitsToken value is 64 then
							set 64 in binaryGen
						*/
						b.bitMode = 64
					default:
						/*
							if other bitToken value then
							return error
						*/
						return fmt.Errorf("invalid bits token `%v`, only support 16, 32 and 64", bitsToken.token)
					}

				} else {
					// not three token avaible
					return fmt.Errorf("invalid modulo tokens for `%v`", tok.token)
				}

			default:
				return fmt.Errorf("invalid modulo operation `%v`", tok.token)
			}

		}

	}

	return nil
}
