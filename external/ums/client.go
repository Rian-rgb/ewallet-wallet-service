package ums

import (
	"context"
	internalErrors "ewallet-wallet/internal/errors"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"github.com/Rian-rgb/ewallet-common-lib/response"
	"github.com/Rian-rgb/ewallet-common-lib/security"
	pb "github.com/Rian-rgb/ewallet-proto/gen/token_validation/v1"
	"time"
)

type Client struct {
	client  pb.TokenValidationServiceClient
	timeout time.Duration
}

func NewClient(client pb.TokenValidationServiceClient) *Client {
	return &Client{
		client:  client,
		timeout: 5 * time.Second,
	}
}

func (c *Client) ValidateToken(token string) (*security.ClaimToken, error) {
	ctx := context.Background()

	rpcCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	result, err := c.client.ValidateToken(rpcCtx, &pb.ValidateTokenRequest{Token: token})
	if err != nil {
		return nil, err
	}

	if result.GetMessage() != response.SuccessMessage {
		logger.WithContext(ctx).Error("response message is not success: ", result.GetMessage())
		return nil, internalErrors.ErrInternalServerError
	}

	dataResult := result.GetData()
	resp := &security.ClaimToken{
		UserID:   int(dataResult.GetUserId()),
		Username: dataResult.GetUsername(),
		FullName: dataResult.GetFullName(),
	}

	return resp, nil
}
