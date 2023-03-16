package enums

type ServiceStatus int

const (
	SUCCESS           = 200
	ERROR             = 500
	INVALID_PARAMS    = 400
	ERROR_AUTH_FAILED = 10001
	SERVER_ERROR      = 10002

	USER_NAME_DUPLICATED = 10003
	USER_EXIST           = 10004

	EMAIL_VALIDATION_ERROR = 1000
)
