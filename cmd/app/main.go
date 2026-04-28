package main

import (
	"bufio"
	"fmt"
	"os"

	"go-calculator/internal/calculator"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Go Calculator 🚀")
	fmt.Println("Enter expression (example: 10 + 5)")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')

		exp, err := calculator.Parse(input)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		result, err := calculator.Execute(exp)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("Result:", result)
	}
}