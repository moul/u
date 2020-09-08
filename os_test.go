package u_test

import (
	"fmt"
	"io/ioutil"

	"moul.io/u"
)

func ExampleTempfileWithContent() {
	f, cleanup, err := u.TempfileWithContent([]byte("AAA\nBBB\nCCC"))
	if err != nil {
		panic(err)
	}
	defer cleanup()

	out, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	// Output:
	// AAA
	// BBB
	// CCC
}

func ExampleMustTempfileWithContent() {
	f, cleanup := u.MustTempfileWithContent([]byte("AAA\nBBB\nCCC"))
	defer cleanup()

	out, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	// Output:
	// AAA
	// BBB
	// CCC
}
