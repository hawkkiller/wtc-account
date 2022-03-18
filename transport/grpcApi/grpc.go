package grpcApi

import (
	"fmt"
	api "github.com/hawkkiller/wtc-account-service-api/api"
	"google.golang.org/grpc"
	"net"
	"os"
)

func NewServerGRPC() (s *AccountServerGRPC) {
	server := grpc.NewServer()
	api.RegisterUserServiceServer(server, &AccountService{})
	port := os.Getenv("GPORT")
	if port == "" {
		port = "8000"
	}
	return &AccountServerGRPC{server, port}
}

func (s *AccountServerGRPC) StartServerGRPC() error {
	if lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", s.port)); err != nil {
		return err
	} else {
		fmt.Println("Started GRPC server on port ", s.port)
		return s.grpc.Serve(lis)
	}
}
