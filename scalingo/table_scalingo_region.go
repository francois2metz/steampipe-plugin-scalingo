package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableScalingoRegion() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_region",
		Description: "A region is a location in which an application can be deployed.",
		List: &plugin.ListConfig{
			Hydrate: listRegion,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Underscore-cased name of the region."},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "How the name of the region should be displayed."},
			{Name: "ssh", Type: proto.ColumnType_STRING, Transform: transform.FromField("SSH"), Description: "SSH Host to git push application code."},
			{Name: "api", Type: proto.ColumnType_STRING, Transform: transform.FromField("API"), Description: "URL to the regional API managing apps."},
			{Name: "dashboard", Type: proto.ColumnType_STRING, Description: "URL to the dashboard of the region."},
			{Name: "database_api", Type: proto.ColumnType_STRING, Transform: transform.FromField("DatabaseAPI"), Description: "URL to the regional API managing databases."},
			{Name: "default", Type: proto.ColumnType_BOOL, Description: "Is the region is the default region."},
		},
	}
}

func listRegion(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_region.listRegion", "connection_error", err)
		return nil, err
	}
	regions, err := client.RegionsList(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_region.listRegion", err)
		return nil, err
	}
	for _, region := range regions {
		d.StreamListItem(ctx, region)
	}

	return nil, nil
}
