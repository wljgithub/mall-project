package errno

import "net/http"

var (
	OK      = &ErrorNo{HTTPStatusCode: http.StatusOK, ServiceCode: 0, Message: "ok"}
	ErrBind = &ErrorNo{HTTPStatusCode: http.StatusBadRequest, ServiceCode: 10001, Message: "参数绑定错误"}

	// user error
	ErrIncorrectPassword   = &ErrorNo{http.StatusUnauthorized, 40001, "密码错误"}
	ErrUserExist           = &ErrorNo{http.StatusForbidden, 40002, "用户已存在"}
	ErrTokenInvalid        = &ErrorNo{http.StatusUnauthorized, 40003, "token无效"}
	ErrTokenExpire         = &ErrorNo{http.StatusUnauthorized, 40004, "token超时"}
	ErrInternalServerError = &ErrorNo{HTTPStatusCode: http.StatusInternalServerError, ServiceCode: 50001, Message: "internal server error"}
)
