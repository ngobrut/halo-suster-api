package usecase

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/ngobrut/halo-suster-api/constant"
	"github.com/ngobrut/halo-suster-api/internal/custom_error"
	"github.com/ngobrut/halo-suster-api/internal/model"
	"github.com/ngobrut/halo-suster-api/internal/types/request"
	"github.com/ngobrut/halo-suster-api/internal/types/response"
	"github.com/ngobrut/halo-suster-api/util"
)

// Register implements IFaceUsecase.
func (u *Usecase) Register(ctx context.Context, req *request.Register) (*response.AuthResponse, error) {
	pwd, err := util.HashPwd(u.cnf.BcryptSalt, []byte(req.Password))
	if err != nil {
		return nil, err
	}

	user := &model.User{
		NIP:      strconv.Itoa(req.NIP),
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

	nip, err := strconv.Atoi(user.NIP)
	if err != nil {
		return nil, custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusInternalServerError,
			Message:  constant.HTTPStatusText(http.StatusInternalServerError),
		})
	}
	res := &response.AuthResponse{
		UserID:      user.UserID,
		Name:        user.Name,
		NIP:         nip,
		AccessToken: token,
	}

	return res, nil
}

// Login implements IFaceUsecase.
func (u *Usecase) Login(ctx context.Context, req *request.Login) (*response.AuthResponse, error) {
	user, err := u.repo.FindOneUserByNIP(ctx, strconv.Itoa(req.NIP))
	if err != nil {
		return nil, err
	}

	if !user.Password.Valid {
		return nil, custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "user is not having access",
		})
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

	nip, err := strconv.Atoi(user.NIP)
	if err != nil {
		return nil, custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusInternalServerError,
			Message:  constant.HTTPStatusText(http.StatusInternalServerError),
		})
	}

	res := &response.AuthResponse{
		UserID:      user.UserID,
		Name:        user.Name,
		NIP:         nip,
		AccessToken: token,
	}

	return res, nil
}

// GetProfile implements IFaceUsecase.
func (u *Usecase) GetProfile(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	return u.repo.FindOneUserByID(ctx, userID)
}

// Login Nurse implements IFaceUsecase.
func (u *Usecase) LoginNurse(ctx context.Context, req *request.LoginNurse) (*response.AuthResponse, error) {
	user, err := u.repo.FindOneUserByNIP(ctx, strconv.Itoa(req.NIP))
	if err != nil {
		return nil, err
	}

	if !user.Password.Valid {
		return nil, custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "user is not having access",
		})
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

	nip, err := strconv.Atoi(user.NIP)
	if err != nil {
		return nil, custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusInternalServerError,
			Message:  constant.HTTPStatusText(http.StatusInternalServerError),
		})
	}

	res := &response.AuthResponse{
		UserID:      user.UserID,
		Name:        user.Name,
		NIP:         nip,
		AccessToken: token,
	}

	return res, nil
}
