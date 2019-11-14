package helpers

const (
	Status_Ok					 = 200
	Status_Created				 = 201
	Status_Accepted				 = 202
	Status_NoContent			 = 204

	Status_BadRequest			 = 400
	Status_Unauthorized			 = 401
	Status_Forbidden			 = 403
	Status_NotFound				 = 404
	Status_RequestTimeout		 = 408

	Status_InternalServerError	 = 500
	Status_BadGateway			 = 502
	Status_ServiceUnavailable	 = 503
)

var statusText = map[int]string{
	Status_Ok:				"OK",
	Status_Created:			"Created",
	Status_Accepted:		"Accepted",
	Status_NoContent:		"No Content",

	Status_BadRequest:		"Bad Request",
	Status_Unauthorized:	"Unauthorized",
	Status_Forbidden:		"Forbidden",
	Status_NotFound:		"Not Found",
	Status_RequestTimeout:	"Request Timeout",

	Status_InternalServerError:		"Internal Server Error",
	Status_BadGateway:				"Bad Gateway",
	Status_ServiceUnavailable:		"Service Unavailable",
}

func StatusDescription(code int) string {
	return statusText[code]
}