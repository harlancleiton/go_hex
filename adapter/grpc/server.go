package rpcserver

import (
	"net"

	rpcuser "github.com/harlancleiton/gopersonalities/adapter/grpc/user"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	user rpcuser.GRPCUserAdapter
}

func NewAdapter(userAdapter rpcuser.GRPCUserAdapter) *GRPCServer {
	return &GRPCServer{user: userAdapter}
}

func (adapter GRPCServer) StartServer() error {
	listener, err := net.Listen("tcp", ":5001")

	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	adapter.user.Register(grpcServer)

	err = grpcServer.Serve(listener)

	if err != nil {
		return err
	}

	return nil
}
