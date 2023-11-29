package main

import (
	"flag"
	"os"
	"testing"
)

func TestOutputFlag(t *testing.T) {
	// Save the original command-line arguments and restore them at the end of the test.
	origArgs := os.Args
	defer func() { os.Args = origArgs }()

	// Set the command-line arguments for this test.
	os.Args = []string{"cmd", "-output", "test_output"}

	// Parse the flags.
	flag.Parse()

	// Check that the output flag was correctly parsed.
	if *output != "test_output" {
		t.Errorf("Expected output flag to be 'test_output', but got '%s'", *output)
	}
}
