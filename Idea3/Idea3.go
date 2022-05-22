package main

import (
	"fmt"
	"strings"
)

/**
This function formats the display of the calculation beeing made
	calculation - the calculation beeing made
	newString - the new integer, or command to enter the calculation
*/
func concatenateCalculation(calculation *string, newString string) {
	if len(*calculation) == 0 {
		result := fmt.Sprintf("%s%s", (*calculation), newString)
		(*calculation) = result
	} else if len(*calculation) > 0 {
		result := fmt.Sprintf("%s %s", (*calculation), newString)
		(*calculation) = result
	} else {
		fmt.Println("What the fuck is happening")
	}
}

// Prints the actual calculation
func printCalculation(calculation string) {
	fmt.Printf("\nCalculating: %s\n", calculation)
}

func calculateResult(result int, number int, command string) int {
	switch command {
	case "-":
		result -= number
	case "+":
		result += number
	case "*":
		result *= number
	case "/":
		result /= number
	default:
		fmt.Println("Command not supported")
	}
	return result
}

func main() {
	fmt.Println("Welcome to zak's calculator")

	var calculation string
	var command string
	// var exit string
	result := 0
	var number int

	/** This variable defines what's the actual step that's beeing performe
	- Odd steps scan numbers
	- Even steps scan the command to perform
	*/
	step := 1

	for true {
		// fmt.Println(step)
		if step%2 == 1 {
			fmt.Println("Enter a number")
			fmt.Scanln(&number)

			// Calculate the result
			if len(calculation) == 0 {
				result += number
			} else {
				result = calculateResult(result, number, command)
			}

			// Put the number in the display and print the display
			concatenateCalculation(&calculation, fmt.Sprint(number))
			printCalculation(calculation)
			// Print partial result
			fmt.Printf("Partial calculation: %d\n\n", result)

			// Go to the next step
			step += 1
		} else {
			fmt.Println("Enter a command")
			fmt.Scanln(&command)
			//Validate if the command isn't + or - or * or / or
			if strings.Compare(command, "-") != 0 && strings.Compare(command, "+") != 0 && strings.Compare(command, "*") != 0 &&
				strings.Compare(command, "/") != 0 {
				fmt.Println("Invalid command")
				continue
			}

			// Put the number in the display and print the display
			concatenateCalculation(&calculation, command)
			printCalculation(calculation)

			// Go to the next step
			step += 1
		}
	}
}
