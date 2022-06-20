package responses

type ServerResponse struct {
	Success bool        `json:"success"`
	Msg     interface{} `json:"msg"`
}

func NewServerResponse(msg interface{}) *ServerResponse {
	return &ServerResponse{
		Success: true,
		Msg:     msg,
	}
}
