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
		DataSize int
	}{
		{DataSize: 1000},
		{DataSize: 100000},
		{DataSize: 100000000},
	}
	for _, bc := range cases {
		for i := 0; i < b.N; i++ {
			b.Run(fmt.Sprintf("%d", bc.DataSize), func(b *testing.B) {
				u.B64Encode(bytes.Repeat([]byte{'A'}, bc.DataSize))
			})
		}
	}
}

func BenchmarkB64Decode(b *testing.B) {
	cases := []struct {
		DataSize int
	}{
		{DataSize: 1000},
		{DataSize: 100000},
		{DataSize: 100000000},
	}
	for _, bc := range cases {
		for i := 0; i < b.N; i++ {
			b.Run(fmt.Sprintf("%d", bc.DataSize), func(b *testing.B) {
				_, err := u.B64Decode(strings.Repeat("a", bc.DataSize))
				if err != nil {
					b.Error(err)
				}
			})
		}
	}
}
