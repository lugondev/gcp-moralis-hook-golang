package jwt

import (
	"context"
	"moralis-webhook/dto"
)

type Validator interface {
	ValidateToken(ctx context.Context, token string) (*dto.UserClaims, error)
}
