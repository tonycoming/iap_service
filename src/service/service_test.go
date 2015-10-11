package main

import (
	pb "proto/appstore"
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:60003"
)

func TestApplePay(t *testing.T) {
	// Set up a connection to the server.
	grpc.WithInsecure()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAppleIapServiceClient(conn)
	// Contact the server and print out its response.
	r, err := c.ApplePayVerify(context.Background(), &pb.Request{""})
	if err != nil {
		t.Fatalf("could not get next value: %v", err)
	}

	t.Logf("%v", r)
}

func BenchmarkTestApplePay(t *testing.B) {
	// Set up a connection to the server.
	grpc.WithInsecure()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAppleIapServiceClient(conn)

	// Contact the server and print out its response.
	for n := 0; n < t.N; n++ {
		_, err := c.ApplePayVerify(context.Background(), &pb.Request{""})
		if err != nil {
			t.Fatalf("could not get next value: %v, ", err)
		}
	}
}
