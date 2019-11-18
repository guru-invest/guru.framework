package messages

const (
	ERR_MISSIG_KEY = "Missing key:'{{ .Param}}'"
	ERR_NOT_FOUND = "'{{ .Param}}' not found"
	ERR_INVALID_FORMAT = "Invalid format on '{{ .Param}}'"
	ERR_CONNECT_TO = "Error on connect to '{{ .Param}}'"
	ERR_DISCONNECT_FROM = "Error on disconnect from '{{ .Param}}'"
	ERR_UNAUTHORIZED = "You are not authorized to perform this action"
	ERR_ON_EXECUTE = "Error on execute '{{ .Param}}' action"
)
