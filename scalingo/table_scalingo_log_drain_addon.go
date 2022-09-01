package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableScalingoLogDrainAddon() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_log_drain_addon",
		Description: "A log drain send logs from an addon to a log management service.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.AllColumns([]string{"app_name", "id"}),
			Hydrate:    listLogDrainAddon,
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{Name: "app_name", Type: proto.ColumnType_STRING, Transform: transform.FromQual("app_name"), Description: "Name of the app."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromQual("id"), Description: "ID of the addon."},

			{Name: "app_id", Type: proto.ColumnType_STRING, Description: "ID of the application."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "URL to the remote log management service."},
		},
	}
}

func listLogDrainAddon(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_log_drain_addon.listLogDrainAddon", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	appName := quals["app_name"].GetStringValue()
	id := quals["id"].GetStringValue()

	logDrains, err := client.LogDrainsAddonList(ctx, appName, id)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_log_drain_addon.listLogDrainAddon", err)
		return nil, err
	}
	for _, logDrain := range logDrains {
		d.StreamListItem(ctx, logDrain)
	}
	return nil, nil
}
