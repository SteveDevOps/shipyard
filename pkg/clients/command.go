package clients

import (
	"os/exec"
	"time"

	"github.com/hashicorp/go-hclog"
)

type Command interface {
	Execute(string, ...string) error
}

// Command executes local commands
type CommandImpl struct {
	timeout time.Duration
	log     hclog.Logger
}

// NewCommand creates a new command with the given logger and maximum command time
func NewCommand(maxCommandTime time.Duration, l hclog.Logger) Command {
	return &CommandImpl{maxCommandTime, l}
}

// Execute the given command
func (c *CommandImpl) Execute(command string, args ...string) error {

	cmd := exec.Command(
		command,
		args...,
	)

	// set the standard out and error to the logger
	cmd.Stdout = c.log.StandardWriter(&hclog.StandardLoggerOptions{InferLevels: true})
	cmd.Stderr = c.log.StandardWriter(&hclog.StandardLoggerOptions{InferLevels: true})

	/*
		// wait for timeout
		t := time.AfterFunc(c.timeout, func() {
			// kill the running process
			cmd.Process.Kill()
			return fmt.Errorf("Command timed out before completing")
		})
	*/

	err := cmd.Run()
	if err != nil {
		return err
	}

	// command has completed clear the timeout timer
	/*
		if t!= nil {
			t.Stop()
		}
	*/

	return nil
}
