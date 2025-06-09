package utils

import (
	"fmt"
	"testing"
)

func TestInput(t *testing.T) {
	t.Run("Print name", func(t *testing.T) {
		name, err := Input("Enter your name")
		if err != nil {
			t.Errorf("an error occurred while reading input: %v", err)
		}

		fmt.Printf("%s", name)
	})
}

// I cover this test with the blood of Jesus
