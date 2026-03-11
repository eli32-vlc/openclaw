package process

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"
)

// RunResult holds the result of a command run.
type RunResult struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

// RunCommand runs a command and returns its output.
func RunCommand(name string, args ...string) (*RunResult, error) {
	cmd := exec.Command(name, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	result := &RunResult{
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		ExitCode: 0,
	}
	if exitErr, ok := err.(*exec.ExitError); ok {
		result.ExitCode = exitErr.ExitCode()
		return result, nil
	}
	return result, err
}

// RunCommandWithTimeout runs a command with a timeout.
func RunCommandWithTimeout(timeoutMs int, name string, args ...string) (*RunResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutMs)*time.Millisecond)
	defer cancel()
	cmd := exec.CommandContext(ctx, name, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if ctx.Err() == context.DeadlineExceeded {
		return nil, fmt.Errorf("command timed out after %dms", timeoutMs)
	}
	result := &RunResult{
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		ExitCode: 0,
	}
	if exitErr, ok := err.(*exec.ExitError); ok {
		result.ExitCode = exitErr.ExitCode()
		return result, nil
	}
	return result, err
}

// ShellRun runs a shell command string.
func ShellRun(command string) (*RunResult, error) {
	return RunCommand("sh", "-c", command)
}
