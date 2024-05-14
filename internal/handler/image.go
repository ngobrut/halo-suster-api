package handler

import "net/http"

func (h Handler) UploadImage(w http.ResponseWriter, r *http.Request) {
	h.ResponseOK(w, http.StatusOK, nil)
}
