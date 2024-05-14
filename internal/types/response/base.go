package response

type JsonResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    interface{}    `json:"data"`
	Error   *ErrorResponse `json:"error"`
}

type ErrorResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
}
