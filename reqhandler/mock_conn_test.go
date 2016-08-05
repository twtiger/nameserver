package reqhandler

import (
	"net"
	"time"
)

type mockConn struct{}

func (m *mockConn) Read(b []byte) (n int, err error) { return 0, nil }

func (m *mockConn) Write(b []byte) (n int, err error) { return 0, nil }

func (m *mockConn) Close() error { return nil }

func (m *mockConn) LocalAddr() net.Addr { return nil }

func (m *mockConn) RemoteAddr() net.Addr { return nil }

func (m *mockConn) SetDeadline(t time.Time) error { return nil }

func (m *mockConn) SetReadDeadline(t time.Time) error { return nil }

func (m *mockConn) SetWriteDeadline(t time.Time) error { return nil }
