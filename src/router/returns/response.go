package returns

type Response struct {
	IsFailure bool
	Messages  []string
	Result    interface{}
}

func (r Response) ToResponse(messages []string, result interface{}) Response {
	isFailure := false
	if messages != nil {
		isFailure = true
	}

	r.IsFailure = isFailure
	r.Messages = messages
	r.Result = result
	return r
}
