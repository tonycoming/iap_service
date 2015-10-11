package main

import (
	"log"
	"net"
	"os"

	pb "proto"

	"google.golang.org/grpc"
)

//	log "github.com/gonet2/libs/nsq-logger"

//	_ "github.com/gonet2/libs/statsd-pprof"

const (
	_port = ":60003"
)

func main() {
	log.SetPrefix(SERVICE)
	lis, err := net.Listen("tcp", _port)
	if err != nil {
		//	log.Critical(err)
		os.Exit(-1)
	}
	//	log.Info("listening on ", lis.Addr())
	s := grpc.NewServer()
	ins := New(1, 10, 10)
	pb.RegisterIAPServiceServer(s, ins)
	s.Serve(lis)
}
