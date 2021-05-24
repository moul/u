package u_test

import (
	"bytes"
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

func BenchmarkB64Encode(b *testing.B) {
	cases := []struct {
		Name string
		Data []byte
	}{
		{"1", bytes.Repeat([]byte{'A'}, 1)},
		{"1000", bytes.Repeat([]byte{'A'}, 1000)},
		{"1000000000", bytes.Repeat([]byte{'A'}, 1000000)},
	}
	b.ResetTimer()
	for _, bc := range cases {
		b.Run(bc.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				u.B64Encode(bc.Data)
			}
		})
	}
}

func BenchmarkB64Decode(b *testing.B) {
	cases := []struct {
		Name string
		Data string
	}{
		{"1000", strings.Repeat("a", 1000)},
		{"10000", strings.Repeat("a", 10000)},
		{"100000", strings.Repeat("a", 100000)},
	}
	b.ResetTimer()
	for _, bc := range cases {
		b.Run(bc.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := u.B64Decode(bc.Data)
				if err != nil {
					b.Error(err)
				}
			}
		})
	}
}
