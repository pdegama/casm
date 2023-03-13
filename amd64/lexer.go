// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// lexer - tokenizer

package amd64

import (
	"fmt"
	"os"
	"strings"

	"hellocomputers/casm/utils"
)

// open file
func openFile(filePath string) ([]byte, error) {

	// read file
	fileBytes, err := os.ReadFile(filePath)

	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("%s file not found: %s", utils.Error, filePath)
		} else {
			return nil, fmt.Errorf("unknow")
		}
	}

	return fileBytes, nil

}

// tokenizer
func lexer(filePath string) ([]asmLine, error) {

	lines := []asmLine{} // asm lines

	// open file
	fileBytes, err := openFile(filePath)
	if err != nil {
		return nil, err
	}

	// asm file string
	asmStr := string(fileBytes)

	linesStr := strings.Split(asmStr, "\n") // asm lines
	for lineIndex, lineStr := range linesStr {
		toks, err := lineLexer(lineStr)
		if err != nil {
			fmt.Printf("%s %v:%v %v\n", utils.Error, filePath, lineIndex+1, err)
		}
		//fmt.Println(toks)

		line := asmLine{
			tokens:   toks,
			index:    lineIndex,
			filePath: &filePath,
			lineType: lineUnknow,
		}

		lines = append(lines, line)

	}

	return lines, nil

}

// line parser
func lineLexer(asmLine string) ([]token, error) {
	tok := ""         // token
	toks := []token{} // tokens
	isStr := false    // string is start

stringLoop:
	for _, t := range asmLine {

		// if str in start
		if isStr {
			if t == '"' {
				// if rune is " then str stop
			} else {
				// add any rune in token and continue
				tok += string(t)
				continue
			}
		}

		if validCharInWord(t) {
			tok += string(t)
			continue
		}

		switch t {
		case ' ': // space
			tokenPrint[rune](&tok, nil, tokenUnknow, &toks)
		case ':': // colon
			tokenPrint(&tok, &t, tokenColon, &toks)
		case ';': // semicolon
			tokenPrint[rune](&tok, nil, tokenUnknow, &toks)
			break stringLoop
		case '%':
			tokenPrint(&tok, &t, tokenModulo, &toks)
		case ',': // comma
			tokenPrint(&tok, &t, tokenComma, &toks)
		case '"': // double quote
			tokenPrint(&tok, &t, tokenDoubleQuote, &toks)
			isStr = !isStr
		default: // other
			return toks, fmt.Errorf("invalid char '%v'", string(t))
		}

	}

	tokenPrint[rune](&tok, nil, tokenUnknow, &toks)

	return toks, nil
}

// print tokne - tmp
func tokenPrint[T string | rune](tok *string, newTok *T, newTokType tokenType, toks *[]token) {
	if *tok != "" {
		// fmt.Println(*token)
		*toks = append(*toks, token{
			token:     *tok,
			tokenType: tokenUnknow,
		})
	}
	if newTok != nil {
		// fmt.Println(string(*token2))
		*toks = append(*toks, token{
			token:     string(*newTok),
			tokenType: newTokType,
		})
	}
	*tok = ""
}

// check rune is valid in word
func validCharInWord(t rune) bool {
	return (t >= 'A' && t <= 'Z') || (t >= 'a' && t <= 'z') || (t >= '0' && t <= '9') || t == '_'
}
