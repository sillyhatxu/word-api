package response

type ResponseEntity struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"message"`
}

const SUCCESS string = "SUCCESS"

const ERROR string = "Server error"

const PARAMS_VALIDATE_ERROR string = "system.validate.error"

func Success(data interface{}) *ResponseEntity {
	return &ResponseEntity{Code: SUCCESS, Data: data, Msg: SUCCESS}
}

func Error(data interface{}, msg string) *ResponseEntity {
	return &ResponseEntity{Code: ERROR, Data: data, Msg: msg}
}

func ErrorParamsValidate(data interface{}, msg string) *ResponseEntity {
	return &ResponseEntity{Code: PARAMS_VALIDATE_ERROR, Data: data, Msg: msg}
}
