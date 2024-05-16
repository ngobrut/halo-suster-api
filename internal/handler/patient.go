package handler

import (
	"net/http"
	"strconv"

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

func (h Handler) GetListPatient(w http.ResponseWriter, r *http.Request) {
	qp := r.URL.Query()

	params := &request.ListPatientQuery{
		Name:           StringPtr(qp.Get("name")),
		Phone:          StringPtr(qp.Get("phone")),
		CreatedAt:      StringPtr(qp.Get("createdAt")),
		IdentityNumber: StringPtr(qp.Get("identityNumber")),
	}

	if limit, err := strconv.Atoi(qp.Get("limit")); err == nil {
		params.Limit = &limit
	}
	if offset, err := strconv.Atoi(qp.Get("offset")); err == nil {
		params.Offset = &offset
	}

	res, err := h.uc.GetListPatient(r.Context(), params)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	h.ResponseOK(w, http.StatusOK, res)
}
