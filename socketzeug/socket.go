package socketzeug

import (
	"github.com/plombardi89/gozeug/closezeug"
	"net"
)

// GetFreeTCPPort returns an available TCP port. A retrieved port is not guaranteed to be available by the time of
// use.
func GetFreeTCPPort() (int, error) {
	address, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}

	listener, err := net.ListenTCP("tcp", address)
	if err != nil {
		return 0, err
	}

	defer closezeug.CloseOrSwallowError(listener)

	return listener.Addr().(*net.TCPAddr).Port, nil
}

// GetFreeTCPPorts returns as many desired TCP ports as specified in count.
func GetFreeTCPPorts(count int) ([]int, error) {
	result := make([]int, 0)
	for idx := 0; idx < count; idx++ {
		port, err := GetFreeTCPPort()
		if err != nil {
			return result, err
		}

		result = append(result, port)
	}

	return result, nil
}
