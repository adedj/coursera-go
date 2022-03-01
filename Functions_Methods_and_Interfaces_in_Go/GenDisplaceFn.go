/*
Let us assume the following formula for
displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

s = ½ a t2 + vot + so

Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function
called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial
displacement so. GenDisplaceFn()
should return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume
the following values for acceleration, initial velocity, and initial
displacement: a = 10, vo = 2, so = 1. I can use the
following statement to call GenDisplaceFn() to
generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to
print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print
the displacement after 5 seconds.

fmt.Println(fn(5))
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
func Read() {
	scanner := bufio.NewScanner(os.Stdin) // Create a scanner for user input
	fmt.Println("Enter values for acceleration, initial velocity, and initial displacement (ex: 10 5 4)")
	scanner.Scan()                                 // Scans a line from Stdin(Console)
	strSlice := strings.Split(scanner.Text(), " ") // Return a slice with a number for each index

	// Limit the user input to 3 characters
	if len(strSlice) > 3 {
		strSlice = strSlice[:3]
		fmt.Println("You entered more than three values, only the first three where used.")
	}

	floatSlice := make([]float64, 0) // Create an empty float64 slice

	// Convert the string slice to an float slice
	for _, v := range strSlice {

		x, _ := strconv.ParseFloat(v, 64)  // Convert each value of the string slice into a float64
		floatSlice = append(floatSlice, x) // Add the value to the float slice
	}

	acc := floatSlice[0]
	vel := floatSlice[1]
	dis := floatSlice[2]

	fmt.Printf("You entered acceleration a=%v, initial velocity v=%v and initial displacement s=%v\n", acc, vel, dis)

	// Using Scanln to store the time input
	fmt.Println("Enter the time")
	var time int

	fmt.Scanln(&time) // Taking input from user
	fmt.Printf("You entered time t=%v\n", time)

}
*/
func GenDisplaceFn(a, v, s float64) func(float64) float64 {

	displacement := func(t float64) float64 {

		return 0.5*a*math.Pow(t, 2) + v*t + s
	}

	return displacement
}

func main() {

	scanner := bufio.NewScanner(os.Stdin) // Create a scanner for user input
	fmt.Println("Enter values for acceleration, initial velocity, and initial displacement (ex: 10 5 4)")
	scanner.Scan()                                 // Scans a line from Stdin(Console)
	strSlice := strings.Split(scanner.Text(), " ") // Return a slice with a number for each index

	// Limit the user input to 3 characters
	if len(strSlice) > 3 {
		strSlice = strSlice[:3]
		fmt.Println("You entered more than three values, only the first three where used.")
	}

	floatSlice := make([]float64, 0) // Create an empty float64 slice

	// Convert the string slice to an float slice
	for _, v := range strSlice {

		x, _ := strconv.ParseFloat(v, 64)  // Convert each value of the string slice into a float64
		floatSlice = append(floatSlice, x) // Add the value to the float slice
	}

	acc := floatSlice[0]
	vel := floatSlice[1]
	dis := floatSlice[2]

	fmt.Printf("You entered acceleration a=%v, initial velocity v=%v and initial displacement s=%v\n", acc, vel, dis)

	// Using Scanln to store the time input
	fmt.Println("Enter the time")
	var time float64

	fmt.Scanln(&time) // Taking input from user
	fmt.Printf("You entered time t=%v\n", time)

	fn := GenDisplaceFn(acc, vel, dis)
	fmt.Println(fn(time))
}
