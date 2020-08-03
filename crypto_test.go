package u_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"moul.io/u"
)

func TestSha1(t *testing.T) {
	require.NotPanics(t, func() { u.Sha1(nil) })
	require.NotPanics(t, func() { u.Sha1Hex(nil) })
}

func ExampleSha1() {
	fmt.Println(u.Sha1([]byte("hello world!")))
	// Output:
	// [67 12 227 77 2 7 36 237 117 161 150 223 194 173 103 199 119 114 209 105]
}

func ExampleSha1Hex() {
	fmt.Println(u.Sha1Hex([]byte("hello world!")))
	// Output:
	// 430ce34d020724ed75a196dfc2ad67c77772d169
}
