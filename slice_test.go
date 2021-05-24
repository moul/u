package u_test

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"

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

func BenchmarkUniqueStrings(b *testing.B) {
	cases := []struct {
		Name string
		Data []string
	}{
		{"slice1", []string{"foofoo", "barbar", "foo", "bazbaz", "foo", "bar", "foobar", "baz", "foobaz"}},
		{"slice2", bigSliceOfString()},
	}

	for _, bc := range cases {
		b.Run(bc.Name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				u.UniqueStrings(bc.Data)
			}
		})
		b.Run(bc.Name+"-parallel", func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					u.UniqueStrings(bc.Data)
				}
			})
		})
	}
}

func BenchmarkUniqueInts(b *testing.B) {
	cases := []struct {
		Name string
		Data []int
	}{
		{"slice1", []int{4, 5, 2, 6, 7, 7, 2, 4, 2, 1, 1, 2, 3, 4}},
		{"slice2", bigSliceOfInt()},
	}

	for _, bc := range cases {
		b.Run(bc.Name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				u.UniqueInts(bc.Data)
			}
		})
		b.Run(bc.Name+"-parallel", func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					u.UniqueInts(bc.Data)
				}
			})
		})
	}
}

func BenchmarkUniqueInterfaces(b *testing.B) {
	cases := []struct {
		Name string
		Data []interface{}
	}{
		{"slice1", []interface{}{1, 2, 2, 1, 3, 4, "foo", "bar", "foo", "foo", "barfoo"}},
		{"slice2", bigSliceOfInterface()},
	}

	for _, bc := range cases {
		b.Run(bc.Name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				u.UniqueInterfaces(bc.Data)
			}
		})
		b.Run(bc.Name+"-parallel", func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					u.UniqueInterfaces(bc.Data)
				}
			})
		})
	}
}

func bigSliceOfInt() []int {
	var ns []int
	rand.Seed(42)
	for i := 0; i < 1000; i++ {
		nb := rand.Intn(500)
		ns = append(ns, nb)
	}
	return ns
}

func bigSliceOfString() []string {
	var ss []string
	rand.Seed(42)
	for i := 0; i < 1000; i++ {
		rn := rand.Intn(500)
		ss = append(ss, strconv.Itoa(rn))
	}
	return ss
}

func bigSliceOfInterface() []interface{} {
	var is []interface{}
	rand.Seed(42)
	for i := 0; i < 1000; i++ {
		nb := rand.Intn(500)
		is = append(is, nb)
	}
	for i := 0; i < 1000; i++ {
		rn := rand.Intn(500)
		is = append(is, strconv.Itoa(rn))
	}
	return is
}
