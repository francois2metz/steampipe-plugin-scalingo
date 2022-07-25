package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableScalingoDatabaseTypeVersion() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_database_type_version",
		Description: "A database type is a version of a database.",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"app_name", "addon_id", "id"}),
			Hydrate:    getDatabaseTypeVersion,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isTokenError,
			},
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			{Name: "app_name", Type: proto.ColumnType_STRING, Transform: transform.FromQual("app_name"), Description: "Name of the app."},
			{Name: "addon_id", Type: proto.ColumnType_STRING, Transform: transform.FromQual("addon_id"), Description: "ID of the addon."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID identifying the database type version."},
			{Name: "database_type_id", Type: proto.ColumnType_STRING, Description: "Unique ID identifying the database type (e.g. PostgreSQL, MySQL, etc)."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the database type version was created."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the database type version was updated."},
			{Name: "features", Type: proto.ColumnType_JSON, Description: "List of available features for this version."},
			{Name: "next_upgrade_id", Transform: transform.FromField("NextUpgrade.ID"), Type: proto.ColumnType_STRING, Description: "Next upgrade ID identifying the database type version."},
			{Name: "next_upgrade_database_type_id", Transform: transform.FromField("NextUpgrade.DatabaseTypeID"), Type: proto.ColumnType_STRING, Description: "Next upgrade unique ID identifying the database type (e.g. PostgreSQL, MySQL, etc)."},
			{Name: "next_upgrade_created_at", Transform: transform.FromField("NextUpgrade.CreatedAt"), Type: proto.ColumnType_TIMESTAMP, Description: "Next upgrade creation date."},
			{Name: "next_upgrade_updated_at", Transform: transform.FromField("NextUpgrade.UpdatedAt"), Type: proto.ColumnType_TIMESTAMP, Description: "Next upgrade update date."},
			{Name: "next_upgrade_major", Transform: transform.FromField("NextUpgrade.Major"), Type: proto.ColumnType_INT, Description: "Next upgrade major version number."},
			{Name: "next_upgrade_minor", Transform: transform.FromField("NextUpgrade.Minor"), Type: proto.ColumnType_INT, Description: "Next upgrade minor version number."},
			{Name: "next_upgrade_patch", Transform: transform.FromField("NextUpgrade.Patch"), Type: proto.ColumnType_INT, Description: "Next upgrade patch version number."},
			{Name: "next_upgrade_build", Transform: transform.FromField("NextUpgrade.Build"), Type: proto.ColumnType_INT, Description: "Next upgrade build version number."},
			{Name: "major", Type: proto.ColumnType_INT, Description: "Major version number."},
			{Name: "minor", Type: proto.ColumnType_INT, Description: "Minor version number."},
			{Name: "patch", Transform: transform.FromField("Patch"), Type: proto.ColumnType_INT, Description: "Patch version number."},
			{Name: "build", Type: proto.ColumnType_INT, Description: "Build version number."},
		},
	}
}

func getDatabaseTypeVersion(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_database_type_version.getDatabaseTypeVersion", "connection_error", err)
		return nil, err
	}
	appName := d.KeyColumnQuals["app_name"].GetStringValue()
	addon := d.KeyColumnQuals["addon_id"].GetStringValue()
	id := d.KeyColumnQuals["id"].GetStringValue()

	dbVersion, err := client.DatabaseTypeVersion(appName, addon, id)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_database_type_version.getDatabaseTypeVersion", err)
		return nil, err
	}
	return dbVersion, nil
}
