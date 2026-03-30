package scalingo

import (
	"context"
	"net/http"

	scalingohttp "github.com/Scalingo/go-scalingo/v11/http"
	"github.com/Scalingo/go-utils/errors/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func isNotFoundError(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
	var requestFailed *scalingohttp.RequestFailedError
	if errors.As(err, &requestFailed) {
		errorsCode := []int{http.StatusNotFound, http.StatusUnauthorized}
		for _, code := range errorsCode {
			if code == requestFailed.Code {
				return true
			}
		}
		return false
	}

	return false
}
