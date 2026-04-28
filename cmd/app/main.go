package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"go-calculator/internal/calculator"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Go Calculator 🚀")
	fmt.Println("Examples:")
	fmt.Println(" 10 + 5")
	fmt.Println(" sqrt 25")
	fmt.Println(" type 'exit' to quit")

	for {
		fmt.Print("> ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// exit condition
		if input == "exit" {
			fmt.Println("Goodbye 👋")
			break
		}

		// parse input → Expression
		exp, err := calculator.Parse(input)
		if err != nil {
			fmt.Println("❌ Invalid input format")
			continue
		}

		// execute using new engine
		result, err := calculator.Execute(exp.Operator, exp.Operands...)
		if err != nil {
			fmt.Println("⚠️ Error:", err)
			continue
		}

		fmt.Println("✅ Result:", result)
	}
}