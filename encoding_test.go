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

func ExamplePrettyJSON() {
	fmt.Println(u.PrettyJSON([]string{"hello", "world"}))
	fmt.Println(u.PrettyJSON(42))
	fmt.Println(u.PrettyJSON(nil))
	// Output:
	// [
	//   "hello",
	//   "world"
	// ]
	// 42
	// null
}

func ExampleJSON() {
	fmt.Println(u.JSON([]string{"hello", "world"}))
	fmt.Println(u.JSON(42))
	fmt.Println(u.JSON(nil))
	// Output:
	// ["hello","world"]
	// 42
	// null
}
