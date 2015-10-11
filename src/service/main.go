package main

import (
	"log"
	"net"
	"os"

	apple "service/appstore"
	google "service/playstore"

	"google.golang.org/grpc"
	//	log "github.com/gonet2/libs/nsq-logger"
	//	_ "github.com/gonet2/libs/statsd-pprof"
)

const (
	_port   = ":60003"
	SERVICE = "IAP SERVICE"
)

func main() {
	log.SetPrefix(SERVICE)
	lis, err := net.Listen("tcp", _port)
	if err != nil {
		os.Exit(-1)
	}
	s := grpc.NewServer()
	appleService := apple.New(1, 10, 10)
	googleService := google.New()
	apple.RegisterService(s, appleService)
	google.RegisterService(s, googleService)
	s.Serve(lis)
}
