package u

import (
	"bytes"
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
