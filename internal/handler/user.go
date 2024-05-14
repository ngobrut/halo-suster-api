package handler

import "net/http"

func (h Handler) GetListUser(w http.ResponseWriter, r *http.Request) {
	h.ResponseOK(w, http.StatusCreated, nil)
}
