package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a json file as a command line argument.")
		os.Exit(1)
	}

	jsonFile, err := os.Open(os.Args[1])
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

	outputFileName := strings.TrimSuffix(os.Args[1], filepath.Ext(os.Args[1])) + ".ndjson"
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
