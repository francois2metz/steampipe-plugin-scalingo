package scalingo

import (
	"context"

	"github.com/Scalingo/go-scalingo/v8"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableScalingoDatabase() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_database",
		Description: "A database is associated to an application.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.AllColumns([]string{"app_name", "addon_id"}),
			Hydrate:    listDatabase,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isAddonTokenError,
			},
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name: "app_name", Type: proto.ColumnType_STRING,
				Transform: transform.FromQual("app_name"), Description: "Name of the app.",
			},
			{
				Name:        "addon_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("addon_id"),
				Description: "ID of the addon.",
			},
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "Unique ID identifying the database.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Creation date of the database."},
			{
				Name:        "resource_id",
				Type:        proto.ColumnType_STRING,
				Description: "Resource reference."},
			{
				Name:        "encryption_at_rest",
				Type:        proto.ColumnType_BOOL,
				Description: "Is encryption at rest enabled on this database."},
			{
				Name:        "force_ssl",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromP(featureValueToBool, "force-ssl"),
				Description: "Is SSL encryption is required.",
			},
			{
				Name:        "publicly_available",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromP(featureValueToBool, "publicly-available"),
				Description: "Is the instance publicly available.",
			},
			{
				Name:        "plan",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the application plan.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "Status of the current database.",
			},
			{
				Name:        "type_id",
				Type:        proto.ColumnType_STRING,
				Description: "Database type ID.",
			},
			{
				Name:        "type_name",
				Type:        proto.ColumnType_STRING,
				Description: "Database type Name.",
			},
			{
				Name:        "version_id",
				Type:        proto.ColumnType_STRING,
				Description: "Database version ID.",
			},
			{
				Name:        "instances",
				Type:        proto.ColumnType_STRING,
				Description: "List of all database instances",
			},
			{
				Name:        "readable_version",
				Type:        proto.ColumnType_STRING,
				Description: "Human readable database version",
			},
			{
				Name:        "periodic_backups_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "True if periodic backups are enabled.",
			},
			{
				Name:        "periodic_backups_scheduled_at",
				Type:        proto.ColumnType_STRING,
				Description: "Hours of the day of the periodic backup (UTC).",
			},
			{
				Name:        "maintenance_window_weekday_utc",
				Type:        proto.ColumnType_INT,
				Description: "Week day of the maintenance window.",
				Transform:   transform.FromField("MaintenanceWindow.WeekdayUTC"),
			},
			{
				Name:        "maintenance_window_starting_hour_utc",
				Type:        proto.ColumnType_INT,
				Description: "Hour of the maintenance window.",
				Transform:   transform.FromField("MaintenanceWindow.StartingHourUTC"),
			},
			{
				Name:        "maintenance_window_duration_in_hour",
				Type:        proto.ColumnType_INT,
				Description: "Duration in hour of the maintenance window.",
				Transform:   transform.FromField("MaintenanceWindow.DurationInHour"),
			},
		},
	}
}

func listDatabase(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_database.listDatabase", "connection_error", err)
		return nil, err
	}
	appName := d.EqualsQuals["app_name"].GetStringValue()
	addon := d.EqualsQuals["addon_id"].GetStringValue()

	db, err := client.DatabaseShow(ctx, appName, addon)
	if err != nil {
		plugin.Logger(ctx).Error("scalingo_database.listDatabase", err)
		return nil, err
	}
	d.StreamListItem(ctx, db)
	return nil, nil
}

func featureValueToBool(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	dbInfo := d.HydrateItem.(scalingo.Database)
	param := d.Param.(string)
	if dbInfo.Features == nil {
		return false, nil
	}
	for i := range dbInfo.Features {
		if dbInfo.Features[i].Name == param {
			return dbInfo.Features[i].Status == "ACTIVATED", nil
		}
	}
	return false, nil
}
