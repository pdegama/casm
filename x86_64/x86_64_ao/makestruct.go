// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

/*
	make arch opcode structure
*/

package main

import (
	"bytes"
	"fmt"
	"text/template"
)

// arch opcode template structure
type archOpodeTemplateStrct struct {
	Name           string // instruction mnemonic name
	Valid32BitMode bool   // instruction is valid in 32-bit mode
	Valid64BitMode bool   // instruction is valid in 64-bit mode
}

// make structure
func makeArchOpcodeStruct(mName string, mOperands []string, valid32bit bool, valid64bit bool) string {

	templateContent := archOpodeTemplateStrct{
		Name:           mName,
		Valid32BitMode: valid32bit,
		Valid64BitMode: valid64bit,
	}

	structStrTemplate := `	{name: "{{.Name}}", operands: []archOperand{}, opcode: []int{}, valid32BitMode: {{.Valid32BitMode}}, valid64BitMode: {{.Valid64BitMode}}},`

	acTemplate := template.New("acTemplate")
	acTemplate, err := acTemplate.Parse(structStrTemplate)
	if err != nil {
		panic(err)
	}

	templateBuf := bytes.NewBufferString("") // structure string buffer
	acTemplate.Execute(templateBuf, templateContent)

	fmt.Println(templateBuf.String())

	// return structure string
	return templateBuf.String()
}
