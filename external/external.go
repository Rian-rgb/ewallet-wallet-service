package external

import (
	"context"
	"ewallet-wallet/external/ums"
)

type UmsClient interface {
	ValidateToken(ctx context.Context, token string) (ums.Token, error)
}
