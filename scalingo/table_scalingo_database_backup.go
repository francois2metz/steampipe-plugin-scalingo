package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableScalingoDatabaseBackup() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_database_backup",
		Description: "Backup from a database.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.AllColumns([]string{"app_name", "addon_id"}),
			Hydrate:    listDatabaseBackup,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isAddonTokenError,
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"app_name", "addon_id", "id"}),
			Hydrate:    getDatabaseBackup,
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{Name: "app_name", Type: proto.ColumnType_STRING, Transform: transform.FromQual("app_name"), Description: "Name of the app."},
			{Name: "addon_id", Type: proto.ColumnType_STRING, Transform: transform.FromQual("addon_id"), Description: "ID of the addon."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID identifying the backup."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Creation date of the backup."},
			{Name: "started_at", Type: proto.ColumnType_TIMESTAMP, Description: "Start date of the backup."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of backup."},
			{Name: "size", Type: proto.ColumnType_INT, Description: "Size of backup."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Status of the current backup."},
			{Name: "method", Type: proto.ColumnType_STRING, Description: "The method of the backup (periodic or manual)."},
		},
	}
}

func listDatabaseBackup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)

	if err != nil {
		plugin.Logger(ctx).Error("scalingo_database_backup.listDatabaseBackup", "connection_error", err)
		return nil, err
	}
	appName := d.EqualsQuals["app_name"].GetStringValue()
	addon := d.EqualsQuals["addon_id"].GetStringValue()

	backups, err := client.BackupList(ctx, appName, addon)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_database_backup.listDatabaseBackup", err)
		return nil, err
	}

	for _, backup := range backups {
		d.StreamListItem(ctx, backup)
	}
	return nil, nil
}

func getDatabaseBackup(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)

	if err != nil {
		plugin.Logger(ctx).Error("scalingo_database_backup.getDatabaseBackup", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals

	id := quals["id"].GetStringValue()
	appName := quals["app_name"].GetStringValue()
	addon := quals["addon_id"].GetStringValue()

	result, err := client.BackupShow(ctx, appName, addon, id)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_database_backup.getDatabaseBackup", err)
		return nil, err
	}
	return result, nil
}
