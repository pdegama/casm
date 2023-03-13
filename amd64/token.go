// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// token

package amd64

// token type
type tokenType string

const ( // value is tmp
	tokenUnknow      tokenType = "Unknow"
	tokenComma                 = "Comma"        // ,
	tokenColon                 = "Colon"        // :
	tokenDoubleQuote           = "Double Quote" // "
)

// token structure
type token struct {
	token     string    // token
	tokenType tokenType // token type
}
