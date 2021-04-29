package u_test

import (
	"fmt"
	"math/rand"

	"moul.io/u"
)

func ExampleRandomLetters() {
	rand.Seed(42)
	fmt.Println(u.RandomLetters(8))
	fmt.Println(u.RandomLetters(8))
	fmt.Println(u.RandomLetters(8))
	fmt.Println(u.RandomLetters(42))
	// Output:
	// HRukpTTu
	// eZPtNeuv
	// unhuksqV
	// GzAdxlgghEjkMVeZJpmKqakmTRgKfBSWYjUNGkdmdt
}
