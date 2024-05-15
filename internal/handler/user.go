package handler

import (
	"net/http"
	"strconv"

	"github.com/ngobrut/halo-suster-api/internal/types/request"
)

func (h Handler) GetListUser(w http.ResponseWriter, r *http.Request) {
	qp := r.URL.Query()

	params := &request.ListUserQuery{
		UserID:    StringPtr(qp.Get("userId")),
		Name:      StringPtr(qp.Get("name")),
		NIP:       StringPtr(qp.Get("nip")),
		Role:      StringPtr(qp.Get("role")),
		CreatedAt: StringPtr(qp.Get("createdAt")),
	}

	if limit, err := strconv.Atoi(qp.Get("limit")); err == nil {
		params.Limit = &limit
	}
	if offset, err := strconv.Atoi(qp.Get("offset")); err == nil {
		params.Offset = &offset
	}

	res, err := h.uc.GetListUser(r.Context(), params)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	h.ResponseOK(w, http.StatusOK, res)
}
