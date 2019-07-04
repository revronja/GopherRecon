package main

import (
	"fmt"
	"os"
	"os/exec"
)

func run_cmd(program string, args ...string) {
	cmd := exec.Command(program, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
func checkExeExists(program string) {
	path, err := exec.LookPath(program)
	if err != nil {
		fmt.Printf("didn't find 'ls' executable\n")
	} else {
		fmt.Printf("'ls' executable is in '%s'\n", path)
	}
}


func main() {
	for _, arg := range os.Args[1:] {
		run_cmd(os.Args[1], arg)
	}

}
