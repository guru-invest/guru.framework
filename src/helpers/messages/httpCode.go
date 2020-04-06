package messages

type httpCode = int

type listCode struct {
	//200
	Ok        httpCode
	Created   httpCode
	Accepted  httpCode
	NoContent httpCode
	//400
	BadRequest     httpCode
	Unauthorized   httpCode
	Forbidden      httpCode
	NotFound       httpCode
	RequestTimeout httpCode
	//500
	InternalServerError httpCode
	BadGateway          httpCode
	ServiceUnavailable  httpCode
}

var HttpCode = &listCode{
	Ok:        200,
	Created:   201,
	Accepted:  202,
	NoContent: 204,

	BadRequest:     400,
	Unauthorized:   401,
	Forbidden:      403,
	NotFound:       404,
	RequestTimeout: 408,

	InternalServerError: 500,
	BadGateway:          502,
	ServiceUnavailable:  503,
}
