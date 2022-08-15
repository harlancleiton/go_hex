package rpcuser

import (
	"github.com/harlancleiton/gopersonalities/adapter/grpc/pb"
	apiport "github.com/harlancleiton/gopersonalities/core/ports/left/api"
	"google.golang.org/grpc"
)

type GRPCUserAdapter struct {
	userApi apiport.UserAPIPort
	pb.UnimplementedUserServiceServer
}

func NewUserAdapter(userApi apiport.UserAPIPort) *GRPCUserAdapter {
	return &GRPCUserAdapter{userApi: userApi}
}

func (adapter GRPCUserAdapter) Register(grpcServer *grpc.Server) {
	pb.RegisterUserServiceServer(grpcServer, adapter)
}
