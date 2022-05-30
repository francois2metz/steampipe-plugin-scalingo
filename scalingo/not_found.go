package scalingo

import (
	"context"
	"net/http"

	scalingohttp "github.com/Scalingo/go-scalingo/v4/http"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"gopkg.in/errgo.v1"
)

func isNotFoundError(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
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
