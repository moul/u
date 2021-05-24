package u_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"moul.io/u"
)

func ExampleCombineFuncs() {
	cleanup := func() { fmt.Print("A") }
	cleanup = u.CombineFuncs(cleanup, func() { fmt.Print("B") })
	cleanup = u.CombineFuncs(func() { fmt.Print("C") }, cleanup)
	cleanup()
	// Output: CAB
}

func ExampleCheckErr() {
	_, err := http.Get("http://foo.bar")
	u.CheckErr(err) // panic
}

func ExampleFuture() {
	future := u.Future(func() (interface{}, error) {
		time.Sleep(100 * time.Millisecond)
		return "foobar", nil
	})

	// here, we can do some stuff

	ret := <-future
	fmt.Println("Ret:", ret.Ret)
	fmt.Println("Err:", ret.Err)

	// Output:
	// Ret: foobar
	// Err: <nil>
}

func BenchmarkCombineFuncs(b *testing.B) {
	f1 := func() {}
	f2 := func() {
		fmt.Println("A")
	}
	f3 := func() {
		fmt.Println("B")
		fmt.Println("C")
		fmt.Println("D")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.CombineFuncs(f1, f2, f3)
	}
}

func BenchmarkFuture(b *testing.B) {
	f1 := func() (interface{}, error) {
		return nil, nil
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.Future(f1)
	}
}
