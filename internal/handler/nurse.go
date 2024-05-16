package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/ngobrut/halo-suster-api/constant"
	"github.com/ngobrut/halo-suster-api/internal/custom_error"
	"github.com/ngobrut/halo-suster-api/internal/types/request"
)

func (h Handler) CreateNurse(w http.ResponseWriter, r *http.Request) {
	var req request.CreateNurse
	err := h.ValidateStruct(r, &req)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	res, err := h.uc.CreateNurse(r.Context(), &req)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	h.ResponseOK(w, http.StatusCreated, res)
}

func (h Handler) UpdateNurse(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(r.PathValue("userId"))
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  constant.HTTPStatusText(http.StatusNotFound),
		})
		h.ResponseError(w, err)
		return
	}

	var req request.UpdateNurse
	err = h.ValidateStruct(r, &req)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	req.UserID = userID
	req.Role = constant.StrUserRoleNurse

	err = h.uc.UpdateNurse(r.Context(), &req)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	h.ResponseOK(w, http.StatusOK, nil)
}

func (h Handler) DeleteNurse(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(r.PathValue("userId"))
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "userId is not a nurse (nip not starts with 303)",
		})
		h.ResponseError(w, err)
		return
	}
	err = h.uc.DeleteNurse(r.Context(), userID)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	h.ResponseOK(w, http.StatusOK, nil)
}

func (h Handler) GrantNurseAccess(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(r.PathValue("userId"))
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "userId is not a nurse (nip not starts with 303)",
		})
		h.ResponseError(w, err)
		return
	}

	var req request.GrantNurseAccess
	err = h.ValidateStruct(r, &req)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	req.UserID = userID
	req.Role = constant.StrUserRoleNurse

	err = h.uc.GrantNurseAccess(r.Context(), &req)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	h.ResponseOK(w, http.StatusOK, nil)
}
