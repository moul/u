package u_test

import (
	"fmt"
	"testing"

	"moul.io/u"
)

func ExampleBoolPtr() {
	truePtr := u.BoolPtr(true)
	falsePtr := u.BoolPtr(false)
	fmt.Println("true ptr:  ", *truePtr)
	fmt.Println("false ptr: ", *falsePtr)
	// Output:
	// true ptr:   true
	// false ptr:  false
}

func BenchmarkBoolPtr(b *testing.B) {
	b.Run("serial", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			u.BoolPtr(true)
		}
	})
	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				u.BoolPtr(true)
			}
		})
	})
}
