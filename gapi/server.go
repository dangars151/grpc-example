package gapi

import "grpc-example/pb"

type Server struct {
	pb.UnimplementedSimpleBankServer
}

func NewServer() (*Server, error) {
	server := &Server{}

	return server, nil
}
