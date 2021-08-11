# joined

Provides functions for joining lists into natural language strings.

## Functions

### func [And](/joined.go#L8)

`func And(list []string) string`

And joins the last 2 elements of a list with " and ", while joining
preceding elements with a comma.

```golang

list := []string{"one", "two"}
fmt.Println(And(list)) // "one and two"

list = []string{"one", "two", "three"}
fmt.Println(And(list)) // "one, two and three"

```

### func [Arbitrary](/joined.go#L22)

`func Arbitrary(wordSeparator string, conjunction string, list []string) string`

Arbitrary joins the last 2 elements of a list with the supplied conjunction,
while joining preceding elements with the supplied word separator.  Note
that you'll need to provide your own padding for both the word separator and
the conjunction.

```golang

list := []string{"eins", "zwei", "drei"}
fmt.Println(Arbitrary(", ", " und ", list)) // "eins, zwei und drei"

```

### func [Or](/joined.go#L14)

`func Or(list []string) string`

Or joins the last 2 elements of a list with " or ", while joining
preceding elements with a comma.

```golang

list := []string{"one", "two"}
fmt.Println(Or(list)) // "one or two"

list = []string{"one", "two", "three"}
fmt.Println(Or(list)) // "one, two or three"

```

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
