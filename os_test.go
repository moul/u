package u_test

import (
	"fmt"
	"io/ioutil"
	"os"

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

func ExampleExpandUser() {
	os.Setenv("HOME", "/home/foo") // just for example
	ret, err := u.ExpandUser("~/hello-world.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)
	// Output: /home/foo/hello-world.txt
}

func ExampleMustExpandUser() {
	os.Setenv("HOME", "/home/foo") // just for example
	ret := u.MustExpandUser("~/hello-world.txt")
	fmt.Println(ret)
	// Output: /home/foo/hello-world.txt
}
