package u_test

import (
	"fmt"

	"moul.io/u"
)

func ExampleB64Encode() {
	fmt.Println(u.B64Encode([]byte("hello world!")))
	// Output: aGVsbG8gd29ybGQh
}

func ExampleB64Decode() {
	ret, _ := u.B64Decode("aGVsbG8gd29ybGQh")
	fmt.Println(string(ret))
	// Output: hello world!
}
