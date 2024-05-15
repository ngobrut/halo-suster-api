package handler

import (
	"net/http"

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
	h.ResponseOK(w, http.StatusOK, nil)
}

func (h Handler) DeleteNurse(w http.ResponseWriter, r *http.Request) {
	h.ResponseOK(w, http.StatusOK, nil)
}

func (h Handler) GrantNurseAccess(w http.ResponseWriter, r *http.Request) {
	h.ResponseOK(w, http.StatusOK, nil)
}
