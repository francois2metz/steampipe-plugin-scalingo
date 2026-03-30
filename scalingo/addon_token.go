package scalingo

import (
	"context"
	"regexp"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func isAddonTokenError(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
	const addonTokenErrorMsg = "get addon token"
	matched, _ := regexp.MatchString(addonTokenErrorMsg, err.Error())
	if matched {
		return true
	}

	return isNotFoundError(ctx, d, h, err)
}
