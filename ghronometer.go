package main

import (
	"flag"
	"fmt"

	"github.com/sendyHalim/ghronometer/executor"
	"github.com/sendyHalim/ghronometer/timer"
)

const CLR_G = "\x1b[32;1m"
const CLR_N = "\x1b[0m"

func main() {
	flag.Parse()
	var input []string = flag.Args()

	// if no input then just return
	if len(input) < 1 {
		return
	}

	var processName string = input[0]
	var args []string = input[1:]

	t := &timer.Timer{}
	// start the timer
	t.Start()

	executor.Execute(processName, args...)

	fmt.Printf("\n%selapsed time is %d ms%s\n", CLR_G, t.Counter(), CLR_N)
}
