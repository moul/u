package u

import (
	"bytes"
	"fmt"
	"os/exec"
)

// ExecStandaloneOutputs runs the command and returns its standard output and standard error.
func ExecStandaloneOutputs(cmd *exec.Cmd) ([]byte, []byte, error) {
	var (
		outbuf = &bytes.Buffer{}
		errbuf = &bytes.Buffer{}
	)
	cmd.Stdout = outbuf
	cmd.Stderr = errbuf
	err := cmd.Run()
	return outbuf.Bytes(), errbuf.Bytes(), err
}

// SafeExec runs a command and return a string containing the combined standard output and standard error.
// If the program fails, the result of `err` is appended to the output.
func SafeExec(cmd *exec.Cmd) string {
	outBytes, err := cmd.CombinedOutput()
	out := string(outBytes)
	if err != nil {
		if out != "" {
			out += "\n"
		}
		out += fmt.Sprintf("error: %v\n", err)
	}
	return out
}

// CommandExists checks whether a command is available in the $PATH.
func CommandExists(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}
