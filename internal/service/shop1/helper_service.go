package shop1

import (
	"context"
	"strconv"

	loginPb "github.com/Shemistan/uzum_shop/pkg/login_v1"
	"google.golang.org/grpc/metadata"
)

func (s *shopSystemService) GetUserIdFromLoginServ(ctx context.Context) (int, error) {
	emp := &loginPb.Check_Request{EndpointAddress: ""}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ctx = metadata.NewOutgoingContext(ctx, md)
	}

	// TODO: get the ID from auth_service
	// replace _ with the current response
	_, err := s.loginClient.Check(ctx, emp)
	if err != nil {
		return 0, err
	}

	// Place the actual ID converted to int
	userId, err := strconv.Atoi("dasdads")
	if err != nil {
		return 0, err
	}

	return userId, nil
}
