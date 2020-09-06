package u_test

import (
	"fmt"

	"moul.io/godev"
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

func ExamplePrettyJSON() {
	fmt.Println(godev.PrettyJSON([]string{"hello", "world"}))
	fmt.Println(godev.PrettyJSON(42))
	fmt.Println(godev.PrettyJSON(nil))
	// Output:
	// [
	//   "hello",
	//   "world"
	// ]
	// 42
	// null
}

func ExampleJSON() {
	fmt.Println(godev.JSON([]string{"hello", "world"}))
	fmt.Println(godev.JSON(42))
	fmt.Println(godev.JSON(nil))
	// Output:
	// ["hello","world"]
	// 42
	// null
}
