/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop.
Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop,
the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice,
and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides
to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	slice := make([]int, 0, 3) // Initialize an integer slice of size 3

	fmt.Printf("len: %d, cap: %d, s: %v\n", len(slice), cap(slice), slice)

	for true {
		fmt.Println("Enter an integer to be added to the slice")
		scanner := bufio.NewScanner(os.Stdin) // Make a newscanner object that has the user input
		scanner.Scan()                        // Scan the line and store the text
		input := scanner.Text()               // Store the user input into a variable
		number, _ := strconv.Atoi(input)      // Convert the input text to an integer (_ ignore the error)

		if input == "X" {
			fmt.Println("You entered X to exit the program")
			os.Exit(0) // Exit the program
		}

		slice = append(slice, number) // Append the user input to the slice
		sort.Ints(slice)              // Sort the slice
		fmt.Println(slice)            // Print the slice
	}
}
