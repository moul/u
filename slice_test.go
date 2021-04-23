package u_test

import (
	"fmt"

	"moul.io/u"
)

func ExampleUniqueStrings() {
	fmt.Println(u.UniqueStrings([]string{"foo", "bar", "foo", "baz", "foo", "bar", "foobar", "baz", "foobaz"}))
	// Output: [foo bar baz foobar foobaz]
}

func ExampleUniqueInts() {
	fmt.Println(u.UniqueInts([]int{13, 51, 36, 69, 92, 92, 42, 21, 36, 13, 51}))
	// Output: [13 51 36 69 92 42 21]
}

func ExampleUniqueInterfaces() {
	fmt.Println(u.UniqueInterfaces([]interface{}{13, "foo", "bar", 42, 13, 43, "baz"}))
	// Output: [13 foo bar 42 43 baz]
}
