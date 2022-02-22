/* Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character
‘a’. The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are
upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("You must enter a string")

	scanner := bufio.NewScanner(os.Stdin) // Make a newscanner object that has the user input
	scanner.Scan()                        // Scan the line and store the text
	input := scanner.Text()               // Store the user input into a variable

	lowerText := strings.ToLower(input) // Make all text lowercase

	/*
	   For Debugging only
	   fmt.Printf("User input text: %v\n", input)
	   fmt.Println("Lowercase text:", lowerText)
	*/

	prefix := strings.HasPrefix(lowerText, "i")
	sufix := strings.HasSuffix(lowerText, "n")
	contains := strings.ContainsAny(lowerText, "a")

	if prefix && sufix && contains {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
