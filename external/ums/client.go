package ums

import (
	"context"
	internalErrors "ewallet-wallet/internal/errors"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"github.com/Rian-rgb/ewallet-common-lib/response"
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

func (c *Client) ValidateToken(ctx context.Context, token string) (Token, error) {
	var resp Token

	rpcCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	result, err := c.client.ValidateToken(rpcCtx, &pb.ValidateTokenRequest{Token: token})
	if err != nil {
		return resp, err
	}

	if result.GetMessage() != response.SuccessMessage {
		logger.WithContext(ctx).Error("response message is not success: ", result.GetMessage())
		return resp, internalErrors.ErrInternalServerError
	}

	resp.UserID = result.GetData().GetUserId()
	resp.Username = result.GetData().GetUsername()
	resp.FullName = result.GetData().GetFullName()

	return resp, nil
}
