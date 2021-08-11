package joined

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleAnd() {
	list := []string{"one", "two"}
	fmt.Println(And(list)) // "one and two"

	list = []string{"one", "two", "three"}
	fmt.Println(And(list)) // "one, two and three"
}

func ExampleOr() {
	list := []string{"one", "two"}
	fmt.Println(Or(list)) // "one or two"

	list = []string{"one", "two", "three"}
	fmt.Println(Or(list)) // "one, two or three"
}

func ExampleArbitrary() {
	list := []string{"eins", "zwei", "drei"}
	fmt.Println(Arbitrary(", ", " und ", list)) // "eins, zwei und drei"
}

func TestAnd(t *testing.T) {
	assert.Equal(t, "one, two and three", And([]string{"one", "two", "three"}))
	assert.Equal(t, "one", And([]string{"one"}))
	assert.Equal(t, "one and two", And([]string{"one", "two"}))
	assert.Equal(t, "", And([]string{""}))
}

func TestOr(t *testing.T) {
	assert.Equal(t, "one, two or three", Or([]string{"one", "two", "three"}))
	assert.Equal(t, "one", Or([]string{"one"}))
	assert.Equal(t, "one or two", Or([]string{"one", "two"}))
	assert.Equal(t, "", Or([]string{""}))
}

func TestArbitrary(t *testing.T) {
	assert.Equal(
		t,
		"eins, zwei und drei",
		Arbitrary(", ", " und ", []string{"eins", "zwei", "drei"}),
	)
	assert.Equal(t, "eins", Arbitrary(
		", ",
		" und ",
		[]string{"eins"},
	))
	assert.Equal(
		t,
		"eins und zwei",
		Arbitrary(", ", " und ", []string{"eins", "zwei"}),
	)
	assert.Equal(t, "", Arbitrary(
		", ",
		" und ",
		[]string{""},
	))
}
