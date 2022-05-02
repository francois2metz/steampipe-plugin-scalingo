package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin/transform"
)

func tableScalingoLogDrain() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_log_drain",
		Description: "A log drain send logs from an application to a log management service.",
		List: &plugin.ListConfig{
			KeyColumns:        plugin.SingleColumn("app_name"),
			Hydrate:           listLogDrain,
			ShouldIgnoreError: isNotFoundError,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			{Name: "app_name", Type: proto.ColumnType_STRING, Transform: transform.FromQual("app_name"), Description: "Name of the app."},

			{Name: "app_id", Type: proto.ColumnType_STRING, Description: "ID of the application."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "URL to the remote log management service."},
		},
	}
}

func listLogDrain(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	appName := d.KeyColumnQuals["app_name"].GetStringValue()

	logDrains, err := client.LogDrainsList(appName)
	if err != nil {
		return nil, err
	}
	for _, logDrain := range logDrains {
		d.StreamListItem(ctx, logDrain)
	}
	return nil, nil
}
