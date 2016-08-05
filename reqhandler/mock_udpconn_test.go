package reqhandler

import (
	"net"
	"time"
)

type mockUDPConn struct {
	data []byte
}

func (m *mockUDPConn) Read(b []byte) (n int, err error) { return 0, nil }

func (m *mockUDPConn) Write(b []byte) (n int, err error) { return 0, nil }

func (m *mockUDPConn) Close() error { return nil }

func (m *mockUDPConn) LocalAddr() net.Addr { return nil }

func (m *mockUDPConn) RemoteAddr() net.Addr { return nil }

func (m *mockUDPConn) SetDeadline(t time.Time) error { return nil }

func (m *mockUDPConn) SetReadDeadline(t time.Time) error { return nil }

func (m *mockUDPConn) SetWriteDeadline(t time.Time) error { return nil }

func (m *mockUDPConn) ReadFrom(b []byte) (n int, addr net.Addr, err error) {
	copy(b, m.data)
	return 0, nil, nil
}

func (m *mockUDPConn) WriteTo(b []byte, addr net.Addr) (n int, err error) { return 0, nil }
