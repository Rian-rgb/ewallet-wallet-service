package external

import (
	"github.com/Rian-rgb/ewallet-common-lib/security"
)

type UmsClient interface {
	ValidateToken(token string) (*security.ClaimToken, error)
}
