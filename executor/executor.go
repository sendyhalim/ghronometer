package executor

import (
	"os"
	"os/exec"
)

func Execute(processName string, argv ...string) {
	cmd := exec.Command(processName, argv...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}
