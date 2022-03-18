package grpcApi

import (
	api "github.com/hawkkiller/wtc-account-service-api/api"
	"google.golang.org/grpc"
)

type AccountServerGRPC struct {
	grpc *grpc.Server
	port string
}

type AccountService struct {
	api.UnimplementedUserServiceServer
}
