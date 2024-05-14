package constant

import "net/http"

type JwtKey string

const (
	UserIDKey JwtKey = "user_id"
	RoleKey   JwtKey = "role"
)

func HTTPStatusText(code int) string {
	switch code {
	case http.StatusInternalServerError:
		return "something went wrong with our side, please try again"
	case http.StatusNotFound:
		return "data not found"
	case http.StatusUnauthorized:
		return "you are not authorized to access this api"
	case http.StatusConflict:
		return "duplicated data error"
	case http.StatusUnprocessableEntity:
		return "please check your body request"
	case http.StatusBadRequest:
		return "request doesn't pass validation"
	default:
		return ""
	}
}
