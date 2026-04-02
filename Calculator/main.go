package main

import (
	"fmt"
)

func main() {

	for {
		var choice string
		var a int
		var b int

		fmt.Println("Input Choices:")
		fmt.Println("+ a b")
		fmt.Println("- a b")
		fmt.Println("* a b")
		fmt.Println("/ a b")
		fmt.Println("Type 'help' or 'quit'")
		fmt.Print("> ")

		_, err := fmt.Scan(&choice)

		if err != nil {
			fmt.Println("Error: invalid input")
			continue
		}

		if choice == "quit" {
			fmt.Println("Enjoy your day")
			return
		}

		if choice == "help" {

			fmt.Println("Help instructions:")
			fmt.Println("add a b to add use this= + 5 6 ")
			fmt.Println("sub a b to subtract use = - 9 2")
			fmt.Println("mul a b to multiply use = * 7 2")
			fmt.Println("div a b to divide use = / 8 2")
			fmt.Println(" you will see your result deisplayed below the calculation")
			continue
		}

		_, err = fmt.Scan(&a, &b)
		if err != nil {
			fmt.Println("Error: enter valid numbers")
			continue
		}

		if choice == "+" {
			fmt.Println("✦ Result:", a+b)

		} else if choice == "-" {
			fmt.Println("✦ Result:", a-b)

		} else if choice == "*" {
			fmt.Println("✦ Result:", a*b)

		} else if choice == "/" {
			if b == 0 {
				fmt.Println("Error: cannot divide by zero")
				continue
			}
			fmt.Println("✦ Result:", a/b)

		} else {
			fmt.Println("Unknown command. Type 'help'")
		}
	}
}
