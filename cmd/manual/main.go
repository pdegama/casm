// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

/*
	x86_64 / amd64 manual
	information and download
*/

package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

// res links
const (

	// intel
	intelManualName string = "x86"
	intelManualWeb  string = "https://www.intel.com/content/www/us/en/developer/articles/technical/intel-sdm.html"
	intelManualURL  string = "https://cdrdv2.intel.com/v1/dl/getContent/671200"
	intelManualURL2 string = "https://cdrdv2.intel.com/v1/dl/getContent/671110"

	// amd
	amdManualName string = "amd64"
	amdManualWeb  string = "https://www.amd.com/en/support/tech-docs?keyword=AMD64+Architecture+Programmer%27s+Manual"
	amdManualURL  string = "https://www.amd.com/system/files/TechDocs/40332.pdf"
	amdManualURL2 string = "https://www.amd.com/system/files/TechDocs/24594.pdf"
)

func main() {

	info := flag.String("info", "", "[ intel | amd | both ] information of manual")
	download := flag.String("download", "", "[ intel | amd | both ] download manual")
	flag.Parse()

	if *info == "" && *download == "" {
		flag.PrintDefaults()
		return
	}

	if *info != "" {
		// information
		switch *info {
		case "intel", "amd", "both":
			manualInfo(*info)
		default:
			flag.PrintDefaults()
			return
		}
	}

	if *download != "" {
		// download
		switch *download {
		case "intel", "amd", "both":
			manualDownload(*download)
		default:
			flag.PrintDefaults()
			return
		}
	}

}

// information
func manualInfo(manualName string) {
	if manualName == "intel" || manualName == "both" {
		fmt.Println(intelManualName)
		fmt.Printf("    Website:\t%v\n", intelManualWeb)
		fmt.Printf("    Link:\t%v\n", intelManualURL)
		fmt.Printf("    Link:\t%v\n", intelManualURL2)
	}

	if manualName == "amd" || manualName == "both" {
		fmt.Println(amdManualName)
		fmt.Printf("    Website:\t%v\n", amdManualWeb)
		fmt.Printf("    Link:\t%v\n", amdManualURL)
		fmt.Printf("    Link:\t%v\n", amdManualURL2)
	}
}

// download
func manualDownload(manualName string) {
	if manualName == "intel" || manualName == "both" {
		downloadFile(intelManualURL, intelManualName)
		downloadFile(intelManualURL2, intelManualName+"_instruction")
	}

	if manualName == "amd" || manualName == "both" {
		downloadFile(amdManualURL, amdManualName)
		downloadFile(amdManualURL2, amdManualName+"_instruction")
	}
}

func downloadFile(downloadFileURL string, downloadFileName string) {

	fmt.Println("Downloading", downloadFileName)

	// Create blank file
	userHome, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	downloadFileName = userHome + "/" + downloadFileName + ".pdf"

	file, err := os.Create(downloadFileName)
	if err != nil {
		panic(err)
	}

	// send request
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	// Put content on file
	resp, err := client.Get(downloadFileURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	fmt.Printf("Download complete, Downloaded a file location %v\n", downloadFileName)

}
