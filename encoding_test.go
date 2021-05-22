package u_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

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

// BenchmarkB64Encode - repetitive benchmark test for b64 encoding
func BenchmarkB64Encode(b *testing.B) {
	cases := []struct {
		StringSize int
	}{
		{StringSize: 1000},
		{StringSize: 100000},
		{StringSize: 100000000},
	}
	var s string
	for i := 0; i < b.N; i++ {
		for _, bc := range cases {
			b.Run(fmt.Sprintf("String Size = %d", bc.StringSize), func(b *testing.B) {
				s = strings.Repeat("A", bc.StringSize)
				bytes, _ := json.Marshal(s)
				u.B64Encode(bytes)
			})
		}
	}

}
