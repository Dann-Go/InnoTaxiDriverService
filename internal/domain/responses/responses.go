package responses

type ServerInternalError struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type ServerBadRequestError struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type ServerGoodResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

type ServerUnauthorizedResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

func NewServerInternalError(err string) *ServerInternalError {
	return &ServerInternalError{
		Success: false,
		Error:   err,
	}
}

func NewServerBadRequestError(err string) *ServerBadRequestError {
	return &ServerBadRequestError{
		Success: false,
		Error:   err,
	}
}

func NewServerGoodResponse(msg string) *ServerGoodResponse {
	return &ServerGoodResponse{
		Success: true,
		Msg:     msg,
	}
}
func NewServerUnauthorizedResponse(msg string) *ServerUnauthorizedResponse {
	return &ServerUnauthorizedResponse{
		Success: false,
		Msg:     msg,
	}
}
