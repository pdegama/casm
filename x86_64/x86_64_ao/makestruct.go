// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// make arch opcode structure

package main

import (
	"bytes"
	"fmt"
	"text/template"
)

// arch opcode template structure
type archOpodeTemplateStrct struct {
	Name           string // instruction mnemonic name
	Operands       string // instruction mnemonic operands
	Valid32BitMode bool   // instruction is valid in 32-bit mode
	Valid64BitMode bool   // instruction is valid in 64-bit mode
}

// make arch opcode structure string
func makeArchOpcodeStruct(mName string, mOperands string, valid32bit bool, valid64bit bool) string {

	templateContent := archOpodeTemplateStrct{
		Name:           mName,
		Operands:       mOperands,
		Valid32BitMode: valid32bit,
		Valid64BitMode: valid64bit,
	}

	structStrTemplate := `	{name: "{{.Name}}", operands: {{.Operands}}, opcode: []int{}, valid32BitMode: {{.Valid32BitMode}}, valid64BitMode: {{.Valid64BitMode}}},`

	aTemplate := template.New("archOpcodeTemplate")
	aTemplate, err := aTemplate.Parse(structStrTemplate)
	if err != nil {
		panic(err)
	}

	templateBuf := bytes.NewBufferString("") // structure string buffer
	aTemplate.Execute(templateBuf, templateContent)

	fmt.Println(templateBuf.String())

	// return structure string
	return templateBuf.String()
}

// arch opcode template structure
type archOperandTemplateStrct struct {
	T string // operand type
	V string // operand value
}

// make arch operand structure string
func makeArchOperandStruct(t string, v string) string {

	templateContent := archOperandTemplateStrct{
		T: t,
		V: v,
	}

	structStrTemplate := `{{.T}}, {{.V}}`

	aTemplate := template.New("archOperandTemplate")
	aTemplate, err := aTemplate.Parse(structStrTemplate)
	if err != nil {
		panic(err)
	}

	templateBuf := bytes.NewBufferString("") // structure string buffer
	aTemplate.Execute(templateBuf, templateContent)

	// add curly braces
	templateStr := "{" + templateBuf.String() + "}"

	// return structure string
	return templateStr
}
