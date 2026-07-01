package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type menuOption struct {
	text   string
	action func()
}

func menuPrint(options []menuOption, exit bool) {
	for {
		println()
		for i, o := range options {
			fmt.Printf(" %d. %s\n", i+1, o.text)
		}
		if exit {
			println(" 0. Exit")
		} else {
			println(" 0. Back")
		}

		Scanner.Scan()

		input := strings.TrimSpace(Scanner.Text())
		if input == "" {
			continue
		}
		choice, err := strconv.Atoi(input)
		if err != nil {
			println("Please enter a number")
			continue
		}
		if choice < 0 || choice > len(options) {
			println("Invalid choice")
			continue
		}
		if choice == 0 {
			if exit {
				os.Exit(0)
			}
			return
		}
		if err := Scanner.Err(); err != nil {
			fmt.Printf("Error while reading input: %s", err)
			continue
		}
		options[choice-1].action()
		return
	}
}
