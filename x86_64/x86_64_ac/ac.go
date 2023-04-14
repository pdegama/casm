// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

/*
	x86_64.csv (x86_64 architecture code) to go file
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
		fmt.Println(i, v)
	}

	//fmt.Println(row[0])
}
