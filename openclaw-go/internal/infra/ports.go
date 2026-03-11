package infra

import (
	"fmt"
	"net"
)

const DefaultGatewayPort = 18789

// PortInUseError is returned when a port is already in use.
type PortInUseError struct {
	Port    int
	Details string
}

func (e *PortInUseError) Error() string {
	return fmt.Sprintf("port %d is already in use", e.Port)
}

// EnsurePortAvailable checks if a port is available.
func EnsurePortAvailable(port int) error {
	addr := fmt.Sprintf(":%d", port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return &PortInUseError{Port: port}
	}
	ln.Close()
	return nil
}

// DescribePortOwner attempts to describe what process is using a port.
func DescribePortOwner(port int) string {
	return fmt.Sprintf("port %d", port)
}
