package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ngobrut/halo-suster-api/constant"
	"github.com/ngobrut/halo-suster-api/internal/types/response"
	"github.com/ngobrut/halo-suster-api/util"
)

func Authorize(secret string, userRole *constant.UserRole) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			token, err := GetTokenFromHeader(r)
			if err != nil {
				UnauthorizedError(w)
				return
			}

			res, err := jwt.ParseWithClaims(token, &util.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})

			if err != nil {
				UnauthorizedError(w)
				return
			}

			claims, ok := res.Claims.(*util.CustomClaims)
			if !ok && !res.Valid {
				UnauthorizedError(w)
				return
			}

			if userRole != nil {
				if claims.Role != string(*userRole) {
					UnauthorizedError(w)
					return
				}
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, constant.UserIDKey, claims.UserID)
			ctx = context.WithValue(ctx, constant.RoleKey, claims.Role)
			h.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}

func UnauthorizedError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(response.JsonResponse{
		Message: "Error",
		Error: &response.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: constant.HTTPStatusText(http.StatusUnauthorized),
		},
	})
}

func GetTokenFromHeader(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", errors.New("token is empty")
	}

	token := strings.Split(header, " ")
	if len(token) < 2 {
		return "", errors.New("token is invalid")
	}

	return token[1], nil
}

func ParseWithoutVerified(token string) *util.CustomClaims {
	res, _, err := new(jwt.Parser).ParseUnverified(token, &util.CustomClaims{})
	if err != nil {
		return nil
	}

	claims, ok := res.Claims.(*util.CustomClaims)
	if ok && claims.ID != "" {
		return claims
	}

	return nil
}
