package handler

import (
	"net/http"

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

	if ValidateNipNurse(req.NIP) {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  "user is not found / user is not from Nurse (nip not starts with 303)",
		})
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
	h.ResponseOK(w, http.StatusOK, nil)
}

func (h Handler) DeleteNurse(w http.ResponseWriter, r *http.Request) {
	h.ResponseOK(w, http.StatusOK, nil)
}

func (h Handler) GrantNurseAccess(w http.ResponseWriter, r *http.Request) {
	h.ResponseOK(w, http.StatusOK, nil)
}
