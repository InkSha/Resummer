package network

type Response struct {
	Code ResponseStatus `json:"code" example:"200"`
	Msg  string         `json:"msg" example:"Success"`
	Data any            `json:"data,omitempty"`
}

func Success(data any) Response {
	return Response{
		Code: ResponseStatusSuccess,
		Msg:  "Success",
		Data: data,
	}
}

func Failed(msg string, data any) Response {
	return Response{
		Code: ResponseStatusFailed,
		Msg:  msg,
		Data: data,
	}
}
