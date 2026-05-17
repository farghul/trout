package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type ExecOptions struct {
	Stream bool
	Env    []string
	Dir    string
}

// Execute a command with the given arguments and options, returning the output or an error
func execute(task string, args []string, opts ExecOptions) ([]byte, error) {
	cmd := exec.Command(task, args...)
	cmd.Env = append(os.Environ(), opts.Env...)
	cmd.Dir = opts.Dir

	if opts.Stream {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return nil, cmd.Run()
	}

	return cmd.CombinedOutput()
}

// Record a message to a log file
func journal(message string) {
	file, err := os.OpenFile("/data/automation/logs/trout.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	inspect(err)
	log.SetOutput(file)
	log.Println(message)
}

// Check for errors, print the result if found
func inspect(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}

// Confirm the current working directory is correct
func changedir() {
	os.Chdir(bitbucket.WordPress)
	var filePath string = "composer-prod.json"

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		alert("This is not the correct folder,")
	}
}

// Println function for colourized text
func (c Color) Println(text string) {
	fmt.Println(string(c) + text + Reset)
}

// Printf function for colourized text
func (c Color) Printf(format string, a ...any) {
	fmt.Printf(string(c)+format+Reset, a...)
}
