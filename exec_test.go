package u_test

import (
	"fmt"
	"os/exec"
	"testing"

	"moul.io/u"
)

func ExampleExecStandaloneOutputs() {
	stdout, stderr, err := u.ExecStandaloneOutputs(exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr"))
	fmt.Print("stdout: ", string(stdout))
	fmt.Print("stderr: ", string(stderr))
	fmt.Println("err: ", err)
	// Output:
	// stdout: stdout
	// stderr: stderr
	// err:  <nil>
}

func ExampleSafeExec() {
	out := u.SafeExec(exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr; exit 1"))
	fmt.Println(out)
	// Output:
	// stdout
	// stderr
	//
	// error: exit status 1
}

func ExampleCommandExists() {
	fmt.Println(u.CommandExists("go"))
	fmt.Println(u.CommandExists("asldkglsakdjaslkdg"))
	// Output:
	// true
	// false
}

func BenchmarkCommandExists(b *testing.B) {
	cases := []struct {
		Command string
	}{
		{"go"},
		{"asddsa"},
	}
	for _, bc := range cases {
		b.Run(bc.Command, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				u.CommandExists(bc.Command)
			}
		})
		b.Run(bc.Command+"-parallel", func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					u.CommandExists(bc.Command)
				}
			})
		})
	}
}

func BenchmarkSafeExec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u.SafeExec(exec.Command("true"))
	}
}
