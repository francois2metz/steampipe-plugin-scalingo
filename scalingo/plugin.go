package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-scalingo",
		DefaultTransform: transform.FromGo().NullIfZero(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"scalingo_apps":          tableScalingoApps(),
			"scalingo_addons":        tableScalingoAddons(),
			"scalingo_collaborators": tableScalingoCollaborators(),
			"scalingo_app_events":    tableScalingoAppEvents(),
			"scalingo_deployments":   tableScalingoDeployments(),
		},
	}
	return p
}
