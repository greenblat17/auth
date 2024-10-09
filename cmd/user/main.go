package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/greenblat17/auth/internal/model"
	"github.com/greenblat17/auth/internal/repository"
	"github.com/greenblat17/auth/internal/repository/user"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/golang/protobuf/ptypes/empty"
	desc "github.com/greenblat17/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	grpcPort = 50051
	dbDSN    = "host=localhost port=54321 dbname=auth user=auth-user password=auth-password sslmode=disable"
)

type server struct {
	desc.UnimplementedUserV1Server
	userRepository repository.UserRepository
}

// NewServer creates a new server
func NewServer(userRepository repository.UserRepository) *server {
	return &server{
		userRepository: userRepository,
	}
}

// Create creates a new user
func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	now := time.Now()
	createdUser := &model.User{
		Name:      req.GetName(),
		Email:     req.GetEmail(),
		Password:  req.GetPassword(),
		Role:      req.GetRole().String(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	id, err := s.userRepository.Create(ctx, createdUser)
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}

// Get gets a user
func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	getUser, err := s.userRepository.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	var role desc.Role
	switch getUser.Role {
	case model.RoleAdmin:
		role = desc.Role_ADMIN
	case model.RoleUser:
		role = desc.Role_USER
	default:
		role = desc.Role_UNKNOWN
	}

	return &desc.GetResponse{
		Id:        req.GetId(),
		Name:      getUser.Name,
		Email:     getUser.Email,
		Role:      role,
		CreatedAt: timestamppb.New(getUser.CreatedAt),
		UpdatedAt: timestamppb.New(getUser.UpdatedAt),
	}, nil
}

// Update updates a user
func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {
	updateUser := &model.User{
		ID:    req.GetId(),
		Name:  req.GetName().GetValue(),
		Email: req.GetEmail().GetValue(),
		Role:  req.GetRole().String(),
	}

	err := s.userRepository.Update(ctx, updateUser)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

// Delete deletes a user
func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	err := s.userRepository.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func main() {
	ctx := context.Background()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	conn, err := pgxpool.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	db := user.NewRepository(conn)
	server := NewServer(db)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, server)

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
