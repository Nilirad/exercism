// Contains the implementation for the "Two Fer" exercise.
package twofer

import "fmt"

// Returns a string saying:
// "One for (name), one for me.",
// using the provided `name`.
// If `name` is an empty string,
// it returns:
// "One for you, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
