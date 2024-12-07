package access

import (
	"context"
	"errors"
	"strings"

	desc "github.com/greenblat17/auth/pkg/access_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	authorizationHeader = "authorization"
	authPrefix          = "Bearer "
)

// Check has user access to resource
func (i *Implementation) Check(ctx context.Context, req *desc.CheckRequest) (*emptypb.Empty, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata is not provided")
	}
	accessToken, err := getTokenFromMetadata(md)
	if err != nil {
		return nil, err
	}

	err = i.accessService.Check(ctx, accessToken, req.GetEndpointAddress())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func getTokenFromMetadata(md metadata.MD) (string, error) {
	authHeader, ok := md[authorizationHeader]
	if !ok || len(authHeader) == 0 {
		return "", errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return "", errors.New("invalid authorization header format")
	}

	return strings.TrimPrefix(authHeader[0], authPrefix), nil
}
