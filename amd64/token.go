// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// token

package amd64

// token type
type tokenType string

// token type
const ( // value is tmp

	//
	tokenUnknow      tokenType = "Unknow"       // unknow
	tokenComma       tokenType = "Comma"        // ,
	tokenColon       tokenType = "Colon"        // :
	tokenDoubleQuote tokenType = "Double Quote" // "
	tokenModulo      tokenType = "Modulo"       // %
	tokenLabel       tokenType = "Label"        // $*

	//
)

// token structure
type token struct {
	token     string    // token
	tokenType tokenType // token type
}
