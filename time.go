package u

import (
	"fmt"
	"strings"
	"time"
)

// ShortDuration returns a short human-friendly representation of a duration.
// For duration < 100 days, the output length will be <= 7.
func ShortDuration(d time.Duration) string {
	var s string
	switch {
	case d < time.Microsecond:
		return d.String()
	case d < time.Millisecond:
		return d.Round(time.Microsecond / 10).String()
	case d < time.Second:
		return d.Round(time.Millisecond / 10).String()
	case d < time.Minute:
		return d.Round(time.Second / 10).String()
	case d < time.Hour:
		s = d.Round(time.Second).String()
	case d < time.Hour*24:
		s = d.Round(time.Second).String()
	default:
		days := float64(d) / float64(time.Hour*24)
		d %= (time.Hour * 24)
		s = fmt.Sprintf("%dd%s", int(days), d.Round(time.Minute).String())
	}
	if len(s) > 2 {
		s = strings.TrimSuffix(s, "0s")
	}
	if len(s) > 2 {
		s = strings.TrimSuffix(s, "0m")
	}
	return s
}
