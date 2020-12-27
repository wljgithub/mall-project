package errno

type ErrorNo struct {
	HTTPStatusCode int
	ServiceCode    int
	Message        string
}

func (this *ErrorNo) Error() string {
	return this.Message
}

func DecodeErr(err error) (int, int, string) {
	if err == nil {
		return OK.HTTPStatusCode, OK.ServiceCode, OK.Message
	}
	switch typed := err.(type) {
	case *ErrorNo:
		return typed.HTTPStatusCode, typed.ServiceCode, typed.Message
	}

	return ErrInternalServerError.HTTPStatusCode, ErrInternalServerError.ServiceCode, err.Error()
}
