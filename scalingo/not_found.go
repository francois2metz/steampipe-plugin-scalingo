package scalingo

import (
	"net/http"

	scalingohttp "github.com/Scalingo/go-scalingo/v4/http"
	"gopkg.in/errgo.v1"
)

func isNotFoundError(err error) bool {
	errgo, ok := err.(*errgo.Err)
	if !ok {
		return false
	}
	underlyingError := errgo.Underlying()

	requestFailedError, ok := underlyingError.(*scalingohttp.RequestFailedError)
	if !ok {
		return false
	}

	return requestFailedError.Code == http.StatusNotFound
}
