package main

import (
	"encoding/json"
	"fmt"
	"github.com/akamensky/argparse"
	"io"
	"log"
	"net/http"
	"os"
	"portego/src/pkgbuild"
)

// Logger types
var (
	DebugLogger		*log.Logger
	WarningLogger 	*log.Logger
	InfoLogger    	*log.Logger
	ErrorLogger   	*log.Logger
)

var AUR_BASE_URL = "https://aur.archlinux.org/rpc/?v=5&type=search&by=name&arg="

type AURPackage struct {
	ID				uint
	Name			string
	PackageBaseID	uint
	PackageBase		string
	Version			string
	Description		string
	URL				string
	NumVotes		uint
	Popularity		float32
	OutOfDate		bool
	Maintainer		string
	FirstSubmitted	uint
	LastModified	uint
	URLPath			string
}

type AURResult struct {
	Version			int
	Type			string
	ResultCount		int
	Results			[]AURPackage
}

func main() {
	// Build argument parser
	parser := argparse.NewParser("portego", "Portego uses PKGBUILD files form AUR to create Pkgfiles for Crux ")
	var packageName *string = parser.String("q", "query", 
								&argparse.Options{
									Required: true, 
									Help: "Name of queried package from AUR",
								})
	var outputPath *string = parser.String("o", "output", 
								&argparse.Options{
									Required: true,
									Help: "Output path of generated Pkgfile",
									Default: "./",
								})
	var consoleOutput *bool = parser.Flag("c", "console",
								&argparse.Options{
									Required: false,
									Help: "Print Pkgfile to console",
									Default: false,
								})
	var debugLogging *bool = parser.Flag("d", "debug",
								&argparse.Options{
									Required: false,
									Help: "Debug output",
									Default: false,
								})
	
	// Try to parse arguments
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	
	// Build logger
	file, err := os.OpenFile("portego.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

	// Logging instance
	DebugLogger = log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
    InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
    
	// Log application config
	if *debugLogging {
		DebugLogger.Printf("Portego arguments => Query: '%s', Output: '%s', Console: '%t', Debug: '%t'", 
			*packageName, 
			*outputPath, 
			*consoleOutput, 
			*debugLogging,
		)
	}
	
	// Build query URL
	var queryUrl = fmt.Sprintf("%s%s", AUR_BASE_URL, *packageName)
	InfoLogger.Printf("Query package '%s' from URL '%s'", *packageName, queryUrl)
	
	// Query AUR
	resp, err := http.Get(queryUrl)
	if err != nil {
		log.Fatal(err)
	}
	// Read http response body
	httpBody, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	
	if *debugLogging {
		DebugLogger.Printf("Response from AUR: %s", httpBody)
	}
	
	// Unmarshal JSON response to struct
	var results AURResult	
	json.Unmarshal(httpBody, &results)
	
	// List found packages
	InfoLogger.Printf("Found %s packages", results.ResultCount)
	for i, s := range results.Results {
		InfoLogger.Printf("Package %d: %s / %s / %s", 
			i, 
			s.Name, 
			s.Description,
			s.URL,
		)
	}
	
	source := portego.Pkgbuild{}
	source.AddPkgname(results.Results[0].Name)
	
	fmt.Println(source.Pkgname())
}
