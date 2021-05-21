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
	stringSize := 10000000
	s := strings.Repeat("a", stringSize)
	for i := 0; i < b.N; i++ {
		bytes, _ := json.Marshal(s)
		u.B64Encode(bytes)
	}
}

// BenchmarkB64Encode12345 - case benchmark tests
func BenchmarkB64Encode1(b *testing.B) {
	benchmarkB64Encode(100, b)
}
func BenchmarkB64Encode2(b *testing.B) {
	benchmarkB64Encode(10000, b)
}
func BenchmarkB64Encode3(b *testing.B) {
	benchmarkB64Encode(1000000, b)
}
func BenchmarkB64Encode4(b *testing.B) {
	benchmarkB64Encode(100000000, b)
}
func BenchmarkB64Encode5(b *testing.B) {
	benchmarkB64Encode(10000000000, b)
}

func benchmarkB64Encode(size int, b *testing.B) {
	s := strings.Repeat("s", size)
	bytes, _ := json.Marshal(s)
	u.B64Encode(bytes)
}

// BenchmarkB64Decode repetitive benchmark test for b64 decoding
func BenchmarkB64Decode(b *testing.B) {
	stringSize := 10000000
	s := strings.Repeat("YWFh", stringSize)
	for i := 0; i < b.N; i++ {
		_, err := u.B64Decode(s)
		if err != nil {
			return
		}
	}
}

// BenchmarkB64Decode12345 - case benchmark tests
func BenchmarkB64Decode1(b *testing.B) {
	benchmarkB64Decode(100, b)
}
func BenchmarkB64Decode2(b *testing.B) {
	benchmarkB64Decode(10000, b)
}
func BenchmarkB64Decode3(b *testing.B) {
	benchmarkB64Decode(1000000, b)
}
func BenchmarkB64Decode4(b *testing.B) {
	benchmarkB64Decode(100000000, b)
}
func BenchmarkB64Decode5(b *testing.B) {
	benchmarkB64Decode(10000000000, b)
}

func benchmarkB64Decode(size int, b *testing.B) {
	s := strings.Repeat("YWFh", size)
	_, err := u.B64Decode(s)
	if err != nil {
		b.Fail()
	}
}
