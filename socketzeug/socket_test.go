package socketzeug_test

import (
	"fmt"
	"github.com/plombardi89/gozeug/closezeug"
	"github.com/plombardi89/gozeug/socketzeug"
	"github.com/stretchr/testify/assert"
	"io"
	"net"
	"testing"
)

func TestGetFreeTCPPort(t *testing.T) {
	port, err := socketzeug.GetFreeTCPPort()
	assert.NoError(t, err)
	assert.NotZero(t, port)
	assert.True(t, port < 65535)

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		t.Error(err)
	}

	defer closezeug.CloseOrHandleError(listener, func(closer io.Closer, e error) {
		t.Fatal(e)
	})
}

func TestGetFreeTCPPorts(t *testing.T) {
	ports, err := socketzeug.GetFreeTCPPorts(1000)
	assert.NoError(t, err)
	for _, p := range ports {
		listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", p))
		if err != nil {
			t.Error(err)
		}

		closezeug.CloseOrHandleError(listener, func(closer io.Closer, e error) {
			t.Fatal(e)
		})
	}
}
