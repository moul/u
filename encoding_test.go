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

func ExampleIsBinary() {
	fmt.Println(u.IsBinary([]byte{'c', 'h', 'i', 'c', 'k', 'e', 'n'}))
	fmt.Println(u.IsBinary([]byte{'c', 'h', 'i', 0, 'k', 'e', 'n'}))
	// Output:
	// false
	// true
}

func BenchmarkB64Encode(b *testing.B) {
	cases := []struct {
		Name string
		Data []byte
	}{
		{"1", bytes.Repeat([]byte{'A'}, 1)},
		{"1000", bytes.Repeat([]byte{'A'}, 1000)},
		{"1000000", bytes.Repeat([]byte{'A'}, 1000000)},
	}
	for _, bc := range cases {
		b.Run(bc.Name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				u.B64Encode(bc.Data)
			}
		})
		b.Run(bc.Name+"-parallel", func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					u.B64Encode(bc.Data)
				}
			})
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
	for _, bc := range cases {
		b.Run(bc.Name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, err := u.B64Decode(bc.Data)
				if err != nil {
					b.Error(err)
				}
			}
		})
		b.Run(bc.Name+"-parallel", func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_, err := u.B64Decode(bc.Data)
					if err != nil {
						b.Error(err)
					}
				}
			})
		})
	}
}

func BenchmarkIsBinary(b *testing.B) {
	cases := []struct {
		Name string
		Data []byte
	}{
		{"small-valid", bytes.Repeat([]byte{'A'}, 1)},
		{"long-valid", bytes.Repeat([]byte{'A'}, 80000)},
		{"small-invalid", []byte{'c', 'h', 'i', 0, 'k', 'e', 'n'}},
		// long-invalid
	}
	for _, bc := range cases {
		b.Run(bc.Name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				u.IsBinary(bc.Data)
			}
		})
		b.Run(bc.Name+"-parallel", func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					u.IsBinary(bc.Data)
				}
			})
		})
	}
}

func ExampleIsASCII() {
	fmt.Println(u.IsASCII([]byte("hello")))
	// Output: true
}

func BenchmarkIsASCII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u.IsASCII([]byte("hello"))
	}
}
