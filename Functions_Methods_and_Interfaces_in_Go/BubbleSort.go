/*
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing. The BubbleSort() function
should modify the slice so that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing,
but it should swap the contents of the slice in position i with the contents
in position i+1.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func Swap(s1 []int, index int) {

	swapSli := reflect.Swapper(s1) // Swapper() function is used to swap the elements in the provided slice
	swapSli(index, index+1)        // Swap contents of the slice from index with the contents from index+1
}

func BubbleSort(s2 []int) {

	for range s2 { // traverse the slice until the end
		lastInd := len(s2) - 1
		for i := range s2[:lastInd] { // traverse the slice s2 until last element minus one (since the last-1 element is compared with the last element)
			if s2[i] > s2[i+1] { // If element is greater than adjacent element swap them
				Swap(s2, i)
			}
		}
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin) // Create a scanner for user input
	fmt.Println("Enter a sequence of up to 10 integers")
	scanner.Scan()                                 // Scans a line from Stdin(Console)
	strSlice := strings.Split(scanner.Text(), " ") // Return a slice with a number for each index

	// Limit the user input to 10 characters
	if len(strSlice) > 10 {
		strSlice = strSlice[:10]
	}

	intSlice := make([]int, len(strSlice)) // Create an integer slice with the lenght equal to string slice

	// Convert the string slice to an integer slice
	for i, v := range strSlice {
		intSlice[i], _ = strconv.Atoi(v)
	}

	BubbleSort(intSlice) // Sort elements

	fmt.Println(intSlice)

}
