package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// Input takes in an input from a user and prints it out in the terminal
//
// For instance if we are trying to print out a name, it can be done like this
//
//	name, err = utils.Input("Enter your name")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Your name is &s", name)
func Input(stmt string) (string, error) {
	fmt.Printf("%s ", stmt)

	reader := bufio.NewScanner(os.Stdin)

	if ok := reader.Scan(); !ok {
		return "", errors.New("error scanning input")
	}

	return reader.Text(), nil
}
