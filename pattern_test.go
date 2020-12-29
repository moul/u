package u_test

import (
	"fmt"
	"net/http"
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
