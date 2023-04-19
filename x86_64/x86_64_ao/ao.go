// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

/*
	x86_64.csv (x86_64 architecture opcode) to go file
*/

package main

import (
	"fmt"
)

// main function
func main() {

	// read csv file
	row := readCSV()

	for i, v := range row {
		fmt.Printf("%v ", i)
		parseData(v)
	}

	//fmt.Println(row[0])
}
