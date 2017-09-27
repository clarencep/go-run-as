package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"syscall"
)

func main() {
	var err error
	if len(os.Args) < 3 {
		printHelp()
		return
	}

	user := os.Args[1]
	program := os.Args[2]
	programArgs := os.Args[3:]

	uid, err := parseUid(user)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to get uid by `%s`: %v\n", user, err)
		os.Exit(1)
	}

	var exeFile string
	programDir, _ := path.Split(program)
	if programDir == "" {
		exeFile, err = exec.LookPath(program)
		if err != nil {
			exeFile = program
		}
	} else {
		exeFile = program
	}

	cmd := exec.Command(exeFile, programArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = setCmdCredential(cmd, int64(uid), -1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to set uid `%d`, error: %v\n", uid, err)
		os.Exit(1)
	}

	err = cmd.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to start the command: %v\n", err)
		os.Exit(1)
		return
	}

	err = cmd.Wait()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			// This program has exited with exit code != 0
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				retCode := status.ExitStatus()
				os.Exit(retCode)
				return
			}
		}
	}

	return
}

func printHelp() {
	fmt.Println("Usage: ")
	fmt.Println("	run-as <user> <program> [args...]")
	os.Exit(1)
}
