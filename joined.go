// Provides functions for joining lists into natural language strings.
package joined

import "strings"

// And joins the last 2 elements of a list with " and ", while joining
// preceding elements with a comma.
func And(list []string) string {
	return Arbitrary(", ", " and ", list)
}

// Or joins the last 2 elements of a list with " or ", while joining
// preceding elements with a comma.
func Or(list []string) string {
	return Arbitrary(", ", " or ", list)
}

// Arbitrary joins the last 2 elements of a list with the supplied conjunction,
// while joining preceding elements with the supplied word separator.  Note
// that you'll need to provide your own padding for both the word separator and
// the conjunction.
func Arbitrary(wordSeparator string, conjunction string, list []string) string {
	if len(list) == 0 {
		return ""
	}

	if len(list) == 1 {
		return list[0]
	}

	// Don't alter original slice
	x := make([]string, len(list))
	copy(x, list)

	// Replace last two elements with a single element joined by the conjunction.
	// For 3+ items, use an Oxford comma (word separator before the conjunction).
	joinWith := conjunction
	if len(list) >= 3 {
		joinWith = strings.TrimRight(wordSeparator, " ") + conjunction
	}
	x[len(x)-2] = strings.Join(x[len(x)-2:], joinWith)
	x = x[:len(x)-1]

	return strings.Join(x, wordSeparator)
}
