package main

import (
	"context"
	"flag"
	"log"
	"net"
	"time"

	"github.com/greenblat17/auth/internal/config"
	"github.com/greenblat17/auth/internal/config/env"

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

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

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
		UpdatedAt: &now,
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
		UpdatedAt: timestamppb.New(getUser.GetUpdateAt()),
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
	flag.Parse()
	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer pool.Close()

	db := user.NewRepository(pool)
	server := NewServer(db)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, server)

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
