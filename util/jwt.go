package util

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ngobrut/halo-suster-api/constant"
)

type CustomClaims struct {
	UserID string
	Role   string
	jwt.RegisteredClaims
}

const (
	JWT_TTL time.Duration = 8 * time.Hour
)

func GenerateAccessToken(claims *CustomClaims, secret string) (string, error) {
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(JWT_TTL)),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString([]byte(secret))
}

func GetUserIDFromCtx(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	if userID, ok := ctx.Value(constant.UserIDKey).(string); ok {
		return userID
	}

	return ""
}

func GetUserRoleFromCtx(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	if userRole, ok := ctx.Value(constant.RoleKey).(string); ok {
		return userRole
	}

	return ""
}
