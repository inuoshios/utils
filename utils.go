package utils

import (
	"bufio"
	"fmt"
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
	reader := bufio.NewReader(os.Stdin)
	read, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return read
}
