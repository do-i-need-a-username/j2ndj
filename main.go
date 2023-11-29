package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	input  = flag.String("input", "", "Input JSON file")
	output = flag.String("output", "", "Output file. Uses input with extension .ndjson if not provided")
)

func run() error {
	// input := flag.String("input", "", "Input JSON file")
	// output := flag.String("output", "", "Output file. Uses input with extension .ndjson if not provided")

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Println("Converts a json file to a New Line Delimited JSON (ndjson) file.")
		fmt.Println("Flags:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *input == "" {
		flag.Usage()
		os.Exit(1)
	}

	jsonFile, err := os.Open(*input)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	var result []map[string]interface{}
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}

	outputFileName := *output
	if outputFileName == "" {
		outputFileName = strings.TrimSuffix(*input, filepath.Ext(*input)) + ".ndjson"
	}
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	for _, item := range result {
		jsonStr, err := json.Marshal(item)
		if err != nil {
			fmt.Println("Error converting to JSON:", err)
			os.Exit(1)
		}
		_, err = outputFile.WriteString(string(jsonStr) + "\n")
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			os.Exit(1)
		}
	}

	return nil
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
