package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/ngobrut/halo-suster-api/internal/types/request"
	"github.com/ngobrut/halo-suster-api/util"
)

func (h Handler) CreateMedicalRecord(w http.ResponseWriter, r *http.Request) {
	userId, err := uuid.Parse(util.GetUserIDFromCtx(r.Context()))
	if err != nil {
		h.ResponseError(w, err)
		return
	}
	var req request.CreateMedicalRecord
	err = h.ValidateStruct(r, &req)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	req.UserID = userId
	err = h.uc.CreateMedicalRecord(r.Context(), &req)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	h.ResponseOK(w, http.StatusCreated, nil)
}
