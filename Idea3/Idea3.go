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
		// fmt.Println("calculation equal 0")
		result := fmt.Sprintf("%s%s", (*calculation), newString)
		(*calculation) = result
		// fmt.Printf("%s\n", (*calculation))
	} else if len(*calculation) > 0 {
		// fmt.Println("Calculation is bigger than 1")
		result := fmt.Sprintf("%s %s", (*calculation), newString)
		(*calculation) = result
		// fmt.Printf("%s\n", (*calculation))
	} else {
		fmt.Println("What the fuck is happening")
	}
}

// Prints the actual calculation
func printCalculation(calculation string) {
	fmt.Printf("\nCalculating\n%s\n\n", calculation)
}

func main() {
	fmt.Println("Welcome to zak's calculator")

	var calculation string
	var command string
	// var exit string
	var number int
	step := 1

	for true {
		fmt.Println(step)
		if step%2 == 1 {
			fmt.Println("Enter a number")
			fmt.Scanln(&number)
			concatenateCalculation(&calculation, fmt.Sprint(number))
			printCalculation(calculation)
			step += 1
		} else {
			fmt.Println("Enter a command")
			fmt.Scanln(&command)
			//Validate if the command isn't + or - or * or /
			if strings.Compare(command, "-") != 0 && strings.Compare(command, "+") != 0 && strings.Compare(command, "*") != 0 &&
				strings.Compare(command, "/") != 0 {
				fmt.Println("Invalid command")
				continue
			}
			concatenateCalculation(&calculation, command)
			printCalculation(calculation)
			step += 1
		}

	}
}
