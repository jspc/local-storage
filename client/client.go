package local

import (
	"context"
	"crypto/tls"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/j-and-j-global/local-storage/local"
	"github.com/j-and-j-global/storage-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Opts holds configuration options for talking to an
// entities service gRPC server.
//
// It may be used to fine tune connections and stuff
type Opts struct {
	// TLS determines whether or not a gRPC connection
	// should connect with TLS or not.
	TLS *bool
}

// Secure returns a set of Opts which are configured for
// secure access to the entities service, including:
//
// 1. Setting TLS to true
func Secure() Opts {
	t := true

	return Opts{
		TLS: &t,
	}
}

// Insecure returns a set of Opts which are configured for
// insecure access to the entities service, including:
//
// 1. Setting TLS to false
//
// These options should really only be used in development of
// services
func Insecure() Opts {
	f := false

	return Opts{
		TLS: &f,
	}
}

// Client connects to an instance of the Entities service, ahead of
// entity extraction operations.
//
// It is concurrency-safe
type Client struct {
	client local.LocalStorageClient
}

// New creates an Entities service Client, creating all connections
// and testing them.
//
// This allows client code to Fail Fast
func New(addr string, opts Opts) (c Client, err error) {
	conf := grpc.WithInsecure()
	if opts.TLS != nil && *opts.TLS {
		conf = grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{}))
	}

	conn, err := grpc.Dial(addr, conf)
	if err != nil {
		return
	}

	c.client = local.NewLocalStorageClient(conn)

	return
}

func (c Client) Status() (status *storage.Status, err error) {
	return c.client.Status(context.Background(), &empty.Empty{})
}
