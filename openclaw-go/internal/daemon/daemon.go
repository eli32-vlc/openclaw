package daemon

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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

// pidFilePath returns the PID file path under ~/.openclaw/run/ which is
// owned by the user and not world-writable like /tmp.
func pidFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		home = "."
	}
	runDir := filepath.Join(home, ".openclaw", "run")
	return filepath.Join(runDir, "openclaw-gateway.pid")
}

// ensureRunDir creates the run directory with 0700 permissions.
func ensureRunDir() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("cannot determine home directory: %w", err)
	}
	runDir := filepath.Join(home, ".openclaw", "run")
	return os.MkdirAll(runDir, 0o700)
}

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
	return os.Remove(pidFilePath())
}

func readPID() (int, error) {
	data, err := os.ReadFile(pidFilePath())
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
	if err := ensureRunDir(); err != nil {
		return fmt.Errorf("failed to create run directory: %w", err)
	}
	return os.WriteFile(pidFilePath(), []byte(strconv.Itoa(pid)), 0o600)
}
