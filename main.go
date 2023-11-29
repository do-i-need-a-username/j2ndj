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

func main() {
	input := flag.String("input", "", "Input JSON file")
	output := flag.String("output", "", "Output file")
	flag.Parse()

	if *input == "" {
		fmt.Println("Please provide an input json file using the -input flag.")
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
}
