package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	//These lines create regular expressions used for building the acronym and testing for numeric input
	expressionForFirstLetter, _ := regexp.Compile(`^[A-Z]`)
	expressionForNumberTest, _ := regexp.Compile(`[0-9]*`)

	//This creates the scanner object to use for getting input from stdin
	scanner := bufio.NewScanner(os.Stdin)

	//Prompt user for an input string
	promptUser()

	//Infinite loop to continue to read from standard input
	for scanner.Scan() {

		//Used to test for input to quit the program
		inputString := scanner.Text()

		//Creates slices of the input string separated by spaces and captalizes in case the user didn't
		splitString := strings.SplitN(strings.ToUpper(scanner.Text()), " ", -1)

		//Checks the original input string for a condition to exit the program
		if inputString == "/quit" {
			fmt.Printf("Quitting\n")
			os.Exit(3)
			//Checks for numeric input to output an error and prompt user again
		} else if splitString[0] == expressionForNumberTest.FindString(inputString) {
			fmt.Printf("Invalid Input. Please Try Again.\n")
			promptUser()
			//Loops through the splitString slices and uses the regular expression to take and print only the first letter in each slice
		} else {
			fmt.Printf("Acronym: ")
			for i := range splitString {
				fmt.Printf("%s", expressionForFirstLetter.FindString(splitString[i]))
			}
			fmt.Printf("\n")
			promptUser()
		}
	}
}

func promptUser() {
	fmt.Printf("Enter a string to turn into an acronym. Type /quit to exit the program.\nInput: ")
}
