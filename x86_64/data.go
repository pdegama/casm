// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// data bytes

package x86_64

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// data parse
func dataBytes(line asmLine) ([]uint8, error) {

	dataType := line.tokens[0]

	doubleQuoteStart := false

	dataBytes := []uint8{}

	for _, tok := range line.tokens[1:] {

		switch tok.tokenType {
		case tokenUnknow:

			if doubleQuoteStart {

				// str is on then pars this string to byte
				if dataType.token == "str" {
					// data type is string then add byte bo array
					b := []uint8(tok.token)
					dataBytes = append(dataBytes, b...)
				} else {
					// if not str then return err
					return nil, fmt.Errorf("this is string is not support in %v this data type", dataType.token)
				}

			} else {

				// check token is number
				tokenVal, err := strconv.Atoi(tok.token) // try to convert imm
				if err == nil {
					// if err is nil then token is imm

					/*
						parse and assign to oper, operand
						over set true and continue
					*/

					dataBytes = append(dataBytes, dataArray(uint64(tokenVal), dataType.token)...) // add byte

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
						return nil ,fmt.Errorf("hex not valid")
					}

					tokenVal := big.NewInt(0).SetBytes(tokenHex).Uint64() // hex to int

					dataBytes = append(dataBytes, dataArray(tokenVal, dataType.token)...) // add byte

					continue

				}

			}

		case tokenDoubleQuote:

			/*
				if double quote is start then
				end or end then start
			*/
			doubleQuoteStart = !doubleQuoteStart

		case tokenComma:
			// pass

		}

	}

	for _, b := range dataBytes {
		fmt.Printf("%x ", b)
	}

	return dataBytes, nil

}

// data uint8 array
func dataArray(v uint64, dataType string) []uint8 {

	switch dataType {
	case "db":
		return uint8le(uint8(v))
	case "dw":
		return uint16le(uint16(v))
	case "dd":
		return uint32le(uint32(v))
	case "dq":
		return uint64le(uint64(v))
	}

	return nil

}
