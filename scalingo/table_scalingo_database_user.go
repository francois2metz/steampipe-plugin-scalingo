package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableScalingoDatabaseUser() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_database_user",
		Description: "Users from a database.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.AllColumns([]string{"app_name", "addon_id"}),
			Hydrate:    listDatabaseUser,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isAddonTokenError,
			},
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{Name: "app_name", Type: proto.ColumnType_STRING, Transform: transform.FromQual("app_name"), Description: "Name of the app."},
			{Name: "addon_id", Type: proto.ColumnType_STRING, Transform: transform.FromQual("addon_id"), Description: "ID of the addon."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the user."},
			{Name: "read_only", Type: proto.ColumnType_BOOL, Description: "True if the user is readonly."},
			{Name: "protected", Type: proto.ColumnType_BOOL, Description: "True if the user is protected."},
		},
	}
}

func listDatabaseUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)

	if err != nil {
		plugin.Logger(ctx).Error("scalingo_database_user.listDatabaseUser", "connection_error", err)
		return nil, err
	}
	appName := d.EqualsQuals["app_name"].GetStringValue()
	addon := d.EqualsQuals["addon_id"].GetStringValue()

	backups, err := client.DatabaseListUsers(ctx, appName, addon)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_database_user.listDatabaseUser", err)
		return nil, err
	}

	for _, backup := range backups {
		d.StreamListItem(ctx, backup)
	}
	return nil, nil
}
