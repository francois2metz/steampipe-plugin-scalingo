package scalingo

import (
	"context"
	"gopkg.in/errgo.v1"
	"regexp"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func isAddonTokenError(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
	const addonTokenErrorMsg = "fail to get addon token"
	if errorGo, ok := err.(errgo.Wrapper); ok {
		underlyingError := errorGo.Underlying()
		matched, _ := regexp.MatchString(addonTokenErrorMsg, underlyingError.Error())
		if matched {
			return true
		}
	} else {
		matched, _ := regexp.MatchString(addonTokenErrorMsg, err.Error())
		if matched {
			return true
		}
	}

	return isNotFoundError(ctx, d, h, err)
}
