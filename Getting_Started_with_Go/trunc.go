/*
Write a program which prompts the user to enter a floating point number and prints the integer which is a truncated version of
the floating point number that was entered. Truncation is the process of removing the digits to the right of the decimal place.

Submit your source code for the program, “trunc.go”.
*/

package main

import (
	"fmt"
	"math"
)

func main() {

	var number float64 // Declare the number variable

	fmt.Println("You must enter a floating point number:")
	fmt.Scan(&number) // The user's input value is stored into the number variable

	truncNr := math.Trunc(number) // Returns the integer value of the number entered by the user
	fmt.Println("The integer portion of your number is:", truncNr)
}
