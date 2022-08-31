package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableScalingoEnvironment() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_environment",
		Description: "An environment variable is used to configure your app.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("app_name"),
			Hydrate:    listEnvironment,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			{Name: "app_name", Type: proto.ColumnType_STRING, Transform: transform.FromQual("app_name"), Description: "Name of the app."},

			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID of the variable."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the variable."},
			{Name: "value", Type: proto.ColumnType_STRING, Description: "Value of the variable."},
		},
	}
}

func listEnvironment(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_environment.listEnvironment", "connection_error", err)
		return nil, err
	}
	appName := d.KeyColumnQuals["app_name"].GetStringValue()

	variables, err := client.VariablesList(ctx, appName)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_environment.listEnvironment", err)
		return nil, err
	}
	for _, variable := range variables {
		d.StreamListItem(ctx, variable)
	}
	return nil, nil
}
