package daemon

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

// DaemonStatus is the status of the daemon.
type DaemonStatus string

const (
	DaemonStatusRunning DaemonStatus = "running"
	DaemonStatusStopped DaemonStatus = "stopped"
	DaemonStatusUnknown DaemonStatus = "unknown"
)

const pidFile = "/tmp/openclaw-gateway.pid"

// GetStatus returns the current daemon status.
func GetStatus() DaemonStatus {
	pid, err := readPID()
	if err != nil {
		return DaemonStatusStopped
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		return DaemonStatusStopped
	}
	if err := proc.Signal(syscall.Signal(0)); err != nil {
		return DaemonStatusStopped
	}
	return DaemonStatusRunning
}

// StartDaemon starts the gateway daemon in the background.
func StartDaemon(port int, args []string) error {
	if GetStatus() == DaemonStatusRunning {
		return fmt.Errorf("daemon is already running")
	}
	cmdArgs := append([]string{"gateway", "run", "--port", strconv.Itoa(port)}, args...)
	cmd := exec.Command(os.Args[0], cmdArgs...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start daemon: %w", err)
	}
	return writePID(cmd.Process.Pid)
}

// StopDaemon stops the gateway daemon.
func StopDaemon() error {
	pid, err := readPID()
	if err != nil {
		return fmt.Errorf("daemon is not running")
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("failed to find daemon process: %w", err)
	}
	if err := proc.Kill(); err != nil {
		return fmt.Errorf("failed to stop daemon: %w", err)
	}
	return os.Remove(pidFile)
}

func readPID() (int, error) {
	data, err := os.ReadFile(pidFile)
	if err != nil {
		return 0, err
	}
	pid, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		return 0, err
	}
	return pid, nil
}

func writePID(pid int) error {
	return os.WriteFile(pidFile, []byte(strconv.Itoa(pid)), 0o600)
}
