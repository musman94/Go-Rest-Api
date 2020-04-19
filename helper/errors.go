package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//HttpError Struct defining the error response of a request
type HttpError struct {
	Code int
	Msg  string
}

func (httpError *HttpError) Error() string {
	return fmt.Sprintf("%v %v", httpError.Code, httpError.Msg)
}

//RespondWithError helper function to return an error
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJson(w, code, map[string]string{"error": msg})
}

//RespondWithJson helper function to return the result of an http
func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
