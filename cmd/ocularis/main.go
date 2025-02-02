package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"os"

	"ocularis/internal/core"
	"ocularis/internal/inputs"
	"ocularis/internal/outputs"
)

func usage() {
	fmt.Println(`Usage:
  
  To optimize JSON data:
      ocularis optimize -input <input.json> -output <optimized.json>
  
  To generate the report:
      ocularis generate -data <data.json> -output <report.html> -template <template.gohtml> [-key <base64_key>]

  If the data file does not exist, sample data will be generated.`)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	cmd := os.Args[1]
	switch cmd {
	case "optimize":
		optimizeCmd := flag.NewFlagSet("optimize", flag.ExitOnError)
		inputFile := optimizeCmd.String("input", "", "Input JSON file to optimize")
		outputFile := optimizeCmd.String("output", "", "Output file for optimized JSON")
		optimizeCmd.Parse(os.Args[2:])
		if *inputFile == "" || *outputFile == "" {
			fmt.Println("Both -input and -output flags are required for optimize.")
			usage()
		}
		if err := core.OptimizeJSONData(*inputFile, *outputFile); err != nil {
			fmt.Printf("Error optimizing JSON: %v\n", err)
			os.Exit(1)
		}

	case "generate":
		genCmd := flag.NewFlagSet("generate", flag.ExitOnError)
		dataFile := genCmd.String("input", "", "Path to input JSON data file")
		outputFile := genCmd.String("output", "", "Path to output HTML report file")
		templateFile := genCmd.String("template", "", "Path to HTML template file")
		keyStr := genCmd.String("key", "", "Base64-encoded encryption key (optional)")
		genCmd.Parse(os.Args[2:])

		if *dataFile == "" || *outputFile == "" || *templateFile == "" {
			fmt.Println("Flags -input, -output, and -template are required for generate.")
			usage()
		}

		var data []core.Entry
		if _, err := os.Stat(*dataFile); os.IsNotExist(err) {
			fmt.Printf("Input data file '%s' not found. Generating sample data.\n", *dataFile)
			data = inputs.GenerateSampleData()
		} else {
			var err error
			data, err = inputs.LoadSubfinderData(*dataFile)
			if err != nil {
				fmt.Printf("Failed to load data: %v\n", err)
				os.Exit(1)
			}
		}

		var encryptionKey []byte
		if *keyStr != "" {
			var err error
			encryptionKey, err = base64.StdEncoding.DecodeString(*keyStr)
			if err != nil {
				fmt.Printf("Invalid encryption key: %v\n", err)
				os.Exit(1)
			}
		} else {
			encryptionKey = make([]byte, 16)
			if _, err := rand.Read(encryptionKey); err != nil {
				fmt.Printf("Failed to generate encryption key: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("No encryption key provided. A random key has been generated.")
		}

		if err := outputs.GenerateHTMLReport(data, *outputFile, *templateFile, encryptionKey); err != nil {
			fmt.Printf("Error generating report: %v\n", err)
			os.Exit(1)
		}

	default:
		usage()
	}
}
