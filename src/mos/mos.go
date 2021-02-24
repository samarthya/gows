package mos

import (
	"fmt"
	"os"
)

// Cmd Starts a process for listing
func Cmd() {
	// Environ returns a copy of strings representing the environment,
	env := os.Environ()
	if len(env) > 0 {
		for i, v := range env {
			fmt.Println(i, ": ", v)

		}
	}

	// ProcAttr holds the attributes that will be applied to a new process
	attrs := os.ProcAttr{
		Dir: "/Users/ss670121/sourcebox",
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}

	fmt.Println("List files...")
	pid, err := os.StartProcess("/bin/ls", []string{"ls", "-ls"}, &attrs)

	if err != nil {
		fmt.Printf("Error %v starting process!", err) //
		os.Exit(1)
	}

	fmt.Printf("The process id is %d\n[%v]\n", pid.Pid, pid)

	os.Exit(0)
}
