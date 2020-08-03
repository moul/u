package u

import "io"

// SilentClose calls an io.Closer.Close() function and ignore potential errors.
//
// You can use it as `defer SilenceClose(f)`
func SilentClose(closer io.Closer) {
	if closer != nil {
		_ = closer.Close()
	}
}
