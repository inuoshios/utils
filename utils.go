package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Input takes in an input from a user and prints it out in the terminal
//
// For instance if we are trying to print out a name, it can be done like this
//
//	name = utils.Input("Enter your name")
//	fmt.Printf("Your name is &s", name)
func Input(stmt string) string {
	fmt.Printf("%s ", stmt)

	reader := bufio.NewScanner(os.Stdin)

	if ok := reader.Scan(); !ok {
		log.Fatal("error scanning input")
	}

	return reader.Text()
}
