// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

/*
	read csv file
*/

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// read csv file
func readCSV() [][]string {

	// remove comment in /x86_64/x86_64_ao/x86_64.csv and create new file /x86_64/x86_64_ao/x86_64_clear.csv
	removeComment()

	// open file
	csvfile, err := os.Open("./x86_64/x86_64_ao/x86_64_clear.csv")
	if err != nil {
		panic(err)
	}
	defer csvfile.Close()

	csvReader := csv.NewReader(csvfile) // new csv reader
	csvData := [][]string{}             // csv data

	for {
		csvDataRow, err := csvReader.Read() // read data row
		if err != nil {
			if err.Error() == "EOF" {
				// if end of file then break loop
				break
			}
			fmt.Println(err)
		}
		csvData = append(csvData, csvDataRow) // append row to csvData
	}

	// return
	return csvData
}

func removeComment() {

	// ./x86_64/x86_64_ao/x86_64_clear.csv is already exist then delete file
	fi, _ := os.Stat("./x86_64/x86_64_ao/x86_64_clear.csv")
	if fi != nil {
		err := os.Remove("./x86_64/x86_64_ao/x86_64_clear.csv")
		if err != nil {
			panic(err)
		}
	}

	// create ./x86_64/x86_64_ao/x86_64_clear.csv
	f, err := os.OpenFile("./x86_64/x86_64_ao/x86_64_clear.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// read file "./x86_64/x86_64_ao/x86_64.csv"
	readFile, err := os.Open("./x86_64/x86_64_ao/x86_64.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile) // make new scanner
	fileScanner.Split(bufio.ScanLines)        // split new line

	for fileScanner.Scan() {
		// if line is empty or start then # then ignore
		if fileScanner.Text() != "" && !strings.HasPrefix(fileScanner.Text(), "#") && strings.ReplaceAll(fileScanner.Text(), " ", "") != "" {
			f.WriteString(fileScanner.Text()) // append line
			f.WriteString("\n")               // append new line
		}
	}

}
