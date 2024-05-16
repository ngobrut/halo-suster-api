package handler

import (
	"net/http"

	"github.com/ngobrut/halo-suster-api/internal/types/request"
)

func (h Handler) CreatePatient(w http.ResponseWriter, r *http.Request) {
	var req request.CreatePatient
	err := h.ValidateStruct(r, &req)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	err = h.uc.CreatePatient(r.Context(), &req)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	h.ResponseOK(w, http.StatusCreated, nil)
}
