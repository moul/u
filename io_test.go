package u_test

import (
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
