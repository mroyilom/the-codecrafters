package main

import (
	"fmt"
	"strconv"
	"strings"
)

func BaseConverter(word string, base int) (int64, error) {
	return strconv.ParseInt(word, base, 64)
}

func dex(n int64, base int) string {
	result := strconv.FormatInt(n, base)
	return strings.ToUpper(result)
}

func main() {

	for {
		var choice string
		var word string
		var base string

		fmt.Println("Type 'help' for instructions or 'quit' to exit")
		fmt.Print("> ")

		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Error: invalid input")
			continue
		}

		choice = strings.ToLower(choice)

		if choice == "quit" {
			fmt.Println("Goodbye")
			return
		}

		if choice == "help" {
			fmt.Println("Commands:")
			fmt.Println("convert <number> <base>")

			fmt.Println("\nExamples:")
			fmt.Println("convert 1E hex ")
			fmt.Println("convert 10 bin ")
			fmt.Println("convert 255 dec")

			continue
		}

		if choice != "convert" {
			fmt.Println("Unknown command")
			continue
		}

		_, err = fmt.Scan(&word, &base)
		if err != nil {
			fmt.Println("Usage: convert <number> <base>")
			continue
		}

		base = strings.ToLower(base)

		if base == "hex" {
			result, err := BaseConverter(word, 16)
			if err != nil {
				fmt.Println("Error: invalid hex")
				continue
			}
			fmt.Println("✦ Decimal:", result)

		} else if base == "bin" {
			result, err := BaseConverter(word, 2)
			if err != nil {
				fmt.Println("Error: invalid binary")
				continue
			}
			fmt.Println("✦ Decimal:", result)

		} else if base == "dec" {
			n, err := strconv.ParseInt(word, 10, 64)
			if err != nil {
				fmt.Println("Error: invalid decimal")
				continue
			}

			fmt.Println("✦ Binary:", dex(n, 2))
			fmt.Println("✦ Hex:", dex(n, 16))

		} else {
			fmt.Println("Error: base must be hex, bin, or dec")
		}
	}
}
