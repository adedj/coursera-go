/*
Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a
series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string
of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create
a struct which contains the first and last names found in the file. Each struct created will be added to a slice, and after all lines have
been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file,
your program should iterate through your slice of structs and print the first and last names found in each struct.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {

	sliceName := make([]Name, 0) // Declare a slice of type Name and lenght 0
	var objectName Name          // Declare the objectName of type Name

	fileName := bufio.NewScanner(os.Stdin) // Create a scanner for user input
	fmt.Println("You must enter the name of the text file")
	fileName.Scan()
	textFile := fileName.Text()

	f, err := os.Open(textFile) // Open the file

	if err != nil { // If opening the file results in an error print it and close the file
		fmt.Println("error:", err)
	}

	scanner := bufio.NewScanner(f) // Create a scanner that reads the file
	for scanner.Scan() {           // For each line in the file
		s := strings.Split(scanner.Text(), " ") // For each line in the text file it returns a slice with each index containing a word

		// If lenght of first name (s[0]) is > 20, use only the first 20 chars. (s is []strings)
		if len(s[0]) > 20 {
			s[0] = s[0][0:19]
		}

		// If lenght of last name (s[1]) is > 20, use only the first 20 chars
		if len(s[1]) > 20 {
			s[1] = s[1][0:19]
		}

		objectName = Name{fname: s[0], lname: s[1]} // Create an object of type Name containg the first and last name
		sliceName = append(sliceName, objectName)   // add the resulting object to the slice
	}

	for _, v := range sliceName { // Iterate the slice and print the values of the Name object
		fmt.Println(v.fname, v.lname)
	}

}
