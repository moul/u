package u

// CombineFuncs create a chain of functions.
// This can be particularly useful for creating cleanup function progressively.
// It solves the infinite loop you can have when trying to do it manually: https://play.golang.org/p/NQem8UJ500t.
func CombineFuncs(left func(), right ...func()) func() {
	return func() {
		left()
		for _, fn := range right {
			fn()
		}
	}
}

// CheckErr panics if the passed error is not nil.
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Future starts running the given function in background and return a chan that will return the result of the execution.
func Future(fn func() (interface{}, error)) <-chan FutureRet {
	c := make(chan FutureRet, 1)
	go func() {
		ret, err := fn()
		c <- FutureRet{Ret: ret, Err: err}
	}()
	return c
}

// FutureRet is a generic struct returned by Future.
type FutureRet struct {
	Ret interface{}
	Err error
}
