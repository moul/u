package u_test

import (
	"fmt"
	"os/exec"

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
