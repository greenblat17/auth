package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	desc "github.com/greenblat17/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
)

const (
	grpcPort = 50051
)

type server struct {
	desc.UnimplementedUserV1Server
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	return &desc.CreateResponse{
		Id: 1,
	}, nil
}
func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	now := timestamppb.Now()

	return &desc.GetResponse{
		Id:        req.GetId(),
		Name:      "name",
		Email:     "email",
		Role:      desc.Role_USER,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}