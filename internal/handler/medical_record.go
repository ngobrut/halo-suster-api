package handler

import (
	"net/http"
	"strconv"

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

func (h Handler) GetListMedicalRecord(w http.ResponseWriter, r *http.Request) {
	qp := r.URL.Query()

	params := &request.ListMedicalRecordQuery{
		IdentityNumber: StringPtr(qp.Get("identityDetail.identityNumber")),
		UserID:         StringPtr(qp.Get("createdBy.userId")),
		NIP:            StringPtr(qp.Get("createdBy.nip")),
		CreatedAt:      StringPtr(qp.Get("createdAt")),
	}

	if limit, err := strconv.Atoi(qp.Get("limit")); err == nil {
		params.Limit = &limit
	}
	if offset, err := strconv.Atoi(qp.Get("offset")); err == nil {
		params.Offset = &offset
	}

	res, err := h.uc.GetListMedicalRecord(r.Context(), params)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	h.ResponseOK(w, http.StatusOK, res)
}
