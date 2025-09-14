package routes

// ResponseHTTP represents response body of this API
type ResponseHTTP struct {
	Success bool        `json:"success"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// NewResponseHTTP creates a new response body
func NewResponseHTTP(success bool, message string, data interface{}) *ResponseHTTP {
	return &ResponseHTTP{
		Success: success,
		Message: message,
		Data:    data,
	}
}
