package main

import (
	"flag"
	"fmt"

	"github.com/go-task/task"
)

func main() {
	flag.CommandLine.Usage = func() {
		fmt.Println(`
task [target1 target2 ...]: Runs commands under targets like make.

Example: 'task hello' with the following 'Taskfile.yml' file will generate
an 'output.txt' file.
'''
hello:
  cmds:
    - echo "I am going to write a file named 'output.txt' now."
    - echo "hello" > output.txt
  generates:
    - output.txt
'''
`)
	}
	flag.BoolVar(&task.Force, "f", false, "forces execution even when the task is up-to-date")
	flag.Parse()
	task.Run()
}