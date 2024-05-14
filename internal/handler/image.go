package handler

import (
	"errors"
	"net/http"

	"github.com/ngobrut/halo-suster-api/constant"
	"github.com/ngobrut/halo-suster-api/internal/custom_error"
)

func (h Handler) UploadImage(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, constant.FILE_UPLOAD_MAX_SIZE)
	if err := r.ParseMultipartForm(constant.FILE_UPLOAD_MAX_SIZE); err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "ukuran file melebihi 2MB",
		})

		h.ResponseError(w, err)
		return
	}

	_, header, err := r.FormFile("file")
	if err != nil {
		if errors.Is(err, http.ErrMissingFile) {
			err = custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusBadRequest,
				Message:  "file tidak boleh kosong",
			})

			h.ResponseError(w, err)
			return
		}

		h.ResponseError(w, err)
		return
	}

	allowedext := map[string]bool{
		"image/jpg":  true,
		"image/jpeg": true,
	}

	if !allowedext[header.Header.Get("Content-Type")] {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "file must be in `.jpg` or `.jpeg` format",
		})

		h.ResponseError(w, err)
		return
	}

	if header.Size/1000 < 10 {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "file size must be at least 10KB",
		})

		h.ResponseError(w, err)
		return
	}

	res, err := h.uc.UploadImage(r.Context(), header)
	if err != nil {
		h.ResponseError(w, err)
		return
	}

	h.ResponseOK(w, http.StatusOK, res)
}
