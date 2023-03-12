// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// token

package amd64

// token type
const (
	tokenUnknow      = iota
	tokenComma       // ,
	tokenColon       // :
	tokenDoubleQuote // "
)

// token structure
type token struct {
	token     string // token
	tokenType int    // token type
}
