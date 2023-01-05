package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableScalingoContainerType() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_container_type",
		Description: "Container type from an application.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("app_name"),
			Hydrate:    listContainerType,
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{Name: "app_name", Type: proto.ColumnType_STRING, Transform: transform.FromQual("app_name"), Description: "Name of the app."},

			{Name: "app_id", Type: proto.ColumnType_STRING, Description: "ID of the application."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Type of container (web, worker, etc.)."},
			{Name: "amount", Type: proto.ColumnType_INT, Description: "Amount of containers of the given type."},
			{Name: "command", Type: proto.ColumnType_STRING, Description: "Command used to run the container."},
			{Name: "size", Type: proto.ColumnType_STRING, Description: "Size of the containers of this type (S/M/XL/..)."},
		},
	}
}

func listContainerType(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_container_type.listContainerType", "connection_error", err)
		return nil, err
	}
	appName := d.EqualsQuals["app_name"].GetStringValue()

	containers, err := client.AppsContainerTypes(ctx, appName)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_container_type.listContainerType", err)
		return nil, err
	}
	for _, container := range containers {
		d.StreamListItem(ctx, container)
	}
	return nil, nil
}
