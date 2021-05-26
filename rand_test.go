package u_test

import (
	"fmt"
	"math/rand"
	"testing"

	"moul.io/u"
)

func ExampleRandomLetters() {
	rand.Seed(42)
	fmt.Println(u.RandomLetters(8))
	fmt.Println(u.RandomLetters(8))
	fmt.Println(u.RandomLetters(8))
	fmt.Println(u.RandomLetters(42))
	// Output:
	// HRukpTTu
	// eZPtNeuv
	// unhuksqV
	// GzAdxlgghEjkMVeZJpmKqakmTRgKfBSWYjUNGkdmdt
}

func BenchmarkRandomLetters(b *testing.B) {
	cases := []struct {
		Name            string
		NumberOfLetters int
	}{
		{"1000", 100},
		{"10000", 1000},
		{"100000", 10000},
	}
	for _, bc := range cases {
		b.Run(bc.Name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				u.RandomLetters(bc.NumberOfLetters)
			}
		})
		b.Run(bc.Name+"-parallel", func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					u.RandomLetters(bc.NumberOfLetters)
				}
			})
		})
	}
}
