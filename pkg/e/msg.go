package e

var MsgFlags = map[int]string{
	SUCCESS:                        "Ok",
	ERROR:                          "Fail",
	INVALID_PARAMS:                 "Invalid params",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token fail",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token timeout",
	ERROR_AUTH_TOKEN:               "Token not success",
	ERROR_AUTH:                     "Token wrong",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
