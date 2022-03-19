/*
Write a program which allows the user to get information about a predefined set of animals. Three animals are predefined, cow, bird, and snake.
Each animal can eat, move, and speak. The user can issue a request to find out one of three things about an animal: 1) the food that it eats,
2) its method of locomotion, and 3) the sound it makes when it speaks. The following table contains the three animals and their associated data
which should be hard-coded into your program.
Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your program accepts one request at a time
from the user, prints out the answer to the request, and prints out a new prompt. Your program should continue in this loop forever.
Every request from the user must be a single line containing 2 strings. The first string is the name of an animal, either “cow”, “bird”, or “snake”.
The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. Your program should process
each request by printing out the requested data.
You will need a data structure to hold the information about each animal. Make a type called Animal which is a struct containing three fields:
food, locomotion, and noise, all of which are strings. Make three methods called Eat(), Move(), and Speak(). The receiver type of all of your
methods should be your Animal type. The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion,
and the Speak() method should print the animal’s spoken sound. Your program should call the appropriate method when the user makes a request.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Declare Animal struct type
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Method that prints the animal food
func (a Animal) Eat() {
	fmt.Println(a.food)
}

// Method that prints the animal locomotion
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

// Method that prints the animal spoken sound
func (c Animal) Speak() {
	fmt.Println(c.noise)
}

func main() {

	// Declare the animal variables
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hss"}

	fmt.Println("Enter the name of the animal (cow, bird or snake) and the information requested (eat, move or speak).")

	// Loop forever
	for true {
		scanner := bufio.NewScanner(os.Stdin) // Create a scanner for user input
		fmt.Print("> ")
		scanner.Scan()                                 // Scans a line from Stdin(Console)
		strSlice := strings.Split(scanner.Text(), " ") // Return a slice with a number for each index

		// Limit the user input to two characters
		if len(strSlice) > 2 {
			strSlice = strSlice[:2]
		}

		animal := strSlice[0]
		data := strSlice[1]

		// Check if user input is correct
		if animal != "cow" || animal != "bird" || animal != "snake" || data != "eat" || data != "move" || data != "speak" {
			fmt.Println("You did not enter the correct data")
		}

		// Process the user request based on animal type and information
		switch {
		case animal == "cow" && data == "eat":
			cow.Eat()
		case animal == "cow" && data == "move":
			cow.Move()
		case animal == "cow" && data == "speak":
			cow.Speak()
		case animal == "bird" && data == "eat":
			bird.Eat()
		case animal == "bird" && data == "move":
			bird.Move()
		case animal == "bird" && data == "speak":
			bird.Speak()
		case animal == "snake" && data == "eat":
			snake.Eat()
		case animal == "snake" && data == "move":
			snake.Move()
		case animal == "snake" && data == "speak":
			snake.Speak()
		}
	}
}
