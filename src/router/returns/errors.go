package returns

import "github.com/guru-invest/guru.framework/src/helpers/messages"

func CreateReturningError(error string) map[string]interface{} {
	m := make(map[string]interface{})
	m["error"] = error
	return m
}

func buildMessage(message string, param string) string {
	return messages.ParseTemplate(message, param)
}

func MissingKeyError(param string) map[string]interface{} {
	return CreateReturningError(buildMessage(messages.ERR_MISSIG_KEY, param))
}

func NotFoundError(param string) map[string]interface{} {
	return CreateReturningError(buildMessage(messages.ERR_NOT_FOUND, param))
}

func InvalidFormatError(param string) map[string]interface{} {
	return CreateReturningError(buildMessage(messages.ERR_INVALID_FORMAT, param))
}

func UnauthorizedError() map[string]interface{} {
	return CreateReturningError(buildMessage(messages.ERR_UNAUTHORIZED, ""))
}
