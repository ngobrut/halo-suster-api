package usecase

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/ngobrut/halo-sus-api/constant"
	"github.com/ngobrut/halo-sus-api/internal/custom_error"
	"github.com/ngobrut/halo-sus-api/internal/model"
	"github.com/ngobrut/halo-sus-api/internal/types/request"
	"github.com/ngobrut/halo-sus-api/internal/types/response"
	"github.com/ngobrut/halo-sus-api/util"
)

// Register implements IFaceUsecase.
func (u *Usecase) Register(ctx context.Context, req *request.Register) (*response.AuthResponse, error) {
	pwd, err := util.HashPwd(u.cnf.BcryptSalt, []byte(req.Password))
	if err != nil {
		return nil, err
	}

	user := &model.User{
		NIP:      req.NIP,
		Name:     req.Name,
		Password: sql.NullString{String: pwd, Valid: true},
		Role:     constant.UserRoleIT,
	}

	err = u.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	claims := &util.CustomClaims{
		UserID: user.UserID.String(),
		Role:   string(user.Role),
	}

	token, err := util.GenerateAccessToken(claims, u.cnf.JWTSecret)
	if err != nil {
		return nil, err
	}

	res := &response.AuthResponse{
		UserID:      user.UserID,
		Name:        user.Name,
		NIP:         user.NIP,
		AccessToken: token,
	}

	return res, nil
}

// Login implements IFaceUsecase.
func (u *Usecase) Login(ctx context.Context, req *request.Login) (*response.AuthResponse, error) {
	user, err := u.repo.FindOneUserByNIP(ctx, req.NIP)
	if err != nil {
		return nil, err
	}

	if user.Password.String == "" || (req.UserRole != user.Role) {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusUnauthorized,
			Message:  "you don't have the right to access this api",
		})

		return nil, err
	}

	err = util.ComparePwd([]byte(user.Password.String), []byte(req.Password))
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "wrong password",
		})

		return nil, err
	}

	claims := &util.CustomClaims{
		UserID: user.UserID.String(),
		Role:   string(user.Role),
	}

	token, err := util.GenerateAccessToken(claims, u.cnf.JWTSecret)
	if err != nil {
		return nil, err
	}

	res := &response.AuthResponse{
		UserID:      user.UserID,
		Name:        user.Name,
		NIP:         user.NIP,
		AccessToken: token,
	}

	return res, nil
}

// GetProfile implements IFaceUsecase.
func (u *Usecase) GetProfile(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	return u.repo.FindOneUserByID(ctx, userID)
}
