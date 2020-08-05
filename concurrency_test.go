package u_test

import (
	"context"
	"fmt"
	"time"

	"moul.io/u"
)

func ExampleUniqueChild() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	parent := u.NewUniqueChild(ctx)

	parent.SetChild(func(ctx context.Context) {
		select {
		case <-time.After(50 * time.Millisecond):
			fmt.Print("A")
		case <-ctx.Done():
		}
	})

	time.Sleep(100 * time.Millisecond)

	parent.SetChild(func(ctx context.Context) {
		select {
		case <-time.After(50 * time.Millisecond):
			fmt.Print("B")
		case <-ctx.Done():
		}
	})

	parent.SetChild(func(ctx context.Context) {
		select {
		case <-time.After(50 * time.Millisecond):
			fmt.Print("C")
		case <-ctx.Done():
		}
	})

	time.Sleep(100 * time.Millisecond)

	// Output: AC
}

func ExampleUniqueChild_CloseChild() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	parent := u.NewUniqueChild(ctx)

	parent.SetChild(func(ctx context.Context) {
		select {
		case <-time.After(50 * time.Millisecond):
			fmt.Print("A")
		case <-ctx.Done():
		}
	})

	time.Sleep(100 * time.Millisecond)

	parent.SetChild(func(ctx context.Context) {
		select {
		case <-time.After(50 * time.Millisecond):
			fmt.Print("B")
		case <-ctx.Done():
		}
	})

	parent.SetChild(func(ctx context.Context) {
		select {
		case <-time.After(50 * time.Millisecond):
			fmt.Print("C")
		case <-ctx.Done():
		}
	})

	parent.CloseChild()

	time.Sleep(100 * time.Millisecond)

	// Output: A
}
