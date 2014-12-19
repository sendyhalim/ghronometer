package executor

import (
	"os"
	"os/exec"
)

func Execute(processName string, argv ...string) {
	cmd := exec.Command(processName, argv...)
	// change the stdin, stdout, and stderr so the cmd will print output in the main process output 
	// http://stackoverflow.com/questions/8875038/redirect-stdout-pipe-of-child-process-in-golang
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}
