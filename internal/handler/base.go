package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/ngobrut/halo-suster-api/constant"
	"github.com/ngobrut/halo-suster-api/internal/custom_error"
	"github.com/ngobrut/halo-suster-api/internal/types/response"
	"github.com/ngobrut/halo-suster-api/internal/usecase"
)

type ValidatorError struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details"`
}

func (e ValidatorError) Error() string {
	return e.Message
}

type Handler struct {
	uc usecase.IFaceUsecase
}

func (h Handler) ValidateStruct(r *http.Request, data interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	err = json.Unmarshal(body, data)
	if err != nil {
		fmt.Println("[error-parse-body]", err.Error())
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "please check your body request",
		})

		return err
	}

	validate := validator.New()
	eng := en.New()
	uni := ut.New(eng, eng)
	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(validate, trans)
	validate.RegisterValidation("nipLen", validateNipLen)

	err = validate.Struct(data)
	if err == nil {
		return nil
	}

	var message string
	var details = make([]string, 0)
	for _, field := range err.(validator.ValidationErrors) {
		message = field.Translate(trans)

		switch field.Tag() {
		case "nipLen":
			message = "NIP should be 13 char"
		}

		details = append(details, message)
	}

	err = ValidatorError{
		Code:    http.StatusBadRequest,
		Message: message,
		Details: details,
	}

	return err
}

func validateNipLen(fl validator.FieldLevel) bool {
	nip := strconv.Itoa(fl.Field().Interface().(int))
	return len(nip) == 13
}

func ValidateNipNurse(n int) bool {
	nip := strconv.Itoa(n)
	if nip[0:3] != "303" {
		return true
	}
	if nip[3:4] != "1" && nip[3:4] != "2" {
		return true
	}
	if nip[4:8] < "2000" || nip[4:8] > "2024" {
		return true
	}
	if nip[8:10] < "01" || nip[8:10] > "12" {
		return true
	}
	if nip[10:13] < "000" || nip[10:13] > "999" {
		return true
	}
	return false
}

func ValidateNipIt(n int) bool {
	nip := strconv.Itoa(n)
	if nip[0:3] != "615" {
		return true
	}
	if nip[3:4] != "1" && nip[3:4] != "2" {
		return true
	}
	if nip[4:8] < "2000" || nip[4:8] > "2024" {
		return true
	}
	if nip[8:10] < "01" || nip[8:10] > "12" {
		return true
	}
	if nip[10:13] < "000" || nip[10:13] > "999" {
		return true
	}
	return false
}

func (h Handler) ResponseOK(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response.JsonResponse{
		Success: true,
		Message: "Success",
		Data:    data,
	})
}

func (h Handler) ResponseError(w http.ResponseWriter, err error) {
	v, isValidationErr := err.(ValidatorError)
	if isValidationErr {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.JsonResponse{
			Message: "ValidationError",
			Error: &response.ErrorResponse{
				Code:    v.Code,
				Message: v.Message,
				Details: v.Details,
			},
		})

		return
	}

	e, isCustomErr := err.(*custom_error.CustomError)
	if !isCustomErr {
		if err != nil && !errors.Is(err, context.Canceled) {
			fmt.Println(err.Error(), "[unhandled-error]")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response.JsonResponse{
			Message: http.StatusText(http.StatusInternalServerError),
			Error: &response.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: constant.HTTPStatusText(http.StatusInternalServerError),
			},
		})

		return
	}

	httpCode := http.StatusInternalServerError
	msg := constant.HTTPStatusText(httpCode)

	if e.ErrorContext != nil && e.ErrorContext.HTTPCode > 0 {
		httpCode = e.ErrorContext.HTTPCode
		msg = constant.HTTPStatusText(httpCode)

		if e.ErrorContext.Message != "" {
			msg = e.ErrorContext.Message
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(response.JsonResponse{
		Message: http.StatusText(httpCode),
		Error: &response.ErrorResponse{
			Code:    httpCode,
			Message: msg,
		},
	})
}
