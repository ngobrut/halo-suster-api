package handler

import "net/http"

func (h Handler) CreateNurse(w http.ResponseWriter, r *http.Request) {
	h.ResponseOK(w, http.StatusCreated, nil)
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
