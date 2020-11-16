package returns

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// ErrorResponse is Error response template
type ErrorResponse struct {
	Message string `json:"reason"`
	Error   error  `json:"-"`
}

type Response struct {
	IsFailure bool
	Message   string
	Result    interface{}
}

// printDebugf behaves like log.Printf only in the debug env
func printDebugf(format string, args ...interface{}) {
	if env := os.Getenv("GO_SERVER_DEBUG"); len(env) != 0 {
		log.Printf("[DEBUG] "+format+"\n", args...)
	}
}

func (r Response) ToJson() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

func (e *ErrorResponse) string() string {
	return fmt.Sprintf("reason: %s, error: %s", e.Message, e.Error.Error())
}

// Respond is response write to ResponseWriter
func respond(w http.ResponseWriter, code int, src interface{}) {
	var body []byte
	var err error
	switch s := src.(type) {
	case []byte:
		if !json.Valid(s) {
			Error(w, http.StatusInternalServerError, err, "invalid json")
			return
		}
		body = s
	case string:
		body = []byte(s)
	case *ErrorResponse, ErrorResponse:
		// avoid infinite loop
		if body, err = json.Marshal(src); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"reason\":\"failed to parse json\"}"))
			return
		}
	default:
		if body, err = json.Marshal(src); err != nil {
			Error(w, http.StatusInternalServerError, err, "failed to parse json")
			return
		}
	}
	w.WriteHeader(code)
	w.Write(body)
}

// Error is wrapped Respond when error response
func Error(w http.ResponseWriter, code int, err interface{}, msgFriendly string) {
	e := struct {
		ErrorFriendly string      `json:"error_friendly"`
		Error         interface{} `json:"error"`
	}{
		ErrorFriendly: msgFriendly,
		Error:         err,
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	respond(w, code, e)
}

// JSON is wrapped Respond when success response
//src is message success or struct/json/interface result
func Ok(w http.ResponseWriter, code int, src interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	respond(w, code, src)
}
