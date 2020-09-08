package u_test

import (
	"fmt"
	"time"

	"moul.io/u"
)

func ExampleShortDuration() {
	fmt.Println(u.ShortDuration(time.Nanosecond * 0))
	fmt.Println(u.ShortDuration(time.Nanosecond))
	fmt.Println(u.ShortDuration(time.Nanosecond * 12))
	fmt.Println(u.ShortDuration(time.Nanosecond * 123))
	fmt.Println(u.ShortDuration(time.Nanosecond * 1234))
	fmt.Println(u.ShortDuration(time.Nanosecond * 12345))
	fmt.Println(u.ShortDuration(time.Nanosecond * 123456))
	fmt.Println(u.ShortDuration(time.Nanosecond * 1234567))
	fmt.Println(u.ShortDuration(time.Nanosecond * 12345678))
	fmt.Println(u.ShortDuration(time.Nanosecond * 123456789))
	fmt.Println(u.ShortDuration(time.Nanosecond * 1234567891))
	fmt.Println(u.ShortDuration(time.Nanosecond * 12345678912))
	fmt.Println(u.ShortDuration(time.Nanosecond * 123456789123))
	fmt.Println(u.ShortDuration(time.Nanosecond * 1234567891234))
	fmt.Println(u.ShortDuration(time.Nanosecond * 12345678912345))
	fmt.Println(u.ShortDuration(time.Nanosecond * 123456789123456))
	fmt.Println(u.ShortDuration(time.Nanosecond * 1234567891234567))
	fmt.Println(u.ShortDuration(time.Nanosecond * 12345678912345678))
	fmt.Println(u.ShortDuration(time.Nanosecond * 123456789123456789))
	fmt.Println(u.ShortDuration(time.Nanosecond * 1234567891234567891))
	fmt.Println("-------")
	fmt.Println(u.ShortDuration(time.Nanosecond))
	fmt.Println(u.ShortDuration(time.Microsecond))
	fmt.Println(u.ShortDuration(time.Millisecond))
	fmt.Println(u.ShortDuration(time.Second))
	fmt.Println(u.ShortDuration(time.Minute))
	fmt.Println(u.ShortDuration(time.Hour))
	fmt.Println(u.ShortDuration(time.Hour * 24))
	fmt.Println(u.ShortDuration(time.Hour + time.Second))
	// Output:
	// 0s
	// 1ns
	// 12ns
	// 123ns
	// 1.2µs
	// 12.3µs
	// 123.5µs
	// 1.2ms
	// 12.3ms
	// 123.5ms
	// 1.2s
	// 12.3s
	// 2m3s
	// 20m35s
	// 3h25m46s
	// 1d10h18m
	// 14d6h56m
	// 142d21h21m
	// 1428d21h33m
	// 14288d23h32m
	// -------
	// 1ns
	// 1µs
	// 1ms
	// 1s
	// 1m
	// 1h
	// 1d
	// 1h0m1s
}
