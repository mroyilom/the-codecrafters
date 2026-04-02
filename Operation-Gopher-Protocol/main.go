package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Rule 1: Trim spaces
func trimSpaces(line string) string {
	return strings.TrimSpace(line)
}

// Rule 2: Replace TODO
func replaceTODO(line string) string {
	return strings.ReplaceAll(line, "TODO:", "✦ ACTION:")
}

// Rule 3: Replace CLASSIFIED
func replaceClassified(line string) string {
	return strings.ReplaceAll(line, "CLASSIFIED:", "[REDACTED]:")
}

// Rule 4: Remove blank or dash lines
func isRemovable(line string) bool {
	if line == "" {
		return true
	}
	if strings.Trim(line, "-") == "" {
		return true
	}
	return false
}

// Rule 5: Add line numbers
func addLineNumber(line string, index int) string {
	return fmt.Sprintf("%03d. %s", index, line)
}

func main() {

	// Check arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Prevent same file input/output
	if inputFile == outputFile {
		fmt.Println("✗ Input and output cannot be the same file.")
		return
	}

	// Open input file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("✗ File not found:", inputFile)
		return
	}
	defer file.Close()

	// Check if output is directory
	info, err := os.Stat(outputFile)
	if err == nil && info.IsDir() {
		fmt.Println("✗ Cannot write to output: path is a directory.")
		return
	}

	// Create output file (overwrite)
	out, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("✗ Failed to write output file")
		return
	}
	defer out.Close()

	writer := bufio.NewWriter(out)

	// Write header
	writer.WriteString("SENTINEL FIELD REPORT — PROCESSED\n\n")

	scanner := bufio.NewScanner(file)

	linesRead := 0
	linesWritten := 0
	linesRemoved := 0

	lineIndex := 1

	// Read line by line
	for scanner.Scan() {
		linesRead++

		line := scanner.Text()

		// Apply rules in order

		line = trimSpaces(line)
		line = replaceTODO(line)
		line = replaceClassified(line)

		if isRemovable(line) {
			linesRemoved++
			continue
		}

		line = addLineNumber(line, lineIndex)
		lineIndex++

		writer.WriteString(line + "\n")
		linesWritten++
	}

	// Handle empty file
	if linesRead == 0 {
		fmt.Println("⚠ Input file is empty. Nothing to process.")
	}

	writer.Flush()

	// Terminal summary
	fmt.Println("✦ Lines read    :", linesRead)
	fmt.Println("✦ Lines written :", linesWritten)
	fmt.Println("✦ Lines removed :", linesRemoved)
	fmt.Println("✦ Rules applied : trim, TODO replace, CLASSIFIED replace, remove blanks, numbering")
}
