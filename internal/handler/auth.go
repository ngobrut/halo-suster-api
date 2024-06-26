package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/ngobrut/halo-suster-api/constant"
	"github.com/ngobrut/halo-suster-api/internal/types/request"
	"github.com/ngobrut/halo-suster-api/util"
)

func (h Handler) RegisterIT(w http.ResponseWriter, r *http.Request) {
	var req request.Register
	err := h.ValidateStruct(r, &req)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	res, err := h.uc.Register(r.Context(), &req)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	h.ResponseOK(w, http.StatusCreated, res)
}

func (h Handler) LoginIT(w http.ResponseWriter, r *http.Request) {
	var req request.Login
	err := h.ValidateStruct(r, &req)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	req.UserRole = constant.UserRoleIT

	res, err := h.uc.Login(r.Context(), &req)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	h.ResponseOK(w, http.StatusOK, res)
}

func (h Handler) GetProfileIT(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, err := uuid.Parse(util.GetUserIDFromCtx(ctx))
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	res, err := h.uc.GetProfile(r.Context(), userID)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	h.ResponseOK(w, http.StatusOK, res)
}

func (h Handler) LoginNurse(w http.ResponseWriter, r *http.Request) {
	var req request.LoginNurse
	err := h.ValidateStruct(r, &req)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	req.UserRole = constant.UserRoleNurse

	res, err := h.uc.LoginNurse(r.Context(), &req)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	h.ResponseOK(w, http.StatusOK, res)
}

func (h Handler) GetProfileNurse(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, err := uuid.Parse(util.GetUserIDFromCtx(ctx))
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	res, err := h.uc.GetProfile(r.Context(), userID)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	h.ResponseOK(w, http.StatusCreated, res)
}
