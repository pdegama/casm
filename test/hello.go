// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// simple hello test file
// this code is compare expected output and nasm gen output

package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"
)

// test structure
type testStruct struct {
	input          string
	expectedOutput []byte
}

// test task
var testTask []testStruct = []testStruct{
	{"bits 64\nadd [0x74], dword 0x700\n", []byte{0x81, 0x04, 0x25, 0x74, 0x00, 0x00, 0x00, 0x00, 0x07, 0x00, 0x00}},
	{"bits 32\nadd [0x74], dword 0x700\n", []byte{0x81, 0x05, 0x74, 0x00, 0x00, 0x00, 0x00, 0x07, 0x00, 0x00}},
	{"bits 16\nadd [0x74], dword 0x700\n", []byte{0x66, 0x81, 0x06, 0x74, 0x00, 0x00, 0x07, 0x00, 0x00}},
	{"bits 16\nadd [si], dword 0x700\n", []byte{0x66, 0x81, 0x04, 0x00, 0x07, 0x00, 0x00}},
	{"bits 16\nadd [esi], dword 0x700\n", []byte{0x66, 0x67, 0x81, 0x06, 0x00, 0x07, 0x00, 0x00}},
	{"bits 32\nadd [si], dword 0x700\n", []byte{0x67, 0x81, 0x04, 0x00, 0x07, 0x00, 0x00}},
	{"bits 32\nadd [esi], dword 0x700\n", []byte{0x81, 0x06, 0x00, 0x07, 0x00, 0x00}},
	{"bits 64\nadd [esi], dword 0x700\n", []byte{0x67, 0x81, 0x06, 0x00, 0x07, 0x00, 0x00}},
}

// main function
func main() {

	ctime := time.Now().Unix()                     // get current time
	inputFile := fmt.Sprintf("tmp/%d.test", ctime) // input file
	outputFile := fmt.Sprintf("tmp/%d.out", ctime) // output file

	// if tmp directory is not exist then make directory
	_, err := os.ReadDir("tmp")
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir("tmp", 0775)
		} else {
			panic(err)
		}
	}

	fmt.Println("---------------------------------")

	// loop to test task
	for _, t := range testTask {

		dbyte := []byte(t.input) // file data byte

		// write or re-write asm file
		os.WriteFile(inputFile, dbyte, 0644)

		// assemble file
		cmd := exec.Command("nasm", inputFile, "-o", outputFile)
		err = cmd.Run() //
		if err != nil {
			panic(err)
		}

		obyte, err := os.ReadFile(outputFile)
		if err != nil {
			panic(nil)
		}

		fmt.Printf("%s\nExpected:\t[% x]\nOutput:\t\t[% x]\nPass:\t\t%v\n", t.input, t.expectedOutput, obyte, bytes.Equal(t.expectedOutput, obyte))

		fmt.Println("---------------------------------")

		// one second delay
		time.Sleep(time.Second)
	}

	fmt.Println("Test Complete") // test complete

}
