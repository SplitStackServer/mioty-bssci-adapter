package bssci_v1

import (
	"net"
	"time"

	"github.com/pkg/errors"
)

type TcpKeepAliveListener struct {
	*net.TCPListener
	keepAlivePeriod time.Duration
}

func NewTcpKeepAliveListener(addr string, keepAlivePeriod   time.Duration) (*TcpKeepAliveListener, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, errors.Wrap(err, "invalid bind address")
	}

	ln, err := net.ListenTCP("tcp", tcpAddr)

	if err != nil {
		return nil, errors.Wrap(err, "create tcp listener error")
	}
	return &TcpKeepAliveListener{ln, keepAlivePeriod}, nil
}

func (ln TcpKeepAliveListener) Accept() (c net.Conn, err error) {
	tc, err := ln.AcceptTCP()

	if err != nil {
		return
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(ln.keepAlivePeriod)
	return tc, nil
}
