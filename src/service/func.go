package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"net"
	"time"
)

/*
	common functions
*/

func LoadPublicKey(key string) (*rsa.PublicKey, error) {
	data, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, errors.New("Decode base64 error !")
	}
	pk, err := x509.ParsePKIXPublicKey(data)
	if err != nil {
		return nil, errors.New("Parse PublickKey error !")

	}

	ret, ok := pk.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("Type convert to PublicKey error !")
	}
	return ret, nil
}

func TimeoutDialer(dailTimeout, rwTimeout int64) func(network, addr string) (net.Conn, error) {
	return func(network, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(network, addr, time.Duration(dailTimeout)*time.Second)
		if nil != conn {
			conn.SetDeadline(time.Now().Add(time.Duration(rwTimeout) * time.Second))
		}
		return conn, err
	}
}
