package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func runCmd(program string, args ...string) {
	cmd := exec.Command(program, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
func checkExeExists() {
	path, err := exec.LookPath("ls")
	if err != nil {
		fmt.Printf("didn't find 'ls' executable\n")
	} else {
		fmt.Printf("'ls' executable is in '%s'\n", path)
	}
}

func runCmdChannels(program string, ch chan<- error, args ...string) {
	cmd := exec.Command(program, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		ch <- errors.New(fmt.Sprintf("err %v ", err))
		fmt.Printf("%v\n", err)
	}

}

func main() {
	for _, arg := range os.Args[1:] {
		runCmd(os.Args[1], arg)
	}

	pingErrorChan := make(chan error)
	if len(os.Args[1:]) >= 3 {
		for _, arg := range os.Args[1:] {
			go runCmdChannels(os.Args[1], pingErrorChan, arg)
		}
	}

}
