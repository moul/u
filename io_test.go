package u_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"moul.io/u"
)

type closer struct {
	called bool
}

func (c *closer) Close() error {
	c.called = true
	return nil
}

func TestSilentClose(t *testing.T) {
	c := &closer{}
	require.False(t, c.called)
	u.SilentClose(c)
	require.True(t, c.called)

	require.NotPanics(t, func() { u.SilentClose(nil) })
}

func ExampleSilentClose() {
	f, _ := os.Open("file.txt")
	defer u.SilentClose(f)
}

func ExampleCaptureStdout() {
	fmt.Println("AAA")
	closer, err := u.CaptureStdout()
	if err != nil {
		panic(err)
	}
	fmt.Println("BBB")
	ret := closer()
	fmt.Println("CCC")
	fmt.Println(ret)
	// Output:
	// AAA
	// CCC
	// BBB
}

func ExampleMustCaptureStdout() {
	fmt.Println("AAA")
	closer := u.MustCaptureStdout()
	fmt.Println("BBB")
	ret := closer()
	fmt.Println("CCC")
	fmt.Println(ret)
	// Output:
	// AAA
	// CCC
	// BBB
}

func ExampleCaptureStderr() {
	os.Stderr = os.Stdout // hack to run this test as an example test

	fmt.Fprintln(os.Stderr, "AAA")
	closer, err := u.CaptureStderr()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stderr, "BBB")
	ret := closer()
	fmt.Fprintln(os.Stderr, "CCC")
	fmt.Fprintln(os.Stderr, ret)
	// Output:
	// AAA
	// CCC
	// BBB
}

func ExampleMustCaptureStderr() {
	os.Stderr = os.Stdout // hack to run this test as an example test

	fmt.Fprintln(os.Stderr, "AAA")
	closer := u.MustCaptureStderr()
	fmt.Fprintln(os.Stderr, "BBB")
	ret := closer()
	fmt.Fprintln(os.Stderr, "CCC")
	fmt.Fprintln(os.Stderr, ret)
	// Output:
	// AAA
	// CCC
	// BBB
}

func ExampleCaptureStdoutAndStderr() {
	os.Stderr = os.Stdout // hack to run this test as an example test

	fmt.Fprintln(os.Stderr, "AAA")
	fmt.Println("BBB")
	closer, err := u.CaptureStdoutAndStderr()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stderr, "CCC")
	fmt.Println("DDD")
	ret := closer()
	fmt.Fprintln(os.Stderr, "EEE")
	fmt.Println("FFF")
	fmt.Println(ret)
	// Output:
	// AAA
	// BBB
	// EEE
	// FFF
	// CCC
	// DDD
}

func ExampleMustCaptureStdoutAndStderr() {
	os.Stderr = os.Stdout // hack to run this test as an example test

	fmt.Fprintln(os.Stderr, "AAA")
	fmt.Println("BBB")
	closer := u.MustCaptureStdoutAndStderr()
	fmt.Fprintln(os.Stderr, "CCC")
	fmt.Println("DDD")
	ret := closer()
	fmt.Fprintln(os.Stderr, "EEE")
	fmt.Println("FFF")
	fmt.Println(ret)
	// Output:
	// AAA
	// BBB
	// EEE
	// FFF
	// CCC
	// DDD
}

func TestCaptureStderr(t *testing.T) {
	closer := u.MustCaptureStderr()
	fmt.Fprintln(os.Stderr, "hello world")
	ret := closer()
	require.Equal(t, ret, "hello world\n")
}
