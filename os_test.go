package u_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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

func ExamplePathExists() {
	file, err := ioutil.TempFile("", "bar")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(file.Name())
	fmt.Println(u.PathExists("/laksjdflkasdjflaksdjfalskdfjasdlfkj")) // should not exist
	fmt.Println(u.PathExists(file.Name()))
	fmt.Println(u.PathExists(filepath.Dir(file.Name())))
	// Output:
	// false
	// true
	// true
}

func ExampleDirExists() {
	file, err := ioutil.TempFile("", "bar")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(file.Name())
	fmt.Println(u.DirExists("/laksjdflkasdjflaksdjfalskdfjasdlfkj")) // should not exist
	fmt.Println(u.DirExists(file.Name()))
	fmt.Println(u.DirExists(filepath.Dir(file.Name())))
	// Output:
	// false
	// false
	// true
}

func ExampleFileExists() {
	file, err := ioutil.TempFile("", "bar")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(file.Name())
	fmt.Println(u.FileExists("/laksjdflkasdjflaksdjfalskdfjasdlfkj")) // should not exist
	fmt.Println(u.FileExists(file.Name()))
	fmt.Println(u.FileExists(filepath.Dir(file.Name())))
	// Output:
	// false
	// true
	// false
}
