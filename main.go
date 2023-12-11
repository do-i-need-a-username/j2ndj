// run converts JSON to NDJSON format.
// It reads the input JSON file or standard input if no input file is provided.
// It writes the output to the specified file or standard output if no output file is provided.
// The function returns an error if any error occurs during the conversion process.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	input  = flag.String("input", "", "Input JSON file")
	output = flag.String("output", "", "Output file. Uses input with extension .ndjson if not provided")
)

// Convert json to ndjson
func run() error {
	var reader io.Reader
	var writer io.Writer

	if *input == "" {
		reader = os.Stdin
	} else {
		file, err := os.Open(*input)
		if err != nil {
			return fmt.Errorf("Error opening input file: %w", err)
		}
		defer file.Close()
		reader = file
	}

	if *output == "" {
		writer = os.Stdout
	} else {
		file, err := os.Create(*output)
		if err != nil {
			return fmt.Errorf("Error creating output file: %w", err)
		}
		defer file.Close()
		writer = file
	}

	byteValue, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("Error reading from input: %w", err)
	}

	var result []map[string]interface{}
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return fmt.Errorf("Error decoding JSON from input: %w", err)
	}

	for _, item := range result {
		jsonStr, err := json.Marshal(item)
		if err != nil {
			return fmt.Errorf("Error converting to JSON: %w", err)
		}
		_, err = writer.Write([]byte(string(jsonStr) + "\n"))
		if err != nil {
			return fmt.Errorf("Error writing to output: %w", err)
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
