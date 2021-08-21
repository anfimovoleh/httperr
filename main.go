package httperr

import (
	"fmt"
	"net/http"
)

const responseErrorFormat = `{"code": %d, "error": "%s"}`

func ErrResponse(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	_, _ = w.Write([]byte(fmt.Sprintf(responseErrorFormat, code, err)))
}

func InternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(fmt.Sprintf(responseErrorFormat, http.StatusInternalServerError, ErrInternalServer)))
}

func BadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte(fmt.Sprintf(responseErrorFormat, http.StatusBadRequest, err)))
}

func NotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte(fmt.Sprintf(responseErrorFormat, http.StatusNotFound, ErrNotFound)))
}