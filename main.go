package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/nlink-jp/nlk/jsonfix"
)

// The following variables are set during build time by the linker.
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// processInput reads all data from standard input and returns it as a single string.
func processInput() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var input strings.Builder
	for scanner.Scan() {
		input.WriteString(scanner.Text())
		input.WriteString("\n")
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("Error reading from stdin: %w", err)
	}
	return input.String(), nil
}

// extractAndValidateJSON extracts a JSON string from the input, repairs it if
// needed, and returns prettified output. Powered by nlk/jsonfix.
func extractAndValidateJSON(input string) (string, error) {
	result, err := jsonfix.Extract(input)
	if err != nil {
		if errors.Is(err, jsonfix.ErrNoJSON) {
			return "", fmt.Errorf("No valid JSON found in the input.")
		}
		return "", fmt.Errorf("Could not parse or fix the extracted JSON. Original output: %s", strings.TrimSpace(input))
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(result), "", "  "); err != nil {
		return result, nil
	}
	return prettyJSON.String(), nil
}

// handleOutput prints the result or handles the error based on the bypass flag.
func handleOutput(result string, err error, bypass bool, originalInput string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		if bypass {
			fmt.Print(originalInput)
		}
		os.Exit(1)
	} else {
		fmt.Println(result)
	}
}

// main is the entry point of the application.
func main() {
	// Define a boolean flag for the bypass mode.
	bypassMode := flag.Bool("bypass", false, "Bypass mode: if JSON parsing fails, output the original input instead of skipping.")
	
	// Add a version flag.
	showVersion := flag.Bool("version", false, "Print version information")
	
	flag.Parse()

	// Handle version flag.
	if *showVersion {
		fmt.Printf("json-filter-cli version: %s, commit: %s, built on: %s\n", version, commit, date)
		return
	}

	// Process the input from stdin.
	rawOutput, err := processInput()
	if err != nil {
		handleOutput("", err, *bypassMode, rawOutput)
	}

	// Extract and validate the JSON from the input.
	extractedJSON, err := extractAndValidateJSON(rawOutput)
	
	// Handle the final output based on the result and bypass mode.
	handleOutput(extractedJSON, err, *bypassMode, rawOutput)
}
