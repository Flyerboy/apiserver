package errormsg

type ErrorMsg struct {
	Code int
	Message string
}

func (e ErrorMsg) Error() string {
	return e.Message
}

func DecodeError(e error) (int, string) {
	if e == nil {
		return OK.Code, OK.Message
	}

	switch t := e.(type) {
	case *ErrorMsg:
		return t.Code, t.Message
	}

	return ServerError.Code, ServerError.Message
}
