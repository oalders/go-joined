// Provides functions for joining lists into natural language strings.
package joined

import "strings"

// And joins 2 element lists with " and ", while joining lists with 3 or more
// elements with ", ", followed by " and " in place of the final comma.
// e.g. []string{"one", "two"} becomes "one and two".
// e.g. []string{"one", "two", "three"} becomes "one, two and three".
func And(y []string) string {
	return Arbitrary(", ", " and ", y)
}

// Or joins 2 element lists with " or ", while joining lists with 3 or more
// elements with ", ", followed by " or " in place of the final comma.
// e.g. []string{"one", "two"} becomes "one or two".
// e.g. []string{"one", "two", "three"} becomes "one, two or three".
func Or(y []string) string {
	return Arbitrary(", ", " or ", y)
}

// Arbitrary joins 2 element lists with the second string passed to the
// function, while joining lists with 3 or more elements with the first string
// followed by the second string in place of the final comma. Note that you'll
// need to provide your own padding for each string you wish to join with.
// e.g. join.Arbitrary(", ", " und ", []string{"eins", "zwei"} returns "eins und zwei".
// e.g. []string{"one", "two", "three"} becomes "one, two or three".
func Arbitrary(a string, b string, y []string) string {
	if len(y) == 0 {
		return ""
	}

	if len(y) == 1 {
		return y[0]
	}

	// Don't alter original slice
	x := make([]string, len(y))
	copy(x, y)

	// replace last two elements with a single element joined by "and"
	x[len(x)-2] = strings.Join(x[len(x)-2:], b)
	x = x[:len(x)-1]

	return strings.Join(x, a)
}
