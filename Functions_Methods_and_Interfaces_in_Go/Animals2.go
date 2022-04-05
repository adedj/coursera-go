/*
Write a program which allows the user to create a set of animals and to get information about those animals. Each animal has a name and can be
either a cow, bird, or snake. With each command, the user can either create a new animal of one of the three types, or the user can request
information about an animal that he/she has already created. Each animal has a unique name, defined by the user. Note that the user can define
animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake. The following table contains the three types of
animals and their associated data.
Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your program should accept one command at
a time from the user, print out a response, and print out a new prompt on a new line. Your program should continue in this loop forever.
Every command from the user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”. The second string is an arbitrary string
which will be the name of the new animal. The third string is the type of the new animal, either “cow”, “bird”, or “snake”.  Your program should
process each newanimal command by creating the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”. The second string is the name of the animal.
The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal. Specifically, the Animal interface should contain the methods
Eat(), Move(), and Speak(), which take no arguments and return no values. The Eat() method should print the animal’s food, the Move() method should
print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. Define three types Cow, Bird, and Snake.
For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface.
When the user creates an animal, create an object of the appropriate type. Your program should call the appropriate method when the user issues
a query command.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define the animal interface with 3 methods
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Define the Cow type
type Cow struct {
	name       string
	food       string
	locomotion string
	sound      string
}

// Define the Name method for the Cow type (Cow is a reciever type for the Name method)
func (c Cow) Name() {
	fmt.Println(c.name)
}

// Define the Eat method for the Cow type (Cow is a reciever type for the Eat method)
func (c Cow) Eat() {
	fmt.Println(c.food)
}

// Define the Move method for the Cow type (Cow is a reciever type for the Move method)
func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

// Define the Speak method for the Cow type (Cow is a reciever type for the Speak method)
func (c Cow) Speak() {
	fmt.Println(c.sound)
}

// Define the Bird type
type Bird struct {
	food       string
	locomotion string
	sound      string
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b Bird) Speak() {
	fmt.Println(b.sound)
}

// Define the Snake type
type Snake struct {
	food       string
	locomotion string
	sound      string
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s Snake) Speak() {
	fmt.Println(s.sound)
}

func main() {

	sli := make([]Animal, 0)
	//animal = Cow{"test", "grass", "walk", "moo"}
	//animal = test
	//animal.Speak()

	//test, ok := animal(Cow)
	//fmt.Println(test, ok)

	//animal = Cow{"grass", "walk", "moo"}
	//var cow = Cow{"grass", "walk", "moo"}
	//var bird = Bird{"worms", "fly", "peep"}
	//var snake = Snake{"mice", "slither", "hsss"}

	for true {
		scanner := bufio.NewScanner(os.Stdin) // Create a scanner for user input
		fmt.Print("> ")
		scanner.Scan()                                 // Scans a line from Stdin(Console)
		strSlice := strings.Split(scanner.Text(), " ") // Return a slice with a number for each index

		// Limit the user input to three strings
		if len(strSlice) > 3 {
			strSlice = strSlice[:3]
		}

		command := strSlice[0]
		aninalName := strSlice[1]

		if command != "newanimal" || command != "query" {
			fmt.Println("Each command must start with newanimal or query")

		} else if command == "newanimal" {
			animalType := strSlice[2]

			switch animalType {
			case "cow":
				sli = append(Cow{aninalName, "grass", "walk", "moo"})
				//animal = Cow{aninalName, "grass", "walk", "moo"}
				fmt.Println("Created it!")
			}

		} else if command == "query" {
			method := strSlice[2]

			switch method {
			case "eat":

			}

		} else {
			fmt.Println("wrong animal")
		}

	}

}
