package responses

type ServerResponse struct {
	Success bool        `json:"success"`
	Msg     interface{} `json:"msg"`
}

func NewServerResponse(success bool, msg interface{}) *ServerResponse {
	return &ServerResponse{
		Success: success,
		Msg:     msg,
	}
}
