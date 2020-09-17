package u_test

import (
	"context"

	"moul.io/u"
)

func ExampleWaitForCtrlC() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		// do your stuff
		<-ctx.Done()
	}()

	u.WaitForCtrlC()
}
