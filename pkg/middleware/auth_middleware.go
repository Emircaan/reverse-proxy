package middleware

import (
	"context"
	"fmt"

	contracts "github.com/Emircaan/reverse-proxy/ethereum/contracts/auth"
	"github.com/Emircaan/reverse-proxy/pkg/proto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"google.golang.org/grpc/metadata"
)

type AuthMiddleware struct {
	client *ethclient.Client
}

func NewAuthMiddleware(client *ethclient.Client) *AuthMiddleware {
	return &AuthMiddleware{
		client: client,
	}
}

func (a *AuthMiddleware) Handle(ctx context.Context, request *proto.Request) (*proto.Request, error) {
	contractAdress := common.HexToAddress("0x30522ae286597dfced8d123a8f560b3b7b094aed")
	authRegistery, err := contracts.NewContracts(contractAdress, a.client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate contract: %v", err)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata in context")
	}

	publicKeys := md["x-public-key"]
	if len(publicKeys) == 0 {
		return nil, fmt.Errorf("missing x-public-key in metadata")
	}

	publicKey := publicKeys[0]
	userAddress := common.HexToAddress(publicKey)

	isAuthorized, err := authRegistery.IsAuthorized(nil, userAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to check authorization: %v", err)
	}

	if !isAuthorized {
		return nil, fmt.Errorf("unauthorized: user %s is not authorized", userAddress.Hex())
	}

	return request, nil
}
