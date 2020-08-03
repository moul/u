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
