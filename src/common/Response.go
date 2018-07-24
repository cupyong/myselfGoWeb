package common

import (
	"net/http"
	"fmt"
)

type ResponseWriter interface {
	Header() http.Header

	WriteJson(v interface{}) error

	WriteErrorJson(code int) error

	EncodeJson(v interface{}) ([]byte, error)

	WriteHeader(int)
}

var ErrorFieldName = "Error"

func Error(w ResponseWriter, error string, code int) {
	w.WriteHeader(code)
	fmt.Println(code)
	err := w.WriteJson(map[string]string{ErrorFieldName: error})
	if err != nil {
		panic(err)
	}
}
func NotFound(w ResponseWriter, r *Request) {
	Error(w, "Resource not found", http.StatusNotFound)
}

func NotPermission(w ResponseWriter, r *Request) {
	Error(w, "not Permission", http.StatusForbidden)
}



