package main

import (
	"fmt"
	"os"

	cli "github.com/y44k0v/grpc-rest-api-ykv/gen/grpc/cli/shop"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func doGRPC(scheme, host string, timeout int, debug bool) (goa.Endpoint, any, error) {
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to gRPC server at %s: %v\n", host, err)
	}
	return cli.ParseEndpoint(conn)
}
