package main

import (
	"fmt"
	"net/http"
	"testing-area/pkg"
	// "strings"
)

func main() {
	http.HandleFunc("/", pkg.PathHandler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
	// // Store argument value in input variable
	// Input := strings.Split(pkg.GetArgString(), "\\n")
	// // Store string value from specified txt file in lines variable
	// lines := strings.Split(pkg.GetBannerFile("banners/standard.txt"), "\n")
	// // Printing ascii art word
	// pkg.PrintAsciiArtWord(Input, lines)
}
