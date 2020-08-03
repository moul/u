package u_test

import (
	"fmt"

	"moul.io/u"
)

func ExampleCombineFuncs() {
	cleanup := func() { fmt.Print("A") }
	cleanup = u.CombineFuncs(cleanup, func() { fmt.Print("B") })
	cleanup = u.CombineFuncs(func() { fmt.Print("C") }, cleanup)
	cleanup()
	// Output: CAB
}
