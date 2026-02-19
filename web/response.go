package web

type Response struct {
	Ok    bool  `json:"ok"`
	Data  any   `json:"data"`
	Error error `json:"error"`
}

func NewResponse(data any, err error) Response {
	if err != nil {
		return Response{
			Ok:    false,
			Data:  nil,
			Error: err,
		}
	}

	return Response{
		Ok:    true,
		Data:  &data,
		Error: nil,
	}
}

func NewOkResponse(data any) Response {
	return Response{
		Ok:    true,
		Data:  data,
		Error: nil,
	}
}

func NewErrResponse(err error) Response {
	return Response{
		Ok:    false,
		Data:  nil,
		Error: err,
	}
}
