package u_test

import (
	"fmt"
	"net/http"

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
