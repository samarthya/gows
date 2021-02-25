package idea

import (
	"fmt"
	"log"
	"os/exec"
)

// Cmd Starts a process for listing
func Cmd() {
	// Command returns the Cmd struct to execute the named program with the given arguments.
	cmd := exec.Command("idea")

	if cmd != nil {
		//Cmd structure
		fmt.Printf("%v\n", cmd.Env)
		fmt.Printf("%v\n", cmd.Dir)
	}

	if err := cmd.Run(); err == nil {
		log.Fatal(err)
	}
}
