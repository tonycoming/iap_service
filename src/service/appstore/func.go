package appstore

import (
	"net"
	"time"
)

func TimeoutDialer(dailTimeout, rwTimeout int64) func(network, addr string) (net.Conn, error) {
	return func(network, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(network, addr, time.Duration(dailTimeout)*time.Second)
		if nil != conn {
			conn.SetDeadline(time.Now().Add(time.Duration(rwTimeout) * time.Second))
		}
		return conn, err
	}
}
