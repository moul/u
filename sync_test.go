package u_test

import (
	"fmt"
	"sync"
	"time"

	"moul.io/u"
)

func ExampleMutexMap() {
	var wg sync.WaitGroup
	var m u.MutexMap

	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("abc%d", i)
		wg.Add(1)
		go func(i int) {
			fmt.Printf("init %s\n", name)
			defer m.Lock("abc-X")()
			sleep := time.Duration(100/(i+1)) * time.Millisecond
			time.Sleep(sleep)
			fmt.Printf("run %s (after sleep of %s)\n", name, sleep)
			wg.Done()
		}(i)
		time.Sleep(1 * time.Millisecond)
	}

	wg.Wait()

	// Output:
	// init abc0
	// init abc1
	// init abc2
	// init abc3
	// init abc4
	// init abc5
	// init abc6
	// init abc7
	// init abc8
	// init abc9
	// run abc0 (after sleep of 100ms)
	// run abc1 (after sleep of 50ms)
	// run abc2 (after sleep of 33ms)
	// run abc3 (after sleep of 25ms)
	// run abc4 (after sleep of 20ms)
	// run abc5 (after sleep of 16ms)
	// run abc6 (after sleep of 14ms)
	// run abc7 (after sleep of 12ms)
	// run abc8 (after sleep of 11ms)
	// run abc9 (after sleep of 10ms)
}
