package utils

import (
	"fmt"
	"testing"
)

func TestInput(t *testing.T) {
	t.Run("Print name", func(t *testing.T) {
		name := Input("Enter your name")

		fmt.Printf("%s", name)
	})
}

// I cover this test with the blood of Jesus
