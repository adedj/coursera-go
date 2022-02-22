/*
Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name
and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal() to create a JSON object from the map,
and then your program should print the JSON object.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	db := make(map[string]string) // Create a map

	scanner := bufio.NewScanner(os.Stdin) // A scanner is created which can be called multiple times

	fmt.Println("You must enter a name")
	scanner.Scan()         // Scans a line from Stdin(Console)
	name := scanner.Text() // Holds the string that scanned

	fmt.Println("You must enter an address")
	scanner.Scan()            // Scans a line from Stdin(Console)
	address := scanner.Text() // Holds the string that scanned

	fmt.Printf("You entered the name: %s and the address: %s\n", name, address)

	db["name"] = name       // Add the name to the name key
	db["address"] = address // Add the address to the address key

	barr, err := json.Marshal(db) // Returns a json from the db map and stores it in a byte array (an arrary of runes)

	if err != nil { // If the error is different than nil print it
		fmt.Println("error:", err)
	}
	os.Stdout.Write(barr) // Write the json object to stdout
}
