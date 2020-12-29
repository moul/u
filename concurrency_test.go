package u_test

import (
	"context"
	"fmt"
	"sort"
	"strings"
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

func ExampleFanIn() {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})
	merged := u.FanIn(ch1, ch2, ch3)
	done := make(chan bool)
	received := []string{}

	go func() {
		for item := range merged {
			fmt.Println("tick")
			received = append(received, fmt.Sprintf("%v", item))
		}
		done <- true
	}()

	ch1 <- 1
	ch2 <- 2
	ch3 <- 3
	close(ch1)
	ch2 <- 4
	ch2 <- 5
	ch3 <- 6
	close(ch2)
	ch3 <- 7
	close(ch3)

	<-done

	sort.Strings(received)
	fmt.Println(strings.Join(received, ", "))

	// Output:
	// tick
	// tick
	// tick
	// tick
	// tick
	// tick
	// tick
	// 1, 2, 3, 4, 5, 6, 7
}
