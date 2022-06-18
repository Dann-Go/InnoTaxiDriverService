package responses

type ServerError struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type ServerGoodResponse struct {
	Success bool        `json:"success"`
	Msg     interface{} `json:"msg"`
}

func NewServerError(err string) *ServerError {
	return &ServerError{
		Success: false,
		Error:   err,
	}
}

func NewServerGoodResponse(msg interface{}) *ServerGoodResponse {
	return &ServerGoodResponse{
		Success: true,
		Msg:     msg,
	}
}
