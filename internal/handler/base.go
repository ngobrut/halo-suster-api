package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

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
	validate.RegisterValidation("nipIt", validateNipIt)
	validate.RegisterValidation("nipNurse", validateNipNurse)
	validate.RegisterValidation("validUrl", validateURL)
	validate.RegisterValidation("idNum", validateIDNum)
	validate.RegisterValidation("phone", validatePhone)
	validate.RegisterValidation("dateFormat", validateDate)
	validate.RegisterValidation("gender", validateGender)

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
		case "nipIt":
			message = "NIP should be it Format"
		case "nipNurse":
			message = "NIP should be nurse format"
		case "validUrl":
			message = "identityCardScanImg should be url"
		case "idNum":
			message = "identityNumber should be 16 digit"
		case "phone":
			message = "phone should be starts with `+62`"
		case "dateFormat":
			message = "should be string with ISO 8601 format"
		case "gender":
			message = "gender should be enum of 'male'|'female'"
		}

		details = append(details, message)
	}

	err = ValidatorError{
		Code:    http.StatusBadRequest,
		Message: "request doesn’t pass validation",
		Details: details,
	}

	return err
}

func validateURL(fl validator.FieldLevel) bool {
	parsedURL, err := url.Parse(fl.Field().String())
	if err != nil {
		return false
	}

	// Check if the scheme is present and it's http or https
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return false
	}

	// Check if the host is present and it has a valid format
	if parsedURL.Host == "" {
		return false
	}

	// Check if the host has a valid domain format
	parts := strings.Split(parsedURL.Host, ".")
	if len(parts) < 2 {
		return false
	}

	// Check if the path, if present, is in a valid format
	if parsedURL.Path != "" && !strings.HasPrefix(parsedURL.Path, "/") {
		return false
	}

	// All checks passed, URL is valid
	return true
}

func validateIDNum(fl validator.FieldLevel) bool {
	idNum := strconv.Itoa(fl.Field().Interface().(int))
	return len(idNum) == 16
}

func validatePhone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	return phone[0:3] == "+62"
}

func validateGender(fl validator.FieldLevel) bool {
	return constant.ValidGender[fl.Field().String()]
}

func validateDate(fl validator.FieldLevel) bool {
	date := fl.Field().String()

	for _, format := range constant.DateTime {
		if _, err := time.Parse(format, date); err == nil {
			return true
		}
	}
	return false
}

func validateNipLen(fl validator.FieldLevel) bool {
	nip := strconv.Itoa(fl.Field().Interface().(int))
	return len(nip) >= 13 && len(nip) <= 15
}

func validateNipNurse(fl validator.FieldLevel) bool {
	nip := strconv.Itoa(fl.Field().Interface().(int))
	if len(nip) < 13 || len(nip) > 15 {
		return false
	}
	if nip[0:3] != "303" {
		return false
	}
	if nip[3:4] != "1" && nip[3:4] != "2" {
		return false
	}

	yearStr := nip[4:8]
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return false
	}
	currentYear := time.Now().Year()
	if year < 2000 || year > currentYear {
		return false
	}

	monthStr := nip[8:10]
	month, err := strconv.Atoi(monthStr)
	if err != nil {
		return false
	}
	if month < 1 || month > 12 {
		return false
	}

	randomDigits := nip[10:]
	randomDigitsPattern := `^\d{3,5}$`
	match, _ := regexp.MatchString(randomDigitsPattern, randomDigits)
	return match
}

func validateNipIt(fl validator.FieldLevel) bool {
	nip := strconv.Itoa(fl.Field().Interface().(int))
	if len(nip) < 13 || len(nip) > 15 {
		return false
	}
	if nip[0:3] != "615" {
		return false
	}
	if nip[3:4] != "1" && nip[3:4] != "2" {
		return false
	}

	yearStr := nip[4:8]
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return false
	}
	currentYear := time.Now().Year()
	if year < 2000 || year > currentYear {
		return false
	}

	monthStr := nip[8:10]
	month, err := strconv.Atoi(monthStr)
	if err != nil {
		return false
	}
	if month < 1 || month > 12 {
		return false
	}

	randomDigits := nip[10:]
	randomDigitsPattern := `^\d{3,5}$`
	match, _ := regexp.MatchString(randomDigitsPattern, randomDigits)
	return match

}

func StringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
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
