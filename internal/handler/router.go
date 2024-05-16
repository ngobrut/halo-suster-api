package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/ngobrut/halo-suster-api/config"
	"github.com/ngobrut/halo-suster-api/constant"
	"github.com/ngobrut/halo-suster-api/internal/middleware"
	"github.com/ngobrut/halo-suster-api/internal/types/response"
	"github.com/ngobrut/halo-suster-api/internal/usecase"
)

func InitHTTPHandler(cnf config.Config, uc usecase.IFaceUsecase) http.Handler {
	h := Handler{
		uc: uc,
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestLogger)
	r.Use(middleware.Recover)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response.JsonResponse{
			Message: "Error",
			Error: &response.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "please check url",
			},
		})
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.JsonResponse{
			Success: true,
			Message: "Success",
			Data: map[string]interface{}{
				"app-name": "eniqilo-store-api-api",
			},
		})
	})

	r.Route("/v1", func(r chi.Router) {
		r.Route("/user", func(user chi.Router) {
			user.Group(func(manageUser chi.Router) {
				manageUser.Use(middleware.Authorize(cnf.JWTSecret, &constant.StrUserRoleIT))
				manageUser.Get("/", h.GetListUser)
			})

			user.Route("/it", func(it chi.Router) {
				it.Post("/register", h.RegisterIT)
				it.Post("/login", h.LoginIT)

				it.Group(func(profile chi.Router) {
					profile.Use(middleware.Authorize(cnf.JWTSecret, &constant.StrUserRoleIT))
					profile.Get("/profile", h.GetProfileIT)
				})
			})

			user.Route("/nurse", func(nurse chi.Router) {
				nurse.Post("/login", h.LoginNurse)

				nurse.Group(func(profile chi.Router) {
					profile.Use(middleware.Authorize(cnf.JWTSecret, &constant.StrUserRoleNurse))
					profile.Get("/profile", h.GetProfileNurse)
				})

				nurse.Group(func(manageNurse chi.Router) {
					manageNurse.Use(middleware.Authorize(cnf.JWTSecret, &constant.StrUserRoleIT))
					manageNurse.Post("/register", h.CreateNurse)
					manageNurse.Put("/{userId}", h.UpdateNurse)
					manageNurse.Delete("/{userId}", h.DeleteNurse)
					manageNurse.Post("/{userId}/access", h.GrantNurseAccess)
				})
			})
		})

		r.Route("/medical", func(medical chi.Router) {
			medical.Use(middleware.Authorize(cnf.JWTSecret, nil))

			medical.Route("/patient", func(patient chi.Router) {
				patient.Post("/", h.CreatePatient)
				patient.Get("/", h.GetListPatient)
			})

			medical.Route("/record", func(record chi.Router) {
				// todo:
			})
		})

		r.Route("/image", func(image chi.Router) {
			image.Use(middleware.Authorize(cnf.JWTSecret, nil))
			image.Post("/", h.UploadImage)
		})
	})

	return r
}
